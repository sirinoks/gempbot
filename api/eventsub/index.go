package eventsub

import (
	"net/http"

	"github.com/gempir/gempbot/pkg/config"
	"github.com/gempir/gempbot/pkg/eventsub"
	"github.com/gempir/gempbot/pkg/helix"
	"github.com/gempir/gempbot/pkg/store"
	nickHelix "github.com/nicklaw5/helix"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cfg := config.FromEnv()
	db := store.NewDatabase(cfg)
	helixClient := helix.NewClient(cfg, db)
	eventSubManager := eventsub.NewEventSubManager(cfg, helixClient, db)

	if r.URL.Query().Get("type") == nickHelix.EventSubTypeChannelPointsCustomRewardRedemptionAdd {
		var redemption nickHelix.EventSubChannelPointsCustomRewardRedemptionEvent
		done, err := eventSubManager.HandleWebhook(w, r, redemption)
		if err != nil || done {
			if err != nil {
				http.Error(w, err.Error(), err.Status())
			}
			return
		}
		eventSubManager.HandleChannelPointsCustomRewardRedemption(redemption)
		return
	}

	http.Error(w, "Invalid event type", http.StatusBadRequest)
}
