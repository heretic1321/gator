package rss

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
)

type Client struct {
	http *http.Client
}


func New(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		http: httpClient,
	}

}

func (c *Client) FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error){
	
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)

	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Set("User-Agent", "gator")

	res, err := c.http.Do(req)

	if err != nil {
		return &RSSFeed{}, err
	}

	content, err := io.ReadAll(res.Body)

	if err != nil {
		return &RSSFeed{}, err
	}

	feed := RSSFeed{}
	
	err = xml.Unmarshal(content, &feed)
	
	if err != nil {
		return &RSSFeed{}, err
	}

	cleanData(&feed)

	return &feed, nil
}
