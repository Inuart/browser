package main

import "github.com/Inuart/browser"

func main() {
	browser.Incognito = true
	if err := browser.AppOrTab("http://google.com"); err != nil {
		println(err.Error())
	}
}
