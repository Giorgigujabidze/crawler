package main

import (
	"net/url"
	"path"
)

func normalizeUrl(rawUrl string) (string, error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	retVal, _ := url.JoinPath(parsedUrl.Host, path.Clean(parsedUrl.Path))
	return retVal, nil
}
