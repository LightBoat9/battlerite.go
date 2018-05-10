# battlerite.go
Golang wrapper around the official Battlerite API

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/LightBoat9/battlerite.go/blob/master/LICENSE) file for details

## Installing

Run this in console to install

```go 
go get github.com/LightBoat9/battlerite.go/battlerite
```

# Usage

## Import

```go
import (
  "github.com/LightBoat9/battlerite.go/battlerite"
)
```

## Client

Gather data from the API using the Client type. Instance a Client type using an API key.
Register an app and get an API key [here](https://developer.battlerite.com/).

```go
client := battlerite.Client{ APIKey }
```

## Reference

### **Status**

See [Gamelocker Documentation](https://battlerite-docs.readthedocs.io/en/master/status/status.html)

- GetStatus()(Status, error)

Return a Status instance, and an error if one occurs.

#### **Status struct**

Contains information about the current status of the gamelocker API.

- Fields
```go
Type    string  
ID      string  
Release string  
Version string  
```

- Example Use
```go
status, err := client.GetStatus()
if err != nil {
  fmt.Println("Error:", err)
}

fmt.Printf("Current Version %s was released on %s", status.Version, status.Release)
```

### **Player**

See [Gamelocker Documentation](https://battlerite-docs.readthedocs.io/en/master/players/players.html)

- GetPlayer(id string) (Player, error)
  - id string - The user ID of the player
  
Returns a Player by their ID, and an error if one occurs.
  
- GetPlayersFiltered(filter PlayerFilter) ([]Player, error)
  - filter PlayerFilter - The filter to search for players see PlayerFilter struct below
  
Returns a slice of Players based on the filter, and an error if one occurs.

#### **PlayerFilter struct**
Contains filters for searching players with GetPlayersFiltered. 
Note that only one filter parameter should be used at a time.

- Fields
```go
  Names    []string
  UserIDs  []int
  SteamIDs []int
```

#### **Player struct**
Contains information about a single battlerite user.

- **Fields**
```go
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
  AccountXP                    int
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
```

Every "Character" map[string]int contains the following keys representing the battlerite champions.
```go
"Lucie", "Sirius", "Iva", "Jade", "RuhKaan", 
"Oldur", "Ashka", "Varesh", "Pearl", "Taya", 
"Poloma", "Croak", "Freya", "Jumong", "Shifu", 
"Ezmo", "Bakko", "Rook", "Pestilus", "Destiny", 
<<<<<<< HEAD
"Raigon", "Blossum", "Thorn", "Alysia", "Jamila", 
"Ulric"
=======
"Raigon", "Blossum", "Thorn", "Alysia", "Zander"
>>>>>>> aab3b0dd2d087cb6d2da6105b8a16f6fbe7513f2
```

**Example Use**
```go
// Get one player by ID
player, err := client.GetPlayer(776450744541908992)
if err != nil {
  fmt.Println("Error:", err)
}

fmt.Printf("\nPlayer %s has %d total wins! With Taya they have %d wins!", 
  player.Name, player.Wins, player.CharacterWins["Taya"])
  
// Get multiple players by Name
filter := battlerite.PlayerFilter{
  Names: []string{"Averse", "ProsteR18", "Aldys"},
}

players, err := client.GetPlayersFiltered(filter)
if err != nil {
  fmt.Println("Error:", err)
}

playerNames := []string{}
for _, plr := range players {
  playerNames = append(playerNames, plr.Name)
}

fmt.Printf("\n%d players found! Their names are %s", len(players), strings.Join(playerNames, ", "))
```

### **Teams**

See [Gamelocker Documentation](https://battlerite-docs.readthedocs.io/en/master/teams/teams.html)

- GetTeamsFiltered(filter TeamFilter) ([]Team, error)
  - filter TeamFilter - the filter to search for teams see TeamFilter below.
  
 Returns a slice of Teams based on the filter, and an error if one occurs.
 
#### **TeamFilter struct**
Contains filters for searching tean with GetTeamsFiltered.
Note that both filter parameters are required.

- Fields
```go
	Season    int
	PlayerIDs []int
```

#### **Team struct**
 
Contains information about a battlerite team.
 
 - Fields
```go
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
```
  
- Example Use
```go
teamFilter := battlerite.TeamFilter{
  Season:    6,
  PlayerIDs: []int{5983},
}

teams, err := client.GetTeamsFiltered(teamFilter)
if err != nil {
  fmt.Println("Error:", err)
}

fmt.Printf("%+v\n", teams)
```

### **Match**

See [Gamelocker Documentation](https://battlerite-docs.readthedocs.io/en/master/matches/matches.html)

- GetMatch(id string) (Match, error)
  id string - id of the match to return
  
Returns a Match and an error if one occurs
  
- GetMatchesFiltered(filter MatchFilter) ([]Match, error)
  filter MatchFilter - the filter to search for matches
  
Returns a slice of Matches and an error if one occurs

- MatchFilter struct

Contains filters for searching for Matches using GetMatchesFiltered.

```go
PageOffset     int
PageLimit      int
Sort           string
CreatedAtStart string
CreatedAtEnd   string
PlayerIDs      []string
PatchVersion   []string
```

- Match struct

Contains information about a single battlerite Match.

```go
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
```
  
- Example Use
```go
// Get a single match by ID
match, err := client.GetMatch("3A7A00AF2CFF4356B94951CD4A59610B")
if err != nil {
  log.Fatal(err)
}

fmt.Printf("\nThe match %s was created on %s", match.ID, match.CreatedAt)

// Get a slice of Matches by MatchFilter
matchFilter := battlerite.MatchFilter{
  PageLimit: 2,
}

matches, err := client.GetMatchesFiltered(matchFilter)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("Found %d matches!", len(matches))
```

## **Telementry Data**

Each match contains events that are stored as telementry data.

- GetTelementry(URL string) (Telementry, error)
  - URL string - The URL of the data, can be found from a match under <match>.Assets.URL

- Example Use

Once you have a match you can search for telementry data using the matches URL in Assets.URL

```go
match, err := client.GetMatch("3A7A00AF2CFF4356B94951CD4A59610B")
if err != nil {
  log.Fatal(err)
}
  
telemetry, err := client.GetTelemetry(match.Asset.URL)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("%+v\n", telemetry) // Prints data with keys
```

### **Telementry**

Stores all of the events of a given match. More information on each event type below.

- Fields
```go
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
```

### **MatchStart**

A telemetry event containing information at a matches start.

- Fields
```go
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
```
### **RoundEvent**

A telemetry event containing information about various events during a round.

- Fields
```go
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
```

### **UserRoundSpell**

A telemetry event containing information about a characters ability use.

- Fields
```go
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
```

### **DeathEvent**

A telemetry event containing information about a characters death.

- Fields
```go
Type            string
Cursor          int
Time            int
MatchID         string
ExternalMatchID string
UserID          string
```

### **MatchReservedUser**

A telemetry event containing information about a match user.

- Fields
```go
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
```

### **QueueEvent**

A telemetry event containing information about a user's queue.
It contains RegionSample information, see below for more details.

- Fields
```go
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
```

### **RegionSample**

Contains information about a user's region during a QueueEvent.

- Fields
```go
Region    string
LatencyMS int
```

### **TeamUpdateEvent**

A telemetry event containing information about a team data update.

- Fields
```go
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
```

### **ServerShutdown**

A telemetry event containing information about a battlerite server's closing.

- Fields
```go
Type            string
Cursor          int
Time            int
MatchID         string
ExternalMatchID string
MatchTime       int
Reason          string
```

### **RoundFinishedEvent**

A telemetry event containing information about the end of a round.
Contains PlayerStats information, more details below.

- Fields
```go
Type            string
Cursor          int
Time            int
MatchID         string
ExternalMatchID string
Round           int
RoundLength     int
WinningTeam     int
PlayerStats     []PlayerStats
```

### **PlayerStats**

Contains information about a players stats at the end of a round as part of a RoundFinishedEvent.

- Fields
```go
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
```

### **MatchFinishedEvent**

A telemetry event containing information about the end of a match.

- Fields
```go
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
```
