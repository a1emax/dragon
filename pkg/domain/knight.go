package domain

import (
	"iter"
	"math/rand"

	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/clock/driver/ebitenclock"

	"dragon/pkg/domain/space"
)

const (
	knightSpawnInterval = 10 * ebitenclock.Second
	knightMoveInterval  = ebitenclock.Second / 3
	knightMaxStopCount  = 3
)

type KnightSet struct {
	config     knightSetConfig
	knights    []Knight
	len        int
	spawnSince clock.Time
}

type knightSetConfig struct {
	random        *rand.Rand
	clock         clock.Clock
	tileMap       *TileMap
	doors         []space.TileVec
	onKnightEnter func(event KnightEnterEvent)
}

func newKnightSet(config knightSetConfig) *KnightSet {
	return &KnightSet{
		config:     config,
		knights:    nil,
		spawnSince: config.clock.Now(),
	}
}

func (k *KnightSet) Len() int {
	return k.len
}

func (k *KnightSet) Get(id KnightID) *Knight {
	i := k.decodeID(id)
	if i < 0 || i >= len(k.knights) {
		return nil // Panic?
	}

	knight := &k.knights[i]
	if knight.id == 0 {
		return nil // Panic?
	}

	return knight
}

func (k *KnightSet) All() iter.Seq[*Knight] {
	return func(yield func(*Knight) bool) {
		for i := range k.knights {
			if k.knights[i].id == 0 {
				continue
			}

			if !yield(&k.knights[i]) {
				continue
			}
		}
	}
}

func (k *KnightSet) update() {
	if clock.CheckInterval(k.config.clock, k.spawnSince, knightSpawnInterval) {
		k.spawn()
	}

	for i := range k.knights {
		if k.knights[i].id == 0 {
			continue
		}

		k.knights[i].update()
	}
}

func (k *KnightSet) spawn() {
	if k.len-1 == k.decodeID(MaxKnightID) || len(k.config.doors) == 0 {
		return
	}
	door := k.config.doors[k.config.random.Intn(len(k.config.doors))]

	path := k.findPathToTreasure(door, false)
	if len(path) == 0 {
		return
	}

	i := -1
	for j := range k.knights {
		if k.knights[j].id == 0 {
			i = j

			break
		}
	}
	if i == -1 {
		i = len(k.knights)
		k.knights = append(k.knights, Knight{})
	}

	k.knights[i].init(k, k.encodeID(i), door, path)
	k.len++
	k.config.onKnightEnter(KnightEnterEvent{})
}

func (k *KnightSet) findPathToTreasure(from space.TileVec, dragon bool) []space.TileVec {
	return space.FindPath(from, func(t space.TileVec) bool {
		return k.config.tileMap.TileAt(t).IsPassableForKnight(dragon)
	}, func(t space.TileVec) bool {
		tile := k.config.tileMap.TileAt(t)

		return tile != nil && tile.ObjectTag() == Treasure
	})
}

func (k *KnightSet) findPathToDoor(from space.TileVec, dragon bool) []space.TileVec {
	return space.FindPath(from, func(t space.TileVec) bool {
		return k.config.tileMap.TileAt(t).IsPassableForKnight(dragon)
	}, func(t space.TileVec) bool {
		tile := k.config.tileMap.TileAt(t)

		return tile != nil && tile.ObjectTag() == Door
	})
}

func (k *KnightSet) encodeID(i int) KnightID {
	return KnightID(i + 1)
}

func (k *KnightSet) decodeID(id KnightID) int {
	return int(id - 1)
}

type Knight struct {
	set       *KnightSet
	id        KnightID
	position  space.TileVec
	path      []space.TileVec
	moveSince clock.Time
	stopCount int
	isLoaded  bool
}

func (k *Knight) init(set *KnightSet, id KnightID, position space.TileVec, path []space.TileVec) {
	*k = Knight{
		set:       set,
		id:        id,
		position:  position,
		path:      path,
		moveSince: set.config.clock.Now(),
		stopCount: 0,
		isLoaded:  false,
	}
	k.setTiling()
}

func (k *Knight) ID() KnightID {
	return k.id
}

func (k *Knight) Position() space.TileVec {
	return k.position
}

func (k *Knight) IsLoaded() bool {
	return k.isLoaded
}

func (k *Knight) Load() {
	if k.id == 0 {
		return // Panic?
	}

	k.isLoaded = true
	k.path = k.set.findPathToDoor(k.position, false)
}

func (k *Knight) update() {
	if !clock.CheckInterval(k.set.config.clock, k.moveSince, knightMoveInterval) {
		return
	}

	target := len(k.path) - 1
	if target < 0 || !k.set.config.tileMap.TileAt(k.path[target]).IsPassableForKnight(true) {
		k.stopCount++
		if k.stopCount > knightMaxStopCount {
			if k.isLoaded {
				k.path = k.set.findPathToDoor(k.position, true)
			} else {
				k.path = k.set.findPathToTreasure(k.position, true)
			}

			k.stopCount = 0
		}

		return
	}

	k.stopCount = 0

	k.unsetTiling()
	k.position = k.path[target]
	k.path = k.path[:target]
	k.setTiling()
}

func (k *Knight) dispose() {
	if k.id == 0 {
		return // Panic?
	}

	k.unsetTiling()
	k.set.len--
	*k = Knight{}
}

func (k *Knight) setTiling() {
	k.set.config.tileMap.TileAt(k.position).setKnight(k.id)
}

func (k *Knight) unsetTiling() {
	k.set.config.tileMap.TileAt(k.position).unsetKnight()
}
