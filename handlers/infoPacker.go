package handlers

import (
	"../newssources"
	"errors"
	log "github.com/sirupsen/logrus"
)

type indexHeader struct {
	APNews  newssources.NewsSource
	NYTNews newssources.NewsSource
}

func IndexPackager(nytkey string) (*indexHeader, error) {

	nytdata, err1 := newssources.Nytapiconnect(nytkey, "home")
	apdata, err2 := newssources.APHome("https://apnews.com/apf-topnews")

	if err1 != nil || err2 != nil {
		log.Error("Data retrieval error")
		newError := errors.New("Error: could not recieve data from a news source (IndexPackager func)")
		return nil, newError
	}

	combinedHeader := indexHeader{APNews: apdata, NYTNews: nytdata}
	combinedHeader.convertTimes()
	return &combinedHeader, nil
}

func (header *indexHeader) convertTimes() {
	header.APNews.FormatDate()
	header.NYTNews.FormatDate()

}
