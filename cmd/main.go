package main

import "bitbucket.org/Inuart/browser"

func main() {
	if err := browser.App("http://google.com"); err != nil {
		println(err.Error())
	}
}
