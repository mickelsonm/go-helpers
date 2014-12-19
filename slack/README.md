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

		msg := slack.Message{
			 //channels can be prefixed with a hashtag, but it isn't required
			Channel: "Debugging",
			 //username can also be a service name
			Username: "matt",
			 //text is the message you want to be displayed
			Text: "Yo...your servers be crashing yo!"
		}
		//send returns nil on success and error if there is one that comes up
		msg.Send()
	}
```
