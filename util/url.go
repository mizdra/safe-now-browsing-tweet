package util

import "net/http"

func IsPublicURL(url string) bool {
	resp, err := http.Get(url) // TODO: set UA
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}
