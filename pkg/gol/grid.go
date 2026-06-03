package gol

// Can be negative and is not constrained by boundaries
type Cell struct {
	X int
	Y int
}

// Sort by Y axis first, if equal sort by X axis
func SortCell(c1, c2 Cell) int {
	if c1.Y != c2.Y {
		return c1.Y - c2.Y
	}
	return c1.X - c2.X
}

// Only contains live cells
type Grid map[Cell]any

// Returns the list of 8 surrounding neighbours
func (c Cell) Neighbours() []Cell {
	return []Cell{
		{c.X - 1, c.Y - 1},
		{c.X - 1, c.Y},
		{c.X - 1, c.Y + 1},

		{c.X, c.Y - 1},
		{c.X, c.Y + 1},

		{c.X + 1, c.Y - 1},
		{c.X + 1, c.Y},
		{c.X + 1, c.Y + 1},
	}
}
