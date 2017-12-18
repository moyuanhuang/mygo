package main

import (
    "fmt"
    "regexp"
    "io/ioutil"
    "net/http"
    "html/template"
)

type Page struct {
    Title string
    Body []byte
}
// validation
// A "raw string literal" is quoted by ` backtick characters. There are no special characters in a raw string literal
var validPath = regexp.MustCompile(`^\/(view|edit|save)\/([a-zA-Z0-9]+)$`)

const DATA_PATH = "./data/"
const TEMPLATE_PATH = "./templates/"
// The function template.Must is a convenience wrapper that panics when passed a non-nil error value, and otherwise returns the *Template unaltered.
// ** A panic is appropriate here; if the templates can't be loaded the only sensible thing to do is EXIT the program.
var templates = template.Must(template.ParseFiles(TEMPLATE_PATH + "edit.html", TEMPLATE_PATH + "view.html"))

func (p *Page) Save() error {
    fileName := DATA_PATH + p.Title + ".txt"
    return ioutil.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    body, err := ioutil.ReadFile(DATA_PATH + title + ".txt")
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

// handler functions
func handler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Received request for %s", req.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, req *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        // func Redirect(w ResponseWriter, r *Request, url string, code int)
        http.Redirect(w, req, "/edit/" + title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, req *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, req *http.Request, title string) {
    // func (r *Request) FormValue(key string) string
    // FormValue踩坑: https://yanyiwu.com/work/2015/01/07/golang-http-formvalue.html
    body := req.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.Save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, req, "/view/"+title, http.StatusFound)
}

// helper
func renderTemplate(w http.ResponseWriter, temp string, p *Page) {
    err := templates.ExecuteTemplate(w, temp + ".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        m := validPath.FindStringSubmatch(req.URL.Path)
        if m == nil {
            fmt.Println(req.URL.Path)
            http.NotFound(w, req)
            return
        }
        fn(w, req, m[2])
    }
}

func main() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))
    http.ListenAndServe(":8080", nil)
}
