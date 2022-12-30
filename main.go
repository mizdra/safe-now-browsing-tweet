package main

import (
	"fmt"
	"net/http"
)

func IsPublicURL(url string) bool {
	resp, err := http.Get(url) // TODO: set UA
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

func main() {
	fmt.Println("200", IsPublicURL("https://httpstat.us/200"))
	fmt.Println("201", IsPublicURL("https://httpstat.us/201"))
	fmt.Println("404", IsPublicURL("https://httpstat.us/404"))
}
