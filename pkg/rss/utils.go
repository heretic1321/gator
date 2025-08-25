package rss

import "html"

func cleanData(feed *RSSFeed){
	
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	
	for i, item := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(item.Title) 
		feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}
}
