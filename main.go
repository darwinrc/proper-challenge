package main

import (
	"log"
	"proper-challenge/app/web"
)

func main() {
	web := web.Web{
		Url: "https://icanhas.cheezburger.com",
	}
	if err := web.FetchDocument(); err != nil {
		log.Fatal(err)
	}

	images, err := web.GetImages(".mu-content-card", 10)
	if err != nil {
		log.Fatal(err)
	}

	for _, image := range images {
		if err := image.Store("img/"); err != nil {
			log.Fatal(err)
		}
	}
}
