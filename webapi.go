package crawlers

import (
    "net/http"
    "io/ioutil"
    "log"
)

/*
    Define web api interface
*/
type I_WebAPI interface {
    Get(url string, result chan interface{})
}


type C_WebAPI struct {
}

/*
    Retrieve data from http.get url
*/
func (webapi *C_WebAPI) Get(url string, result chan interface{}) {
    resp, weberr := http.Get(url)
    if weberr != nil {
        log.Fatal(weberr)
    }
    body, ioerr := ioutil.ReadAll(resp.Body)
    if ioerr != nil {
        log.Fatal(ioerr)
    }
    //target := string(body)
    result <- body
}


