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
	return exec.Command("google-chrome", "--app="+url).Start()
}
