package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("hex.txt")
	if err != nil {
		println(err)
		return
	}

	hexData := strings.Split(string(data), "\n")

	findEncryptedString(hexData)
}

var englishFreq = map[rune]float64{
	'a': 0.0651738, 'b': 0.0124248, 'c': 0.0217339, 'd': 0.0349835, 'e': 0.1041442, 'f': 0.0197881,
	'g': 0.0158610, 'h': 0.0492888, 'i': 0.0558094, 'j': 0.0009033, 'k': 0.0050529, 'l': 0.0331490,
	'm': 0.0202124, 'n': 0.0564513, 'o': 0.0596302, 'p': 0.0137645, 'q': 0.0008606, 'r': 0.0497563,
	's': 0.0515760, 't': 0.0729357, 'u': 0.0225134, 'v': 0.0082903, 'w': 0.0171272, 'x': 0.0013692,
	'y': 0.0145984, 'z': 0.0007836, ' ': 0.1918182,
}

func singleCharXOR(input []byte, key byte) []byte {
	output := make([]byte, len(input))
	for i := range input {
		output[i] = input[i] ^ key
	}
	return output
}

func scoreText(text []byte) float64 {
	score := 0.0
	for _, b := range text {
		score += englishFreq[unicode.ToLower(rune(b))]
	}
	return score
}

func findEncryptedString(hexStrings []string) {
	bestScore := 0.0
	var bestDecrypted string
	var bestHex string

	for _, hexString := range hexStrings {
		encryptedBytes, err := hex.DecodeString(hexString)
		if err != nil {
			fmt.Println("Error decoding hex string:", err)
			continue
		}

		for key := 0; key < 256; key++ {
			decryptedBytes := singleCharXOR(encryptedBytes, byte(key))
			if isReadable(decryptedBytes) {
				score := scoreText(decryptedBytes)
				if score > bestScore {
					bestScore = score
					bestDecrypted = string(decryptedBytes)
					bestHex = hexString
				}
			}
		}
	}

	fmt.Printf("Best Decrypted Text: %s\n", bestDecrypted)
	fmt.Printf("Original Hex: %s\n", bestHex)
}

func isReadable(text []byte) bool {
	for _, b := range text {
		if !unicode.IsPrint(rune(b)) && !unicode.IsSpace(rune(b)) {
			return false
		}
	}
	return true
}
