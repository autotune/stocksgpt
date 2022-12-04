package slackbot

import (
	"fmt"

	"github.com/slack-go/slack"
)

// PostToSlack posts a message to a slack channel
func PostToSlack(client *slack.Client, channel string, message string) error {
	_, _, err := client.PostMessage(channel, slack.MsgOptionText(message, false))
	if err != nil {
		return fmt.Errorf("failed to post message to slack: %s", err)
	}
	return nil
}
