package main

import (
	"flag"
	"fmt"
	"log"
	"proper-challenge/app/file"
	"proper-challenge/app/web"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	const (
		perPage = 16
	)

	var (
		amount  = flag.Int("amount", 10, "amount")
		threads = flag.Int("threads", 1, "threads")
	)

	flag.Parse()

	if *threads < 1 || *threads > 5 {
		log.Fatal("threads should be between 1 and 5")
	}

	images := getImages(*amount, *threads, perPage)

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	storeImages(images, *threads)

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// getImages returns a slice of images depending on the amount and items per page
// It executes using the number of concurrent goroutines limited by 'thread'
func getImages(amount, threads, perPage int) (images []*file.File) {
	baseUrl := "https://icanhas.cheezburger.com"

	pages := amount / perPage
	if amount%perPage != 0 {
		pages += 1
	}

	wg := new(sync.WaitGroup)
	ch := make(chan struct{}, threads)

	for p := 1; p <= pages; p++ {
		url := baseUrl
		if p > 1 {
			url = fmt.Sprintf("%s/page/%d", baseUrl, p)
		}
		web := web.Web{
			Url: url,
		}

		ch <- struct{}{}
		wg.Add(1)

		go func(wg *sync.WaitGroup, images *[]*file.File) {
			defer wg.Done()
			if err := web.FetchPage(); err != nil {
				log.Fatal(err)
			}

			pageImages, err := web.GetImages(".mu-content-card", amount, p, perPage)
			if err != nil {
				log.Fatal(err)
			}
			*images = append(*images, pageImages...)
			<-ch
		}(wg, &images)
	}

	wg.Wait()

	return
}

// storeImages gets the images stream data and saves them to local file system as files
// It executes using the number of concurrent goroutines limited by 'thread'
func storeImages(images []*file.File, threads int) {
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
