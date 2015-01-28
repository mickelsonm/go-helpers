package hackernews

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type (
	//The RSS Feed
	Feed struct {
		XMLName xml.Name `xml:"rss" json:"-"`
		Channel *Channel `xml:"channel" json:"channel"`
	}
	//The RSS Channel
	Channel struct {
		XMLName     xml.Name `xml:"channel" json:"-"`
		Title       string   `xml:"title" json:"title"`
		Link        string   `xml:"link" json:"link"`
		Description string   `xml:"description" json:"description"`
		Items       []Item   `xml:"item" json:"item"`
	}
	//An individual item within the channel
	Item struct {
		XMLName     xml.Name `xml:"item" json:"-"`
		Title       string   `xml:"title" json:"title"`
		Link        string   `xml:"link" json:"link"`
		Comments    string   `xml:"comments" json:"comments"`
		Description string   `xml:"description" json:"description"`
	}
)

//Utility method to give us json output
func (f *Feed) ToJSON() (js []byte, err error) {
	return json.Marshal(f)
}

//Gets the actual Hacker News RSS feed
func GetRssFeed() (feed Feed, err error) {
	resp, err := http.Get("https://news.ycombinator.com/rss")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if err = xml.Unmarshal(buf, &feed); err != nil {
		return
	}
	return
}
