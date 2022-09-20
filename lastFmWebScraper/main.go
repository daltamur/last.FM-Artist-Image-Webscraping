package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)
import (
	"github.com/gocolly/colly"
)

func main() {
	var imageAddress string
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
		fmt.Println(r.StatusCode)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnHTML("div[class=header-new-gallery-outer]", func(e *colly.HTMLElement) {
		if e.ChildAttr("a", "class") == "header-new-gallery\n                            header-new-gallery--link\n                            hidden-xs\n                            link-block-target" {
			imageAddress = e.ChildAttr("a", "href")
			positionOfID := strings.LastIndex(imageAddress, "/") + 1
			imageAddress = "https://lastfm.freetls.fastly.net/i/u/174s/" + imageAddress[positionOfID:]
		}

	})
	lastFmAddress := new(string)
	fmt.Println("Put in address")
	fmt.Scanf("%s", lastFmAddress)
	c.Visit(*lastFmAddress)
	if imageAddress == "" {
		fmt.Println("Last.fm does not have an image for this artist")
	} else {
		err := exec.Command("xdg-open", imageAddress).Start()
		if err != nil {
			fmt.Println(err)
		}
	}
}
