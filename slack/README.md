Slack Helper
==========

This is a pretty basic slack notification helper. When I wrote it, its purpose was for quick and dirty real-time debugging.

Example:
```go
	package main

	import(
		"github.com/mickelsonm/go-helpers/slack"
	)

	func main(){
		slack.SLACK_API = "https://myorg.slack.com/services/hooks/incoming-webhook"
		slack.SLACK_TOKEN = "MYSECRETTOKEN1234"

		m := slack.Message{
			Channel:  "random",
			Username: "Randomness",
			Text:     "Some guy is testing some new functionality thingy on slack",
		}

		m.Attachments = append(m.Attachments, slack.Attachment{
			//Pretext:    "BEDROCK NEWS:",
			Text:       "Who is Fred Flintstone?",
			Color:      "#CCC",
			Fallback:   "This is some fallback text...in case something funny happens.",
			Title:      "Man this sure is random:",
			AuthorName: "Barney Rubble Trivia",
			AuthorIcon: "http://upload.wikimedia.org/wikipedia/en/thumb/a/ad/Fred_Flintstone.png/165px-Fred_Flintstone.png",
			AuthorLink: "http://upload.wikimedia.org/wikipedia/en/thumb/a/ad/Fred_Flintstone.png/165px-Fred_Flintstone.png",
			Fields: slack.AttachmentFields{
				slack.AttachmentField{Title: "#1", Value: "Some guy that used to play in the band."},
				slack.AttachmentField{Title: "#2", Value: "A fellow that got lost at Wal-Mart."},
				slack.AttachmentField{Title: "#3", Value: "A girl friend's sister's cousin."},
				slack.AttachmentField{Title: "#4", Value: "That one guy that sold that guy a car."},
				slack.AttachmentField{Title: "#5", Value: "Just a caveman."},
			},
		})

		m.Send()
	}
```
