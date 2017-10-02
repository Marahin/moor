package moor

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"github.com/rs/cors"
)

func Credits() {
	fmt.Printf("Moor Server v%v\n", VERSION)
	fmt.Printf("Created by ~marahin\n")
	fmt.Println("---- variables ----")
	fmt.Printf("BLOCKER_CHARACTERS_AMOUNT=%v\n", BlockerCharactersAmount())
	fmt.Printf("IGNORE_ENDPOINTS=%v\n", IGNORE_ENDPOINTS)
	fmt.Printf("allowed origins: %v\n", allowedOrigins())
	if len(allowedOrigins()) == 0 {
		fmt.Println("allowed origins is empty, so we default to *")
	}
	fmt.Printf("Listening on %s\n", webAddr())
}

func Start() {
	mux := goji.NewMux()

	var handler http.Handler

	if len(allowedOrigins()) == 0 {
		handler = cors.Default().Handler(mux)
	} else {
		c := cors.New(cors.Options{
			AllowedOrigins: allowedOrigins(),
		})
		handler = c.Handler(mux)
	}
	mux.HandleFunc(pat.Get("/:url"), handle)
	http.ListenAndServe(webAddr(), handler)
}