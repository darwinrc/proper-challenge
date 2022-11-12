package main

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	res, err := http.Get("http://icanhas.cheezburger.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	path := "img"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	doc.Find(".mu-content-card").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i > 9 {
			return false
		}

		src, _ := s.Find("img").Attr("data-src")

		log.Printf("%v: %v \n", i, src)

		file := "img/" + strconv.Itoa(i+1) + ".jpg"

		f, err := os.Create(file)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		res, err := http.Get(src)

		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		_, err = io.Copy(f, res.Body)

		if err != nil {
			log.Fatal(err)
		}

		return true
	})
}
