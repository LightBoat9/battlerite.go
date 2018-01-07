package battlerite

import (
	"log"
	"strconv"
)

// Player contains information about a battlerite user.
// See: https://battlerite-docs.readthedocs.io/en/latest/players/players.html
type Player struct {
	Type                         string
	ID                           int
	LinkSelf                     string
	TitleID                      string
	Name                         string
	Picture                      int
	Wins                         int
	Losses                       int
	GradeScore                   int
	TimePlayed                   int
	Ranked2v2Wins                int
	Ranked2v2Loses               int
	Ranked3v3Wins                int
	Ranked3v3Losses              int
	Unranked2v2Wins              int
	Unranked2v2Losses            int
	Unranked3v3Wins              int
	Unranked3v3Losses            int
	BrawlWins                    int
	BrawlLosses                  int
	BattlegroundsWins            int
	BattlegroundsLosses          int
	AcctounXP                    int
	AccountLevel                 int
	TwitchAccountLinked          int
	VsAiPlayed                   int
	RatingMean                   int
	RatingDev                    int
	CharacterXP                  map[string]int
	CharacterWins                map[string]int
	CharacterLosses              map[string]int
	CharacterKills               map[string]int
	CharacterDeaths              map[string]int
	CharacterTimePlayed          map[string]int
	CharacterRanked2v2Wins       map[string]int
	CharacterRanked2v2Losses     map[string]int
	CharacterRanked3v3Wins       map[string]int
	CharacterRanked3v3Losses     map[string]int
	CharacterUnranked2v2Wins     map[string]int
	CharacterUnranked2v2Losses   map[string]int
	CharacterUnranked3v3Wins     map[string]int
	CharacterUnranked3v3Losses   map[string]int
	CharacterBrawlWins           map[string]int
	CharacterBrawlLosses         map[string]int
	CharacterBattlegroundsWins   map[string]int
	CharacterBattlegroundsLosses map[string]int
	CharacterLevels              map[string]int
}

// PlayerFilter contains filters for searching for battlerite users using GetPlayerFilter() in client.go.
// It should be noted if more than one field is used to sorting the API will return only the first.
// In most cases only one field should be used for filtering at one time.
// See: https://battlerite-docs.readthedocs.io/en/master/players/players.html#get-a-collection-of-players
type PlayerFilter struct {
	Names    []string
	UserIDs  []int
	SteamIDs []int
}

// Champions contains information about each battlerite champion.
// The string array contains the character index based on the battlerite mappings json.
// It also contains the names of BLC characters that battlerite characters were based on also used in the mappings json.
// See: https://github.com/gamelocker/battlerite-assets/tree/master/mappings
var Champions = map[string][2]string{
	"Lucie":    [2]string{"1", "Alchemist"},
	"Sirius":   [2]string{"2", "Astronomer"},
	"Iva":      [2]string{"3", "Engineer"},
	"Jade":     [2]string{"4", "Gunner"},
	"RuhKaan":  [2]string{"5", "Harbinger"},
	"Oldur":    [2]string{"6", "Herald"},
	"Ashka":    [2]string{"7", "Igniter"},
	"Varesh":   [2]string{"8", "Inhibitor"},
	"Pearl":    [2]string{"9", "Inquisitor"},
	"Taya":     [2]string{"10", "Nomad"},
	"Poloma":   [2]string{"11", "Psychopomp"},
	"Croak":    [2]string{"12", "Ranid"},
	"Freya":    [2]string{"13", "Ravener"},
	"Jumong":   [2]string{"14", "Seeker"},
	"Shifu":    [2]string{"15", "Spearmaster"},
	"Ezmo":     [2]string{"16", "Stormcaller"},
	"Bakko":    [2]string{"17", "Vanguard"},
	"Rook":     [2]string{"18", "Glutton"},
	"Pestilus": [2]string{"19", "BloodPriest"},
	"Destiny":  [2]string{"20", "MetalWarden"},
	"Raigon":   [2]string{"21", "Swordmaster"},
	"Blossum":  [2]string{"22", "Druid"},
	"Thorn":    [2]string{"25", "Thorn"},
	"Alysia":   [2]string{"41", "FrostMage"},
}

// GetChampionData returns a map of the data for each battlerite champion.
// The indexes are based on the mappings json.
// See: https://github.com/gamelocker/battlerite-assets/tree/master/mappings
func GetChampionData(stats map[string]interface{}, startIndex int) map[string]int {
	champXP := make(map[string]int)

	for k := range Champions {
		champID, err := strconv.Atoi(Champions[k][0])
		if err != nil {
			log.Fatal(err)
		}

		if stats[strconv.Itoa(startIndex+champID)] == nil {
			champXP[string(k)] = 0
		} else {
			champXP[string(k)] = int(stats[strconv.Itoa(startIndex+champID)].(float64))
		}

	}

	return champXP
}

// SinglePlayerFromData creates a player out of the data of a single battlerite user
func SinglePlayerFromData(data map[string]interface{}) Player {
	links := data["links"].(map[string]interface{})
	attributes := data["attributes"].(map[string]interface{})
	stats := attributes["stats"].(map[string]interface{})

	id, _ := strconv.Atoi(data["id"].(string))
	brawlWins := 0
	if stats["18"] != nil {
		brawlWins = int(stats["18"].(float64))
	}
	brawlLosses := 0
	if stats["19"] != nil {
		brawlLosses = int(stats["19"].(float64))
	}
	battlegroundsWins := 0
	if stats["22"] != nil {
		battlegroundsWins = int(stats["22"].(float64))
	}
	battlegroundsLosses := 0
	if stats["23"] != nil {
		battlegroundsLosses = int(stats["23"].(float64))
	}
	twitchAccountLinked := 0
	if stats["27"] != nil {
		twitchAccountLinked = int(stats["27"].(float64))
	}

	return Player{
		Type:                         data["type"].(string),
		ID:                           id,
		LinkSelf:                     links["self"].(string),
		TitleID:                      attributes["titleId"].(string),
		Name:                         attributes["name"].(string),
		Picture:                      int(stats["picture"].(float64)),
		Wins:                         int(stats["2"].(float64)),
		Losses:                       int(stats["3"].(float64)),
		GradeScore:                   int(stats["4"].(float64)),
		TimePlayed:                   int(stats["8"].(float64)),
		Ranked2v2Wins:                int(stats["10"].(float64)),
		Ranked2v2Loses:               int(stats["11"].(float64)),
		Ranked3v3Wins:                int(stats["12"].(float64)),
		Ranked3v3Losses:              int(stats["13"].(float64)),
		Unranked2v2Wins:              int(stats["14"].(float64)),
		Unranked2v2Losses:            int(stats["15"].(float64)),
		Unranked3v3Wins:              int(stats["16"].(float64)),
		Unranked3v3Losses:            int(stats["17"].(float64)),
		BrawlWins:                    brawlWins,
		BrawlLosses:                  brawlLosses,
		BattlegroundsWins:            battlegroundsWins,
		BattlegroundsLosses:          battlegroundsLosses,
		AcctounXP:                    int(stats["25"].(float64)),
		AccountLevel:                 int(stats["26"].(float64)),
		TwitchAccountLinked:          twitchAccountLinked,
		VsAiPlayed:                   int(stats["56"].(float64)),
		RatingMean:                   int(stats["70"].(float64)),
		RatingDev:                    int(stats["71"].(float64)),
		CharacterXP:                  GetChampionData(stats, 11000),
		CharacterWins:                GetChampionData(stats, 12000),
		CharacterLosses:              GetChampionData(stats, 13000),
		CharacterKills:               GetChampionData(stats, 14000),
		CharacterDeaths:              GetChampionData(stats, 15000),
		CharacterTimePlayed:          GetChampionData(stats, 16000),
		CharacterRanked2v2Wins:       GetChampionData(stats, 17000),
		CharacterRanked2v2Losses:     GetChampionData(stats, 18000),
		CharacterRanked3v3Wins:       GetChampionData(stats, 19000),
		CharacterRanked3v3Losses:     GetChampionData(stats, 20000),
		CharacterUnranked2v2Wins:     GetChampionData(stats, 21000),
		CharacterUnranked2v2Losses:   GetChampionData(stats, 22000),
		CharacterUnranked3v3Wins:     GetChampionData(stats, 23000),
		CharacterUnranked3v3Losses:   GetChampionData(stats, 24000),
		CharacterBrawlWins:           GetChampionData(stats, 25000),
		CharacterBrawlLosses:         GetChampionData(stats, 26000),
		CharacterBattlegroundsWins:   GetChampionData(stats, 27000),
		CharacterBattlegroundsLosses: GetChampionData(stats, 28000),
		CharacterLevels:              GetChampionData(stats, 40000),
	}
}

// MultiPlayersFromData creates a slice of players out of a slice of battlerite user datas
func MultiPlayersFromData(data []interface{}) ([]Player, error) {
	playerDatas := []Player{}

	if len(data) < 1 {
		return playerDatas, nil
	}

	for _, tempData := range data {
		player := SinglePlayerFromData(tempData.(map[string]interface{}))
		playerDatas = append(playerDatas, player)
	}

	return playerDatas, nil
}
