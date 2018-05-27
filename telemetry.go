package battleritego

// Telemetry contains a matches telemetry data.
// See https://battlerite-docs.readthedocs.io/en/master/telemetry/telemetry.html
type Telemetry struct {
	MatchStart          MatchStart
	RoundEvents         []RoundEvent
	UserRoundSpells     []UserRoundSpell
	DeathEvents         []DeathEvent
	MatchReservedUsers  []MatchReservedUser
	QueueEvents         []QueueEvent
	TeamUpdateEvents    []TeamUpdateEvent
	ServerShutdown      ServerShutdown
	RoundFinishedEvents []RoundFinishedEvent
	MatchFinishedEvent  MatchFinishedEvent
}

// MatchStart is a telemetry event containing information at a matches start.
type MatchStart struct {
	Type            string
	Cursor          int
	Time            int
	MatchID         string
	ExternalMatchID string
	Version         string
	EventType       string
	GameMode        int
	MapID           string
	TeamSize        int
	Region          string
}

// MatchStartFromData returns a MatchStart from data.
func MatchStartFromData(data map[string]interface{}) MatchStart {
	dataObject := data["dataObject"].(map[string]interface{})

	return MatchStart{
		Type:            data["type"].(string),
		Cursor:          int(data["cursor"].(float64)),
		Time:            int(dataObject["time"].(float64)),
		MatchID:         dataObject["matchID"].(string),
		ExternalMatchID: dataObject["externalMatchID"].(string),
		Version:         dataObject["version"].(string),
		EventType:       dataObject["type"].(string),
		GameMode:        int(dataObject["gameMode"].(float64)),
		MapID:           dataObject["mapID"].(string),
		TeamSize:        int(dataObject["teamSize"].(float64)),
		Region:          dataObject["region"].(string),
	}
}

// RoundEvent is a telemetry event containing information about various events during a round.
type RoundEvent struct {
	Type            string
	Cursor          int
	Time            int
	MatchID         string
	ExternalMatchID string
	UserID          string
	Round           int
	Character       int
	EventType       string
	Value           int
	TimeIntoRound   int
}

// RoundEventFromData returns a RoundEvent from data.
func RoundEventFromData(data map[string]interface{}) RoundEvent {
	dataObject := data["dataObject"].(map[string]interface{})

	return RoundEvent{
		Type:            data["type"].(string),
		Cursor:          int(data["cursor"].(float64)),
		Time:            int(dataObject["time"].(float64)),
		MatchID:         dataObject["matchID"].(string),
		ExternalMatchID: dataObject["externalMatchID"].(string),
		UserID:          dataObject["userID"].(string),
		Round:           int(dataObject["round"].(float64)),
		Character:       int(dataObject["character"].(float64)),
		EventType:       dataObject["type"].(string),
		Value:           int(dataObject["value"].(float64)),
		TimeIntoRound:   int(dataObject["timeIntoRound"].(float64)),
	}
}

// UserRoundSpell is a telemetry event containing information about a characters ability use.
type UserRoundSpell struct {
	Type         string
	Cursor       int
	Time         int
	AccountID    string
	MatchID      string
	Round        int
	Character    int
	TypeID       int
	SourceTypeID int
	ScoreType    string
	Value        int
}

// UserRoundSpellFromData returns a UserRoundSpell from data.
func UserRoundSpellFromData(data map[string]interface{}) UserRoundSpell {
	dataObject := data["dataObject"].(map[string]interface{})

	return UserRoundSpell{
		Type:         data["type"].(string),
		Cursor:       int(data["cursor"].(float64)),
		Time:         int(dataObject["time"].(float64)),
		AccountID:    dataObject["accountId"].(string),
		MatchID:      dataObject["matchId"].(string),
		Round:        int(dataObject["round"].(float64)),
		Character:    int(dataObject["character"].(float64)),
		TypeID:       int(dataObject["typeId"].(float64)),
		SourceTypeID: int(dataObject["sourceTypeId"].(float64)),
		ScoreType:    dataObject["scoreType"].(string),
		Value:        int(dataObject["value"].(float64)),
	}
}

// DeathEvent is a telemetry event containing information about a characters death.
type DeathEvent struct {
	Type            string
	Cursor          int
	Time            int
	MatchID         string
	ExternalMatchID string
	UserID          string
}

// DeathEventFromData returns a single DeathEvent from data.
func DeathEventFromData(data map[string]interface{}) DeathEvent {
	dataObject := data["dataObject"].(map[string]interface{})

	return DeathEvent{
		Type:            data["type"].(string),
		Cursor:          int(data["cursor"].(float64)),
		Time:            int(dataObject["time"].(float64)),
		MatchID:         dataObject["matchID"].(string),
		ExternalMatchID: dataObject["externalMatchID"].(string),
		UserID:          dataObject["userID"].(string),
	}
}

// MatchReservedUser is a telemetry event containing information about a match user.
type MatchReservedUser struct {
	Type                string
	Cursor              int
	Time                int
	AccountID           string
	MatchID             string
	ServerType          string
	CharacterLevel      int
	TeamID              string
	TotalTimePlayed     int
	CharacterTimePlayed int
	Character           int
	Team                int
	RankingType         string
	Mount               int
	Attachment          int
	Outfit              int
	Emote               int
	League              int
	Division            int
	DivisionRating      int
	SeasonID            int
}

// MatchReservedUserFromData returns a single MatchReservedUser from data.
func MatchReservedUserFromData(data map[string]interface{}) MatchReservedUser {
	dataObject := data["dataObject"].(map[string]interface{})

	return MatchReservedUser{
		Type:                data["type"].(string),
		Cursor:              int(data["cursor"].(float64)),
		Time:                int(dataObject["time"].(float64)),
		AccountID:           dataObject["accountId"].(string),
		MatchID:             dataObject["matchId"].(string),
		ServerType:          dataObject["serverType"].(string),
		CharacterLevel:      int(dataObject["characterLevel"].(float64)),
		TeamID:              dataObject["teamId"].(string),
		TotalTimePlayed:     int(dataObject["totalTimePlayed"].(float64)),
		CharacterTimePlayed: int(dataObject["characterTimePlayed"].(float64)),
		Character:           int(dataObject["character"].(float64)),
		Team:                int(dataObject["team"].(float64)),
		RankingType:         dataObject["rankingType"].(string),
		Mount:               int(dataObject["mount"].(float64)),
		Attachment:          int(dataObject["attachment"].(float64)),
		Outfit:              int(dataObject["outfit"].(float64)),
		Emote:               int(dataObject["emote"].(float64)),
		League:              int(dataObject["league"].(float64)),
		Division:            int(dataObject["division"].(float64)),
		DivisionRating:      int(dataObject["divisionRating"].(float64)),
		SeasonID:            int(dataObject["seasonId"].(float64)),
	}
}

// QueueEvent is a telemetry event containing information about a user's queue.
type QueueEvent struct {
	Type                  string
	Cursor                int
	Time                  int
	UserID                string
	TeamID                string
	SessionID             string
	Season                int
	EventType             string
	TimeJoinedQueue       string
	TimeInQueue           float64
	Character             int
	CharacterArchetype    int
	QueueTypes            []interface{}
	LimitMatchmakingRange bool
	RegionSamples         []RegionSample
	PreferedRegion        string
	RankingType           string
	League                int
	Division              int
	DivisionRating        int
	TeamSize              int
	TeamMembers           interface{}
	PlacementGamesLeft    int
	MatchID               string
	MatchRegion           string
	TeamSide              int
	AutoMatchmaking       bool
}

// QueueEventFromData returns a single QueueEvent from data.
func QueueEventFromData(data map[string]interface{}) QueueEvent {
	dataObject := data["dataObject"].(map[string]interface{})

	return QueueEvent{
		Type:                  data["type"].(string),
		Cursor:                int(data["cursor"].(float64)),
		Time:                  int(dataObject["time"].(float64)),
		UserID:                dataObject["userId"].(string),
		TeamID:                dataObject["teamId"].(string),
		SessionID:             dataObject["sessionId"].(string),
		Season:                int(dataObject["season"].(float64)),
		EventType:             dataObject["eventType"].(string),
		TimeJoinedQueue:       dataObject["timeJoinedQueue"].(string),
		TimeInQueue:           dataObject["timeInQueue"].(float64),
		Character:             int(dataObject["character"].(float64)),
		CharacterArchetype:    int(dataObject["characterArchetype"].(float64)),
		QueueTypes:            dataObject["queueTypes"].([]interface{}),
		LimitMatchmakingRange: dataObject["limitMatchmakingRange"].(bool),
		RegionSamples:         MultiRegionSamplesFromData(dataObject["regionSamples"].([]interface{})),
		PreferedRegion:        dataObject["preferredRegion"].(string),
		RankingType:           dataObject["rankingType"].(string),
		League:                int(dataObject["league"].(float64)),
		Division:              int(dataObject["division"].(float64)),
		DivisionRating:        int(dataObject["divisionRating"].(float64)),
		TeamSize:              int(dataObject["teamSize"].(float64)),
		TeamMembers:           dataObject["teamMembers"],
		PlacementGamesLeft:    int(dataObject["placementGamesLeft"].(float64)),
		MatchID:               dataObject["matchId"].(string),
		MatchRegion:           dataObject["matchRegion"].(string),
		TeamSide:              int(dataObject["teamSide"].(float64)),
		AutoMatchmaking:       dataObject["autoMatchmaking"].(bool),
	}
}

// RegionSample contains information about a user's region during a QueueEvent.
// See QueueEvent for more information.
type RegionSample struct {
	Region    string
	LatencyMS int
}

// SingleRegionSampleFromData returns a RegionSample from data.
func SingleRegionSampleFromData(data map[string]interface{}) RegionSample {
	return RegionSample{
		Region:    data["region"].(string),
		LatencyMS: int(data["latencyMS"].(float64)),
	}
}

// MultiRegionSamplesFromData returns a slice of RegionSamples from data.
func MultiRegionSamplesFromData(data []interface{}) []RegionSample {
	regionSamples := []RegionSample{}

	for _, rs := range data {
		regionSamples = append(regionSamples, SingleRegionSampleFromData(rs.(map[string]interface{})))
	}

	return regionSamples
}

// TeamUpdateEvent is a telemetry event containing information about a team data update.
type TeamUpdateEvent struct {
	Type                   string
	Cursor                 int
	Time                   int
	Season                 int
	TeamID                 string
	MatchID                string
	ExternalMatchID        string
	UserIDs                []int
	Mode                   string
	League                 int
	PrevLeague             int
	PrevDivision           int
	Division               int
	PrevDivisionRating     int
	DivisionRating         int
	PrevWins               int
	Wins                   int
	PrevLosses             int
	Losses                 int
	RankingChangeType      string
	PrevPlacementGamesLeft int
	PlacementGamesLeft     int
	MatchRegion            string
}

// TeamUpdateEventFromData returns a TeamUpdateEvent from data.
func TeamUpdateEventFromData(data map[string]interface{}) TeamUpdateEvent {
	dataObject := data["dataObject"].(map[string]interface{})

	userIDs := []int{}
	for _, id := range dataObject["userIDs"].([]interface{}) {
		userIDs = append(userIDs, int(id.(float64)))
	}

	return TeamUpdateEvent{
		Type:                   data["type"].(string),
		Cursor:                 int(data["cursor"].(float64)),
		Time:                   int(dataObject["time"].(float64)),
		Season:                 int(dataObject["season"].(float64)),
		TeamID:                 dataObject["teamID"].(string),
		MatchID:                dataObject["matchID"].(string),
		ExternalMatchID:        dataObject["externalMatchID"].(string),
		UserIDs:                userIDs,
		Mode:                   dataObject["mode"].(string),
		League:                 int(dataObject["league"].(float64)),
		PrevLeague:             int(dataObject["prevLeague"].(float64)),
		PrevDivision:           int(dataObject["prevDivision"].(float64)),
		Division:               int(dataObject["division"].(float64)),
		PrevDivisionRating:     int(dataObject["prevDivisionRating"].(float64)),
		DivisionRating:         int(dataObject["divisionRating"].(float64)),
		PrevWins:               int(dataObject["prevWins"].(float64)),
		Wins:                   int(dataObject["wins"].(float64)),
		PrevLosses:             int(dataObject["prevLosses"].(float64)),
		Losses:                 int(dataObject["losses"].(float64)),
		RankingChangeType:      dataObject["rankingChangeType"].(string),
		PrevPlacementGamesLeft: int(dataObject["prevPlacementGamesLeft"].(float64)),
		PlacementGamesLeft:     int(dataObject["placementGamesLeft"].(float64)),
		MatchRegion:            dataObject["matchRegion"].(string),
	}
}

// ServerShutdown is a telemetry event containing information about a battlerite server's closing.
type ServerShutdown struct {
	Type            string
	Cursor          int
	Time            int
	MatchID         string
	ExternalMatchID string
	MatchTime       int
	Reason          string
}

// ServerShutdownFromData returns a ServerShutdown from data.
func ServerShutdownFromData(data map[string]interface{}) ServerShutdown {
	dataObject := data["dataObject"].(map[string]interface{})

	return ServerShutdown{
		Type:            data["type"].(string),
		Cursor:          int(data["cursor"].(float64)),
		Time:            int(dataObject["time"].(float64)),
		MatchID:         dataObject["matchID"].(string),
		ExternalMatchID: dataObject["externalMatchID"].(string),
		MatchTime:       int(dataObject["matchTime"].(float64)),
		Reason:          dataObject["reason"].(string),
	}
}

// RoundFinishedEvent is a telemetry event containing information about the end of a round.
type RoundFinishedEvent struct {
	Type            string
	Cursor          int
	Time            int
	MatchID         string
	ExternalMatchID string
	Round           int
	RoundLength     int
	WinningTeam     int
	PlayerStats     []PlayerStats
}

// RoundFinishedEventFromData returns a RoundFinishedEvent from data.
func RoundFinishedEventFromData(data map[string]interface{}) RoundFinishedEvent {
	dataObject := data["dataObject"].(map[string]interface{})

	return RoundFinishedEvent{
		Type:            data["type"].(string),
		Cursor:          int(data["cursor"].(float64)),
		Time:            int(dataObject["time"].(float64)),
		MatchID:         dataObject["matchID"].(string),
		ExternalMatchID: dataObject["externalMatchID"].(string),
		Round:           int(dataObject["round"].(float64)),
		RoundLength:     int(dataObject["roundLength"].(float64)),
		WinningTeam:     int(dataObject["winningTeam"].(float64)),
		PlayerStats:     MultiPlayerStatsFromData(dataObject["playerStats"].([]interface{})),
	}
}

// PlayerStats contains information about a players stats at the end of a round as part of a RoundFinishedEvent.
type PlayerStats struct {
	UserID           string
	Kills            int
	Deaths           int
	Score            int
	DamageDone       int
	DamageReceived   int
	HealingDone      int
	HealingReceived  int
	DisablesDone     int
	DisablesReceived int
	EnergyGained     int
	EnergyUsed       int
	TimeAlive        int
	AbilityUses      int
}

// SinglePlayerStatsFromData returns a single PlayerStats from data.
func SinglePlayerStatsFromData(data map[string]interface{}) PlayerStats {
	return PlayerStats{
		UserID:           data["userID"].(string),
		Kills:            int(data["kills"].(float64)),
		Deaths:           int(data["deaths"].(float64)),
		Score:            int(data["score"].(float64)),
		DamageDone:       int(data["damageDone"].(float64)),
		DamageReceived:   int(data["damageReceived"].(float64)),
		HealingDone:      int(data["healingDone"].(float64)),
		HealingReceived:  int(data["healingReceived"].(float64)),
		DisablesDone:     int(data["disablesDone"].(float64)),
		DisablesReceived: int(data["disablesReceived"].(float64)),
		EnergyGained:     int(data["energyGained"].(float64)),
		EnergyUsed:       int(data["energyUsed"].(float64)),
		TimeAlive:        int(data["timeAlive"].(float64)),
		AbilityUses:      int(data["abilityUses"].(float64)),
	}
}

// MultiPlayerStatsFromData returns a slice of PlayerStats from data.
func MultiPlayerStatsFromData(data []interface{}) []PlayerStats {
	playerStats := []PlayerStats{}

	for _, pData := range data {
		stat := SinglePlayerStatsFromData(pData.(map[string]interface{}))
		playerStats = append(playerStats, stat)
	}
	return playerStats
}

// MatchFinishedEvent is a telemetry event containing information about the end of a match.
type MatchFinishedEvent struct {
	Type            string
	Cursor          int
	Time            int
	TeamOneScore    int
	TeamTwoScore    int
	MatchLength     int
	MatchID         string
	ExternalMatchID string
	Leavers         interface{}
	Region          string
}

// MatchFinishedEventFromData returns a single MatchFinishedEvent from data.
func MatchFinishedEventFromData(data map[string]interface{}) MatchFinishedEvent {
	dataObject := data["dataObject"].(map[string]interface{})

	return MatchFinishedEvent{
		Type:            data["type"].(string),
		Cursor:          int(data["cursor"].(float64)),
		Time:            int(dataObject["time"].(float64)),
		TeamOneScore:    int(dataObject["teamOneScore"].(float64)),
		TeamTwoScore:    int(dataObject["teamTwoScore"].(float64)),
		MatchLength:     int(dataObject["matchLength"].(float64)),
		MatchID:         dataObject["matchID"].(string),
		ExternalMatchID: dataObject["externalMatchID"].(string),
		Leavers:         dataObject["leavers"],
		Region:          dataObject["region"].(string),
	}
}
