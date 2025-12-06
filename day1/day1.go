package day1

import (
	"os"
	"strconv"
	"strings"
)

func divmod(a int, b int) (int, int) {
	div := 0
	mod := a
	for mod < 0 || mod >= b {
		if mod < 0 {
			mod += b
			div -= 1
		} else if mod >= b {
			mod -= b
			div += 1
		}
	}
	return div, mod
}

func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

func loadInstructions(path string) []int {
	inst_bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	inst_str := string(inst_bytes)
	line_count := 0
	for i := 0; i < len(inst_str); i++ {
		if inst_str[i] == '\n' {
			line_count++
		}
	}
	out := make([]int, line_count+1)
	position := 0
	for i := 0; i < len(inst_str); i++ {
		rest_str := inst_str[i:]
		if rest_str[0] == 'L' || rest_str[0] == 'R' {
			end_ind := strings.Index(rest_str, "\n")
			if end_ind < 0 {
				end_ind = len(rest_str)
			}
			val, err := strconv.Atoi(rest_str[1:end_ind])
			if err != nil {
				panic(err)
			}
			if rest_str[0] == 'L' {
				val = 0 - val
			}
			out[position] = val
			position++
		}
	}
	return out
}

func countZeroPos(inst []int) int {
	pos := 50
	count := 0
	for i := 0; i < len(inst); i++ {
		pos += inst[i]
		_, pos = divmod(pos, 100)
		if pos == 0 {
			count++
		}
	}
	return count
}

func Problem1(path string) int {
	inst := loadInstructions(path)
	return countZeroPos(inst)
}

func countZeroCrossings(inst []int) int {
	pos := 50
	div := 0
	count := 0
	for i := 0; i < len(inst); i++ {
		// prevent double countings starting from zero turning left
		if pos == 0 && inst[i] < 0 {
			count--
		}
		pos += inst[i]
		div, pos = divmod(pos, 100)
		// prevent missed counts landing on zero turning left
		if pos == 0 && inst[i] < 0 {
			count++
		}
		count += abs(div)
	}
	return count
}

func Problem2(path string) int {
	inst := loadInstructions(path)
	return countZeroCrossings(inst)
}
