package main

import (
	"encoding/hex"
)

func main() {

	// Encrypt it, under the key "ICE", using repeating-key XOR.
	// In repeating-key XOR, you'll sequentially apply each byte of the key; the first byte of plaintext will be XOR'd against I,
	// the next C, the next E, then I again for the 4th byte, and so on.
	// It should come out to:
	// 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272
	// a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f
	words := `Burning 'em, if you ain't quick and nimble
	I go crazy when I hear a cymbal`

	// hexString := strToHex(words)
	plaintext := []byte(words)

	key := "ICE"
	keySlice := []byte(key)


	var res []byte

	// str := "ICE"
	// strSlice := strToHex(str)


	// for i := range hexString{
	// 	hex3 := hexString[i] ^ strSlice[i%len(strSlice)]
	// 	res = append(res, hex3)
	// }

	for i := 0; i < len(plaintext); i++{
		encryptedByte := plaintext[i]^keySlice[i%len(keySlice)]
		res = append(res, encryptedByte)
	}
	println(hex.EncodeToString(res))
}

// func strToHex(str string) string  {
	
// 	byteSlice := []byte{}
// 	for _, char := range str {
// 		byteSlice = append(byteSlice, byte(char))
// 	}
// 	return hex.EncodeToString(byteSlice)
// }