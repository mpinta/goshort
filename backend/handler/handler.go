package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"goshort/backend/data"
	"goshort/backend/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, "goshort running!")
}

func Shorten(c *gin.Context) {
	req, err := GetRequestBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong data format!")
		return
	}

	url := data.Url {
		CreatedAt: time.Now(),
		FullUrl:   req.Url,
		ShortUrl:  utils.GetRandomString(5),
	}

	err = data.Insert(url)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, Body{ Url: url.ShortUrl })
}

func Find(c *gin.Context) {
	url := strings.Replace(c.Request.RequestURI, "/", "", -1)

	urls, err := data.Get(url)
	if err != nil {
		panic(err)
	}

	if len(urls) != 1 {
		c.JSON(http.StatusNotFound, "Url not found!")
		return
	}

	http.Redirect(c.Writer, c.Request, urls[0].FullUrl, http.StatusFound)
}

func GetRequestBody(c *gin.Context) (Body, error) {
	var body Body

	aBody, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		return body, err
	}

	err = json.Unmarshal(aBody, &body)
	if err != nil {
		return body, err
	}

	return body, nil
}
