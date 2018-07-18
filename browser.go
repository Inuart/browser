// +build !appengine

package browser

import (
	"errors"
	"os/exec"
	"runtime"
)

// Tab opens the URL in a default browser tab
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

// App opens the URL in a Chrome app
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

// AppOrTab opens the URL in a Chrome app, a preinstalled
// browser app, or default browser tab
func AppOrTab(url string) error {
	/*var err error
	for _, f := range []func(string) error{App, DefaultApp, Tab} {
		if err = f(url); err != nil {
			continue
		}
	}
	return nil*/
	if err := App(url); err != nil {
		return Tab(url)
	}
	return nil
}

// DefaultApp opens the app in the preinstalled browser
func DefaultApp(url string) error {
	switch runtime.GOOS {
	case "windows": // todo
		return exec.Command("cmd", "/c", "start", "microsoft-edge:"+url).Start()
	case "darwin":
	default:
	}
	return errors.New("not implemented")
}
