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
		target, err := url.Parse(c.URL)
		if err != nil {
			return err
		}
		lh := *target
		lh.Scheme = "http"
		lh.Host = "localhost:8080"
		c.URL = lh.String()
		target.Path = ""
		go func() {
			if err := modifyResponse(ctx, target, c.ModifyResponse); err != nil {
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

func modifyResponse(ctx context.Context, target *url.URL, rm ResponseModifier) error {
	rp := httputil.NewSingleHostReverseProxy(target)
	shrp := rp.Director
	rp.Director = func(req *http.Request) {
		req.Host = target.Host
		shrp(req)
	}
	rp.ModifyResponse = rm
	http.Handle("/", rp)

	addr := ":8080"
	var lc net.ListenConfig
	ln, err := lc.Listen(ctx, "tcp", addr)
	if err != nil {
		return err
	}
	srv := http.Server{Addr: addr, Handler: rp}
	return srv.Serve(ln)
}
