// 1491. Average Salary Excluding the Minimum and Maximum Salary

// You are given an array of unique integers salary where salary[i] is the salary of the ith employee.

// Return the average salary of employees excluding the minimum and maximum salary. 
// Answers within 10^-5 of the actual answer will be accepted.

// Input: salary = [4000,3000,1000,2000]
// Output: 2500.00000
// Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
// Average salary excluding minimum and maximum salary is (2000+3000) / 2 = 2500


func average(salary []int) float64 {
    sort.Ints(salary)
    var count, total float64
    for i := 1; i < len(salary)-1; i++ {
        total += float64(salary[i])
        count += 1
    }
    return total/count
}