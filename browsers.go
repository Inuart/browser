//go:build !appengine
// +build !appengine

package browser

import (
	"errors"
	"runtime"
)

func chrome() (cmd string, args []string) {
	switch runtime.GOOS {
	case "windows":
		return "cmd", []string{"/c", "start", "chrome"}
	case "darwin":
		return "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", nil
	default:
		return "google-chrome", nil
	}
}

// Chrome opens the url in Chrome.
func (c Config) Chrome(url string) error {
	cmd, args := chrome()

	if c.AsApp {
		url = "--app=" + url
	}
	args = append(args, url)

	if c.Private {
		args = append(args, "--incognito")
	}

	return c.start(cmd, args...)
}

// Chrome opens the url in Chrome.
func Chrome(url string) error {
	return DefaultCfg.Chrome(url)
}

func firefox() (cmd string, args []string) {
	switch runtime.GOOS {
	case "windows":
		return "cmd", []string{"/c", "start", "firefox"}
	case "darwin":
		return "/Applications/Firefox.app/Contents/MacOS/Firefox", nil
	default:
		return "firefox", nil
	}
}

// Firefox opens the URL in Firefox. Needs testing.
func (c Config) Firefox(url string) error {
	cmd, args := firefox()

	if c.AsApp {
		args = append(args, "--kiosk") // just a maximized window
	}
	if c.Private {
		args = append(args, "-private-window") // this sometimes works
	}
	args = append(args, url)

	return c.start(cmd, args...)
}

// Firefox opens the URL in Firefox.
func Firefox(url string) error {
	return DefaultCfg.Firefox(url)
}

func edge() (cmd string, args []string) {
	switch runtime.GOOS {
	case "windows":
		return "cmd", []string{"/c", "start", "microsoft-edge"}
	case "darwin":
		return "/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge", nil
	default:
		return "microsoft-edge", nil
	}
}

// Edge opens the URL in Edge.
func (c Config) Edge(url string) error {
	cmd, args := edge()

	if c.AsApp {
		url = "--app=" + url
	}
	args = append(args, url)

	if c.Private {
		args = append(args, "--incognito")
	}

	return c.start(cmd, args...)
}

// Edge opens the URL in Edge.
func Edge(url string) error {
	return DefaultCfg.Edge(url)
}

// Safari opens the URL in Safari.
func (c Config) Safari(url string) error {
	if runtime.GOOS != "darwin" {
		return errors.New("safari not available in " + runtime.GOOS)
	}
	// Safari does not have a cmd line interface.
	return c.start("open", "-a", "Safari", url)
}

// Safari opens the URL in Safari.
func Safari(url string) error {
	return DefaultCfg.Safari(url)
}