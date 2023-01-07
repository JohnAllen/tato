package main

import (
	"fmt"
	"log"
	"net/http"
)
func main() {
    fmt.Println("Now Listening on 443")
    http.HandleFunc("/", serveFiles)
    log.Fatal(http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/rarelanguages.net/fullchain.pem", "/etc/letsencrypt/live/rarelanguages.net/privkey.pem", nil))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    p := "static/" + r.URL.Path
    if p == "./" {
        p = "./static/index.html"
    }
    http.ServeFile(w, r, p)
}
