package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var options = [7]string{
	"なんでやねん！",
	"ありえん！",
	"しょーもないなー",
	"いやいや、あかんあかん",
	"お前なんかおかしいぞ、病院に行ってこい",
	"あっち行け、馬鹿が移ったら困るから",
	"アホたれ"}

func getRandomTukkomi() string {
	// To get a different series of numbers from rand.Intn,
	// "seed" with the current time.
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(len(options))
	return options[randomNumber]
}

func handleTukkomiPath(w http.ResponseWriter, r *http.Request) {
	phrase := r.URL.Query()["phrase"][0]
	if len(phrase) > 0 {
		fmt.Fprint(w, phrase+"って、"+getRandomTukkomi())
	} else {
		fmt.Fprint(w, "おい！ツッコムねぇじゃん")
	}
}

func main() {
	http.HandleFunc("/tukkomi", handleTukkomiPath)
	http.ListenAndServe(":3001", nil)
}
