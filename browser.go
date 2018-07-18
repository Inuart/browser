package browser

import (
	"os/exec"
	"runtime"
)

// Open the URL in a browser
func Open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

// App opens the URL in a browser app
func App(url string) error {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("cmd", "/c", "start", "chrome", "--app="+url).Start()
	case "darwin":
	}
	// "linux", "freebsd", "openbsd", "netbsd"
	return exec.Command("google-chrome", "--app="+url).Start()
}

// AppOrTab opens the URL in a browser app or, if it fails, a tab
func AppOrTab(url string) error {
	if err := App(url); err != nil {
		return Open(url)
	}
	return nil
}
