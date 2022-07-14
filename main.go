package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	whatIsIt = string(sd)

	fmt.Println(reverse(whatIsIt))

}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
