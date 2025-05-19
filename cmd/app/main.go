package main

import (
	cc "comick-dl/internal/comickclient"
	"context"
	"flag"
	"net/http"
	"os"
)

func main() {

	flag.Parse()

	if flag.NArg() != 1 {
		flag.PrintDefaults()
		os.Exit(2)
	}

	url := flag.Arg(0)

	h := http.Header{
		"User-Agent": {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"},
	}
	client := cc.New(
		cc.WithHeaders(h),
	)

	client.Download(context.Background(), url, "./")

}
