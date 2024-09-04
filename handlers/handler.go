package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/FloresI1/lesta/struct"
	"github.com/FloresI1/lesta/util"
)

var (
	playersMap  = make(map[string]structur.Player) // Мапа для хранения игроков
	matchGroups []structur.MatchGroup              // Хранилище групп
	mu          sync.Mutex                         // Мьютекс для синхронизации доступа
	nextGroupID = 0                                // Счетчик ID для групп
	groupSize   = 3                                // Размер группы (количество игроков)
)

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

	if player.Name == "" {
		http.Error(w, "Player name cannot be empty", http.StatusBadRequest)
		return
	}

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

func CreateMatchGroup() {
	var matchGroup structur.MatchGroup
	matchGroup.ID = nextGroupID
	nextGroupID++

	var playersInGroup []structur.Player
	var stats util.GroupStats

	for name := range playersMap {
		player := playersMap[name]
		playersInGroup = append(playersInGroup, player)
		stats.Update(player)

		if len(playersInGroup) == groupSize {
			break
		}
	}

	for _, player := range playersInGroup {
		delete(playersMap, player.Name)
	}

	for _, player := range playersInGroup {
		matchGroup.PlayerNames = append(matchGroup.PlayerNames, player.Name)
	}

	avgSkill, avgLatency, avgWaitTime := util.CalculateAverages(stats)

	fmt.Printf("Created Match Group %d with players: %v\n", matchGroup.ID, matchGroup.PlayerNames)
	fmt.Printf("Skill - Min: %.2f, Max: %.2f, Avg: %.2f\n", stats.MinSkill, stats.MaxSkill, avgSkill)
	fmt.Printf("Latency - Min: %.2f, Max: %.2f, Avg: %.2f\n", stats.MinLatency, stats.MaxLatency, avgLatency)
	fmt.Printf("Wait Time - Min: %v, Max: %v, Avg: %v\n", stats.MinWaitTime, stats.MaxWaitTime, avgWaitTime)

	matchGroups = append(matchGroups, matchGroup)
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	json.NewEncoder(w).Encode(matchGroups)
}
