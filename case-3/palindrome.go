package main

import (
	"fmt"
	"strings"
)

func main()  {
	sample := "mAlAyAlaMa"
    // sample  := "racecar"
    // sample  := "hello world"
	result := isCircularPalindrome(strings.ToLower(sample))
	fmt.Println(result)
}

func isCircularPalindrome(param string) bool {
    for i := 0; i< len(param)/2; i++ {
        if param[i] != param[len(param)-i-1]{
          return false
        }
      }
    
    return true
}