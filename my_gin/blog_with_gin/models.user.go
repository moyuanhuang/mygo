package main

import (
    "strings"
    "errors"
)

type User struct {
    Username string
    Password string
}

var userList = []User {
    User{Username: "user0", Password: "password0"},
    User{Username: "user1", Password: "password1"},
    User{Username: "user2", Password: "password2"},
}

func isUserValid(username string, password string) bool {
    for _, u := range userList {
        if u.Username == username && u.Password == password {
            return true
        }
    }
    return false
}

func registerNewUser(username string, password string) (*User, error) {
    if strings.TrimSpace(password) == "" {
        return nil, errors.New("Password can't be blank!")
    }
    if !usernameIsAvailable(username) {
        return nil, errors.New("Username isn't available!")
    }

    newUser := User{Username: username, Password: password}
    userList = append(userList, newUser)

    return &newUser, nil
}

func usernameIsAvailable (username string) bool {
    for _, v := range userList {
        if v.Username == username {
            return false
        }
    }
    return true
}
