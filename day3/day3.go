package day3

import (
	"os"
	"strconv"
	"strings"
)

func loadJolts(path string) [][]int {
	jolts_bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	jolts_str := string(jolts_bytes)
	bank_length := strings.Index(jolts_str, "\n")
	bank_count := len(jolts_str) / bank_length
	banks := make([][]int, bank_count)
	for i := 0; i < len(banks); i++ {
		banks[i] = make([]int, bank_length)
	}
	i_bank := 0
	i_battery := 0
	for i := 0; i < len(jolts_str); i++ {
		if jolts_str[i] == '\n' {
			i_bank++
			i_battery = 0
		} else {
			val, err := strconv.Atoi(jolts_str[i : i+1])
			if err != nil {
				panic(err)
			}
			banks[i_bank][i_battery] = val
			i_battery++
		}
	}
	return banks
}

func findMaxTwoComb(bank []int) int {
	max := 0
	for i := 0; i < len(bank)-1; i++ {
		for j := i + 1; j < len(bank); j++ {
			comb := bank[i]*10 + bank[j]
			if comb > max {
				max = comb
			}
		}
	}
	return max
}

func Problem1(path string) int {
	banks := loadJolts(path)
	out := 0
	for i := 0; i < len(banks); i++ {
		max_comb := findMaxTwoComb(banks[i])
		out += max_comb
	}
	return out
}

func pow(x int, y int) int {
	out := 1
	for i := 0; i < y; i++ {
		out *= x
	}
	return out
}

func nextBinNum(combo []int) []int {
	for i := 0; i < len(combo); i++ {
		combo[i]++
		if combo[i] == 1 {
			return combo
		}
		combo[i] = 0
	}
	return combo
}

func crossSum(combo []int) int {
	sum := 0
	for i := 0; i < len(combo); i++ {
		sum += combo[i]
	}
	return sum
}

func nextCombo(combo []int, size int) ([]int, bool) {
	cross_sum := crossSum(combo)
	for cross_sum > 0 {
		combo = nextBinNum(combo)
		cross_sum = crossSum(combo)
		if cross_sum == size {
			return combo, false
		}
	}
	return combo, true
}

func findMaxJoltageCombos(bank []int, size int) int {
	combo := make([]int, len(bank))
	combo[0] = 1
	end := false
	max := 0
	combo, end = nextCombo(combo, size)
	for end == false {
		joltage := 0
		digit := size - 1
		for i := 0; i < len(bank); i++ {
			if combo[i] == 1 {
				joltage += bank[i] * pow(10, digit)
				digit--
			}
		}
		if joltage > max {
			max = joltage
		}
		combo, end = nextCombo(combo, size)
	}
	return max
}

func maxRange(bank_sample []int) (int, int) {
	max := 0
	i_max := 0
	for i := 0; i < len(bank_sample); i++ {
		if bank_sample[i] > max {
			max = bank_sample[i]
			i_max = i
		}
	}
	return max, i_max
}

func findMaxJoltage(bank []int, size int) int {
	i_start := 0
	out := 0
	for digit := size; digit > 0; digit-- {
		i_end := len(bank) - digit + 1
		max, i_max := maxRange(bank[i_start:i_end])
		i_start = i_start + i_max + 1
		out += max * pow(10, digit-1)
	}
	return out
}

func Problem2(path string) int {
	banks := loadJolts(path)
	out := 0
	for i := 0; i < len(banks); i++ {
		max_comb := findMaxJoltage(banks[i], 12)
		out += max_comb
	}
	return out
}
