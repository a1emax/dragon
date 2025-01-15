package domain

import (
	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/clock/driver/ebitenclock"

	"dragon/pkg/domain/space"
)

const dragonMoveInterval = ebitenclock.Second / 6

type Dragon struct {
	config    dragonConfig
	parts     []space.TileVec
	direction space.Direction
	moveSince clock.Time
}

type dragonConfig struct {
	clock   clock.Clock
	tileMap *TileMap
}

func newDragon(config dragonConfig) *Dragon {
	// TODO: Add dragon parts to input parameters.
	d := &Dragon{
		config: config,
		parts: []space.TileVec{
			{6, 1},
			{5, 1},
			{4, 1},
			{3, 1},
			{2, 1},
			{1, 1},
		},
		direction: space.East,
		moveSince: clock.Time{},
	}
	d.setTiling()

	return d
}

func (d *Dragon) Len() int {
	return len(d.parts)
}

func (d *Dragon) Head() space.TileVec {
	return d.parts[0]
}

func (d *Dragon) grow() {
	last := len(d.parts) - 1
	if last == d.decodePartID(MaxDragonPartID) {
		return
	}

	var growDirection space.Direction
	if last > 0 {
		var ok bool
		growDirection, ok = space.GetDirection(d.parts[last-1], d.parts[last])
		if !ok {
			return // Wrong parts.
		}
	} else {
		growDirection = d.direction.Invert()
	}

	for i := 0; i < 4; i++ {
		growTarget := space.AddDirection(d.parts[last], growDirection)
		if d.config.tileMap.TileAt(growTarget).IsPassableForDragon() {
			d.unsetTiling()
			d.parts = append(d.parts, growTarget)
			d.setTiling()

			break
		}

		growDirection = growDirection.TurnCW()
		if last == 0 && growDirection == d.direction {
			i++
			growDirection = growDirection.TurnCW()
		}
	}
}

func (d *Dragon) chop(partID DragonPartID) {
	i := d.decodePartID(partID)
	if i < 0 || i >= len(d.parts) {
		return // Panic?
	}

	d.unsetTiling()
	d.parts = d.parts[:i]
	d.setTiling()
}

func (d *Dragon) SetDirection(value space.Direction) {
	if value < space.North || value > space.West {
		return // Panic?
	}

	d.direction = value
}

func (d *Dragon) Move() {
	if d.moveSince.IsZero() {
		d.moveSince = d.config.clock.Now()
	}
}

func (d *Dragon) Stop() {
	d.moveSince = clock.Time{}
}

func (d *Dragon) update() {
	if !clock.CheckInterval(d.config.clock, d.moveSince, dragonMoveInterval) {
		return
	}

	headTarget := space.AddDirection(d.parts[0], d.direction)
	if !d.config.tileMap.TileAt(headTarget).IsPassableForDragon() {
		return
	}

	d.unsetTiling()
	for i := len(d.parts) - 1; i > 0; i-- {
		d.parts[i] = d.parts[i-1]
	}
	d.parts[0] = headTarget
	d.setTiling()
}

func (d *Dragon) setTiling() {
	bodyLength := max(2, (len(d.parts)-1)/2)

	for i := range d.parts {
		var tag DragonPartTag
		switch {
		case i == 0:
			tag = DragonHead
		case i <= bodyLength:
			tag = DragonBody
		default:
			tag = DragonTail
		}

		d.config.tileMap.TileAt(d.parts[i]).setDragonPart(tag, d.encodePartID(i))
	}
}

func (d *Dragon) unsetTiling() {
	for i := range d.parts {
		d.config.tileMap.TileAt(d.parts[i]).unsetDragonPart()
	}
}

func (d *Dragon) encodePartID(i int) DragonPartID {
	return DragonPartID(i + 1)
}

func (d *Dragon) decodePartID(id DragonPartID) int {
	return int(id - 1)
}
