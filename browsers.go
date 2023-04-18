package browser

import (
	"runtime"
)

// Preferred opens the url in the user's prefered browser.
func Preferred(c Config) (cmd string, args []string) {
	switch runtime.GOOS {
	case "windows":
		return "cmd", []string{"/c", "start", c.URL}
	case "darwin":
		return "open", []string{c.URL}
	default: // "linux", "freebsd", "openbsd", "netbsd", ...
		return "xdg-open", []string{c.URL}
	}
}

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

// Chrome opens chrome using cfg.
func Chrome(cfg Config) (cmd string, args []string) {
	cmd, args = chrome()

	if cfg.AsApp {
		cfg.URL = "--app=" + cfg.URL
	}

	args = append(args, cfg.URL)

	if cfg.Private {
		args = append(args, "--incognito")
	}

	return cmd, args
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

// Firefox opens firefox using cfg.
func Firefox(c Config) (cmd string, args []string) {
	cmd, args = firefox()

	if c.AsApp {
		args = append(args, "--kiosk") // just a maximized window
	}

	if c.Private {
		args = append(args, "-private-window") // this sometimes works
	}

	args = append(args, c.URL)

	return cmd, args
}

func edge() (cmd string, args []string) {
	switch runtime.GOOS {
	case "windows":
		return "cmd", []string{"/c", "start", "msedge"}
	case "darwin":
		return "/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge", nil
	default:
		return "microsoft-edge", nil
	}
}

// Edge opens edge using cfg.
func Edge(c Config) (cmd string, args []string) {
	cmd, args = edge()

	if c.AsApp {
		c.URL = "--app=" + c.URL
	}
	args = append(args, c.URL)

	if c.Private {
		args = append(args, "--incognito")
	}

	return cmd, args
}

/*
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
*/
