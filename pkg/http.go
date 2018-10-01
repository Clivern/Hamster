package pkg

import (
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
)

func Request(method string, url string, body string, token string) (string, error) {

    client := &http.Client{}
    req, err := http.NewRequest(method, url, bytes.NewBufferString(body))

    if err != nil{
        return "", err
    }

    req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

    resp, err := client.Do(req)

    if err != nil{
        return "", err
    }

    defer resp.Body.Close()

    body_byte, err := ioutil.ReadAll(resp.Body)

    if err != nil{
        return "", err
    }

    return string(body_byte), nil
}