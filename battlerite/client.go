// Package battlerite provides access to the battlerite game data service.
package battlerite

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// BaseURL for the Gamelocker API using their current DC01 Datacenter.
// See: https://battlerite-docs.readthedocs.io/en/latest/datacenters/datacenters.html
const BaseURL = "https://api.dc01.gamelockerapp.com/shards/global/"

// http request client with a timeout of 10 seconds.
var request = &http.Client{Timeout: 10 * time.Second}

// Response represents a response of a single piece of data from the http client request.
type Response struct {
	Data     interface{} `json:"data,omitempty"`
	Errors   interface{} `json:"errors,omitempty"`
	Links    interface{} `json:"links,omitempty"`
	Included interface{} `json:"included,omitempty"`
	Meta     interface{} `json:"meta,omitempty"`
}

// Client stores an API key.
type Client struct {
	APIKey string
}

// getPageBytes retrieves the bites slice of a page.
func (client Client) getPageBytes(URL string) ([]byte, error) {
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("Authorization", client.APIKey)
	req.Header.Set("Accept", "application/vnd.api+json")

	r, err := request.Do(req)
	if err != nil {
		return nil, err
	}

	if len(r.Header["X-Ratelimit-Remaining"]) > 0 && r.Header["X-Ratelimit-Remaining"][0] == "0" {
		return nil, errors.New("Request rate limit hit 0, wait for more requests; " +
			"Learn more: https://battlerite-docs.readthedocs.io/en/master/ratelimits/ratelimits.html ")
	}

	defer r.Body.Close()

	page, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return page, nil
}

// getData returns data from the request URL.
func (client Client) getData(URL string) (Response, error) {
	page, err := client.getPageBytes(URL)
	if err != nil {
		return Response{}, err
	}

	res := Response{}
	jsonErr := json.Unmarshal(page, &res)
	if jsonErr != nil {
		return Response{}, jsonErr
	}

	return res, nil
}

// GetStatus receives the Status of the Gamelocker battlerite API.
func (client Client) GetStatus() (Status, error) {
	URL := "https://api.dc01.gamelockerapp.com/status"

	res, err := client.getData(URL)

	if err != nil {
		return Status{}, err
	}

	attributes := res.Data.(map[string]interface{})["attributes"].(map[string]interface{})
	return Status{
		res.Data.(map[string]interface{})["type"].(string),
		res.Data.(map[string]interface{})["id"].(string),
		attributes["releasedAt"].(string),
		attributes["version"].(string),
	}, nil
}

// GetPlayer receives a single Player using the players battlerite ID.
func (client Client) GetPlayer(id int) (Player, error) {
	URL := fmt.Sprintf("%splayers/%s", BaseURL, strconv.Itoa(id))

	res, err := client.getData(URL)
	if err != nil {
		return Player{}, err
	}

	return SinglePlayerFromData(res.Data.(map[string]interface{})), nil
}

// GetPlayersFiltered receives a slice of players using the passed in PlayerFilter.
func (client Client) GetPlayersFiltered(filter PlayerFilter) ([]Player, error) {
	URL := fmt.Sprintf("%splayers?", BaseURL)

	if filter.Names != nil {
		URL += fmt.Sprintf("&filter[playerNames]=%s", strings.Join(filter.Names, ","))
	}
	if filter.UserIDs != nil {
		// Convert []int -> []string
		strUserIDs := []string{}
		for _, id := range filter.UserIDs {
			strUserIDs = append(strUserIDs, strconv.Itoa(id))
		}

		URL += fmt.Sprintf("&filter[playerIds]=%s", strings.Join(strUserIDs, ","))
	}
	if filter.SteamIDs != nil {
		// Convert []int -> []string
		strSteamIDs := []string{}
		for _, id := range filter.SteamIDs {
			strSteamIDs = append(strSteamIDs, strconv.Itoa(id))
		}

		URL += fmt.Sprintf("&filter[steamIds]=%s", strings.Join(strSteamIDs, ","))
	}

	res, err := client.getData(URL)
	if err != nil {
		return []Player{}, err
	}

	return MultiPlayersFromData(res.Data.([]interface{}))
}

// GetTeamsFiltered returns a slice of teams using the TeamFilter.
// See TeamFilter in team.go.
func (client Client) GetTeamsFiltered(filter TeamFilter) ([]Team, error) {
	// Ensure TeamFilter contains Season and PlayerIDs
	if filter.Season == 0 {
		return []Team{}, errors.New("TeamFilter must contain a Season")
	}
	if filter.PlayerIDs == nil {
		return []Team{}, errors.New("TeamFilter must contain PlayerIDs")
	}

	season := fmt.Sprintf("&tag[season]=%s", strconv.Itoa(filter.Season))

	// Convert []int -> []string
	strPlayerIDs := []string{}
	for _, id := range filter.PlayerIDs {
		strPlayerIDs = append(strPlayerIDs, strconv.Itoa(id))
	}

	playerIDs := fmt.Sprintf("&tag[playerIds]=%s", strings.Join(strPlayerIDs, ","))

	URL := fmt.Sprintf("%steams?%s%s", BaseURL, season, playerIDs)

	res, err := client.getData(URL)
	if err != nil {
		return []Team{}, err
	}

	return MultiTeamsFromData(res.Data.([]interface{}))
}

// GetMatch returns a single match filtered by ID.
func (client Client) GetMatch(id string) (Match, error) {
	URL := fmt.Sprintf("%smatches/%s", BaseURL, id)
	res, err := client.getData(URL)
	if err != nil {
		return Match{}, nil
	}

	return SingleMatchFromResponse(res), nil
}

// GetMatchesFiltered returns a slice of matches filtered by MatchFilter.
// See MatchFilter in match.go
func (client Client) GetMatchesFiltered(filter MatchFilter) ([]Match, error) {
	URL := fmt.Sprintf("%smatches?", BaseURL)

	if filter.PageOffset != 0 {
		URL += "&page[offset]=" + strconv.Itoa(filter.PageOffset)
	}
	if filter.PageLimit != 0 {
		URL += "&page[limit]=" + strconv.Itoa(filter.PageLimit)
	}
	if filter.Sort != "" {
		URL += fmt.Sprintf("&sort=%s", filter.Sort)
	}
	if filter.CreatedAtStart != "" {
		URL += fmt.Sprintf("&filter[createdAt-start]=%s", filter.CreatedAtStart)
	}
	if filter.CreatedAtEnd != "" {
		URL += fmt.Sprintf("&filter[createdAt-end]=%s", filter.CreatedAtEnd)
	}
	if filter.PlayerIDs != nil {
		URL += "&filter[playerIds]=" + strings.Join(filter.PlayerIDs, ",")
	}
	if filter.PatchVersion != nil {
		URL += "&filter[patchVersion]=" + strings.Join(filter.PatchVersion, ",")
	}

	res, err := client.getData(URL)
	if err != nil {
		return []Match{}, err
	}

	return MultiMatchesFromResponse(res), nil
}

// GetTelemetry returns telemetry data relating to match.
// The URL can be found from Match.Asset.URL after searching for a match using
// either GetMactch or GetMatchesFiltered.
// See match.go for more information about Matches
func (client Client) GetTelemetry(URL string) (Telemetry, error) {
	page, err := client.getPageBytes(URL)
	if err != nil {
		return Telemetry{}, err
	}

	var data []interface{}
	jsonErr := json.Unmarshal(page, &data)
	if jsonErr != nil {
		return Telemetry{}, jsonErr
	}

	matchStart := MatchStart{}
	roundEventList := []RoundEvent{}
	userRoundSpellList := []UserRoundSpell{}
	deathEventList := []DeathEvent{}
	matchReservedUserList := []MatchReservedUser{}
	queueEventList := []QueueEvent{}
	teamUpdateEventList := []TeamUpdateEvent{}
	serverShutdown := ServerShutdown{}
	roundFinishedEventList := []RoundFinishedEvent{}
	matchFinishedEvent := MatchFinishedEvent{}

	for _, event := range data {
		switch event.(map[string]interface{})["type"] {
		case "Structures.MatchStart":
			matchStart = MatchStartFromData(event.(map[string]interface{}))
		case "Structures.RoundEvent":
			roundEvent := RoundEventFromData(event.(map[string]interface{}))
			roundEventList = append(roundEventList, roundEvent)
		case "Structures.UserRoundSpell":
			userRoundSpell := UserRoundSpellFromData(event.(map[string]interface{}))
			userRoundSpellList = append(userRoundSpellList, userRoundSpell)
		case "Structures.DeathEvent":
			deathEvent := DeathEventFromData(event.(map[string]interface{}))
			deathEventList = append(deathEventList, deathEvent)
		case "Structures.MatchReservedUser":
			matchReservedUser := MatchReservedUserFromData(event.(map[string]interface{}))
			matchReservedUserList = append(matchReservedUserList, matchReservedUser)
		case "com.stunlock.service.matchmaking.avro.QueueEvent":
			queueEvent := QueueEventFromData(event.(map[string]interface{}))
			queueEventList = append(queueEventList, queueEvent)
		case "com.stunlock.battlerite.team.TeamUpdateEvent":
			teamUpdateEvent := TeamUpdateEventFromData(event.(map[string]interface{}))
			teamUpdateEventList = append(teamUpdateEventList, teamUpdateEvent)
		case "Structures.ServerShutdown":
			serverShutdown = ServerShutdownFromData(event.(map[string]interface{}))
		case "Structures.RoundFinishedEvent":
			roundFinishedEvent := RoundFinishedEventFromData(event.(map[string]interface{}))
			roundFinishedEventList = append(roundFinishedEventList, roundFinishedEvent)
		case "Structures.MatchFinishedEvent":
			matchFinishedEvent = MatchFinishedEventFromData(event.(map[string]interface{}))
		}
	}

	return Telemetry{
		MatchStart:          matchStart,
		RoundEvents:         roundEventList,
		UserRoundSpells:     userRoundSpellList,
		DeathEvents:         deathEventList,
		MatchReservedUsers:  matchReservedUserList,
		QueueEvents:         queueEventList,
		TeamUpdateEvents:    teamUpdateEventList,
		ServerShutdown:      serverShutdown,
		RoundFinishedEvents: roundFinishedEventList,
		MatchFinishedEvent:  matchFinishedEvent,
	}, nil
}
