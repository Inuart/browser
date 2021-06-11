package main

import "github.com/Inuart/browser"

func main() {
	if err := browser.Preferred("https://google.com"); err != nil {
		println(err.Error())
	}
}
