package main

import "github.com/Inuart/browser"

func main() {
	cfg := browser.Config{
		Private: true,
		AsApp:   true,
	}
	if err := cfg.Chrome("http://google.com"); err != nil {
		println(err.Error())
	}
}
