package arrs

// Sum unction calcs total sum of arr
func Sum(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return sum
}

// SumAll function takes range of slice, calcs sum
// of inner element for each, and returs total slice,
// where each element - sum of provided slices elements
func SumAll(nums ...[]int) []int {
	res := make([]int, 0)
	for _, n := range nums {
		res = append(res, Sum(n))
	}
	return res
}

// SumAllTails ...
func SumAllTails(nums ...[]int) []int {
	res := make([]int, len(nums))
	for i, n := range nums {
		if len(n) > 0 {
			res[i] = Sum(n[1:])
		}
	}
	return res
}
