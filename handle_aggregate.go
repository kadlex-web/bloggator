package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"

	r "github.com/kadlex-web/bloggator/internal/rss"
)

// function fetchs a feed from a given url, and assuming that nothing goes wrong, returns a fill-out RSSFeed struct
func fetchFeed(ctx context.Context, feedURL string) (*r.RSSFeed, error) {
	// create a new HTTP Request with background
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &r.RSSFeed{}, fmt.Errorf("error creating request.")
	}
	// add identification header to request
	req.Header.Set("User-Agent", "bloggator")
	// create new http client for sending requests
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &r.RSSFeed{}, fmt.Errorf("error getting response")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &r.RSSFeed{}, fmt.Errorf("error reading response body")
	}
	feed := r.RSSFeed{}
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return &r.RSSFeed{}, fmt.Errorf("error unmarshaling xml data into feed.")
	}
	return &feed, nil
}

func cleanXML(data *r.RSSFeed) {
	fmt.Println(html.UnescapeString(data.Channel.Title))
	fmt.Println(html.UnescapeString(data.Channel.Description))
	fmt.Println(html.UnescapeString(data.Channel.Link))
	for _, rssItem := range data.Channel.Item {
		fmt.Println(html.UnescapeString(rssItem.Title))
		fmt.Println(html.UnescapeString(rssItem.Description))
		fmt.Println(html.UnescapeString(rssItem.Link))
		fmt.Println(html.UnescapeString(rssItem.PubDate))
		fmt.Println()
	}
}

func aggregate(s *state, cmd command) error {
	// right now aggregate will only access ONE feed -- but eventually will do more
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("aggregate command doesn't contain arguments")
	}
	url := "https://www.wagslane.dev/index.xml"
	data, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	cleanXML(data)
	return nil
}
