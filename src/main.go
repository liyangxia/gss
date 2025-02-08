// Copyright 2025 liyangxia.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// loggingMiddleware logs request details and disables caching
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s - - [%s] \"%s %s %s\" %d %d \"%s\" \"%s\"",
			r.RemoteAddr,
			start.Format("02/Jan/2006:15:04:05 -0700"),
			r.Method,
			r.RequestURI,
			r.Proto,
			http.StatusOK,
			r.ContentLength,
			r.Referer(),
			r.UserAgent(),
		)
	})
}

func main() {
	port := flag.String("port", "8080", "Port to serve on")
	lan := flag.Bool("lan", false, "Allow LAN access")
	flag.Usage = func() {
		fmt.Printf("Usage: gss [--port PORT] [--lan]\n\nOptions:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	if _, err := os.Open(currentDir); err != nil {
		log.Fatalf("Cannot read current directory: %v", err)
	}

	fileServer := http.FileServer(http.Dir(currentDir))
	http.Handle("/", loggingMiddleware(fileServer))

	address := "localhost:" + *port
	if *lan {
		address = "0.0.0.0:" + *port
	}

	fmt.Printf("Serving directory: %s on http://%s\n", currentDir, address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Server encountered an error: %v", err)
	}
}
