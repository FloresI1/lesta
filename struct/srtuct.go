package structur

import "time"

type Player struct {
	Name     string `json:"name"`    
	Skill    float64 `json:"skill"`  
	Latency  float64 `json:"latency"`      
	JoinTime time.Time 
}
type MatchGroup struct {
	ID           int           
	PlayerNames  []string   
}
type GroupStats struct {
	MinSkill       float64
	MaxSkill       float64
	TotalSkill     float64
	MinLatency     float64
	MaxLatency     float64
	TotalLatency   float64
	MinWaitTime     time.Duration
	MaxWaitTime     time.Duration
	TotalWaitTime   time.Duration
	PlayerCount     int
}