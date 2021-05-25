package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
)

func main ()  {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err :%q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		//colly.AllowedDomains("youtube.com"), // NOTE : YOUTUBE
		colly.AllowedDomains("internshala.com"),
		)

	//NOTE : WEBSITE
	c.OnHTML(".internship_meta", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("a"),
			e.ChildText("span"),
		})
	})

	//NOTE : YOUTUBE
	//c.OnHTML("#style-scope ytd-expanded-shelf-contents-renderer", func(e *colly.HTMLElement) {
	//	writer.Write([]string{
	//		e.ChildText("aria-label"),
	//		e.ChildText("h3"),
	//		e.ChildText("title-and-badge style-scope ytd-video-renderer"),
	//		e.ChildText("title"),
	//		e.ChildText("aria-label"),
	//		e.ChildText("a"),
	//		e.ChildText("video-title"),
	//		//e.ChildText("span"),
	//	})
	//})

	for i := 0; i <= 20; i++ {
		fmt.Printf("Scrape page : %d\n", i)
		//c.Visit("https://www.youtube.com/feed/trending?bp=6gQJRkVleHBsb3Jl")
		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}

	log.Printf("Scrapping Complete\n")
	log.Println(c)
}