// 925. Long Pressed Name

// Your friend is typing his name into a keyboard. Sometimes, when typing a character c, 
// the key might get long pressed, and the character will be typed 1 or more times.

// You examine the typed characters of the keyboard. 
// Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.

// Input: name = "alex", typed = "aaleex"
// Output: true
// Explanation: 'a' and 'e' in 'alex' were long pressed.


func isLongPressedName(name string, typed string) bool { 
    p1, p2 := 0, 0
    if string(name[p1]) == string(typed[p2]) {
        p1 += 1
        p2 += 1
    } else {
        return false
    }
    for p2 < len(typed) {
        if p1 == len(name) {
            if string(typed[p2]) == string(name[p1-1]) {
                p2 += 1
                continue
            }
            return false
        }
        if string(typed[p2]) == string(name[p1]) {
            p1 += 1
            p2 += 1
            continue
        }
        if string(typed[p2]) == string(name[p1-1]) {
            p2 += 1
            continue
        }
        return false
    }
    if p1 < len(name) || p2 < len(typed) {
        return false
    }
    return true
}