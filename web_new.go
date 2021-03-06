package main

import (
  "encoding/xml"
  "fmt"
  "io/ioutil"
  "net/http"
)

type Sitemapindex struct {
  Locations []string `xml:"sitemap>loc"`
}

type News struct{
    Titles []string `xml:"url>news>title"`
    Keywords []string `xml:"url>news>keywords"`
    Locations []string `xml:"url>loc"`
}


type NewsMap struct{
    Keyword string
    Location string
}

func main() {
  var s Sitemapindex
  var n News
  news_map := make(map[string]NewsMap)
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  xml.Unmarshal(bytes, &s)

  for _, i := range s.Locations{
    resp, _ := http.Get(i)
    bytes, _ := ioutil.ReadAll(resp.Body)
    xml.Unmarshal(bytes, &n)
    for idx, _ := range n.Titles{
        news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
    }
  }
  for idx, data := range news_map{
      fmt.Println("\n\n\n", idx)
      fmt.Println("\n", data.Keyword)
      fmt.Println("\n", data.Location)
  }

}
