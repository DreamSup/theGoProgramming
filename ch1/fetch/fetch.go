package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	uu := url.URL{}
	proxy, _ := uu.Parse("http://127.0.0.1:10809")
	for _, u := range os.Args[1:] {
		if !(strings.HasPrefix(u, "http://") || strings.HasPrefix(u, "https://")) {
			u = "http://" + u
		}
		myHttp := http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
		resp, err := myHttp.Get(u)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(resp.Status)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", u, err)
		}
		defer resp.Body.Close()
	}
}
