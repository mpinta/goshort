package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"goshort/backend/config"
	"goshort/backend/data"
	"goshort/backend/exception"
	"goshort/backend/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const MessageRunning = "Goshort running!"
const MessageNotFound = "URL not found!"
const MessageIncorrectUrl = "Incorrect URL format!"
const MessageInternalError = "Internal server error!"
const MessageIncorrectFormat = "Incorrect request format!"
const MessageIncorrectPeriod = "Incorrect validity period!"

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, MessageRunning)
}

func Shorten(c *gin.Context) {
	cfg := config.GetConfig()

	req, err := GetRequestBody(c)
	if err != nil {
		exception.Http(c, http.StatusBadRequest, MessageIncorrectFormat)
		return
	}

	err = utils.CheckUrlStructure(req.FullUrl)
	if err != nil {
		exception.Http(c, http.StatusBadRequest, MessageIncorrectUrl)
		return
	}

	if req.MinutesValid < 1 {
		exception.Http(c, http.StatusBadRequest, MessageIncorrectPeriod)
		return
	}

	shorten, err := GetShortUrl(cfg.ShortUrl.Length)
	if err != nil {
		exception.Http(c, http.StatusInternalServerError, MessageInternalError)
		exception.Internal(err)
		return
	}

	url := data.Url{
		FullUrl:      req.FullUrl,
		ShortUrl:     shorten,
		CreatedAt:    time.Now(),
		ValidUntil:   time.Now().Add(time.Minute * time.Duration(req.MinutesValid)),
		MinutesValid: req.MinutesValid,
	}

	err = data.Insert(url)
	if err != nil {
		exception.Http(c, http.StatusInternalServerError, MessageInternalError)
		exception.Internal(err)
		return
	}

	c.JSON(http.StatusCreated, Response{
		ShortUrl:   url.ShortUrl,
		CreatedAt:  url.CreatedAt,
		ValidUntil: url.ValidUntil,
	})
}

func Find(c *gin.Context) {
	url := strings.Replace(c.Request.RequestURI, "/", "", -1)

	urls, err := data.Get(url)
	if err != nil {
		exception.Http(c, http.StatusInternalServerError, MessageInternalError)
		exception.Internal(err)
		return
	}

	if len(urls) != 1 {
		exception.Http(c, http.StatusNotFound, MessageNotFound)
		return
	}

	if urls[0].ValidUntil.Before(time.Now()) {
		exception.Http(c, http.StatusNotFound, MessageNotFound)
		return
	}

	http.Redirect(c.Writer, c.Request, urls[0].FullUrl, http.StatusFound)
	c.JSON(http.StatusOK, urls[0])
}

func GetRequestBody(c *gin.Context) (Request, error) {
	var request Request

	aBody, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		return request, err
	}

	err = json.Unmarshal(aBody, &request)
	if err != nil {
		return request, err
	}

	return request, nil
}

func GetShortUrl(l int) (string, error) {
	var url string

	for {
		url = utils.GetRandomString(l)
		urls, err := data.Get(url)
		if err != nil {
			return url, err
		}

		if len(urls) == 0 {
			break
		}
	}

	return url, nil
}
