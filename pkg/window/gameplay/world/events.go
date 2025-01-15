package world

import (
	"dragon/pkg/domain"
	"dragon/pkg/global/vars"
)

func (w *worldImpl[R]) onKnightEnter(event domain.KnightEnterEvent) {
}

func (w *worldImpl[R]) onKnightLoad(event domain.KnightLoadEvent) {
}

func (w *worldImpl[R]) onKnightKill(event domain.KnightKillEvent) {
}

func (w *worldImpl[R]) onKnightExit(event domain.KnightExitEvent) {
}

func (w *worldImpl[R]) onDragonChop(event domain.DragonChopEvent) {
}

func (w *worldImpl[R]) onGameOver(event domain.GameOverEvent) {
	vars.Window.Page = vars.WindowPageGameOver
}
