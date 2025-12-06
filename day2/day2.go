package day2

import (
	"os"
	"strconv"
	"strings"
)

func getRanges(path string) [][]int {
	ranges_bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	ranges_str := string(ranges_bytes)
	range_count := 1
	for i := 0; i < len(ranges_str); i++ {
		if ranges_str[i] == ',' {
			range_count++
		}
	}
	ranges := make([][]int, range_count)
	for i := 0; i < len(ranges); i++ {
		ranges[i] = make([]int, 2)
	}
	rest_str := ranges_str
	start := true // true when reading range start, false when reading range end
	i := 0
	start_ind := 0
	end_ind := 0
	for start_ind < len(rest_str) {
		rest_str = rest_str[start_ind:]
		if start {
			end_ind = strings.Index(rest_str, "-")
			val, err := strconv.Atoi(rest_str[:end_ind])
			if err != nil {
				panic(err)
			}
			ranges[i][0] = val
			start = false
		} else {
			end_ind = strings.Index(rest_str, ",")
			if end_ind == -1 {
				end_ind = len(rest_str)
			}
			val, err := strconv.Atoi(rest_str[:end_ind])
			if err != nil {
				panic(err)
			}
			ranges[i][1] = val
			start = true
			i++
		}
		start_ind = end_ind + 1
	}
	return ranges
}

func containsRepeat(x int) bool {
	str := strconv.Itoa(x)
	if len(str)%2 == 0 {
		first_half := str[:len(str)/2]
		second_half := str[len(str)/2:]
		if first_half == second_half {
			return true
		}
	}
	return false
}

func Problem1(path string) int {
	ranges := getRanges(path)
	out := 0
	for i := 0; i < len(ranges); i++ {
		start := ranges[i][0]
		end := ranges[i][1]
		for num := start; num <= end; num++ {
			if containsRepeat(num) {
				out += num
			}
		}
	}
	return out
}

func containsRepeats(x int) bool {
	str := strconv.Itoa(x)
	for chunk_size := 1; chunk_size <= len(str)/2; chunk_size++ {
		if len(str)%chunk_size == 0 {
			match := true
			ref := str[:chunk_size]
			for i := chunk_size; i < len(str); i += chunk_size {
				chunk := str[i : i+chunk_size]
				if ref != chunk {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

func Problem2(path string) int {
	ranges := getRanges(path)
	out := 0
	for i := 0; i < len(ranges); i++ {
		start := ranges[i][0]
		end := ranges[i][1]
		for num := start; num <= end; num++ {
			if containsRepeats(num) {
				out += num
			}
		}
	}
	return out
}
