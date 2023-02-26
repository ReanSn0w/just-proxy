package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	proxyURL := os.Getenv("PROXY_URL")
	if proxyURL == "" {
		log.Fatal("PROXY_URL environment variable is empty")
	}

	url, err := url.Parse(proxyURL)
	if err != nil {
		log.Fatalln("PROXY_URL error: ", err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handle request:" + r.URL.RequestURI())
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":"+"8080", nil))
}
