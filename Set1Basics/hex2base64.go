package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// Convert hex to base64
	// str := "aaa"

	str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	// answer =SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
	// hint : Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.

	// base64:: A -> Z  0-25 | a -> z 26-51 | 0 - 9 52 - 61 | + 62 | / 63

	// hex >> bytes
	// bytes >> base64

	byt,err := hex.DecodeString(str)
	if err != nil{
		println(err)
	}
	base := base64.RawStdEncoding.EncodeToString(byt)

	fmt.Println(base)

// 	//step 2
// 	Convert the hexadecimal string to binary.
// Group the binary string into 6-bit chunks.
// Convert each 6-bit chunk to its corresponding base64 character.


}
