// 383. Ransom Note

// Given two strings ransomNote and magazine, return true if 
// ransomNote can be constructed by using the letters from magazine and false otherwise.

// Each letter in magazine can only be used once in ransomNote.

// Input: ransomNote = "a", magazine = "b"
// Output: false


func canConstruct(ransomNote string, magazine string) bool {
    magazineMap := map[string]int{}
    for _, val := range magazine {
        magazineMap[string(val)] += 1
    }
    for _, val := range ransomNote {
        v, ok := magazineMap[string(val)]
        if ok && v > 0 {
            magazineMap[string(val)] -= 1
            continue
        }
        return false
    }
    return true
}