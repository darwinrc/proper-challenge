package web

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"proper-challenge/app/file"
)

var (
	Error = errors.New("web error: %v")
)

// Web represents a web page to be crawled
type Web struct {
	Url string
	Doc *goquery.Document
}

// FetchPage requests the web URL and generates a goquery.Document from
// the body response with the corresponding page structure
func (w *Web) FetchPage() error {
	log.Printf("fetching page: %v", w.Url)

	body, err := HttpGet(w.Url)
	defer body.Close()
	if err != nil {
		return fmt.Errorf(Error.Error(), err)
	}

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return fmt.Errorf(Error.Error(), err)
	}

	w.Doc = doc

	return nil
}

// GetPageImages returns a slice of 'amount' images (name, url and data for a file)
// after processing the goquery.Document and extracting them from the '<img>'
// children from the specified selector
func (w *Web) GetPageImages(selector string, amount, page, perPage int) []*file.File {
	var images []*file.File

	w.Doc.Find(selector).EachWithBreak(func(i int, s *goquery.Selection) bool {
		item := (page-1)*perPage + i + 1
		if item > amount {
			return false
		}

		url, _ := s.Find("img").Attr("data-src")
		name := fmt.Sprintf("%d.jpg", item)
		image := &file.File{
			Name: name,
			Url:  url,
		}
		log.Printf("getting image: %v", image)
		images = append(images, image)

		return true
	})

	return images
}

// HttpGet is just a wrapper for http.Get
func HttpGet(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf(Error.Error(), err)
	}

	return res.Body, nil
}
