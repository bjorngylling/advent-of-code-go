package cave

type Cave struct {
	m      []bool
	Width  int
	Height int
}

func New(w, h int) *Cave {
	return &Cave{m: make([]bool, w*h), Width: w, Height: h}
}

func (c *Cave) Blocked(x, y int) bool {
	return c.m[x+y*c.Width]
}

func (c *Cave) SetBlocked(x, y int, val bool) {
	c.m[x+y*c.Width] = val
}
