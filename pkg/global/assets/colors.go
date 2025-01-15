package assets

import (
	"github.com/a1emax/youngine/asset/format/rgba"
	"github.com/a1emax/youngine/x/scope"
)

var Colors struct {
	DebugInfoBackground   rgba.Asset
	DebugInfoText         rgba.Asset
	GameOverBackground    rgba.Asset
	GameOverButtonPressed rgba.Asset
	GameOverButtonPrimary rgba.Asset
	GameOverButtonText    rgba.Asset
	GamePlayBackground    rgba.Asset
	GamePlayDirectionPad  rgba.Asset
	MainMenuBackground    rgba.Asset
	MainMenuButtonPressed rgba.Asset
	MainMenuButtonPrimary rgba.Asset
	MainMenuButtonText    rgba.Asset
}

func initColors(lc scope.Lifecycle) {
	Colors.DebugInfoBackground = load[rgba.Asset](lc, "colors/debug-info-background.rgba")
	Colors.DebugInfoText = load[rgba.Asset](lc, "colors/debug-info-text.rgba")
	Colors.GameOverBackground = load[rgba.Asset](lc, "colors/game-over-background.rgba")
	Colors.GameOverButtonPressed = load[rgba.Asset](lc, "colors/game-over-button-pressed.rgba")
	Colors.GameOverButtonPrimary = load[rgba.Asset](lc, "colors/game-over-button-primary.rgba")
	Colors.GameOverButtonText = load[rgba.Asset](lc, "colors/game-over-button-text.rgba")
	Colors.GamePlayBackground = load[rgba.Asset](lc, "colors/game-play-background.rgba")
	Colors.GamePlayDirectionPad = load[rgba.Asset](lc, "colors/game-play-direction-pad.rgba")
	Colors.MainMenuBackground = load[rgba.Asset](lc, "colors/main-menu-background.rgba")
	Colors.MainMenuButtonPressed = load[rgba.Asset](lc, "colors/main-menu-button-pressed.rgba")
	Colors.MainMenuButtonPrimary = load[rgba.Asset](lc, "colors/main-menu-button-primary.rgba")
	Colors.MainMenuButtonText = load[rgba.Asset](lc, "colors/main-menu-button-text.rgba")
}
