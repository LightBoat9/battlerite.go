package battleritego

import "strconv"

// Participant contains information about a participant in a match.
// See match.go for more information.
// See https://battlerite-docs.readthedocs.io/en/master/matches/matches.html#participants
type Participant struct {
	Type             string
	ID               string
	Actor            int
	ShardID          string
	DamageDone       int
	DamageReceived   int
	Deaths           int
	EnergyGained     int
	EnergyUsed       int
	Kills            int
	Score            int
	TimeAlive        int
	UserID           int
	AbilityUses      int
	DisablesDone     int
	DisablesReceived int
	Emote            int
	Mount            int
	Outfit           int
	Attachment       int
	HealingDone      int
	HealingReceived  int
	Side             int
	Relationships    interface{}
}

// SingleParticipantFromData returns a single participant from data.
func SingleParticipantFromData(data map[string]interface{}) Participant {
	attributes := data["attributes"].(map[string]interface{})
	stats := attributes["stats"].(map[string]interface{})
	relationships := data["relationships"].(map[string]interface{})

	actor, _ := strconv.Atoi(attributes["actor"].(string))
	userID, _ := strconv.Atoi(stats["userID"].(string))

	return Participant{
		Type:             data["type"].(string),
		ID:               data["id"].(string),
		Actor:            actor,
		ShardID:          attributes["shardId"].(string),
		UserID:           userID,
		DamageDone:       int(stats["damageDone"].(float64)),
		DamageReceived:   int(stats["damageReceived"].(float64)),
		Deaths:           int(stats["deaths"].(float64)),
		EnergyGained:     int(stats["energyGained"].(float64)),
		EnergyUsed:       int(stats["energyUsed"].(float64)),
		Kills:            int(stats["kills"].(float64)),
		Score:            int(stats["score"].(float64)),
		TimeAlive:        int(stats["timeAlive"].(float64)),
		AbilityUses:      int(stats["abilityUses"].(float64)),
		DisablesDone:     int(stats["disablesDone"].(float64)),
		DisablesReceived: int(stats["disablesReceived"].(float64)),
		Emote:            int(stats["emote"].(float64)),
		Mount:            int(stats["mount"].(float64)),
		Outfit:           int(stats["outfit"].(float64)),
		Attachment:       int(stats["attachment"].(float64)),
		HealingDone:      int(stats["healingDone"].(float64)),
		HealingReceived:  int(stats["healingReceived"].(float64)),
		Side:             int(stats["side"].(float64)),
		Relationships:    relationships,
	}
}
