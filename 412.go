// 412. Fizz Buzz

// Given an integer n, return a string array answer (1-indexed) where:

// answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
// answer[i] == "Fizz" if i is divisible by 3.
// answer[i] == "Buzz" if i is divisible by 5.
// answer[i] == i (as a string) if none of the above conditions are true.

// Input: n = 3
// Output: ["1","2","Fizz"]

func fizzBuzz(n int) []string {
    var answer []string
    for k := 1; k <= n; k++ {
        switch {
            case k % 15 == 0:
                answer = append(answer, "FizzBuzz")
            case k % 3 == 0:
                answer = append(answer, "Fizz")
            case k % 5 == 0:
                answer = append(answer, "Buzz")
            default: 
                answer = append(answer, strconv.Itoa(k))
        }
    }
    return answer
}