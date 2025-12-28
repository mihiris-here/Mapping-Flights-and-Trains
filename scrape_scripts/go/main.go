package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/PuerkitoBio/goquery"
)


func ExampleScrape() {
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://en.wikipedia.org/wiki/Indianapolis_International_Airport", nil)
  if err != nil {
    log.Fatal(err)
  }

  req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

  res, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()

  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  // Get the page title
  title := doc.Find("h1#firstHeading").Text()
  fmt.Printf("Page title: %s\n", title)

  // Get the first paragraph
  firstPara := doc.Find("#mw-content-text .mw-parser-output > p").First().Text()
  fmt.Printf("\nFirst paragraph: %s\n", firstPara)
}

func main() {
  ExampleScrape()
}


