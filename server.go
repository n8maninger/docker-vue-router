package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"
)

func serveContent(w http.ResponseWriter, r *http.Request, name string, f *os.File) {
	s, err := f.Stat()
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, name, s.ModTime(), f)
}

func main() {
	base := os.Args[1]
	indexPage, err := os.Open(filepath.Join(base, "index.html"))
	if err != nil {
		log.Fatalf("unable to open index page: %s", err)
	}
	defer indexPage.Close()

	log.Println("server started")

	errChan := make(chan error, 1)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		errChan <- nil
	}()

	go func() {
		errChan <- http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			if !strings.HasPrefix(r.URL.Path, "/") {
				r.URL.Path = "/" + r.URL.Path
			}

			if r.URL.Path == "/" {
				serveContent(w, r, "index.html", indexPage)
				return
			}

			path := filepath.Clean(filepath.Join(base, r.URL.Path))
			log.Println(path)

			f, err := os.Open(path)
			if err != nil {
				serveContent(w, r, "index.html", indexPage)
				return
			}
			defer f.Close()

			serveContent(w, r, r.URL.Path, f)
			log.Printf("%s (%s)", r.URL.Path, time.Since(start))
		}))
	}()

	if err := <-errChan; err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server error: %s", err)
	}

	log.Println("server shutdown")
}
