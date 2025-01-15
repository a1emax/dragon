package space

func FindPath(from TileVec, isPassable, isTerminal func(t TileVec) bool) []TileVec {
	horizon := []TileVec{from}
	before := map[TileVec]*TileVec{from: nil}
	var to *TileVec
	for len(horizon) > 0 {
		current := horizon[0]
		horizon = horizon[1:]

		if isTerminal(current) {
			to = &current

			break
		}

		for d := North; d <= West; d++ {
			if next := AddDirection(current, d); isPassable(next) {
				if _, ok := before[next]; !ok {
					horizon = append(horizon, next)
					before[next] = &current
				}
			}
		}
	}

	var path []TileVec
	if to != nil && before[*to] != nil {
		for t := *to; t != from; t = *before[t] {
			path = append(path, t)
		}
	}

	return path
}
