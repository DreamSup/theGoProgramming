package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

func main() {
	urli := url.URL{}
	urlproxy, _ := urli.Parse("http://127.0.0.1:10809")
	c := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}
	resp, err := c.Get("https://www.google.com")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Body)
}
