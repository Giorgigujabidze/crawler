package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHtml(rawUrl string) (string, error) {
	req, err := http.NewRequest("GET", rawUrl, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errors.New(res.Status)
	}
	if !strings.Contains(res.Header.Get("content-type"), "text/html") {
		fmt.Println(res.Header.Get("content-type"))
		return "", errors.New(res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
