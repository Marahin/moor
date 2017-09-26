package moor

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strconv"
	"os"
)

func BlockerCharactersAmount() (int) {
	var err error
	var amount int
	amountStr := os.Getenv("MOOR_BLOCKER_CHARACTERS_AMOUNT")
	if len(amountStr) <= 0 {
		amount = BLOCKER_CHARACTERS_AMOUNT
	} else {
		amount, err = strconv.Atoi(amountStr)
		if err != nil {
			fmt.Println("[moor] MOOR_BLOCKER_CHARACTERS_AMOUNT could not be converted to int")
			fmt.Print("[moor]")
			fmt.Print(err)
			fmt.Printf("[moor] Defaulting to BLOCKER_CHARACTERS_AMOUNT=%v\n", BLOCKER_CHARACTERS_AMOUNT)
			amount = BLOCKER_CHARACTERS_AMOUNT
		}
	}
	return amount
}

func Get(url string) (text string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("[moor] Error during GET\n")
		fmt.Printf("[moor]")
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)[BlockerCharactersAmount():]
}