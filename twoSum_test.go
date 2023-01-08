package leetcode

import (
	"fmt"
	"testing"
)

// func twoSum(nums []int, target int) []int {
// 	for idx0, num0 := range nums {
// 		for idx1, num1 := range nums {
// 			if idx0 == idx1 {
// 				continue
// 			}
// 			if num0+num1 == target {
// 				return []int{idx0, idx1}
// 			}
// 		}
// 	}
// 	return nil
// }

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		m[num] = idx
	}

	for idx, num := range nums {
		diff := target - num
		i, ok := m[diff]
		if ok && i != idx {
			return []int{idx, i}
		}
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	testcases := []struct {
		name   string
		nums   []int
		target int
		out    []int
	}{
		{
			name:   "Example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			out:    []int{0, 1},
		}, {
			name:   "Example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			out:    []int{1, 2},
		}, {
			name:   "Example 3",
			nums:   []int{3, 3},
			target: 6,
			out:    []int{0, 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			o := twoSum(tc.nums, tc.target)
			if o[len(o)-1] != tc.out[len(tc.out)-1] || o[0] != tc.out[0] {
				fmt.Printf("tc.out: %v ------ real.out: %v\n", tc.out, o)
				t.Fail()
			}
		})
	}
}
