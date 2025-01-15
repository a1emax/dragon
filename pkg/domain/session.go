package domain

import (
	"math/rand"

	"github.com/a1emax/youngine/clock"

	"dragon/pkg/domain/space"
)

type Session struct {
	config        SessionConfig
	tileMap       *TileMap
	treasureCount int
	knightSet     *KnightSet
	dragon        *Dragon
}

type SessionConfig struct {
	Random        *rand.Rand
	Clock         clock.Clock
	OnKnightEnter func(event KnightEnterEvent)
	OnKnightLoad  func(event KnightLoadEvent)
	OnKnightKill  func(event KnightKillEvent)
	OnKnightExit  func(event KnightExitEvent)
	OnDragonChop  func(event DragonChopEvent)
	OnGameOver    func(event GameOverEvent)
}

func NewSession(config SessionConfig) *Session {
	tileMap, doors, treasureCount := newLevel()

	return &Session{
		config:        config,
		tileMap:       tileMap,
		treasureCount: treasureCount,
		knightSet: newKnightSet(knightSetConfig{
			random:        config.Random,
			clock:         config.Clock,
			tileMap:       tileMap,
			doors:         doors,
			onKnightEnter: config.OnKnightEnter,
		}),
		dragon: newDragon(dragonConfig{
			clock:   config.Clock,
			tileMap: tileMap,
		}),
	}
}

func (s *Session) TileMap() *TileMap {
	return s.tileMap
}

func (s *Session) KnightSet() *KnightSet {
	return s.knightSet
}

func (s *Session) Dragon() *Dragon {
	return s.dragon
}

func (s *Session) Update() {
	s.knightSet.update()
	s.dragon.update()

	s.checkKnights()
	s.checkDragon()
	s.checkGameOver()
}

func (s *Session) checkKnights() {
	for knight := range s.knightSet.All() {
		tile := s.tileMap.TileAt(knight.Position())

		switch tile.ObjectTag() {
		case Door:
			if knight.IsLoaded() {
				knight.dispose()
				s.treasureCount--
				s.config.OnKnightExit(KnightExitEvent{})
			}
		case Treasure:
			if !knight.IsLoaded() {
				knight.Load()
				tile.unsetObject()
				s.config.OnKnightLoad(KnightLoadEvent{})
			}
		}

		if tile.DragonPartTag() == DragonTail {
			s.dragon.chop(tile.DragonPartID())
			s.config.OnDragonChop(DragonChopEvent{})
		}
	}
}

func (s *Session) checkDragon() {
	tile := s.tileMap.TileAt(s.dragon.Head())

	if tile.KnightID() != 0 {
		knight := s.knightSet.Get(tile.KnightID())
		if knight.IsLoaded() {
			if tile.ObjectTag() == 0 {
				tile.setObject(Treasure)
			} else {
				s.treasureCount--
				// TODO: Find tile to drop treasure instead?
			}
		}
		knight.dispose()
		s.dragon.grow()
		s.config.OnKnightKill(KnightKillEvent{})
	}
}

func (s *Session) checkGameOver() {
	if s.treasureCount == 0 {
		s.config.OnGameOver(GameOverEvent{})

		return
	}

	var dragonCanMove bool
	for d := space.North; d <= space.West; d++ {
		if s.tileMap.TileAt(space.AddDirection(s.dragon.Head(), d)).IsPassableForDragon() {
			dragonCanMove = true

			break
		}
	}
	if !dragonCanMove {
		s.config.OnGameOver(GameOverEvent{})

		return
	}
}
