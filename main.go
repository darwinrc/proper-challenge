package main

import (
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
		amount int
	)

	if len(os.Args) > 1 {
		val := strings.Split(os.Args[1], "=")[1]
		amount, err = strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		amount = 10
	}

	web := web.Web{
		Url: "https://icanhas.cheezburger.com",
	}
	if err := web.FetchDocument(); err != nil {
		log.Fatal(err)
	}

	images, err := web.GetImages(".mu-content-card", amount)
	if err != nil {
		log.Fatal(err)
	}

	dir := "img/"
	if err := file.MkDir(dir); err != nil {
		log.Fatal(err)
	}

	for _, image := range images {
		if err := image.Store(dir); err != nil {
			log.Fatal(err)
		}
	}
}
