package day4

import (
	"os"
	"strings"
)

func loadGrid(path string) [][]int {
	grid_bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	grid_str := string(grid_bytes)
	length := strings.Index(grid_str, "\n")
	width := len(grid_str) / length
	grid := make([][]int, width)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, length)
	}
	row := 0
	col := 0
	for i := 0; i < len(grid_str); i++ {
		if grid_str[i] == '.' {
			grid[row][col] = 0
			col++
		} else if grid_str[i] == '@' {
			grid[row][col] = 1
			col++
		} else {
			col = 0
			row++
		}
	}
	return grid
}

func countNeighbors(grid [][]int, row int, col int) int {
	sum := 0
	max_row := len(grid) - 1
	max_col := len(grid[0]) - 1
	// top
	if row > 0 {
		if col > 0 {
			sum += grid[row-1][col-1]
		}
		sum += grid[row-1][col]
		if col < max_col {
			sum += grid[row-1][col+1]
		}
	}
	// middle
	if col > 0 {
		sum += grid[row][col-1]
	}
	if col < max_col {
		sum += grid[row][col+1]
	}
	// bottom
	if row < max_row {
		if col > 0 {
			sum += grid[row+1][col-1]
		}
		sum += grid[row+1][col]
		if col < max_col {
			sum += grid[row+1][col+1]
		}
	}
	return sum
}

func Problem1(path string) int {
	grid := loadGrid(path)
	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid); col++ {
			if grid[row][col] == 1 && countNeighbors(grid, row, col) < 4 {
				count++
			}
		}
	}
	return count
}

func removeRolls(grid [][]int) ([][]int, int) {
	count := 0
	new_grid := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		new_grid[i] = make([]int, len(grid[0]))
	}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid); col++ {
			if grid[row][col] == 1 && !(countNeighbors(grid, row, col) < 4) {
				new_grid[row][col] = 1
			}
			if grid[row][col] == 1 && countNeighbors(grid, row, col) < 4 {
				count++
			}
		}
	}
	return new_grid, count
}

func Problem2(path string) int {
	grid := loadGrid(path)
	count := -1
	total := 0
	for count != 0 {
		grid, count = removeRolls(grid)
		total += count
	}
	return total
}
