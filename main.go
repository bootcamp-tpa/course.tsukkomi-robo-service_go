package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var options = []string{
	"なんでやねん！",
	"ありえん！",
	"しょーもないなー",
	"いやいや、あかんあかん",
	"お前なんかおかしいぞ、病院に行ってこい",
	"あっち行け、馬鹿が移ったら困るから",
	"アホたれ"}

func getRandomTsukkomi() string {
	randomNumber := rand.Intn(len(options))
	return options[randomNumber]
}

func handleTsukkomi(w http.ResponseWriter, r *http.Request) {
	phrase := r.URL.Query()["phrase"][0]
	if len(phrase) > 0 {
		fmt.Fprint(w, phrase + "って、" + getRandomTsukkomi())
	} else {
		fmt.Fprint(w, "おい！ツッコムねぇじゃん")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/tsukkomi", handleTsukkomi)
	http.ListenAndServe(":3001", nil)
}
