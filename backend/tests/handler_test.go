package tests

import (
	"bytes"
	"encoding/json"
	"goshort/backend/config"
	"goshort/backend/handler"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

var shorten string

func TestStatus(t *testing.T) {
	cfg := config.GetConfig()

	res, err := http.Get("http://" + cfg.Server.RootEndpoint + cfg.Server.Port + cfg.Server.StatusEndpoint)
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
	var reqBody = []byte (`{"full_url": "https://www.github.com/", "minutes_valid": 1}`)

	res, err := http.Post("http://" + cfg.Server.RootEndpoint + cfg.Server.Port + cfg.Server.ShortenEndpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var resBody handler.Response
	err = json.Unmarshal(response, &resBody)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusCreated)
	}

	if len(resBody.ShortUrl) != cfg.ShortUrl.Length {
		t.Errorf("Short URL length is incorrect, got: %d, want: %d.", len(resBody.ShortUrl), cfg.ShortUrl.Length)
	}

	validityPeriod := resBody.CreatedAt.Add(time.Minute * time.Duration(1))
	if validityPeriod.Equal(resBody.ValidUntil) {
		t.Errorf("Validity period is incorrect, got: %v, want: %v.", validityPeriod, resBody.ValidUntil)
	}

	shorten = resBody.ShortUrl
}

func TestShortenIncorrectFormat(t *testing.T) {
	cfg := config.GetConfig()
	var reqBody = []byte (`{"url": "https://www.github.com/", "minutes": 1}`)

	res, err := http.Post("http://" + cfg.Server.RootEndpoint + cfg.Server.Port + cfg.Server.ShortenEndpoint,
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
	var reqBody = []byte (`{"full_url": "www.github.com", "minutes_valid": 1}`)

	res, err := http.Post("http://" + cfg.Server.RootEndpoint + cfg.Server.Port + cfg.Server.ShortenEndpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusBadRequest)
	}
}

func TestShortenIncorrectPeriod(t *testing.T) {
	cfg := config.GetConfig()
	var reqBody = []byte (`{"full_url": "https://www.github.com/", "minutes_valid": 0}`)

	res, err := http.Post("http://" + cfg.Server.RootEndpoint + cfg.Server.Port + cfg.Server.ShortenEndpoint,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusBadRequest)
	}
}

func TestFind(t *testing.T) {
	cfg := config.GetConfig()

	res, err := http.Get("http://" + cfg.Server.RootEndpoint + cfg.Server.Port + "/" + shorten)
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

	res, err := http.Get("http://" + cfg.Server.RootEndpoint + cfg.Server.Port + "/" + "myurl")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusNotFound)
	}
}

func TestFindNotFoundPeriod(t *testing.T) {
	cfg := config.GetConfig()

	time.Sleep(60 * time.Second)

	res, err := http.Get("http://" + cfg.Server.RootEndpoint + cfg.Server.Port + "/" + shorten)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code is incorrect, got: %d, want: %d.", res.StatusCode, http.StatusNotFound)
	}
}
