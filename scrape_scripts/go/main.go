package main

import (
  "fmt"
  "log"
  "net/http"
  "strings"
  "github.com/PuerkitoBio/goquery"
  "strconv"
)

type AirportInfo struct {
  IATA      string
  ICAO      string
  City      string
  Country   string
  Latitude  float64
  Longitude float64
  WikiURL   string
}

func scrapeAirport(url string) (*AirportInfo, error) {
  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }
  
  req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
  
  res, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()
  
  if res.StatusCode != 200 {
    return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
  }
  
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    return nil, err
  }
  
  info := &AirportInfo{
    WikiURL: url,
  }
  
  // Find IATA and ICAO codes
  doc.Find(".infobox .ib-airport-codes .nickname").Each(func(i int, s *goquery.Selection) {
    text := strings.TrimSpace(s.Text())
    parent := s.Parent().Text()
    
    if strings.Contains(parent, "IATA:") {
      info.IATA = text
    } else if strings.Contains(parent, "ICAO:") {
      info.ICAO = text
    }
  })
  
  // Find the city (Serves field)
  doc.Find(".infobox tr").Each(func(i int, s *goquery.Selection) {
    label := s.Find("th").Text()
    if strings.Contains(label, "Serves") {
      info.City = strings.TrimSpace(s.Find("td").Text())
    }
  })
  
  // Extract country from Location field
  doc.Find(".infobox tr").Each(func(i int, s *goquery.Selection) {
    label := s.Find("th").Text()
    if strings.Contains(label, "Location") {
      location := strings.TrimSpace(s.Find("td").Text())
      // Location format: "address, City, State, Country"
      parts := strings.Split(location, ",")
      if len(parts) > 0 {
        info.Country = strings.TrimSpace(parts[len(parts)-1])
      }
    }
  })
  
  // Find coordinates (Latitude and Longitude)
  coordLink := doc.Find(".infobox .geo-dec").First()
  if coordLink.Length() > 0 {
    coords := strings.TrimSpace(coordLink.Text())
    // Format: "39.71722°N 86.29444°W"
    parts := strings.Fields(coords)
    if len(parts) >= 2 {
      info.Latitude = parts[0]
      info.Longitude = parts[1]
    }
  }
  
  return info, nil
}

func main() {
  url := "https://en.wikipedia.org/wiki/Indianapolis_International_Airport"
  
  airport, err := scrapeAirport(url)
  if err != nil {
    log.Fatal(err)
  }
  
  fmt.Printf("IATA Code: %s\n", airport.IATA)
  fmt.Printf("ICAO Code: %s\n", airport.ICAO)
  fmt.Printf("City: %s\n", airport.City)
  fmt.Printf("Country: %s\n", airport.Country)
  fmt.Printf("Latitude: %s\n", airport.Latitude)
  fmt.Printf("Longitude: %s\n", airport.Longitude)
  fmt.Printf("Wiki URL: %s\n", airport.WikiURL)
}
