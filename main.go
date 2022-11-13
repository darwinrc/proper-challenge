package main

import (
	"fmt"
	"log"
	"os"
	"proper-challenge/app/file"
	"proper-challenge/app/web"
	"strconv"
	"strings"
)

func main() {
	var (
		err    error
		amount = 10
	)

	const (
		perPage = 16
	)

	if len(os.Args) > 1 {
		val := strings.Split(os.Args[1], "=")[1]
		amount, err = strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
	}

	images, err := getImages(amount, perPage)
	if err != nil {
		log.Fatal(err)
	}

	if err := storeImages(images); err != nil {
		log.Fatal(err)
	}
}

// getImages returns a slice of images depending on the amount and pages
func getImages(amount, perPage int) (images []*file.File, err error) {
	baseUrl := "https://icanhas.cheezburger.com"

	pages := amount / perPage
	if amount%perPage != 0 {
		pages += 1
	}

	for p := 1; p <= pages; p++ {
		url := baseUrl
		if p > 1 {
			url = fmt.Sprintf("%s/page/%d", baseUrl, p)
		}
		web := web.Web{
			Url: url,
		}
		if err := web.FetchPage(); err != nil {
			return nil, err
		}

		pageImages, err := web.GetImages(".mu-content-card", amount, p, perPage)
		if err != nil {
			return nil, err
		}
		images = append(images, pageImages...)
	}

	return
}

// storeImages saves the images to the local file system
func storeImages(images []*file.File) error {
	dir := "img/"
	if err := file.MkDir(dir); err != nil {
		return err
	}

	for _, image := range images {
		if err := image.Store(dir); err != nil {
			return err
		}
		image.Data.Close()
	}

	return nil
}
