package main


import (
	handlers "./handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)



func main() {
	// TODO: create log file
	log.SetOutput(os.Stdout)
	log.Info("Starting program...")

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	mux := http.NewServeMux()

	// Declare the static file directory
	fs := http.FileServer(http.Dir("frontendassets"))
	mux.Handle("/frontendassets/", http.StripPrefix("/frontendassets/", fs))

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/search", handlers.SearchHandler)
	log.Info("Listening on port ", port)
	http.ListenAndServe(":"+port, mux)


	//err := runnyt()
	//if err!= nil {
	//	os.Exit(1)
	//}
}
/*
// function to get nyt data
func runnyt() error {

	// Format of command line arguments: NEWSSITE SERVICE SECTION(home, arts, finance, etc) TODO: LOGLVL

	// Cmmnd Line Flags to perform particular actions
	newssitePtr := flag.String("site", "nyt",
		"site on which to perform action")

	// TODO: have a set of universal service and section
	servicePtr := flag.String("service", "top",
		"type of service desired from specified news site")
	sectionPtr := flag.String("section", "home",
		"section (arts, finance, etc) of news to get results from")
	tail := flag.Args()
	flag.Parse()

	log.Infof("SITE: %v \t\t SERVICE: %v \t\t SECTION: %v",
		*newssitePtr, *servicePtr, *sectionPtr)
	fmt.Printf("these are the trailing arguments: %v\n", tail)

	var nytdata *NYTResponseHeader
	var nyterr error
	if strings.ToLower(*newssitePtr) == "nyt" {
		nytdata, nyterr = nytapiconnect(*sectionPtr)
		if nyterr != nil {
			os.Exit(1)
		}
		fmt.Println("number of results: ", nytdata.NumResults)
		//for index, result := range nytdata.Results {
		//	fmt.Printf("%d %v \n", index, result.Title)
		//}
	} else {
		fmt.Println("invalid site")
	}

	nytoutputtofile(nytdata)
	return nil
}
*/
