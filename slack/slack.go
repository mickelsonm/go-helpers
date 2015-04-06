package slack

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

var (
	SLACK_API   = "https://<YOUR ORGANIZATION>.slack.com/services/hooks/incoming-webhook"
	SLACK_TOKEN = "<YOUR TOKEN>"
)

type AttachmentFields []AttachmentField
type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Attachments []Attachment
type Attachment struct {
	Fallback   string           `json:"fallback"`
	Color      string           `json:"color,omitempty"`
	Pretext    string           `json:"pretext,omitempty"`
	AuthorName string           `json:"author_name,omitempty"`
	AuthorLink string           `json:"author_link,omitempty"`
	AuthorIcon string           `json:"author_icon,omitempty"`
	Title      string           `json:"title,omitempty"`
	TitleLink  string           `json:"title_link,omitempty"`
	Text       string           `json:"text"`
	ImageURL   string           `json:"image_url,omitempty"`
	Fields     AttachmentFields `json:"fields,omitempty"`
}

type Message struct {
	Channel     string      `json:"channel"`
	Username    string      `json:"username,omitempty"`
	Text        string      `json:"text"`
	Icon        string      `json:"icon_emoji,omitempty"`
	Attachments Attachments `json:"attachments,omitempty"`
}

func (m *Message) Send() error {
	if len(m.Channel) == 0 {
		return errors.New("Must specify a slack channel!")
	}

	if len(m.Text) == 0 {
		return errors.New("Must specifty text for slack message!")
	}

	//added for those who forget or don't want to use the hashtag prefix
	if !strings.HasPrefix(m.Channel, "#") {
		m.Channel = "#" + m.Channel
	}

	js, err := json.Marshal(m)
	if err != nil {
		return err
	}

	resp, err := http.PostForm(SLACK_API, url.Values{
		"token":   {SLACK_TOKEN},
		"payload": {string(js)},
	})
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
