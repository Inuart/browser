package browser

import (
	"os/exec"
	"runtime"
)

type Config struct {
	AsApp   bool
	Private bool
}

var DefaultCfg Config

func (c Config) start(cmd string, args ...string) error {
	return exec.Command(cmd, args...).Start()
}

// Preferred opens the url in the user prefered browser.
func (c Config) Preferred(url string) error {
	switch runtime.GOOS {
	case "windows":
		return c.start("cmd", "/c", "start", url)
	case "darwin":
		return c.start("open", url)
	default: // "linux", "freebsd", "openbsd", "netbsd", ...
		return c.start("xdg-open", url)
	}
}

// Preferred opens the url in the user prefered browser.
func Preferred(url string) error {
	return DefaultCfg.Preferred(url)
}

// Default opens the url in the default browser.
func (c Config) Default(url string) error {
	switch runtime.GOOS {
	case "windows":
		return c.Edge(url)
	case "darwin":
		return c.Safari(url)
	default:
		return c.Firefox(url)
	}
}

// Default opens the url in the default browser.
func Default(url string) error {
	return DefaultCfg.Default(url)
}
