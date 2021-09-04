package main

import (
	"context"
	"net/http"
	"time"

	"github.com/Inuart/browser"
)

func main() {
	/*if err := browser.Open("https://google.com"); err != nil {
		println(err.Error())
	}
	if err := browser.Browse(ctx, browser.Config{URL: "https://google.com"},
		browser.Chrome); err != nil {
		println(err.Error())
	}*/

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := browser.Browse(ctx, browser.Config{
		URL: "https://google.com",

		Private: true,
		ModifyResponse: func(r *http.Response) error {
			println(r.Request.URL.String())
			return nil
		},
	}, browser.Chrome)
	if err != nil {
		println(err.Error())
	}
	<-ctx.Done()
}
