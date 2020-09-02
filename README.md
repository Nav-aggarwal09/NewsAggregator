# NewsAggregator
News Aggregator that pulls news from reliable sources

## How to run
after downloading the project, run main.go from the command line/Terminal and supply a NYT API key and a NewsAPI key. The command should look like this:
`go run main.go -nyt XXXX -news XXXX`

## Overview of Directories & Files
Giving a few notes on each directory in the project and noteworthy aspects of particular files...

### NewsSources
For every news source the project tries to get data from, there is a file for each that handles that

### Handlers
Depending on the url path, main.go calls the appropriate handler function to handle it. The handler functions call the News Source functions to gather data and ultimately pass all the data to the front end template

### Front End
This is where all the html templates reside. They are passed data from the handler functions

### Front End Assets
Anything extra the front end needs to polish it
