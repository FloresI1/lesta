package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/FloresI1/lesta/struct"
)

var playersMap = make(map[string]structur.Player)
var matchGroups []structur.MatchGroup
var mu sync.Mutex
var nextGroupID = 1
var groupSize = 3

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var player structur.Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	mu.Lock()
	defer mu.Unlock()

	if _, exists := playersMap[player.Name]; exists {
		http.Error(w, "Player with this name already exists", http.StatusConflict)
		return
	}

	player.JoinTime = time.Now()
	playersMap[player.Name] = player

	if len(playersMap) >= groupSize {
		CreateMatchGroup()
	}

	fmt.Fprintf(w, "Player %s added.\n", player.Name)
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(playersMap)
}

func CreateMatchGroup() {
	group := structur.MatchGroup{
		ID:           nextGroupID,
		PlayerNames:  make([]string, 0, groupSize),
		MinSkill:     0,
		MaxSkill:     0,
		AvgSkill:     0,
		MinLatency:   0,
		MaxLatency:   0,
		AvgLatency:   0,
		MinQueueTime: 0,
		MaxQueueTime: 0,
		AvgQueueTime: 0,
	}

	for name, player := range playersMap {
		group.PlayerNames = append(group.PlayerNames, name)

		if group.MinSkill == 0 || player.Skill < group.MinSkill {
			group.MinSkill = player.Skill
		}
		if player.Skill > group.MaxSkill {
			group.MaxSkill = player.Skill
		}
		group.AvgSkill += player.Skill

		if group.MinLatency == 0 || player.Latency < group.MinLatency {
			group.MinLatency = player.Latency
		}
		if player.Latency > group.MaxLatency {
			group.MaxLatency = player.Latency
		}
		group.AvgLatency += player.Latency

		queueTime := time.Since(player.JoinTime)
		if group.MinQueueTime == 0 || queueTime < group.MinQueueTime {
			group.MinQueueTime = queueTime
		}
		if queueTime > group.MaxQueueTime {
			group.MaxQueueTime = queueTime
		}
		group.AvgQueueTime += queueTime

		delete(playersMap, name)

		if len(group.PlayerNames) >= groupSize {
			break
		}
	}

	group.AvgSkill /= float64(len(group.PlayerNames))
	group.AvgLatency /= float64(len(group.PlayerNames))
	group.AvgQueueTime /= time.Duration(len(group.PlayerNames))

	matchGroups = append(matchGroups, group)
	nextGroupID++
}