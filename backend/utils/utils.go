package utils

import (
	"go/build"
	"math/rand"
	"net/url"
	"path"
	"time"
)

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetRandomString(l int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, l)
	for i := range b {
		b[i] = Charset[rand.Intn(len(Charset))]
	}

	return string(b)
}

func CheckUrlStructure(u string) error {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return err
	}
	return nil
}

func GetApplicationPath() string {
	return path.Join(build.Default.GOPATH, "/src/github.com/mpinta/goshort/backend")
}
