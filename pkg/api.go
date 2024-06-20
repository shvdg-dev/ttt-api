package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/resources"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/sessions"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/users"
)

// API represents the main entry point to interact with the API functionalities.
type API struct {
	Artists      *artists.API
	Difficulties *difficulties.API
	Instruments  *instruments.API
	Resources    *resources.API
	Sessions     *sessions.API
	Sources      *sources.API
	Tabs         *tabs.API
	Tracks       *tracks.API
	Users        *users.API
}

// NewAPI creates a new instance of the API.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{
		Artists:      artists.NewAPI(database),
		Difficulties: difficulties.NewAPI(database),
		Instruments:  instruments.NewAPI(database),
		Resources:    resources.NewAPI(database),
		Sessions:     sessions.NewAPI(database),
		Sources:      sources.NewAPI(database),
		Tabs:         tabs.NewAPI(database),
		Tracks:       tracks.NewAPI(database),
		Users:        users.NewAPI(database),
	}
}
