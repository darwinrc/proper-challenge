package app

import (
	"fmt"
	"log"
	"proper-challenge/app/file"
	"proper-challenge/app/web"
	"sync"
)

type App struct{}

// GetImages returns a slice of images (files) depending on the 'amount' and items 'per page'
// It executes using the number of concurrent goroutines limited by 'thread'
func (a *App) GetImages(amount, threads, perPage int) []*file.File {
	pages := amount / perPage
	if amount%perPage != 0 {
		pages += 1
	}

	var images []*file.File
	ch := make(chan []*file.File, threads)
	baseUrl := "https://icanhas.cheezburger.com"

	for p := 1; p <= pages; p++ {
		url := baseUrl
		if p > 1 {
			url = fmt.Sprintf("%s/page/%d", baseUrl, p)
		}
		web := web.Web{
			Url: url,
		}

		go func(ch chan []*file.File) {
			if err := web.FetchPage(); err != nil {
				log.Fatal(err)
			}

			pageImages := web.GetPageImages(".mu-content-card", amount, p, perPage)
			ch <- pageImages
		}(ch)

		pageImages := <-ch
		images = append(images, pageImages...)
	}

	return images
}

// StoreImages gets the images stream data and saves them to local file system as files
// It executes using the number of concurrent goroutines limited by 'thread'
func (a *App) StoreImages(images []*file.File, threads int) {
	dir := "img/"
	if err := file.MkDir(dir); err != nil {
		log.Fatal(err)
	}

	wg := new(sync.WaitGroup)
	ch := make(chan struct{}, threads)

	for _, img := range images {
		ch <- struct{}{}
		wg.Add(1)

		go func(wg *sync.WaitGroup, img *file.File) {
			defer wg.Done()
			var err error

			img.Data, err = web.HttpGet(img.Url)
			if err != nil {
				log.Fatal(err)
			}
			if err = img.Store(dir); err != nil {
				log.Fatal(err)
			}
			img.Data.Close()
			<-ch
		}(wg, img)
	}
	wg.Wait()
}
