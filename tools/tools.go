package tools

import (
	"log"
	"math"
	"sort"
)

/*
 * Write a function which would take two variables as input:
 *
 * Variable 1(a list of numbers) - example [1, 2, 4, 7, 9, 11]
 * Variable 2(a target number) - example 11
 *
 * Output -  Return a list of all the combinations of numbers given in input 1
 * which adds up to target. The result of the above inputs would be:
 *
 *     [[4,7], [9,2], [11]]
 */

func deleteDuplicated(digits []int) []int {
	hash := make(map[int]bool)

	for _, v := range digits {
		hash[v] = true
	}

	r := []int{}
	for k := range hash {
		r = append(r, k)
	}

	return r
}

func calcPartialSums(partial, target int, left []int, tested *int) (nums [][]int) {
	// For each digit in the left []int
	for i, digit := range left {
		*tested++
		// adding the current digit to the partial total
		temp := partial + digit

		if temp == target {
			// if we reach the expected result then we need to return the current digit
			nums = append(nums, []int{digit})
			//fmt.Printf("gotcha. %v\n", temp)
			return
		} else if temp > target {
			// If the result is bigger than the target, we returns nothing
			return
		}

		// Now, if the total is lower than the target, then we're gonna test any other num in the array
		r := calcPartialSums(temp, target, left[i+1:], tested)

		// if this returns something, we append the current digit to each result
		for ix := range r {
			r[ix] = append([]int{digit}, r[ix]...)
		}

		// then we append the returned results to the current results
		nums = append(nums, r...)
	}
	return
}

// GetNumbersWhichSums takes a []int and extract all combinations of numbers
// that sums = total in a form of [][]int
func GetNumbersWhichSums(digits []int, total int) (result [][]int) {
	// Cleaning duplicates
	// digits = deleteDuplicated(digits)

	//	Sorting digits, this prevents to check combinations which sums more than the expected result
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] < digits[j]
	})

	testedCases := 0

	// Calling the recursive function that iterates over each combination
	log.Printf("calcPartialSums(partial: %v, total: %v, left: %v)\n", 0, total, digits)
	result = calcPartialSums(0, total, digits, &testedCases)
	log.Printf("Detected %v results\n"+
		"\t\t\t\t\ttested %v of %v possible combinations", len(result), testedCases, math.Pow(2, float64(len(digits)))-1)

	return
}
