package newssources

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func test() {
	getcall := fmt.Sprintf("https://api.nytimes.com/svc/search/v2/articlesearch.json?q=%v&api-key=%v",
		"election", "LAyAA8ZUvR0hAiYkNtOYNLXoZH8IG6VI")
	initialnytresponse, err :=
		http.Get(getcall)

	if err != nil {
		os.Exit(1)
	}

	responsedata, err := ioutil.ReadAll(initialnytresponse.Body)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(responsedata))
}
