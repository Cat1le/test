package main

import (
	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var Url = "https://hypeauditor.com/top-instagram-all-russia/"

func main() {
	resp, err := http.Get(Url)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	open, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(open)
	doc.Find("div.table div.row__top").Map(func(i int, selection *goquery.Selection) string {
		var avatar, _ = selection.Find("img.avatar__img").Attr("src")
		var nick = selection.Find("div.contributor__name-content").Text()
		var name = selection.Find("div.contributor__title").Text()
		var categories = selection.Find("div.tag__content").Map(func(i int, selection *goquery.Selection) string {
			return selection.Text()
		})
		var followers = selection.Find("div.subscribers").Text()
		var audience = selection.Find("div.audience").Text()
		var engAuth = selection.Find("div.authentic").Text()
		var engAvg = selection.Find("div.engagement").Text()
		err := writer.Write([]string{
			strconv.Itoa(i + 1),
			avatar,
			nick,
			name,
			strings.Join(categories, ";"),
			followers,
			audience,
			engAuth,
			engAvg,
		})
		if err != nil {
			panic(err)
		}
		return ""
	})
	writer.Flush()
}
