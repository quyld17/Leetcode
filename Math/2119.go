2119. A Number After a Double Reversal

Reversing an integer means to reverse all its digits.

For example, reversing 2021 gives 1202. Reversing 12300 gives 321 as the leading zeros are not retained.
Given an integer num, reverse num to get reversed1, then reverse reversed1 to get reversed2. 
Return true if reversed2 equals num. Otherwise return false.

Input: num = 526
Output: true
Explanation: Reverse num to get 625, then reverse 625 to get 526, which equals num.

func isSameAfterReversals(num int) bool {
    beforeReversed := num
    reversed1, reversed2 := 0, 0
    for num != 0 {
        remainder := num % 10
        reversed1 = reversed1 * 10 + remainder
        num = num / 10
    }
    for reversed1 != 0 {
        remainder := reversed1 % 10
        reversed2 = reversed2 * 10 + remainder
        reversed1 = reversed1 / 10
    }
    if reversed2 != beforeReversed {
        return false
    }
    return true
}