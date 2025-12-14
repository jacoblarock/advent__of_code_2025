package day8

import (
	"math"
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
		coords[i] = make([]int, 3)
	}
	row := 0
	col := 0
	for i := 0; i < len(data_str)-1; i++ {
		if data_str[i] == ',' {
			col++
		} else if data_str[i] == '\n' {
			col = 0
			row++
		} else {
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

func initCircuits(coords [][]int) []int {
	out := make([]int, len(coords))
	for i := 0; i < len(out); i++ {
		out[i] = i
	}
	return out
}

func getDistances(coords [][]int) [][]float64 {
	out := make([][]float64, len(coords))
	for i := 0; i < len(coords)-1; i++ {
		out[i] = make([]float64, len(coords))
		for j := i + 1; j < len(coords); j++ {
			out[i][j] = math.Sqrt(
				math.Pow(float64(coords[j][0])-float64(coords[i][0]), 2) +
					math.Pow(float64(coords[j][1])-float64(coords[i][1]), 2) +
					math.Pow(float64(coords[j][2])-float64(coords[i][2]), 2))
		}
	}
	return out
}

func findShortest(distances [][]float64, links [][]bool) (int, int) {
	out := []int{0, 1}
	for i := 0; i < len(distances)-1; i++ {
		for j := i + 1; j < len(distances); j++ {
			if !links[i][j] && distances[i][j] < distances[out[0]][out[1]] {
				out[0] = i
				out[1] = j
			}
		}
	}
	return out[0], out[1]
}

func linkNodes(a int, b int, circuits []int, links [][]bool) ([]int, [][]bool) {
	temp := circuits[b]
	if links[a][b] == false {
		links[a][b] = true
		for i := 0; i < len(circuits); i++ {
			if circuits[i] == temp {
				circuits[i] = circuits[a]
			}
		}
	}
	return circuits, links
}

func countCircuits(circuits []int) int {
	count := 0
	for i := 0; i < len(circuits); i++ {
		if circuits[i] != i {
			count++
		}
	}
	return len(circuits) - count
}

func getCircuitLengths(circuits []int) []int {
	out := make([]int, len(circuits))
	for i := 0; i < len(circuits); i++ {
		for j := 0; j < len(circuits); j++ {
			if circuits[j] == i {
				out[i]++
			}
		}
	}
	return out
}

func popLongest(circuits []int) (int, []int) {
	max_val := 0
	max_i := 0
	for i := 0; i < len(circuits); i++ {
		if circuits[i] > max_val {
			max_val = circuits[i]
			max_i = i
		}
	}
	circuits[max_i] = 0
	return max_val, circuits
}

func Problem1(path string) int {
	coords := loadData(path)
	circuits := initCircuits(coords)
	links := make([][]bool, len(coords))
	for i := 0; i < len(links); i++ {
		links[i] = make([]bool, len(links))
	}
	distances := getDistances(coords)
	for i := 0; i < 1000; i++ {
		a, b := findShortest(distances, links)
		circuits, links = linkNodes(a, b, circuits, links)
	}
	lengths := getCircuitLengths(circuits)
	out := 1
	for i := 0; i < 3; i++ {
		val := 0
		val, lengths = popLongest(lengths)
		out *= val
	}
	return out
}

func Problem2(path string) int {
	coords := loadData(path)
	circuits := initCircuits(coords)
	links := make([][]bool, len(coords))
	for i := 0; i < len(links); i++ {
		links[i] = make([]bool, len(links))
	}
	distances := getDistances(coords)
	a, b := 0, 0
	for countCircuits(circuits) > 1 {
		a, b = findShortest(distances, links)
		circuits, links = linkNodes(a, b, circuits, links)
	}
	return coords[a][0] * coords[b][0]
}
