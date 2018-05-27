package battleritego

// MatchFilter contains filter parameters for searching matches.
// See https://battlerite-docs.readthedocs.io/en/master/matches/matches.html#get-a-collection-of-matches
type MatchFilter struct {
	PageOffset     int
	PageLimit      int
	Sort           string
	CreatedAtStart string
	CreatedAtEnd   string
	PlayerIDs      []string
	PatchVersion   []string
}

// Match contains information about a match.
// See https://battlerite-docs.readthedocs.io/en/master/matches/matches.html
type Match struct {
	Type         string
	ID           string
	LinkSelf     string
	CreatedAt    string
	Duration     int
	GameMode     string
	PatchVersion string
	ShardID      string
	TitleID      string
	MapType      string
	MapID        string
	Asset        Asset
	Participants []Participant
	Rosters      []Roster
	MatchPlayers []MatchPlayer
	Rounds       []Round
	Spectators   interface{}
}

// SingleMatchFromResponse returns a single Match from a Response.
// See Response in client.go.
func SingleMatchFromResponse(res Response) Match {
	data := res.Data.(map[string]interface{})
	links := data["links"].(map[string]interface{})
	attributes := data["attributes"].(map[string]interface{})
	stats := attributes["stats"].(map[string]interface{})

	relationships := data["relationships"].(map[string]interface{})

	assetsData := relationships["assets"].(map[string]interface{})["data"].([]interface{})
	rostersData := relationships["rosters"].(map[string]interface{})["data"].([]interface{})
	roundsData := relationships["rounds"].(map[string]interface{})["data"].([]interface{})

	spectators := relationships["spectators"].(map[string]interface{})
	spectatorsData := spectators["data"]

	included := res.Included.([]interface{})

	participantList := []Participant{}
	rosterList := []Roster{}
	matchPlayerList := []MatchPlayer{}
	roundList := []Round{}
	asset := Asset{}

	for _, incl := range included {
		switch incl.(map[string]interface{})["type"] {
		case "roster":
			for _, dat := range rostersData {
				if incl.(map[string]interface{})["id"] == dat.(map[string]interface{})["id"] {
					roster := SingleRosterFromData(incl.(map[string]interface{}))
					rosterList = append(rosterList, roster)
				}
			}
		case "round":
			for _, dat := range roundsData {
				if incl.(map[string]interface{})["id"] == dat.(map[string]interface{})["id"] {
					round := SingleRoundFromData(incl.(map[string]interface{}))
					roundList = append(roundList, round)
				}
			}
		case "asset":
			for _, dat := range assetsData {
				if incl.(map[string]interface{})["id"] == dat.(map[string]interface{})["id"] {
					asset = SingleAssetFromData(incl.(map[string]interface{}))
				}
			}
		}
	}

	for _, incl := range included {
		switch incl.(map[string]interface{})["type"] {
		case "participant":
			for _, rost := range rosterList {
				for _, ply := range rost.Participants.([]interface{}) {
					if ply.(map[string]interface{})["id"] == incl.(map[string]interface{})["id"] {
						partic := SingleParticipantFromData(incl.(map[string]interface{}))
						participantList = append(participantList, partic)
					}
				}
			}
		}
	}

	for _, incl := range included {
		switch incl.(map[string]interface{})["type"] {
		case "player":
			for _, dat := range participantList {
				tempID := dat.Relationships.(map[string]interface{})["player"].(map[string]interface{})["data"].(map[string]interface{})["id"]
				if tempID == incl.(map[string]interface{})["id"] {
					matchPlr := SingleMatchPlayerFromData(incl.(map[string]interface{}))
					matchPlayerList = append(matchPlayerList, matchPlr)
				}
			}
		}
	}

	return Match{
		Type:         data["type"].(string),
		ID:           data["id"].(string),
		LinkSelf:     links["self"].(string),
		CreatedAt:    attributes["createdAt"].(string),
		Duration:     int(attributes["duration"].(float64)),
		GameMode:     attributes["gameMode"].(string),
		PatchVersion: attributes["patchVersion"].(string),
		ShardID:      attributes["shardId"].(string),
		MapType:      stats["type"].(string),
		MapID:        stats["mapID"].(string),
		Asset:        asset,
		Participants: participantList,
		Rosters:      rosterList,
		MatchPlayers: matchPlayerList,
		Rounds:       roundList,
		Spectators:   spectatorsData,
	}
}

// MultiMatchesFromResponse returns a slice of Matches from from a Response.
// See Response in client.go.
func MultiMatchesFromResponse(res Response) []Match {
	responses := []Response{}
	matches := []Match{}

	for _, data := range res.Data.([]interface{}) {
		temp := Response{
			Data: data,
		}
		responses = append(responses, temp)
	}

	for i := range responses {
		responses[i].Included = res.Included
	}

	for _, singRes := range responses {
		match := SingleMatchFromResponse(singRes)
		matches = append(matches, match)
	}

	return matches
}
