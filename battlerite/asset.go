// Package battlerite provides access to the battlerite game data service.
package battlerite

// Asset contains data for match assets, mainly used to pull telemetry data
// for the match.
// See https://battlerite-docs.readthedocs.io/en/master/matches/matches.html
type Asset struct {
	Type        string
	ID          string
	URL         string
	CreatedAt   string
	Description string
	Name        string
}

// SingleAssetFromData returns an Asset from the passed in data
func SingleAssetFromData(data map[string]interface{}) Asset {
	attributes := data["attributes"].(map[string]interface{})

	return Asset{
		Type:        data["type"].(string),
		ID:          data["id"].(string),
		URL:         attributes["URL"].(string),
		CreatedAt:   attributes["createdAt"].(string),
		Description: attributes["description"].(string),
		Name:        attributes["name"].(string),
	}
}
