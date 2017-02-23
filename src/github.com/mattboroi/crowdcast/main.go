package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

type Page struct {
    Title string
    Body  []byte
}

func loadPage(title string) (*Page, error) {
    filename := title + ".html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is a fucking server that's serving you up %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "%s", p.Body)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/view/", viewHandler)
    http.ListenAndServe(":8080", nil)
}
