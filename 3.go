// 3. Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating characters.

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

func lengthOfLongestSubstring(s string) int {
    var strArr []string
    charMap := map[string][]int{}
    for index, value := range s {
        strArr = append(strArr, string(value))
        charMap[string(value)] = append(charMap[string(value)], index)
    }
    fp, sp, count, max, repeat := 0, 0, 0, 0, 0
    for sp < len(strArr) {
		mapCheck := charMap[strArr[sp]]
		for i := 0; i < len(mapCheck); i++ {
			if mapCheck[i] >= fp && mapCheck[i] < sp {
				repeat = 1
				break
			} 
		}
		if repeat == 1 {
			fp += 1 	
			sp = fp	+ 1	
			count = 1
			repeat = 0
			continue
		}
		count += 1
		if count > max {
			max = count
		}
		sp += 1
    }
    return max