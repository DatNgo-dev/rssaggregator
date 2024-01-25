package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title			string		`xml:"title"`
		Link			string		`xml:"link"`
		Description		string		`xml:"description"`
		Language		string		`xml:"language"`
		Item			[]RSSItem	`xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title			string	`xml:"title"`
	Link			string	`xml:"link"`
	PubDate			string	`xml:"pub_date"`
	Guid			string	`xml:"guid"`
	Description		string	`xml:"description"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}

	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil 
}