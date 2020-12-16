package main

import (
	"fmt"
	"net/http"
)

type hotdog int

//  Any other type with this method 'serveHHTTP' is also a Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
