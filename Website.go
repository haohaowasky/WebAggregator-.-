package main

import ("fmt"
        "net/http"
        "html/template"
        "io/ioutil"
        "encoding/xml"
        "sync"
       )

var wg sync.WaitGroup

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

type NewsAggPage struct{
    Title string
    News map[string] NewsMap  //map where the key is a string type and value is a NewsMap struct
}

func indexHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<h1>something</h1>")
}

func newsRoutine(c chan News, Location string){
    defer wg.Done()
    var n News
    resp, _ := http.Get(Location)
    bytes, _ := ioutil.ReadAll(resp.Body)
    xml.Unmarshal(bytes, &n)
    resp.Body.Close()
    c <- n // send value n to the channel
}


func newsAggHnadler(w http.ResponseWriter, r *http.Request){
    var s Sitemapindex
    news_map := make(map[string]NewsMap)
    resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
    bytes, _ := ioutil.ReadAll(resp.Body) // this like formula
    xml.Unmarshal(bytes, &s)
    resp.Body.Close()
    queue := make(chan News, 30)


    for _, i := range s.Locations{
      wg.Add(1)
      go newsRoutine(queue, i)
    }

    wg.Wait()
    close(queue)
    for elem := range queue{
        for idx, _ := range elem.Keywords{
            news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
        }
    }

    p := NewsAggPage{Title: "Ama news", News: news_map} // it passes Title and News to the html page,
    // a dictionary map is passed to the html, the dictionary map has 2 entities, key and value, which value
    // is a struct that has two string values.
    t, _ := template.ParseFiles("basictemplating.html")
    t.Execute(w, p)
}

func main(){
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/agg/", newsAggHnadler)
    http.ListenAndServe(":8000", nil)
}
