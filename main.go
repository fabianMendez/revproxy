package main

import (
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	target, err := url.Parse(os.Getenv("URL"))
	if err != nil {
		println(err)
		os.Exit(1)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		println(r.Method, r.URL.String(), r.Proto)
		for key, values := range r.Header {
			for _, value := range values {
				println(key, value)
			}
		}
		if r.Body != nil {
			r.Body = io.NopCloser(io.TeeReader(r.Body, os.Stdout))
		}
		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
		r.Host = target.Host
		proxy.ServeHTTP(w, r)
		println()
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println("Listening on port", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		println(err)
		os.Exit(1)
	}
}
