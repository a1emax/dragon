package domain

type Tile struct {
	groundTag     GroundTag
	objectTag     ObjectTag
	knightID      KnightID
	dragonPartTag DragonPartTag
	dragonPartID  DragonPartID

	reserved [3]byte
}

type GroundTag uint8

const (
	_ GroundTag = iota
	Floor
	Magma
)

type ObjectTag uint8

const (
	_ ObjectTag = iota
	Wall
	SmallWall
	Door
	Treasure
)

type KnightID uint8

const MaxKnightID KnightID = 255

type DragonPartTag uint8

const (
	_ DragonPartTag = iota
	DragonHead
	DragonBody
	DragonTail
)

type DragonPartID uint8

const MaxDragonPartID DragonPartID = 255

func (t *Tile) GroundTag() GroundTag {
	return t.groundTag
}

func (t *Tile) setGround(tag GroundTag) {
	t.groundTag = tag
}

func (t *Tile) ObjectTag() ObjectTag {
	return t.objectTag
}

func (t *Tile) setObject(tag ObjectTag) {
	t.objectTag = tag
}

func (t *Tile) unsetObject() {
	t.objectTag = 0
}

func (t *Tile) KnightID() KnightID {
	return t.knightID
}

func (t *Tile) setKnight(id KnightID) {
	t.knightID = id
}

func (t *Tile) unsetKnight() {
	t.knightID = 0
}

func (t *Tile) DragonPartTag() DragonPartTag {
	return t.dragonPartTag
}

func (t *Tile) DragonPartID() DragonPartID {
	return t.dragonPartID
}

func (t *Tile) setDragonPart(tag DragonPartTag, id DragonPartID) {
	t.dragonPartTag = tag
	t.dragonPartID = id
}

func (t *Tile) unsetDragonPart() {
	t.dragonPartTag = 0
	t.dragonPartID = 0
}

func (t *Tile) IsPassableForKnight(dragon bool) bool {
	return t != nil &&
		t.groundTag == Floor &&
		(t.objectTag == 0 || t.objectTag == SmallWall || t.objectTag == Door || t.objectTag == Treasure) &&
		(!dragon || t.dragonPartTag == 0 || t.dragonPartTag == DragonTail) &&
		t.knightID == 0
}

func (t *Tile) IsPassableForDragon() bool {
	return t != nil &&
		(t.groundTag == Floor || t.groundTag == Magma) &&
		(t.objectTag == 0 || t.objectTag == Treasure) &&
		t.dragonPartTag == 0
}
