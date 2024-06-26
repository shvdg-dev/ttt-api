package internal

import (
	faker "github.com/brianvoe/gofakeit/v7"
	logic "github.com/shvdg-dev/base-logic/pkg"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	"log"
)

// Seeder helps with deleting data from the database
type Seeder struct {
	Config  *Config
	API     *api.API
	Factory *Factory
}

// NewSeeder creates a new instance of Seeder
func NewSeeder(config *Config, api *api.API) *Seeder {
	return &Seeder{
		API:     api,
		Config:  config,
		Factory: NewFactory(config.Seeding.Dummies, config.Seeding.Instruments, config.Seeding.Difficulties)}
}

// SeedTables attempts to seed the database with the minimally required values and dummy data.
func (s *Seeder) SeedTables() {
	s.minimumSeed()
	s.dummySeed()
}

// minimumSeed when permitted, seeds the database with the minimally required values.
func (s *Seeder) minimumSeed() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseAllowMinimumSeedingCommand) {
		log.Println("It is not allowed to seed the database with the minimally required values.")
		return
	}
	s.seedAdmin()
	s.seedInstruments()
	s.seedDifficulties()
	s.seedSources()
	s.seedEndpoints()
}

// seedAdmin inserts an administrator user into the database.
func (s *Seeder) seedAdmin() {
	email := logic.GetEnvValueAsString(KeyAdminInitialEmail)
	password := logic.GetEnvValueAsString(KeyAdminInitialPassword)
	if email != "" && password != "" {
		s.API.Users.InsertUser(email, password)
	} else {
		log.Println("Did not insert the initial admin account as no credentials were defined")
	}
}

// seedInstruments seeds the instruments table with the default instruments.
func (s *Seeder) seedInstruments() {
	s.API.Instruments.InsertInstruments(s.Config.Seeding.Instruments...)
}

// seedDifficulties seeds the difficulties table with the default difficulties.
func (s *Seeder) seedDifficulties() {
	s.API.Difficulties.InsertDifficulties(s.Config.Seeding.Difficulties...)
}

// seedSources seeds the sources from the config file.
func (s *Seeder) seedSources() {
	s.API.Sources.InsertSources(s.Config.Seeding.Sources...)
}

// seedEndpoints seeds the endpoints from the config file.
func (s *Seeder) seedEndpoints() {
	s.API.Endpoints.InsertEndpoints(s.Config.Seeding.Endpoints...)
}

// dummySeed when permitted, seeds the database with dummy data.
func (s *Seeder) dummySeed() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseAllowDummySeedingCommand) {
		log.Println("It is not allowed to seed the database with dummy data.")
		return
	}
	s.seedDummyArtists()
}

// seedDummyArtists inserts dummy artists, tracks, and tabs into the database.
func (s *Seeder) seedDummyArtists() {
	artists := s.Factory.CreateDummyArtists(
		uint(faker.Number(s.Config.Seeding.Dummies.Tracks.Min, s.Config.Seeding.Dummies.Tracks.Max)))
	// Insert the artist
	s.API.Artists.InsertArtists(artists...)
	for _, artist := range artists {
		// Insert the tracks of an artist
		s.API.Tracks.InsertTracks(artist.Tracks...)
		for _, track := range artist.Tracks {
			// Insert the tabs of a track
			s.API.Tabs.InsertTabs(track.Tabs...)
		}
	}
}
