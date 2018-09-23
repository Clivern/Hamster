package main

import (
    "io/ioutil"
    "log"
    "net/http"
)

// curl localhost:8000 -d '{"name":"Hello"}'
func Cleaner(w http.ResponseWriter, r *http.Request) {
    // Read body
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    w.Header().Set("content-type", "application/json")
    w.Write(b)
    log.Print(string(b))
}

func main() {
    http.HandleFunc("/", Cleaner)
    address := ":8000"
    log.Println("Starting server on address", address)
    err := http.ListenAndServe(address, nil)
    if err != nil {
        panic(err)
    }
}
