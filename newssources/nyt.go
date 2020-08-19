package newssources

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)
//type NYTResponse struct {
//	Status      string `json:"status"`
//	Copyright   string `json:"copyright"`
//	Section     string `json:"section"`
//	LastUpdated string `json:"last_updated"`
//	NumResults  int    `json:"num_results"`
//	Results     []struct {
//		Section           string        `json:"section"`
//		Subsection        string        `json:"subsection"`
//		Title             string        `json:"title"`
//		Abstract          string        `json:"abstract"`
//		URL               string        `json:"url"`
//		URI               string        `json:"uri"`
//		Byline            string        `json:"byline"`
//		ItemType          string        `json:"item_type"`
//		UpdatedDate       string        `json:"updated_date"`
//		CreatedDate       string        `json:"created_date"`
//		PublishedDate     string        `json:"published_date"`
//		MaterialTypeFacet string        `json:"material_type_facet"`
//		Kicker            string        `json:"kicker"`
//		DesFacet          []string      `json:"des_facet"`
//		OrgFacet          []string      `json:"org_facet"`
//		PerFacet          []string      `json:"per_facet"`
//		GeoFacet          []interface{} `json:"geo_facet"`
//		Multimedia        []struct {
//			URL       string `json:"url"`
//			Format    string `json:"format"`
//			Height    int    `json:"height"`
//			Width     int    `json:"width"`
//			Type      string `json:"type"`
//			Subtype   string `json:"subtype"`
//			Caption   string `json:"caption"`
//			Copyright string `json:"copyright"`
//		} `json:"multimedia"`
//		ShortURL string `json:"short_url"`
//	} `json:"results"`
//}

// Header for NYT Response when API called
// Purpose- to map article results
type NYTResponseHeader struct {
	Section		string 		`json:"section"`
	LastUpdated	string 		`json:"last_updated"`
	NumResults 	int 		`json:"num_results"`
	Results 	[]NYTResult `json:"results"`

}

// capture information from returned results
type NYTResult struct {
	Section           string        `json:"section"`
	Subsection        string        `json:"subsection"`
	Title             string        `json:"title"`
	Abstract          string        `json:"abstract"`
	Byline            string        `json:"byline"`
	UpdatedDate       string        `json:"updated_date"`
	PublishedDate     string        `json:"published_date"`
	ShortURL 		  string 		`json:"short_url"`
	DesFacet          []string      `json:"des_facet"`
	OrgFacet          []string      `json:"org_facet"`
	PerFacet          []string      `json:"per_facet"`
	// TODO: Save information from multimedia to display on front end

}

const NYTKEY string = "LAyAA8ZUvR0hAiYkNtOYNLXoZH8IG6VI"

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func FindString(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// connect to NYT API
func nytapiconnect(section string) (*NYTResponseHeader, error){
	// TODO: differentiate between types of calls to be made

	validsections := []string{"arts", "automobiles", "books", "business", "fashion", "food",
		"health", "home", "insider", "magazine", "movies", "nyregion", "obituaries", "opinion",
		"politics", "realestate", "science", "sports", "sundayreview",
		"technology", "theater", "t-magazine", "travel", "upshot", "us", "world"}
	if _, valid := FindString(validsections, section); !valid {
		log.Error("Given section is not valid for NYT")
		return nil, errors.New("Invalid section (NYT)")
	}
	log.Debug("Accepted given section for NYT")

	getcall := fmt.Sprintf("https://api.nytimes.com/svc/topstories/v2/%v.json?api-key=%v",
		section, NYTKEY)
	initialnytresponse, err :=
		http.Get(getcall)

	if err != nil {
		log.Error("Unsuccessful connection to NYT API: ", err)
		return nil, err
	}

	responsedata, err := ioutil.ReadAll(initialnytresponse.Body)
	if err != nil {
		log.Error("Cannot read from response: ", err)
		return nil, err
	}

	var responseObject NYTResponseHeader
	json.Unmarshal(responsedata, &responseObject)
	return &responseObject, nil
}

func nytoutputtofile(data *NYTResponseHeader) error {

	jsonnyt, err := json.MarshalIndent(data, ""," ")
	if err != nil {
		log.Error("Error converting nytResults data to json: ", err)
		return err
	}
	_ = ioutil.WriteFile("/Users/arnav/gocode/GoLand/newsAggregator/testdata.json", jsonnyt, 0644)
	if _,err := os.Stat("/Users/arnav/gocode/GoLand/newsAggregator/testdata.json"); err == nil  {
		log.Info("Created json file from NYT data")
		return nil
	} else if os.IsNotExist(err) {
		log.Error("NYT JSON file does not exist: ", err)
		return errors.New("Writing to NYT JSON file failed")
	}

	return nil
}


func nytparseResults(header NYTResponseHeader) ([]NYTResult, error) {


	var resultdata []NYTResult

	for i:= 0; i < header.NumResults; i++ {
		resultdata = append(resultdata, header.Results[i])
	}

	log.Info("Successfully parsed data from NYT")
	log.Debug("Number of results: ", header.NumResults)

	return resultdata, nil
}
