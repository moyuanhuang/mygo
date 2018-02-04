package main

import (
    "strconv"
    "errors"
)

type Article struct {
    ID int
    Title string
    Content string
}

var articleList = []Article {
  Article{ID: 0, Title: "Article 0", Content: "Article 0 body"},
  Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
}

func getAllArticles() []Article {
    return articleList
}

func getArticleByID(id_ string) (*Article, error) {
    id, _ := strconv.Atoi(id_)
    for _, v := range articleList {
        if v.ID == id {
            return &v, nil
        }
    }
    return nil, errors.New("Aricle NOT FOUND")
}
