package moor

import (
	"os"
	"net/http"
	"net/url"
	"fmt"

	"goji.io/pat"
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
	fmt.Printf("GET → %s\n", parsedParamUrl)
	fmt.Fprintf(w, "%s", Get(parsedParamUrl))
}

func webAddr() (string) {
	var addr = os.Getenv("MOOR_WEBSERVICE_ADDR")
	if addr == "" {
		addr = "0.0.0.0:7999"
	}
	return addr
}