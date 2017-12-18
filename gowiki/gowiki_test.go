package main

import (
    // "fmt"
    "testing"
    "net/http"
    "io/ioutil"
)

// This is the tests for Page related functions
func TestSavePage(t *testing.T){
    page := &Page{Title: "test", Body: []byte("This is a test page!")}
    if err := page.Save(); err != nil {
        t.Error("Can't save page!")
    } else {
        t.Log("Successfully saved page.")
    }
}

func TestPageLoad(t *testing.T){
    if page, err := loadPage("test"); err != nil {
        t.Error("Can't load page file!")
    } else {
        t.Log(string(page.Body))
    }
}

// This is the tests for http handlers
func TestHandler(t *testing.T) {
    http.HandleFunc("/", handler)
    go http.ListenAndServe(":8080", nil)

    resp, err := http.Get("http://localhost:8080/monkey_king")
    defer resp.Body.Close()

    if err != nil {
        t.Error("Can't get response from the server")
    }

    bodyBytes, _ := ioutil.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    if bodyString != "Received request for monkey_king" {
        t.Error("Wrong response!", resp)
    }
    t.Log(bodyString)
}
