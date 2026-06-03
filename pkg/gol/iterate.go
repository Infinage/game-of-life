package gol

// Denotes a single iteration in Conways game of life
func (grid *Grid) Next() {
	// Track alive neighbours around region of interest
	neighbourCounts := make(map[Cell]int)
	for cell := range *grid {
		for _, neighbour := range cell.Neighbours() {
			neighbourCounts[neighbour]++
		}
	}

	// Grid for the next generation
	nextGrid := make(Grid)

	// Conditions to advance a cell:
	// If cell is alive - must contain 2 or 3 alive neighbours
	// If cell is  dead - must contain exactly 3 alive neighbours
	for cell, count := range neighbourCounts {
		_, isAlive := (*grid)[cell]
		if count == 3 || (count == 2 && isAlive) {
			nextGrid[cell] = nil
		}
	}

	// Update grid with next iteration
	*grid = nextGrid
}
