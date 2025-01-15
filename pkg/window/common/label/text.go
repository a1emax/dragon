package label

import (
	"math"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/x/textview"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type textPart struct {
	textKey basic.Opt[textKey]
	text    textview.SingleLine
}

type textKey struct {
	width    basic.Float
	fontFace font.Face
	text     string
}

func (l *labelImpl[R]) setupText() {
	r := l.region.Rect()

	key := textKey{
		width:    r.Width(),
		fontFace: l.state.TextFontFace,
		text:     l.state.Text,
	}

	if l.textKey.IsSet() && l.textKey.Get() == key {
		return
	}

	l.disposeText()

	if key.fontFace == nil || key.text == "" {
		return
	}

	l.textKey = basic.SetOpt(key)
	l.text = textview.NewSingleLine(key.width, key.fontFace, key.text)
}

func (l *labelImpl[R]) drawText(screen *ebiten.Image) {
	if !l.textKey.IsSet() {
		return
	}

	if l.state.TextColor == nil {
		return
	}

	r := l.region.Rect()

	left := r.Left() + (r.Width()-l.text.Width())/2
	top := r.Top() + (r.Height()-l.text.Height())/2

	l.text.Draw(textview.StringDrawerFunc(func(s string, x, y basic.Float, fontFace font.Face) {
		text.Draw(screen, s, fontFace, int(math.Floor(left+x)), int(math.Floor(top+y)), l.state.TextColor)
	}))
}

func (l *labelImpl[R]) disposeText() {
	l.textPart = textPart{}
}
