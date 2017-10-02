package moor

import (
	"os"
	"net/http"
	"net/url"
	"fmt"
	"strings"

	"goji.io/pat"
)

func allowedOrigins() ([]string){
	var origins = []string{}
	for _, allowedUrl := range strings.Split(os.Getenv("MOOR_ALLOWED_ORIGINS"), ",") {
		if len(allowedUrl) > 0 {
			origins = append(origins, allowedUrl)
		}
	}
	return origins
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	paramUrl := pat.Param(r, "url")
	fmt.Printf("GET → %s\n", paramUrl)
	if stringInSlice(paramUrl, IGNORE_ENDPOINTS) {
		fmt.Printf("  url ignored (%s)\n", paramUrl)
		return
	}

	parsedParamUrl, err := url.PathUnescape(paramUrl)
	if err != nil {
		fmt.Printf("Something went dang wrong, yo.")
		fmt.Printf("%s is unparsed paramUrl", paramUrl)
		fmt.Printf("%s is parsed paramUrl", parsedParamUrl)
		fmt.Print(err)
		fmt.Fprint(w, "{ \"error\": \"PathUnescape failed\"}")
	}
	fmt.Printf("GET → parsed=%s\n", parsedParamUrl)
	fmt.Fprintf(w, "%s", Get(parsedParamUrl))
}

func webAddr() (string) {
	var addr = os.Getenv("MOOR_WEBSERVICE_ADDR")
	if addr == "" {
		addr = "0.0.0.0:7999"
	}
	return addr
}