package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	//  a function that takes two equal-length buffers and produces their XOR combination.
	str1 := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	str2 := "4E"

	xor(str1, str2)
}

func xor(str1, str2 string) {
	// hex decoding
	hex1, err := hex.DecodeString(str1)
	if err != nil {
		fmt.Println(err)
	}
	hex2, err := hex.DecodeString(str2)
	if err != nil {
		fmt.Println(err)
	}
	println(hex2[0])
	// check lengths are equal
	// XOR : 1 ^ 1 = 0 : 1 ^ 0 = 1

	// if len(hex1) != len(hex2) {
	// 	fmt.Println("Different lengths")
	// 	return
	// }

	res := []byte {}
	
	for i:= 0; i < len(hex1);i++{

		hex3 := hex1[i]^hex2[0]
		res = append(res, hex3)
		
	}

	fmt.Println(hex.EncodeToString(res))
}
