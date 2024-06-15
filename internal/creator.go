package internal

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	"log"
)

// Creator helps with deleting data from the database
type Creator struct {
	API *api.API
}

// NewCreator creates a new instance of Creator
func NewCreator(API *api.API) *Creator {
	return &Creator{API: API}
}

// CreateTables when permitted, creates tables in the database
func (c *Creator) CreateTables() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseAllowCreatingCommand) {
		log.Fatalf("It is not allowed to create new tables for the database")
	}

	// Lookup table
	c.API.Instruments.CreateInstrumentsTable()
	c.API.Difficulties.CreateDifficultiesTable()

	// Data storage
	c.API.Artists.CreateArtistsTable()
	c.API.IdReferences.CreateIdReferencesTable()
	c.API.Sessions.CreateSessionsTable()
	c.API.Tabs.CreateTabsTable()
	c.API.Tracks.CreateTracksTable()
	c.API.Users.CreateUsersTable()

	// Relation mappings
	c.API.Artists.CreateArtistTrackTable()
	c.API.Tracks.CreateTrackTabTable()
}
