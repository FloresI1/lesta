package util

import (
	"time"
	"github.com/FloresI1/lesta/struct"
)

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

func (stats *GroupStats) Update(player structur.Player) {
	skill := player.Skill
	latency := player.Latency
	waitTime := time.Since(player.JoinTime)

	if stats.PlayerCount == 0 {
		stats.MinSkill = skill
		stats.MaxSkill = skill
		stats.MinLatency = latency
		stats.MaxLatency = latency
		stats.MinWaitTime = waitTime
		stats.MaxWaitTime = waitTime
	} else {
		if skill < stats.MinSkill {
			stats.MinSkill = skill
		}
		if skill > stats.MaxSkill {
			stats.MaxSkill = skill
		}
		if latency < stats.MinLatency {
			stats.MinLatency = latency
		}
		if latency > stats.MaxLatency {
			stats.MaxLatency = latency
		}
		if waitTime < stats.MinWaitTime {
			stats.MinWaitTime = waitTime
		}
		if waitTime > stats.MaxWaitTime {
			stats.MaxWaitTime = waitTime
		}
	}

	stats.TotalSkill += skill
	stats.TotalLatency += latency
	stats.TotalWaitTime += waitTime
	stats.PlayerCount++
}

func CalculateAverages(stats GroupStats) (avgSkill, avgLatency float64, avgWaitTime time.Duration) {
	if stats.PlayerCount > 0 {
		avgSkill = stats.TotalSkill / float64(stats.PlayerCount)
		avgLatency = stats.TotalLatency / float64(stats.PlayerCount)
		avgWaitTime = stats.TotalWaitTime / time.Duration(stats.PlayerCount)
	}
	return
}
