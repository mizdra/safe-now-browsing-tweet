package main

import (
	"fmt"

	"github.com/mizdra/safe-now-browsing-tweet/util"
)

func main() {
	fmt.Println("200", util.IsPublicURL("https://httpstat.us/200"))
	fmt.Println("201", util.IsPublicURL("https://httpstat.us/201"))
	fmt.Println("404", util.IsPublicURL("https://httpstat.us/404"))
}
