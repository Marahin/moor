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
	fmt.Printf("Listening on %s\n", webAddr())
}

func Start() {
	mux := goji.NewMux()
	handler := cors.Default().Handler(mux)
	mux.HandleFunc(pat.Get("/:url"), handle)
	http.ListenAndServe(webAddr(), handler)
}