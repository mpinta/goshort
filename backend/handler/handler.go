package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mpinta/goshort/backend/config"
	"github.com/mpinta/goshort/backend/data"
	"github.com/mpinta/goshort/backend/exception"
	"github.com/mpinta/goshort/backend/utils"
	"io/ioutil"
	"net/http"
	"time"
)

const MessageRunning = "Goshort running!"
const MessageNotFound = "URL not found!"
const MessageIncorrectUrl = "Incorrect URL format!"
const MessageInternalError = "Internal server error!"
const MessageIncorrectFormat = "Incorrect request format!"
const MessageIncorrectPeriod = "Incorrect time limit!"

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, MessageRunning)
}

func Shorten(c *gin.Context) {
	cfg := config.GetConfig()

	req, err := GetFullUrlBody(c)
	if err != nil {
		exception.Http(c, http.StatusBadRequest, MessageIncorrectFormat)
		return
	}

	err = utils.CheckUrlStructure(req.FullUrl)
	if err != nil {
		exception.Http(c, http.StatusBadRequest, MessageIncorrectUrl)
		return
	}

	shorten, err := GetShortUrl(cfg.ShortUrl.Length)
	if err != nil {
		exception.Http(c, http.StatusInternalServerError, MessageInternalError)
		exception.Internal(err)
		return
	}

	url := data.Url{
		FullUrl:   req.FullUrl,
		ShortUrl:  shorten,
		CreatedAt: time.Now(),
	}

	err = data.Insert(url)
	if err != nil {
		exception.Http(c, http.StatusInternalServerError, MessageInternalError)
		exception.Internal(err)
		return
	}

	c.JSON(http.StatusCreated, ShortUrl{
		ShortUrl: url.ShortUrl,
	})
}

func Limit(c *gin.Context) {
	cfg := config.GetConfig()

	req, err := GetLimitBody(c)
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

	c.JSON(http.StatusCreated, LimitRes{
		ShortUrl:   url.ShortUrl,
		CreatedAt:  url.CreatedAt,
		ValidUntil: url.ValidUntil,
	})
}

func Find(c *gin.Context) {
	url := c.Param("url")

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

	if urls[0].ValidUntil.Before(time.Now()) && urls[0].MinutesValid != 0 {
		exception.Http(c, http.StatusNotFound, MessageNotFound)
		return
	}

	c.JSON(http.StatusOK, FullUrl{
		FullUrl: urls[0].FullUrl,
	})
}

func GetFullUrlBody(c *gin.Context) (FullUrl, error) {
	var request FullUrl

	req, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		return request, err
	}

	err = json.Unmarshal(req, &request)
	if err != nil {
		return request, err
	}

	return request, nil
}

func GetLimitBody(c *gin.Context) (LimitReq, error) {
	var request LimitReq

	req, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		return request, err
	}

	err = json.Unmarshal(req, &request)
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
