package day7

import (
	"os"
	"strings"
)

func loadData(path string) ([]string, []bool) {
	data_bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data_str := string(data_bytes)
	line_count := 1
	for i := 0; i < len(data_str); i++ {
		if data_str[i] == '\n' {
			line_count++
		}
	}
	line_length := len(data_str[:strings.Index(data_str, "\n")]) + 1
	lines := make([]string, line_count-1)
	i_line := 0
	for i := 0; i < len(data_str); i += line_length {
		lines[i_line] = data_str[i : i+line_length-1]
		i_line++
	}
	beams := make([]bool, len(lines[0]))
	beams[len(beams)/2] = true
	return lines, beams
}

func processLineCount(line string, beams []bool) ([]bool, int) {
	count := 0
	for i := 1; i < len(beams)-1; i++ {
		if beams[i] && line[i] == '^' {
			beams[i-1] = true
			beams[i] = false
			beams[i+1] = true
			count++
		}
	}
	return beams, count
}

func processBeamPathStart(beams_bools []bool) []int {
	paths := make([]int, len(beams_bools))
	for i := 0; i < len(paths); i++ {
		if beams_bools[i] == true {
			paths[i] = 1
		}
	}
	return paths
}

func processLinePaths(line string, paths []int) []int {
	for i := 1; i < len(paths)-1; i++ {
		val := paths[i]
		if paths[i] > 0 && line[i] == '^' {
			paths[i-1] += val
			paths[i] = 0
			paths[i+1] += val
		}
	}
	return paths
}

func Problem1(path string) int {
	lines, beams := loadData(path)
	total := 0
	for i := 0; i < len(lines); i++ {
		count := 0
		beams, count = processLineCount(lines[i], beams)
		total += count
	}
	return total
}

func Problem2(path string) int {
	lines, beams := loadData(path)
	paths := processBeamPathStart(beams)
	for i := 0; i < len(lines); i++ {
		paths = processLinePaths(lines[i], paths)
	}
	sum := 0
	for i := 0; i < len(paths); i++ {
		sum += paths[i]
	}
	return sum
}
