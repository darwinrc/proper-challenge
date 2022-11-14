package main

import (
	"flag"
	"log"
	"proper-challenge/app"
)

func main() {
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

	app := app.App{}
	images := app.GetImages(*amount, *threads, perPage)
	app.StoreImages(images, *threads)
}
