package main

import "fmt"

func main() {
	var arr = [5]string{"I", "am", "stupid", "and", "weak"} //初始化
	for k, v := range arr {
		fmt.Println(k, v)
	}
	fmt.Println(arr)
	arr[2], arr[4] = "smart", "strong"
	fmt.Println(arr)

}
