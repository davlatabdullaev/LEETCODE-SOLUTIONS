package monotonicarray

func IsMonotonic(nums []int) bool {

	plus, minus := true, true

    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] {
            plus = false
        }
        if nums[i-1] < nums[i] {
            minus = false
        }
    }

    return plus || minus
}
