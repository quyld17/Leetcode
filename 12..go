// 12. Integer to Roman

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. 
// Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

// I can be placed before V (5) and X (10) to make 4 and 9. 
// X can be placed before L (50) and C (100) to make 40 and 90. 
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given an integer, convert it to a roman numeral.

// Input: num = 3
// Output: "III"
// Explanation: 3 is represented as 3 ones.

func intToRoman(num int) string {
    var hangDonVi, hangChuc, hangTram, hangNghin, output string
    var unit, ten, hundred, thousand int
    if num < 10 {
        hangDonVi = units(num)
        output = hangDonVi
    } else if num >= 10 && num < 100 {
        unit = num % 10
        ten = num - unit
        hangDonVi = units(unit)
        hangChuc = tens(ten)
        output = hangChuc + hangDonVi
    } else if num >= 100 && num < 1000 {
        ten = num % 100
        hundred = num - ten
        unit = ten % 10
        ten = ten - unit
        hangDonVi = units(unit)
        hangChuc = tens(ten)
        hangTram = hundreds(hundred)
        output = hangTram + hangChuc + hangDonVi
    } else {
        hundred = num % 1000
        thousand = num - hundred
        ten = hundred % 100
        hundred = hundred - ten
        unit = ten % 10
        ten = ten - unit 
        hangDonVi = units(unit)
        hangChuc = tens(ten)
        hangTram = hundreds(hundred)
        hangNghin = thousands(thousand)
        output = hangNghin + hangTram + hangChuc + hangDonVi
    }
    return output
}

func units(a int) string {
    var str string
    switch {
        case a == 1: str = "I"
        case a == 2: str = "II"
        case a == 3: str = "III"
        case a == 4: str = "IV"
        case a == 5: str = "V"
        case a == 6: str = "VI"
        case a == 7: str = "VII"
        case a == 8: str = "VIII"
        case a == 9: str = "IX"
        default: str = ""
    }
    return str
}

func tens(a int) string {
    var str string
    switch {
        case a == 10: str = "X"
        case a == 20: str = "XX"
        case a == 30: str = "XXX"
        case a == 40: str = "XL"
        case a == 50: str = "L"
        case a == 60: str = "LX"
        case a == 70: str = "LXX"
        case a == 80: str = "LXXX"
        case a == 90: str = "XC"
        default: str = ""
    }
    return str
}

func hundreds(a int) string {
    var str string
    switch {
        case a == 100: str = "C"
        case a == 200: str = "CC"
        case a == 300: str = "CCC"
        case a == 400: str = "CD"
        case a == 500: str = "D"
        case a == 600: str = "DC"
        case a == 700: str = "DCC"
        case a == 800: str = "DCCC"
        case a == 900: str = "CM"
        default: str = ""
    }
    return str
}

func thousands(a int) string {
    var str string
    switch {
        case a == 1000: str = "M"
        case a == 2000: str = "MM"
        case a == 3000: str = "MMM"
        default: str = ""
    }
    return str
}
