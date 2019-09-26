package main

import (
	"fmt"
	"log"
	"os"

	_ "searchengine/matchers"
	 "searchengine/search"

)

//init is called prior to the main
func init() {
	// change the device for logging to stdout
	log.SetOutput(os.Stdout)
}


func main() {
	//perform search for the specified term
	search.Run("president")
}