package event

import (
    "testing"
    "io/ioutil"
)

func TestWatch(t *testing.T) {

    var watch Watch

    dat, err := ioutil.ReadFile("../../../samples/watch.json")

    if err != nil{
        t.Errorf("File samples/watch.json is invalid!")
    }

    ok, _ := watch.LoadFromJSON(dat)

    if !ok {
        t.Errorf("Testing with file samples/watch.json is invalid")
    }

    got := watch.Repository.Owner.Login
    want := "Clivern"

    if !ok || got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}