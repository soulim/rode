package rode

import (
	"encoding/xml"
	"time"
)

type RSS struct {
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	Title string    `xml:"title"`
	Items []RSSItem `xml:"item"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	GUID        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
}

type RSSParser struct{}

func (p *RSSParser) Parse(data []byte, feed *Feed) error {
	rss := &RSS{}

	if err := xml.Unmarshal(data, rss); err != nil {
		return err
	}

	feed.Items = make(map[string]FeedItem)

	for _, item := range rss.Channel.Items {
		pubdate, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			return err
		}

		feed.Items[item.GUID] = FeedItem{
			GUID:        item.GUID,
			Title:       item.Title,
			Description: item.Description,
			PubDate:     pubdate,
		}
	}

	return nil
}
