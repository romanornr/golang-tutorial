package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatal(err)
	}
}
