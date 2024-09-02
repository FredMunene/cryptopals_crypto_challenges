// package main

// import "encoding/hex"

// func main(){

// 	// hex encoded string
// 	str := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

// 	//  the string has beenXOR'd against a single character. Find the key, decrypt the message
// 	//  You can do this by hand. But don't: write code to do it for you.
// 	//  How? Devise some method for "scoring" a piece of English plaintext.
// 	//  Character frequency is a good metric. Evaluate each output and choose the one with the best score. 

// 	byt, err := hex.DecodeString(str)
// 	if err != nil{
// 		println(err)
// 	}

// 	for _, b := range byt{
// 		print(b,",")
// 	}
// 	println()
// }

package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

var englishFreq = map[rune]float64{
	'a': 0.0651738, 'b': 0.0124248, 'c': 0.0217339, 'd': 0.0349835, 'e': 0.1041442, 'f': 0.0197881,
	'g': 0.0158610, 'h': 0.0492888, 'i': 0.0558094, 'j': 0.0009033, 'k': 0.0050529, 'l': 0.0331490,
	'm': 0.0202124, 'n': 0.0564513, 'o': 0.0596302, 'p': 0.0137645, 'q': 0.0008606, 'r': 0.0497563,
	's': 0.0515760, 't': 0.0729357, 'u': 0.0225134, 'v': 0.0082903, 'w': 0.0171272, 'x': 0.0013692,
	'y': 0.0145984, 'z': 0.0007836, ' ': 0.1918182,
}

// Convert the hex string to a byte slice
func hexToBytes(hexStr string) ([]byte, error) {
	return hex.DecodeString(hexStr)
}

// Score a piece of plaintext based on English letter frequency
func scoreText(text string) float64 {
	var score float64
	for _, char := range text {
		score += englishFreq[char]
	}
	return score
}

// Decrypt the message by XORing with each possible byte
func decryptSingleByteXOR(cipherBytes []byte) (byte, string) {
	var bestScore float64
	var bestKey byte
	var bestPlaintext string

	for key := 0; key < 256; key++ {
		var plaintext strings.Builder
		for _, b := range cipherBytes {
			plaintext.WriteByte(b ^ byte(key))
		}
		decodedText := plaintext.String()
		score := scoreText(strings.ToLower(decodedText))
		if score > bestScore {
			bestScore = score
			bestKey = byte(key)
			bestPlaintext = decodedText
		}
	}
	return bestKey, bestPlaintext
}

func main() {
	hexStr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// Convert the hex string to bytes
	cipherBytes, err := hexToBytes(hexStr)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return
	}

	// Decrypt the cipher text
	key, plaintext := decryptSingleByteXOR(cipherBytes)
	fmt.Printf("Key: %c (ASCII %d)\n", key, key)
	fmt.Println("Decrypted message:", plaintext)
}
