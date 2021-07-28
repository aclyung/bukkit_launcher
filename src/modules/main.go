package main

import (
	"fmt"
	"log"
	"net/http"
)

type Bukkit struct {
	version string
	name    string
	api     string
	url     string
}

func main() {
	urls := []string{"https://google.com"}

	fmt.Println(urls)
}

func read_properties() {

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
