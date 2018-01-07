package battlerite

// MatchPlayer holds information about a player from a match
type MatchPlayer struct {
	Type          string
	ID            string
	LinkSelf      string
	Attributes    interface{}
	Relationships interface{}
}

// SingleMatchPlayerFromData returns a MatchPlayer from passed in data
func SingleMatchPlayerFromData(data map[string]interface{}) MatchPlayer {
	attributes := data["attributes"].(map[string]interface{})
	relationships := data["relationships"].(map[string]interface{})
	assets := relationships["assets"].(map[string]interface{})
	links := data["links"].(map[string]interface{})

	return MatchPlayer{
		Type:          data["type"].(string),
		ID:            data["id"].(string),
		LinkSelf:      links["self"].(string),
		Attributes:    attributes,
		Relationships: assets["data"],
	}
}
