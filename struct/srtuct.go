package structur

import "time"

type Player struct {
	Name     string `json:"name"`    // Имя игрока
	Skill    float64 `json:"skill"`   // Уровень навыка игрока
	Latency  float64 `json:"latency"`   // Задержка игрока в миллисекундах
	JoinTime time.Time // Время добавления игрока в очередь
}
type MatchGroup struct {
	ID           int           // Уникальный идентификатор группы
	PlayerNames  []string      // Список имен игроков в группе
	MinSkill     float64       // Минимальный уровень навыка в группе
	MaxSkill     float64       // Максимальный уровень навыка в группе
	AvgSkill     float64       // Средний уровень навыка в группе
	MinLatency   float64       // Минимальная задержка в группе
	MaxLatency   float64       // Максимальная задержка в группе
	AvgLatency   float64       // Средняя задержка в группе
	MinQueueTime time.Duration // Минимальное время в очереди
	MaxQueueTime time.Duration // Максимальное время в очереди
	AvgQueueTime time.Duration // Среднее время в очереди
}
