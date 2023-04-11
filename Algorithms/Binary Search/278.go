278. First Bad Version

You are a product manager and currently leading a team to develop a new product. 
Unfortunately, the latest version of your product fails the quality check. 
Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. 
You should minimize the number of calls to the API.

Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.


/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */


//Method 1: Start from the begining and the end
func firstBadVersion(n int) int {
    l, r := 1, n
    for l <= r {
        if isBadVersion(l) == true {
            return l
        }
        if isBadVersion(r) == false {
            return r+1
        }
        l += 1
        r -= 1
    }
    return l
}


//Method 2: Start from the begining
func firstBadVersion(n int) int {
    for i := 1; i <= n; i++ {
        if isBadVersion(i) == true {
            return i
        }
    }
    return 0
}


//Method 3: Start from the middle
func firstBadVersion(n int) int {
    i := n/2
    if isBadVersion(i) == true {
        for i > 0 {
            if isBadVersion(i) == false {
                return i+1
            }
            i -= 1 
        }
        return i+1
    } else {
        for i <= n {
            if isBadVersion(i) == true {
                return i
            }
            i += 1
        }
    }
    return 0
}