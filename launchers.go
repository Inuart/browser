package browser

import (
	"context"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
)

type Config struct {
	URL            string
	AsApp          bool
	Private        bool
	ModifyResponse ResponseModifier
}

type BrowserCmd func(c Config) (cmd string, args []string)

func Open(url string) error {
	return Browse(context.Background(), Config{URL: url}, nil)
}

func Browse(ctx context.Context, c Config, b BrowserCmd) error {
	if c.ModifyResponse != nil {
		url := c.URL
		c.URL = "http://localhost:8080"
		go func() {
			if err := postprocess(ctx, url, c.ModifyResponse); err != nil {
				println(err.Error())
			}
		}()
	}
	if b == nil {
		b = Preferred
	}

	cmd, args := b(c)
	return exec.Command(cmd, args...).Start()
}

type ResponseModifier func(*http.Response) error

func postprocess(ctx context.Context, target string, rm ResponseModifier) error {
	u, err := url.Parse(target)
	if err != nil {
		return err
	}

	rp := httputil.NewSingleHostReverseProxy(u)
	rp.ModifyResponse = rm
	http.Handle("/", rp)

	addr := ":8080"
	var lc net.ListenConfig
	ln, err := lc.Listen(context.Background(), "tcp", addr)
	if err != nil {
		return err
	}
	srv := &http.Server{Addr: addr, Handler: rp}
	return srv.Serve(ln)
}
