package day6

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func loadData(path string) ([][]int, []byte) {
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
	col_count := 0
	num := false
	row_length := len(data_str[:strings.Index(data_str, "\n")])
	for i := 0; i < row_length; i++ {
		if unicode.IsDigit(rune(data_str[i])) && !num {
			col_count++
			num = true
		}
		if !unicode.IsDigit(rune(data_str[i])) && num {
			num = false
		}
	}
	data := make([][]int, line_count)
	num = false
	for i := 0; i < line_count; i++ {
		data[i] = make([]int, col_count)
	}
	num = false
	row := 0
	col := 0
	for i_row := 0; i_row < len(data_str)-row_length; i_row += row_length + 1 {
		for i := i_row; i < i_row+row_length; i++ {
			if unicode.IsDigit(rune(data_str[i])) {
				num = true
				val, err := strconv.Atoi(data_str[i : i+1])
				if err != nil {
					panic(err)
				}
				data[row][col] *= 10
				data[row][col] += val
			}
			if !unicode.IsDigit(rune(data_str[i])) && num {
				col++
				num = false
			}
		}
		col = 0
		row++
	}
	ops := make([]byte, col_count)
	col = 0
	for i := len(data_str) - row_length; i < len(data_str); i++ {
		if data_str[i] != ' ' {
			ops[col] = data_str[i]
			col++
		}
	}
	return data, ops
}

func performOperation(data [][]int, ops []byte, col int) int {
	out := data[0][col]
	for i := 1; i < len(data); i++ {
		switch ops[col] {
		case '+':
			out += data[i][col]
		case '*':
			out *= data[i][col]
		}
	}
	return out
}

func Problem1(path string) int {
	data, ops := loadData(path)
	out := 0
	for i := 0; i < len(data[0]); i++ {
		out += performOperation(data, ops, i)
	}
	return out
}

func loadVerticalFormat(path string) int {
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
	out := 0
	row_length := len(data_str[:strings.Index(data_str, "\n")]) + 1
	nums := make([]int, 4)
	i_num := 0
	for col := row_length - 1; col > 0; col-- {
		total := 0
		for row := 0; row < line_count; row++ {
			row_i := row_length*row - 1
			if unicode.IsDigit(rune(data_str[col+row_i])) {
				val, err := strconv.Atoi(string(data_str[col+row_i]))
				if err != nil {
					panic(err)
				}
				total *= 10
				total += val
				nums[i_num] = total
			}
			if data_str[col+row_i] == '+' {
				res := nums[0]
				for i := 1; i < len(nums); i++ {
					res += nums[i]
				}
				out += res
				nums = make([]int, 4)
				i_num = 0
			}
			if data_str[col+row_i] == '*' {
				res := nums[0]
				for i := 1; i < len(nums); i++ {
					if nums[i] > 0 {
						res *= nums[i]
					}
				}
				out += res
				nums = make([]int, 4)
				i_num = 0
			}
		}
		if nums[i_num] > 0 {
			i_num++
		}
	}
	return out
}

func Problem2(path string) int {
	return loadVerticalFormat(path)
}
