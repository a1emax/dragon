package button

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

func (b *buttonImpl[R]) setupText() {
	r := b.region.Rect()

	key := textKey{
		width:    r.Width(),
		fontFace: b.state.TextFontFace,
		text:     b.state.Text,
	}

	if b.textKey.IsSet() && b.textKey.Get() == key {
		return
	}

	b.disposeText()

	if key.fontFace == nil || key.text == "" {
		return
	}

	b.textKey = basic.SetOpt(key)
	b.text = textview.NewSingleLine(key.width, key.fontFace, key.text)
}

func (b *buttonImpl[R]) drawText(screen *ebiten.Image) {
	if !b.textKey.IsSet() {
		return
	}

	clr := b.color(b.state.TextPrimaryColor, b.state.TextPressedColor)
	if clr == nil {
		return
	}

	r := b.region.Rect()

	left := r.Left() + (r.Width()-b.text.Width())/2
	top := r.Top() + (r.Height()-b.text.Height())/2

	b.text.Draw(textview.StringDrawerFunc(func(s string, x, y basic.Float, fontFace font.Face) {
		text.Draw(screen, s, fontFace, int(math.Floor(left+x)), int(math.Floor(top+y)), clr)
	}))
}

func (b *buttonImpl[R]) disposeText() {
	b.textPart = textPart{}
}
