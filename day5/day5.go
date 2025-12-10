package day5

import (
	"os"
	"strconv"
	"strings"
)

func loadData(path string) ([][]int, []int) {
	data_bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data_str := string(data_bytes)
	ranges_end := strings.Index(data_str, "\n\n")
	ranges_str := data_str[:ranges_end+1]
	range_count := 0
	for i := 0; i < len(ranges_str); i++ {
		if ranges_str[i] == '\n' {
			range_count++
		}
	}
	ranges := make([][]int, range_count)
	for i := 0; i < len(ranges); i++ {
		ranges[i] = make([]int, 2)
	}
	start := true
	i_range := 0
	for i := 0; i < len(ranges_str); {
		if start {
			next_i := strings.Index(ranges_str[i:], "-") + i
			val, err := strconv.Atoi(ranges_str[i:next_i])
			if err != nil {
				panic(err)
			}
			ranges[i_range][0] = val
			start = false
			i = next_i + 1
		} else {
			next_i := strings.Index(ranges_str[i:], "\n") + i
			val, err := strconv.Atoi(ranges_str[i:next_i])
			if err != nil {
				panic(err)
			}
			ranges[i_range][1] = val
			start = true
			i = next_i + 1
			i_range++
		}
	}
	ids_str := data_str[ranges_end+2:]
	id_count := 0
	for i := 0; i < len(ids_str); i++ {
		if ids_str[i] == '\n' {
			id_count++
		}
	}
	id_count++
	ids := make([]int, id_count)
	i_id := 0
	for i := 0; i < len(ids_str); {
		next_i := strings.Index(ids_str[i:], "\n") + i
		if next_i < i {
			next_i = len(ids_str)
		}
		val, err := strconv.Atoi(ids_str[i:next_i])
		if err != nil {
			panic(err)
		}
		ids[i_id] = val
		i_id++
		i = next_i + 1
	}
	return ranges, ids
}

func valInRange(val int, ranges [][]int) bool {
	for i := 0; i < len(ranges); i++ {
		if val >= ranges[i][0] && val <= ranges[i][1] {
			return true
		}
	}
	return false
}

func Problem1(path string) int {
	ranges, ids := loadData(path)
	count := 0
	for i := 0; i < len(ids); i++ {
		if valInRange(ids[i], ranges) {
			count++
		}
	}
	return count
}

func truncateRanges(ranges [][]int) [][]int {
	out := make([][]int, len(ranges))
	copy(out, ranges)
	for i := 0; i < len(out); i++ {
		for j := 0; j < len(out); j++ {
			if out[i][0] != -1 && out[i][1] != -1 && j != i {
				// top overlap
				if out[i][0] >= out[j][0] && out[i][1] > out[j][1] && out[i][0] <= out[j][1] {
					out[i][0] = out[j][1] + 1
				}
				// bottom overlap
				if out[i][0] < out[j][0] && out[i][1] <= out[j][1] && out[i][1] >= out[j][0] {
					out[i][1] = out[j][0] - 1
				}
				// complete overlap
				if out[i][0] >= out[j][0] && out[i][1] <= out[j][1] {
					out[i][0] = -1
					out[i][1] = -1
				}
			}
		}
	}
	return out
}

func Problem2(path string) int {
	ranges, _ := loadData(path)
	ranges = truncateRanges(ranges)
	count := 0
	for i := 0; i < len(ranges); i++ {
		if ranges[i][0] != -1 {
			count += ranges[i][1] - ranges[i][0] + 1
		}
	}
	return count
}
