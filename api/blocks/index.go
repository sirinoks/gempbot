package blocks

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gempir/gempbot/pkg/api"
	"github.com/gempir/gempbot/pkg/auth"
	"github.com/gempir/gempbot/pkg/config"
	"github.com/gempir/gempbot/pkg/helixclient"
	"github.com/gempir/gempbot/pkg/store"
	"github.com/gempir/gempbot/pkg/user"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cfg := config.FromEnv()
	db := store.NewDatabase(cfg)
	helixClient := helixclient.NewClient(cfg, db)
	userAdmin := user.NewUserAdmin(cfg, db, helixClient, nil)
	auth := auth.NewAuth(cfg, db, helixClient)

	authResp, _, apiErr := auth.AttemptAuth(r, w)
	if apiErr != nil {
		return
	}
	userID := authResp.Data.UserID

	if r.URL.Query().Get("managing") != "" {
		userID, apiErr = userAdmin.CheckEditor(r, userAdmin.GetUserConfig(userID))
		if apiErr != nil {
			http.Error(w, apiErr.Error(), apiErr.Status())
			return
		}
	}

	if r.Method == http.MethodGet {
		page := r.URL.Query().Get("page")
		if page == "" {
			page = "1"
		}

		pageNumber, err := strconv.Atoi(page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		blocks := db.GetEmoteBlocks(userID, pageNumber, api.BLOCKS_PAGE_SIZE)
		api.WriteJson(w, blocks, http.StatusOK)
		return
	}
	if r.Method == http.MethodPatch {
		var req blockRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		toBlock := []string{}
		for _, emote := range strings.Split(req.EmoteIds, ",") {
			toBlock = append(toBlock, strings.TrimSpace(emote))
		}

		err = db.BlockEmotes(userID, toBlock, req.EmoteType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == http.MethodDelete {
		var req deleteRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = db.DeleteEmoteBlock(userID, req.EmoteID, req.Type)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type deleteRequest struct {
	store.EmoteBlock
}

type blockRequest struct {
	EmoteIds  string `json:"emoteIds"`
	EmoteType string `json:"type"`
}
