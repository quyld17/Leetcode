// 989. Add to Array-Form of Integer

// The array-form of an integer num is an array representing its digits in left to right order.

// For example, for num = 1321, the array form is [1,3,2,1].
// Given num, the array-form of an integer, and an integer k, return the array-form of the integer num + k.

// Input: num = [1,2,0,0], k = 34
// Output: [1,2,3,4]
// Explanation: 1200 + 34 = 1234

func addToArrayForm(num []int, k int) []int {
	for i := 0; i < len(num)/2; i++ {
		swap := num[i]
		num[i] = num[len(num)-i-1]
		num[len(num)-i-1] = swap
	}

    arrK := []int{}
    for k > 0 {
        remainder := k % 10
        k = k / 10
        arrK = append(arrK, remainder)
    }

	pointer, plusOne := 0, 0
	if len(num) >= len(arrK) {
		for pointer < len(num) {
			if pointer < len(arrK) {
				num[pointer] = num[pointer] + arrK[pointer]
			}
			if plusOne == 1 {
				num[pointer] += 1
				plusOne = 0
			}
			if num[pointer] >= 10 {
				num[pointer] = num[pointer] - 10
				plusOne = 1
			}
			pointer += 1
		}
        if plusOne == 1 {
            num = append(num, 1)
        }
	} else {
        for pointer < len(arrK) {
			if pointer < len(num) {
				arrK[pointer] = arrK[pointer] + num[pointer]
			}
			if plusOne == 1 {
				arrK[pointer] += 1
				plusOne = 0
			}
			if arrK[pointer] >= 10 {
				arrK[pointer] = arrK[pointer] - 10
				plusOne = 1
			}
			pointer += 1
		}
        if plusOne == 1 {
            arrK = append(arrK, 1)
        }
        num = arrK
    }

	for i := 0; i < len(num)/2; i++ {
		swap := num[i]
		num[i] = num[len(num)-i-1]
		num[len(num)-i-1] = swap
	}
	return num
}