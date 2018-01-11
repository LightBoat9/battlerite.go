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

- **GetStatus()(Status, error)**

Return a Status instance, and an error if one occurs.

#### **Status struct**

Contains information about the current status of the gamelocker API.

- **Fields**
```go
Type    string  
ID      string  
Release string  
Version string  
```

**Example Use**
```go
status, err := client.GetStatus()
if err != nil {
  fmt.Println("Error:", err)
}

fmt.Printf("Current Version %s was released on %s", status.Version, status.Release)
```

### **Player**

See [Gamelocker Documentation](https://battlerite-docs.readthedocs.io/en/master/players/players.html)

- **GetPlayer(id string) (Player, error)**
  - id string - The user ID of the player
  
Returns a Player by their ID, and an error if one occurs.
  
- **GetPlayersFiltered(filter PlayerFilter) ([]Player, error)**
  - filter PlayerFilter - The filter to search for players see PlayerFilter struct below
  
Returns a slice of Players based on the filter, and an error if one occurs.

#### **PlayerFilter struct**
Contains filters for searching players with GetPlayersFiltered. 
Note that only one filter parameter should be used at a time.

- **Fields**
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
"Raigon", "Blossum", "Thorn", "Alysia"
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
