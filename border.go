package termui

import (
	"github.com/boombuler/termui/css"
)

// Border draws a border around a given child element
type Border struct {
	BaseElement
	child Element

	width, height int
}

var _ Element = new(Border) // Interface checking...

// Width returns the width of the border.
func (b *Border) Width() int {
	return b.width
}

// Height returns the height of the border.
func (b *Border) Height() int {
	return b.height
}

// Children returns all nested child elements of the border.
func (b *Border) Children() []css.Styleable {
	if b.child != nil {
		return []css.Styleable{b.child}
	}
	return []css.Styleable{}
}

// Name returns the constant name of the border element for css styling.
func (b *Border) Name() string {
	return "border"
}

// Measure gets the "wanted" size of the element based on the available size.
func (b *Border) Measure(availableWidth, availableHeight int) (width int, height int) {
	if b.child == nil {
		return 2, 2
	}

	cw, ch := MeasureChild(b.child, availableWidth-2, availableHeight-2)
	grav := GravityProperty.Get(b)
	if grav&horizontal == horizontal && availableWidth > 0 { // stretch
		width = availableWidth
	} else {
		width = cw + 2
	}
	if grav&vertical == vertical && availableHeight > 0 {
		height = availableHeight
	} else {
		height = ch + 2
	}
	return
}

// Arrange sets the final size for the Element end tells it to Arrange itself
func (b *Border) Arrange(finalWidth, finalHeight int) {
	if finalWidth > 2 && finalHeight > 2 && b.child != nil {
		ArrangeChild(b.child, finalWidth-2, finalHeight-2)
	}
	b.width, b.height = finalWidth, finalHeight
}

// Render renders the element on the given Renderer
func (b *Border) Render(r Renderer) {
	for x := 1; x < b.width-1; x++ {
		r.Set(x, 0, borderHorizontalLine)
		r.Set(x, b.height-1, borderHorizontalLine)

		for y := 1; y < b.height-1; y++ {
			r.Set(x, y, ' ')
		}
	}
	for y := 1; y < b.height-1; y++ {
		r.Set(0, y, borderVerticalLine)
		r.Set(b.width-1, y, borderVerticalLine)
	}

	r.Set(0, 0, borderTopLeft)
	r.Set(0, b.height-1, borderBottomLeft)
	r.Set(b.width-1, 0, borderTopRight)
	r.Set(b.width-1, b.height-1, borderBottomRight)

	if b.child != nil {
		r.RenderChild(b.child, b.width-2, b.height-2, 1, 1)
	}
}

// NewBorder creates a new border element with the given child
func NewBorder(child Element) *Border {
	b := &Border{
		child: child,
	}
	if child != nil {
		child.SetParent(b)
	}
	return b
}

// TextBorder is a border with a text on the top left corner
type TextBorder struct {
	*Border
	txt []rune
}

var _ Element = new(TextBorder) // Interface checking...

// NewTextBorder creates a new TextBorder with a given text and child
func NewTextBorder(txt string, child Element) *TextBorder {
	return &TextBorder{
		NewBorder(child),
		[]rune(txt),
	}
}

// Render renders the element on the given Renderer
func (b *TextBorder) Render(rn Renderer) {
	b.Border.Render(rn)
	w := b.Width() - 2
	for i, r := range b.txt {
		if i > w {
			break
		}
		rn.Set(i+1, 0, r)
	}
}

// SetText sets the text on the border
func (b *TextBorder) SetText(txt string) {
	b.txt = []rune(txt)
}

// Text returns the current text value
func (b *TextBorder) Text() string {
	return string(b.txt)
}
