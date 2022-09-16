package main

import (
	"encoding/xml"
	"io"
	"fmt"
	"net/http"
	"os"
)

type RSS struct {
	XMLData     xml.Name       `xml:"rss"`
	Channel     *Channel       `xml:"channel"`
}

type Channel struct {
	Title      string          `xml:"title"`
	ItemList   []Item          `xml:"item"`
}  

type Item struct{
	Title			string           `xml:"title"`
	Link      string           `xml:"link"`
	Traffic		string           `xml:"approx_traffic"`
	NewsItem  []News           `xml:"news_item"`
}

type News struct{
	Headline      string       `xml:"news_item_title"`
	HeadlinkLink  string       `xml:"news_item_url"`
}

func main(){
	var r RSS

	data := readGoogleTrends()

  err := xml.Unmarshal(data,&r)
	if err != nil{
		fmt.Println("error: ",err)
	}

	fmt.Println("\n Below are the Google Search Trends For Today !")
	fmt.Println("------------------------------------------------")
  
	for i := range r.Channel.ItemList {
     rank := i + 1
		 fmt.Println("#",rank)
     fmt.Println("Search Term:",r.Channel.ItemList[i].Title)
		 fmt.Println("Link of the Trends:",r.Channel.ItemList[i].Link)

		 for j := range r.Channel.ItemList[i].NewsItem {
			 fmt.Println("Headline:",r.Channel.ItemList[i].NewsItem[j].Headline)
			 fmt.Println("Link to a article:",r.Channel.ItemList[i].NewsItem[j].HeadlinkLink)
		 }
		 fmt.Println("-----------------------------\n")
	}
}

func readGoogleTrends() []byte{
	resp := getGoogleTrends()

	data,err := io.ReadAll(resp.Body)

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}

func getGoogleTrends() *http.Response{
     resp,err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=IN")

		 if err != nil{
				fmt.Println(err)
				os.Exit(1)
		 }
		 return resp
}
