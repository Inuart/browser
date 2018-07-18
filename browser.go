// +build !appengine

package browser

import (
	"os/exec"
	"runtime"
)

// Tab the URL in a browser
func Tab(url string) error {
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
	var args []string
	switch runtime.GOOS {
	case "windows":
		args = []string{"cmd", "/c", "start", "chrome"}
	case "darwin":
		args = []string{`/Applications/Google Chrome.app/Contents/MacOS/Google Chrome`}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		args = []string{"google-chrome"}
	}
	args = append(args, "--app="+url)
	return exec.Command(args[0], args[1:]...).Start()
}

// AppOrTab opens the URL in a browser app or, if it fails, a tab
func AppOrTab(url string) error {
	if err := App(url); err != nil {
		return Tab(url)
	}
	return nil
}
