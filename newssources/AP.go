package newssources

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type App struct {
	Client *http.Client
}

type APResponseHeader struct {
	NumResults int
	Results    []APArticle
}

type APArticle struct {
	Title    string
	Author   string
	Time     string
	Abs      string
	URL      string
	PhotoURL string
}

func APHome(url string) APResponseHeader {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("failed to get response from home: ", err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	//This seems like it works??
	var articlelinks []string
	document.Find("a[data-key='story-link']").Each(func(i int, selection *goquery.Selection) {
		link, _ := selection.Attr("href")
		articlelinks = append(articlelinks, link)
	})

	//Abstracts work
	var abstracts []string
	document.Find("a[data-key='story-link']").Each(func(index int, item *goquery.Selection) {
		abs := item.Children().Text()
		abstracts = append(abstracts, abs)
	})

	//h1[class='Component-h1-0-2-84']
	//body article div div a h1

	var titles []string
	document.Find("a[data-key='card-headline']").Each(func(index int, element *goquery.Selection) {
		articleTitle := element.Children().Text()
		titles = append(titles, articleTitle)
	})

	//THIS DOES NOT WORK FIX THIS
	var authors []string
	document.Find("a[data-key='card-headline']").Next().Each(func(index int, element *goquery.Selection) {
		getElement := element.Children().First()
		attribute, isThere := getElement.Attr("data-key")
		artAuthor := ""
		if attribute == "timestamp" {
			artAuthor = "No author found"
		} else if attribute != "timestamp" || !isThere {
			artAuthor = getElement.Text()
		}
		authors = append(authors, artAuthor)
	})

	//This works
	var timeStamps []string
	document.Find("span[data-key='timestamp']").Each(func(index int, element *goquery.Selection) {
		time, _ := element.Attr("data-source")
		timeStamps = append(timeStamps, time)
	})

	//var pictureURLS []string
	document.Find("div div a div img").Each(func(index int, element *goquery.Selection) {

		//fmt.Println(element)
		//picURL,_ := element.Children()
		//fmt.Println(picURL)

	})

	var articles []APArticle
	for i := 0; i < len(articlelinks); i++ {
		tempArticle := APArticle{
			Title:    titles[i],
			Author:   authors[i],
			Time:     timeStamps[i],
			Abs:      abstracts[i],
			URL:      articlelinks[i],
			PhotoURL: "",
		}
		articles = append(articles, tempArticle)
	}

	return APResponseHeader{NumResults: len(articles), Results: articles}

}
