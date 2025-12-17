package day9

import (
	"os"
	"strconv"
)

func loadData(path string) [][]int {
	data_bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data_str := string(data_bytes)
	line_count := 0
	for i := 0; i < len(data_str); i++ {
		if data_str[i] == '\n' {
			line_count++
		}
	}
	coords := make([][]int, line_count)
	for i := 0; i < len(coords); i++ {
		coords[i] = make([]int, 2)
	}
	row := 0
	col := 0
	for i := 0; i < len(data_str)-1; i++ {
		switch data_str[i] {
		case ',':
			col++
		case '\n':
			col = 0
			row++
		default:
			val, err := strconv.Atoi(string(data_str[i]))
			if err != nil {
				panic(err)
			}
			coords[row][col] *= 10
			coords[row][col] += val
		}
	}
	return coords
}

func getArea(a []int, b []int) int {
	out := (a[0] - b[0] + 1) * (a[1] - b[1] + 1)
	if out > 0 {
		return out
	} else {
		return -out
	}
}

func Problem1(path string) int {
	tiles := loadData(path)
	max_area := 0
	for i := 0; i < len(tiles)-1; i++ {
		for j := i + 1; j < len(tiles); j++ {
			area := getArea(tiles[i], tiles[j])
			if area > max_area {
				max_area = area
			}
		}
	}
	return max_area
}

func containsOtherTile(tiles [][]int, a int, b int) bool {
	for i := 0; i < len(tiles); i++ {
		contains_x := false
		contains_y := false
		if i == a-1 || i == a || i == a+1 {
			continue
		}
		if i == b-1 || i == b || i == b+1 {
			continue
		}
		if tiles[i][0]-tiles[a][0] < 0 && tiles[i][0]-tiles[b][0] > 0 {
			contains_y = true
		}
		if tiles[i][0]-tiles[a][0] > 0 && tiles[i][0]-tiles[b][0] < 0 {
			contains_y = true
		}
		if tiles[i][1]-tiles[a][1] < 0 && tiles[i][1]-tiles[b][1] > 0 {
			contains_x = true
		}
		if tiles[i][1]-tiles[a][1] > 0 && tiles[i][1]-tiles[b][1] < 0 {
			contains_x = true
		}
		if contains_x && contains_y {
			return true
		}
	}
	return false
}

func linesIntersect(a_start []int, a_end []int, b_start []int, b_end []int) bool {
	// parallel lines
	if a_start[0] == a_end[0] && b_start[0] == b_end[0] {
		return false
	}
	if a_start[1] == a_end[1] && b_start[1] == b_end[1] {
		return false
	}
	// exclude shared corners
	if a_start[0] == b_start[0] && a_start[1] == b_start[1] {
		return false
	}
	if a_end[0] == b_end[0] && a_end[1] == b_end[1] {
		return false
	}
	if a_start[0] == a_end[0] && ((a_start[1] > b_start[1] && a_end[1] < b_start[1]) || (a_start[1] < b_start[1] && a_end[1] > b_start[1])) {
		if b_start[0] > a_start[0] && b_end[0] < a_start[0] {
			return true
		}
		if b_start[0] < a_start[0] && b_end[0] > a_start[0] {
			return true
		}
	}
	if a_start[1] == a_end[1] && ((a_start[0] > b_start[0] && a_end[0] < b_start[0]) || (a_start[0] < b_start[0] && a_end[0] > b_start[0])) {
		if b_start[1] > a_start[1] && b_end[1] < a_start[1] {
			return true
		}
		if b_start[1] < a_start[1] && b_end[1] > a_start[1] {
			return true
		}
	}
	return false
}

func rectPoints(a []int, b []int) [][]int {
	out := make([][]int, 4)
	out[0] = a
	out[1] = []int{a[0], b[1]}
	out[2] = b
	out[3] = []int{b[0], a[1]}
	return out
}

func rectIntersect(a []int, b []int, line_start []int, line_end []int) bool {
	points := rectPoints(a, b)
	for i := 0; i < len(points); i++ {
		j := (i + 1) % len(points)
		if linesIntersect(points[i], points[j], line_start, line_end) {
			return true
		}
	}
	return false
}

func existsIntersect(a []int, b []int, tiles [][]int) bool {
	for i := 0; i < len(tiles); i++ {
		j := (i + 1) % len(tiles)
		if rectIntersect(a, b, tiles[i], tiles[j]) {
			return true
		}
	}
	return false
}

func Problem2(path string) int {
	tiles := loadData(path)
	max_area := 0
	for i := 0; i < len(tiles)-1; i++ {
		for j := i + 1; j < len(tiles); j++ {
			if !existsIntersect(tiles[i], tiles[j], tiles) && !containsOtherTile(tiles, i, j) {
				area := getArea(tiles[i], tiles[j])
				if area > max_area {
					max_area = area
				}
			}
		}
	}
	return 0
}
