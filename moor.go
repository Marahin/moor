package main

import (
	"moor/moor"
	"net/http"
	"net/url"
	"fmt"
	"os"

	"goji.io"
	"goji.io/pat"
	"github.com/rs/cors"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	paramUrl := pat.Param(r, "url")
	parsedParamUrl, err := url.PathUnescape(paramUrl)
	if err != nil {
		fmt.Printf("Something went dang wrong, yo.")
		fmt.Printf("%s is unparsed paramUrl", paramUrl)
		fmt.Printf("%s is parsed paramUrl", parsedParamUrl)
		fmt.Print(err)
		fmt.Fprint(w, "{ \"error\": \"PathUnescape failed\"}")
	}
	fmt.Fprintf(w, "%s", moor.Get(parsedParamUrl))
}

func webAddr() (string) {
	var addr = os.Getenv("MOOR_WEBSERVICE_ADDR")
	if addr == "" {
		addr = "localhost:7999"
	}
	return addr
}

func main() {
	moor.Credits()
	mux := goji.NewMux()
	handler := cors.Default().Handler(mux)
	mux.HandleFunc(pat.Get("/:url"), handle)
	http.ListenAndServe(webAddr(), handler)
}