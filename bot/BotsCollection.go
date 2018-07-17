package bot

import "sync"

type BotsCollection struct {
	sync.RWMutex
	bots []*Bot
}

func NewBotsCollection() *BotsCollection {
	return &BotsCollection{bots: make([]*Bot, 0)}
}

func (bc *BotsCollection) Bots() []*Bot {
	bc.Lock()
	existingBots := make([]*Bot, len(bc.bots))
	copy(existingBots, bc.bots)
	bc.Unlock()
	return existingBots
}

func (bc *BotsCollection) AddBot(bot *Bot) {
	bc.Lock()
	bc.bots = append(bc.bots, bot)
	bc.Unlock()
}

func (bc *BotsCollection) BotsCount() int {
	bc.RLock()
	defer bc.RUnlock()
	return len(bc.bots)
}
