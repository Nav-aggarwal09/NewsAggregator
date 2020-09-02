package newssources

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Test() {
	getcall := fmt.Sprintf("https://api.ft.com/content/7348edd8-0403-11de-845b-000077b07658/v1?apiKey=%v",
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
