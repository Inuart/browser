package main

import "bitbucket.org/Inuart/browser"

func main() {
	browser.Incognito = true
	if err := browser.AppOrTab("http://google.com"); err != nil {
		println(err.Error())
	}
}
