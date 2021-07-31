package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const BASEURL = "https://getbukkit.org"

var api_type = []string{"vanilla", "spigot", "craftbukkit"}

type Bukkit map[string]string

func main() {
	fmt.Println(BASEURL)
	var bukkits []Bukkit
	c := make(chan []Bukkit)
	for _, api := range api_type {
		go getBukkit(api, c)
	}
	for range api_type {
		extracted := <-c
		bukkits = append(bukkits, extracted...)
	}
	//fmt.Print(bukkits)
	writeBukkits(bukkits)
}

func getBukkit(api string, bukkitC chan<- []Bukkit) {
	var bukkits []Bukkit
	c := make(chan Bukkit)
	pageUrl := fmt.Sprint(BASEURL, "/download/", api)
	fmt.Println("Requesting:", pageUrl)
	req, _ := http.NewRequest("GET", pageUrl, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	res, err := http.DefaultClient.Do(req)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(res)
	checkErr(err)

	searchCards := doc.Find(".download-pane>div")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractBukkit(card, api, c)
	})
	for i := 0; i < searchCards.Length(); i++ {
		bukkits = append(bukkits, <-c)
	}
	bukkitC <- bukkits

}

func extractBukkit(card *goquery.Selection, api string, c chan<- Bukkit) {
	version := card.Find("div>h2").Text()
	//fmt.Print(version)
	name := fmt.Sprint(api, "-", version, ".jar")
	link, _ := card.Find("#downloadr").Attr("href")

	url := getUrl(link)
	bukkit := make(Bukkit)
	bukkit["version"] = version
	bukkit["name"] = name
	bukkit["url"] = url
	c <- bukkit
}

func getUrl(dl_url string) string {
	req, _ := http.NewRequest("GET", dl_url, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	res, err := http.DefaultClient.Do(req)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(res)
	checkErr(err)

	url, _ := doc.Find(".well>h2>a").Attr("href")
	return url
}

func writeBukkits(bukkits []Bukkit) {

	doc, _ := json.Marshal(bukkits)
	err := ioutil.WriteFile("bukkit.json", doc, 0644)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		if err != nil {
			log.Fatalln(err)
		}
	}
}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
