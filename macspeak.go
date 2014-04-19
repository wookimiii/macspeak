package main

import (
        "fmt"
        "net/http"
        "os/exec"
        "net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/speak":
                query, _ := url.ParseQuery(r.URL.RawQuery)
                words := query["words"][0]
                say := exec.Command("say", words)
                say.Run()
                fmt.Fprintf(w, "Spoke: %s", words)
        case "/ping":
                fmt.Fprintf(w, "200")
        default:
                fmt.Fprintf(w, "404")
        }
}

func main() {
        http.HandleFunc("/", handler);
        http.ListenAndServe(":8080", nil);
}

