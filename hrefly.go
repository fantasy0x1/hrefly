package main 

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func main() {
	urlPtr := flag.String("url", "", "Website URL")
	flag.Parse()

	if *urlPtr == "" {
		fmt.Println("Use the --url parameter to pass the target URL")
		return
	}

	if !strings.HasPrefix(*urlPtr, "http://") && !strings.HasPrefix(*urlPtr, "https://") {
		*urlPtr = "http://" + *urlPtr
	}

	printBanner()

	response, err := http.Get(*urlPtr)
	if err != nil {
		log.Fatalf("Error at sending a GET request: %s", err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalf("Error at loading HTML source code: %s", err)
	}

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			fmt.Println(href)
		}
	})
}

func printBanner() {
	banner := `
 _               __ _
| |__  _ __ ___ / _| |_   _
| '_ \| '__/ _ \ |_| | | | |
| | | | | |  __/  _| | |_| |  [*] Bringing up the links used by href on the page 
|_| |_|_|  \___|_| |_|\__, |
                      |___/
`
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Println(cyan(banner))
}
