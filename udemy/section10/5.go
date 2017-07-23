package main

import "net/http"
import "log"
import "io/ioutil"
import "fmt"

func main() {
	resp, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	test := string(body)
	fmt.Println(test)
}
