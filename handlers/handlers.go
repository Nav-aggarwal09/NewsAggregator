package handlers

import (
	"../constants"
	"../newssources"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//  points to a template definition
	var tpl = template.Must(
		template.ParseFiles("/Users/arnav/gocode/GoLand/newsAggregator/frontend/index.html"))

	nytdata, err := newssources.Nytapiconnect(constants.NytKeyPtr, "home")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		log.Error("Connection to NYT API failed")
	}
	tpl.Execute(w, *nytdata)
}

// put user search's through NewsAPI
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	//  points to a template definition
	var tpl = template.Must(
		template.ParseFiles("/Users/arnav/gocode/GoLand/newsAggregator/frontend/search.html"))
	u, err := url.Parse(r.URL.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		log.Error("Could not parse user search query")
		return
	}

	params := u.Query()
	searchKey := params.Get("q")
	page := params.Get("page")
	if page == "" {
		page = "1"
	}

	log.Debugf("Search Query is: ", searchKey)
	log.Debugf("Results page is: ", page)

	search := &newssources.Search{}
	// set SearchKey to user query
	search.SearchKey = searchKey

	// convert user page number from string to int
	next, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Unexpected server error", http.StatusInternalServerError)
		log.Error("Could not convert user page number to int for Search struct")
		return
	}

	search.NextPage = next
	// number of results NewsAPI will return in response
	pageSize := 20

	var endpoint string
	if searchKey == "" {
		endpoint = fmt.Sprintf("https://newsapi.org/v2/top-headlines?category=general&pageSize=%d&page=%d&apiKey=%s",
			pageSize, search.NextPage, newssources.NEWSAPI)
	} else {
		endpoint = fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&pageSize=%d&page=%d&apiKey=%s&sortBy=publishedAt&language=en",
			url.QueryEscape(search.SearchKey), pageSize, search.NextPage, newssources.NEWSAPI)

	}
	resp, err := http.Get(endpoint)
	log.Info("Making GET request to NewsAPI")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Get request to NewsAPI failed")
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		newError := &newssources.NewsAPIError{}
		err := json.NewDecoder(resp.Body).Decode(newError)
		if err != nil {
			http.Error(w, "Unexpected server error", http.StatusInternalServerError)
			return
		}
		log.Error("Status code was not OK (200)")
		http.Error(w, newError.Message, http.StatusInternalServerError)
		return
	}
	log.Info("Got response status code 200")

	err = json.NewDecoder(resp.Body).Decode(&search.Results)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	search.TotalPages = int(math.Ceil(float64(search.Results.TotalResults / pageSize)))
	log.Debugf("There are %v pages", search.TotalPages)
	if ok := !search.IsLastPage(); ok {
		search.NextPage++
	}
	// execute template
	err = tpl.Execute(w, search)
	if err != nil {
		log.Errorf("Failed to execute search template: ", err)
	}
}
