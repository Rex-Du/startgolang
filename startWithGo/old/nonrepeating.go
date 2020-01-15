package main

import "fmt"

func lenOfNonRepeatingSubStr(s string) int{
	lastOccurred := make(map[byte]int)
	start := 0
	maxLen := 0
	for i, ch := range []byte(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLen
}

func main(){
	fmt.Println(lenOfNonRepeatingSubStr("duqing"))
	fmt.Println(lenOfNonRepeatingSubStr("lifan"))
	fmt.Println(lenOfNonRepeatingSubStr("guoguo"))
}