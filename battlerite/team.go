package battlerite

import "strconv"

// Team contains information about a battlerite team.
// See https://battlerite-docs.readthedocs.io/en/master/teams/teams.html
type Team struct {
	Type               string
	ID                 int
	Name               string
	ShardID            string
	TitleID            string
	PlacementGamesLeft int
	Avatar             int
	Wins               int
	Losses             int
	Members            []int
	Division           int
	DivisionRating     int
	TopDivision        int
	TopDivisionRating  int
	League             int
	TopLeague          int
	Assets             map[string]interface{}
}

// TeamFilter contains parameters for filtering teams using
// GetTeamsFiltered.
// Note that both parameters are required.
// See client.go for more information.
type TeamFilter struct {
	Season    int
	PlayerIDs []int
}

// SingleTeamFromData returns a single Team from data.
func SingleTeamFromData(data map[string]interface{}) Team {
	attributes := data["attributes"].(map[string]interface{})
	relationships := data["relationships"].(map[string]interface{})
	stats := attributes["stats"].(map[string]interface{})
	assets := relationships["assets"].(map[string]interface{})

	id, _ := strconv.Atoi(data["id"].(string))

	members := []int{}
	for _, user := range stats["members"].([]interface{}) {
		members = append(members, int(user.(float64)))
	}

	return Team{
		Type:               data["type"].(string),
		ID:                 id,
		Name:               attributes["name"].(string),
		ShardID:            attributes["shardId"].(string),
		TitleID:            attributes["titleId"].(string),
		PlacementGamesLeft: int(stats["placementGamesLeft"].(float64)),
		Avatar:             int(stats["avatar"].(float64)),
		Wins:               int(stats["wins"].(float64)),
		Losses:             int(stats["losses"].(float64)),
		Members:            members,
		Division:           int(stats["division"].(float64)),
		DivisionRating:     int(stats["divisionRating"].(float64)),
		TopDivision:        int(stats["topDivision"].(float64)),
		TopDivisionRating:  int(stats["topDivisionRating"].(float64)),
		League:             int(stats["league"].(float64)),
		TopLeague:          int(stats["topLeague"].(float64)),
		Assets:             assets,
	}
}

// MultiTeamsFromData returns a slice of teams from the data.
func MultiTeamsFromData(data []interface{}) ([]Team, error) {
	teamData := []Team{}

	if len(data) < 1 {
		return teamData, nil
	}

	for _, tempData := range data {
		team := SingleTeamFromData(tempData.(map[string]interface{}))
		teamData = append(teamData, team)
	}

	return teamData, nil
}
