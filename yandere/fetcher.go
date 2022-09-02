package yandere

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

const baseUrl = "https://yande.re/post/show"

// Download By show Id from Yandere
func DownloadByShowId(showId, targetDir string) {
	// check exsits
	pngName := fmt.Sprintf("%s.png", showId)
	jpgName := fmt.Sprintf("%s.jpg", showId)

	// TODO
	fmt.Println(pngName, jpgName)

	url := fmt.Sprintf("%s/%s", baseUrl, showId)
	c := colly.NewCollector(
		colly.UserAgent(""),
	)

	c.OnHTML("", func(h *colly.HTMLElement) {

	})

	c.Visit(url)

}
