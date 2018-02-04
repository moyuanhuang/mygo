package main

import (
    "io/ioutil"
    "crypto/rsa"
    "errors"
    "time"

    jwt "github.com/dgrijalva/jwt-go"
)

const RSA_PATH = "./demo.rsa"

var SECRET_KEY *rsa.PrivateKey

func setSecretKey() error {
    secretKey, err := ioutil.ReadFile(RSA_PATH)
    if err != nil {
        return errors.New("No secret key found!")
    }

    key, err := jwt.ParseRSAPrivateKeyFromPEM(secretKey)
    if err != nil {
        return errors.New("Error parsing private key!")
    }

    SECRET_KEY = key
    return nil
}

func generateSignedJWTToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), jwt.MapClaims{
        "iss": "dahuang",
        "aud": username,
        "exp": time.Now().Add(time.Minute * 1).Unix(),
    })

    // DONNOT store your secret key inside the program!
    signedToken, err := token.SignedString(SECRET_KEY)
    if err != nil {
        return "", errors.New("Generating signed token failed!")
    }

    return signedToken, nil
}
