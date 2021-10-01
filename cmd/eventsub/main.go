package main

import (
	"strings"

	"github.com/gempir/gempbot/pkg/channelpoint"
	"github.com/gempir/gempbot/pkg/chat"
	"github.com/gempir/gempbot/pkg/config"
	"github.com/gempir/gempbot/pkg/emotechief"
	"github.com/gempir/gempbot/pkg/eventsub"
	"github.com/gempir/gempbot/pkg/helix"
	"github.com/gempir/gempbot/pkg/log"
	"github.com/gempir/gempbot/pkg/store"
)

var (
	cfg         *config.Config
	db          *store.Database
	helixClient *helix.Client
)

func main() {
	cfg = config.FromEnv()
	db = store.NewDatabase(cfg)
	helixClient = helix.NewClient(cfg, db)
	emotechief := emotechief.NewEmoteChief(cfg, db)
	chatClient := chat.NewClient(cfg)
	cpm := channelpoint.NewChannelPointManager(cfg, helixClient, db, emotechief, chatClient)
	esm := eventsub.NewEventSubManager(cfg, helixClient, db, cpm, chatClient)

	esm.RemoveAllEventSubSubscriptions("")

	for _, token := range db.GetAllUserAccessToken() {
		log.Infof("Correcting subscriptions for %s", token.OwnerTwitchID)
		esm.SubscribeChannelPoints(token.OwnerTwitchID)

		if strings.Contains(token.Scopes, "predictions") {
			esm.SubscribePredictions(token.OwnerTwitchID)
		}
	}
}
