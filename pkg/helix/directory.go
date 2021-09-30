package helix

import (
	"github.com/gempir/gempbot/pkg/log"
	nickHelix "github.com/nicklaw5/helix/v2"
)

func (c *Client) GetTopChannels() []string {
	response, err := c.Client.GetStreams(&nickHelix.StreamsParams{Type: "live", First: 100})
	if err != nil {
		log.Error(err.Error())
	}

	var ids []string
	for _, stream := range response.Data.Streams {
		ids = append(ids, stream.UserID)
	}

	return ids
}
