package search

import (
	"log"
	"sync"
)

// Map of registered matchers for searching

var matchers = make(map[string]Matcher)

// search logic

func Run(searchTerm string) {

	// Retrieve list of feeds to search through
	feeds, err := RetrieveFeeds()
	if (err != nil) {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results
	results := make(chan *Result)

	// create wait group so as to process all feeds
	var waitGroup sync.WaitGroup

	// set the number of goroutines we need to wait for while they process individual feed
	waitGroup.Add(len(feeds))

	// launch go routine for each feed to find their results
	for _, feed := range feeds {
		// Retrieve a matcher for the search
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		// launch go routine to perform the search
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)

	}

	// launch goroutine to monitor when all work is done
	go func() {
		// wait for everything is processed
		waitGroup.Wait()
		// close channel to signal the display
		// functin that we exit the program
		close(results)

	}
	{
	}

	// start displaying the results as they are available and return after the final result is displayed
	Display(resutls)

}
