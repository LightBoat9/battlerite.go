package battleritego

import (
	"strconv"
)

// Roster contains information about a roster of players in a match.
// See match.go for more information.
// See https://battlerite-docs.readthedocs.io/en/master/matches/matches.html#rosters
type Roster struct {
	Type         string
	ID           string
	ShardID      string
	Won          bool
	Score        int
	Participants interface{}
	Team         interface{}
}

// SingleRosterFromData returns a single Roster from data.
func SingleRosterFromData(data map[string]interface{}) Roster {
	attributes := data["attributes"].(map[string]interface{})
	stats := attributes["stats"].(map[string]interface{})
	relationships := data["relationships"].(map[string]interface{})
	participants := relationships["participants"].(map[string]interface{})
	partData := participants["data"]
	team := relationships["team"].(map[string]interface{})
	teamData := team["data"]

	won, _ := strconv.ParseBool(attributes["won"].(string))

	return Roster{
		Type:         data["type"].(string),
		ID:           data["id"].(string),
		ShardID:      attributes["shardId"].(string),
		Won:          won,
		Score:        int(stats["score"].(float64)),
		Participants: partData,
		Team:         teamData,
	}
}
