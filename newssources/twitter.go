package newssources

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// this W.I.P file is to get the trending hashtags from twitter

func Test1() {
	getcall := fmt.Sprintf("https://api.twitter.com/1.1/trends/place.json?id=1apiKey=%v",
		"")
	initialftresponse, err :=
		http.Get(getcall)

	if err != nil {
		os.Exit(1)
	}

	responsedata, err := ioutil.ReadAll(initialftresponse.Body)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(responsedata))

}
