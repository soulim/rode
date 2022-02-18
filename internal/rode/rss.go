package rode

import (
	"encoding/xml"
)

type RSS struct {
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	Title string    `xml:"title"`
	Items []RSSItem `xml:"item"`
}

type RSSItem struct {
	Title string `xml:"title"`
	GUID  string `xml:"guid"`
}

type RSSParser struct{}

func (p *RSSParser) Parse(data []byte, feed *Feed) error {
	rss := &RSS{}

	if err := xml.Unmarshal(data, rss); err != nil {
		return err
	}

	for _, item := range rss.Channel.Items {
		feed.Items = append(feed.Items, FeedItem{GUID: item.GUID, Title: item.Title})
	}

	return nil
}
