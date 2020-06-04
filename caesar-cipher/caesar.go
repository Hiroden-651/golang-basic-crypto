package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

// Functions for checking ascii bounds of bytes:

func isUpper(b byte) bool {
	if b >= 65 && b <= 90 {
		return true
	}
	return false
}

func isLower(b byte) bool {
	if b >= 97 && b <= 122 {
		return true
	}
	return false
}

// Functions for encrypting cleartext:

func rotateUpper(b, r byte) byte {
	b = b + r
	if b > 90 {
		b -= 26
	}
	return b
}

func rotateLower(b, r byte) byte {
	b += r
	if b > 122 {
		b -= 26
	}
	return b
}

func encrypt(data []byte, r byte) []byte {
	for index := range data {
		if isUpper(data[index]) {
			data[index] = rotateUpper(data[index], r)
		}
		if isLower(data[index]) {
			data[index] = rotateLower(data[index], r)
		}
	}
	return data
}

// Functions for decrypting ciphertext.

func rotateUpperDec(b, r byte) byte {
	b -= r
	if b < 65 {
		b += 26
	}
	return b
}

func rotateLowerDec(b, r byte) byte {
	b -= r
	if b < 97 {
		b += 26
	}
	return b
}

func decrypt(data []byte, r byte) []byte {
	for index := range data {
		if isUpper(data[index]) {
			data[index] = rotateUpperDec(data[index], r)
		}
		if isLower(data[index]) {
			data[index] = rotateLowerDec(data[index], r)
		}
	}
	return data
}

// Sanity check function used to compare decrypted cipher text to original plaintext. Returns true if

func sanityCheck(plaintext, ciphertext []byte) bool {
	if bytes.Compare(plaintext, ciphertext) == 0 {
		return true
	}
	return false
}

func createRotationValue(r int) byte {
	return byte(r % 26)
}

func main() {

	rotationValue := flag.Int("r", 0, "Rotation value used to determine how many units to rotate alphabet.")
	textString := flag.String("s", "", "Text string to be encrypted.")

	flag.Parse()

	if *rotationValue == 0 || *textString == "" {
		fmt.Println("Error! No plaintext provided or bad rotationValue.")
		os.Exit(1)
	}

	fmt.Printf("Caeser Cipher with %d rotation(s).\n", *rotationValue%26)
	fmt.Printf("Plaintext:\t%s\n", *textString)

	rotate := byte(*rotationValue % 26)
	var textCopy = make([]byte, len(*textString))
	copy(textCopy, *textString)

	textCopy = encrypt(textCopy, rotate)
	fmt.Printf("Ciphertext:\t%s\n\n", string(textCopy))

}
