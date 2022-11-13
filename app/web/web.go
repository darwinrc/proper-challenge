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

// FetchDocument requests the web URL and generates a goquery.Document from the body response
// with the corresponding page structure
func (w *Web) FetchDocument() error {
	body, err := httpGet(w.Url)
	if err != nil {
		return fmt.Errorf(Error.Error(), err)
	}
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return fmt.Errorf(Error.Error(), err)
	}

	w.Doc = doc

	return nil
}

// GetImages returns an array of 'amount' images (name, url and data for a file) after processing the goquery.Document
// and extracting them from the '<img>' children from the specified selector
func (w *Web) GetImages(selector string, amount int) (images []*file.File, err error) {
	w.Doc.Find(selector).EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i+1 > amount {
			return false
		}

		url, _ := s.Find("img").Attr("data-src")
		body, err := httpGet(url)
		if err != nil {
			err = fmt.Errorf(Error.Error(), err)
		}

		name := fmt.Sprintf("%d.jpg", i+1)
		image := file.File{
			Name: name,
			Url:  url,
			Data: body,
		}

		log.Printf("Getting image: %v", image)
		images = append(images, &image)

		return true
	})

	return
}

func httpGet(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf(Error.Error(), err)
	}

	return res.Body, nil
}
