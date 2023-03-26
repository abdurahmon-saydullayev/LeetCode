package main

import "fmt"

func main() {
 
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
    var pre rune
    var result = 0
    
    for _, value := range s {
        switch value {
        case 'I':
            result += 1
        case 'V':
            if pre == 'I' {
                result += 3
            } else {
                result += 5
            }
        case 'X':
            if pre == 'I' {
                result += 8
            } else {
                result += 10
            }
        case 'L':
            if pre == 'X' {
                result += 30
            } else {
                result += 50
            }
        case 'C':
            if pre == 'X' {
                result += 80
            } else {
                result += 100
            }
        case 'D':
            if pre == 'C' {
                result += 300
            } else {
                result += 500
            }
        case 'M':
            if pre == 'C' {
                result += 800
            } else {
                result += 1000
            }
        }
        
        pre = value
    }
    
    return result
}