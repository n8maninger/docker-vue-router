package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	base := os.Args[1]
	indexPage, err := os.Open(filepath.Join(base, "index.html"))
	if err != nil {
		log.Fatalf("unable to open index page: %s", err)
	}
	defer indexPage.Close()

	err = http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/") {
			r.URL.Path = "/" + r.URL.Path
		}
		if r.URL.Path == "/" {
			http.ServeContent(w, r, "index.html", time.Time{}, indexPage)
			return
		}

		path := filepath.Clean(filepath.Join(base, r.URL.Path))
		f, err := os.Open(path)
		if err != nil {
			http.ServeContent(w, r, "index.html", time.Time{}, indexPage)
			return
		}
		defer f.Close()
		http.ServeContent(w, r, r.URL.Path, time.Time{}, f)
	}))
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server error: %s", err)
	}
}
