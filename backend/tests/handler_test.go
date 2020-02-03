package tests

import (
	"bytes"
	"encoding/json"
	"github.com/mpinta/goshort/backend/config"
	"github.com/mpinta/goshort/backend/data"
	"github.com/mpinta/goshort/backend/handler"
	"io/ioutil"
	"net/http"
	"path"
	"testing"
	"time"
)

var shorten string
var limit string

func TestStatus(t *testing.T) {
	cfg := config.GetConfig()

	res, err := http.Get(cfg.Server.Host + cfg.Server.Port + cfg.Server.RootEndpoint + cfg.Server.StatusEndpoint)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}

func TestShorten(t *testing.T) {
	cfg := config.GetConfig()
	var reqBody = []byte (`{"full_url": "https://www.github.com/"}`)

	res, err := http.Post(cfg.Server.Host+cfg.Server.Port+cfg.Server.RootEndpoint+cfg.Server.ShortenEndpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var resBody handler.ShortUrl
	err = json.Unmarshal(response, &resBody)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusCreated)
	}

	if len(path.Base(resBody.ShortUrl)) != cfg.ShortUrl.Length {
		t.Errorf("Short URL length is incorrect, got: %d, want: %d.", len(resBody.ShortUrl), cfg.ShortUrl.Length)
	}

	shorten = resBody.ShortUrl
}

func TestShortenIncorrectFormat(t *testing.T) {
	cfg := config.GetConfig()
	var reqBody = []byte (`{"url": "https://www.github.com/"}`)

	res, err := http.Post(cfg.Server.Host+cfg.Server.Port+cfg.Server.RootEndpoint+cfg.Server.ShortenEndpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusBadRequest)
	}
}

func TestShortenIncorrectUrl(t *testing.T) {
	cfg := config.GetConfig()
	var reqBody = []byte (`{"full_url": "www.github.com"}`)

	res, err := http.Post(cfg.Server.Host+cfg.Server.Port+cfg.Server.RootEndpoint+cfg.Server.ShortenEndpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusBadRequest)
	}
}

func TestLimit(t *testing.T) {
	cfg := config.GetConfig()
	var reqBody = []byte (`{"full_url": "https://www.github.com/", "minutes_valid": 1}`)

	res, err := http.Post(cfg.Server.Host+cfg.Server.Port+cfg.Server.RootEndpoint+cfg.Server.LimitEndpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var resBody handler.LimitRes
	err = json.Unmarshal(response, &resBody)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusCreated)
	}

	if len(path.Base(resBody.ShortUrl)) != cfg.ShortUrl.Length {
		t.Errorf("Short URL length is incorrect, got: %d, want: %d.", len(resBody.ShortUrl), cfg.ShortUrl.Length)
	}

	timeLimit := resBody.CreatedAt.Add(time.Minute * time.Duration(1))
	if timeLimit.Equal(resBody.ValidUntil) {
		t.Errorf("Limit period is incorrect, got: %v, want: %v.", timeLimit, resBody.ValidUntil)
	}

	limit = resBody.ShortUrl
}

func TestLimitIncorrectFormat(t *testing.T) {
	cfg := config.GetConfig()
	var reqBody = []byte (`{"url": "https://www.github.com/", "minutes": 1}`)

	res, err := http.Post(cfg.Server.Host+cfg.Server.Port+cfg.Server.RootEndpoint+cfg.Server.LimitEndpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusBadRequest)
	}
}

func TestLimitIncorrectUrl(t *testing.T) {
	cfg := config.GetConfig()
	var reqBody = []byte (`{"full_url": "www.github.com", "minutes_valid": 1}`)

	res, err := http.Post(cfg.Server.Host+cfg.Server.Port+cfg.Server.RootEndpoint+cfg.Server.LimitEndpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusBadRequest)
	}
}

func TestLimitIncorrectPeriod(t *testing.T) {
	cfg := config.GetConfig()
	var reqBody = []byte (`{"full_url": "https://www.github.com/", "minutes_valid": 0}`)

	res, err := http.Post(cfg.Server.Host+cfg.Server.Port+cfg.Server.RootEndpoint+cfg.Server.LimitEndpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusBadRequest)
	}
}

func TestFindShorten(t *testing.T) {
	cfg := config.GetConfig()

	res, err := http.Get(cfg.Server.Host + cfg.Server.Port + cfg.Server.RootEndpoint +
		cfg.Server.FindEndpoint + "/" + shorten)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}

func TestFindLimit(t *testing.T) {
	cfg := config.GetConfig()

	res, err := http.Get(cfg.Server.Host + cfg.Server.Port + cfg.Server.RootEndpoint +
		cfg.Server.FindEndpoint + "/" + limit)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}

func TestFindNotFoundUrl(t *testing.T) {
	cfg := config.GetConfig()

	res, err := http.Get(cfg.Server.Host + cfg.Server.Port + cfg.Server.RootEndpoint +
		cfg.Server.FindEndpoint + "/" + "myurl")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusNotFound)
	}
}

func TestFindNotFoundLimit(t *testing.T) {
	cfg := config.GetConfig()
	time.Sleep(65 * time.Second)

	res, err := http.Get(cfg.Server.Host + cfg.Server.Port + cfg.Server.RootEndpoint +
		cfg.Server.FindEndpoint + "/" + limit)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusNotFound)
	}
}

func TestDataFoundShorten(t *testing.T) {
	cfg := config.GetConfig()
	var reqBody = []byte (`{"full_url": "https://www.github.com/"}`)

	res, err := http.Post(cfg.Server.Host+cfg.Server.Port+cfg.Server.RootEndpoint+cfg.Server.ShortenEndpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	res, err = http.Get(cfg.Server.Host + cfg.Server.Port + cfg.Server.RootEndpoint +
		cfg.Server.FindEndpoint + "/" + shorten)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}

func TestDataUrlsCount(t *testing.T) {
	var urls []data.Url

	db, err := data.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Model(&data.Url{}).Select("*").Rows()
	if err != nil {
		t.Fatal(err)
	}

	urls, err = data.GetUrlsFromRows(db, rows)
	if err != nil {
		t.Fatal(err)
	}

	if len(urls) != 2 {
		t.Errorf("Urls count is incorrect, got: %d, want: %d.", len(urls), 2)
	}
}
