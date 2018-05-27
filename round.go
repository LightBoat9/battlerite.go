package battleritego

// Round contains information about a round in a match.
// See match.go for more information.
// See https://battlerite-docs.readthedocs.io/en/master/matches/matches.html#rounds
type Round struct {
	Type        string
	ID          string
	WinningTeam int
	Duration    int
	Ordinal     int
}

// SingleRoundFromData returns a single Round from data.
func SingleRoundFromData(data map[string]interface{}) Round {
	attributes := data["attributes"].(map[string]interface{})
	stats := attributes["stats"].(map[string]interface{})

	return Round{
		Type:        data["type"].(string),
		ID:          data["id"].(string),
		WinningTeam: int(stats["winningTeam"].(float64)),
		Duration:    int(attributes["duration"].(float64)),
		Ordinal:     int(attributes["ordinal"].(float64)),
	}
}
