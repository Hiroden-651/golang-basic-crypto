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

// Return true if byte value is an uppercase ASCII value. Else, return false.
func isUpper(b byte) bool {
	if b >= 65 && b <= 90 {
		return true
	}
	return false
}

// Return true if byte value is a lowercase ASCII value. Else, return false.
func isLower(b byte) bool {
	if b >= 97 && b <= 122 {
		return true
	}
	return false
}

// Functions for encrypting cleartext:

// Rotate byte b by value of byte r. Subtract by 26 if out of bounds.
func rotateUpper(b, r byte) byte {
	b = b + r
	if b > 90 {
		b -= 26
	}
	return b
}

// Rotate byte b by value of byte r. Subtract by 26 if out of bounds
func rotateLower(b, r byte) byte {
	b += r
	if b > 122 {
		b -= 26
	}
	return b
}

// Iterate through byte slice. Perform actions if a byte is an uppercase or lowercase ASCII value.
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

// Inverse of rotateUpper function.
func rotateUpperDec(b, r byte) byte {
	b -= r
	if b < 65 {
		b += 26
	}
	return b
}

// Inverse of rotateLower function.
func rotateLowerDec(b, r byte) byte {
	b -= r
	if b < 97 {
		b += 26
	}
	return b
}

// Iterate through byte slice. Perform actions if a byte is an uppercase or lowercase ASCII value.
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

// Sanity check function used to compare decrypted cipher text to original plaintext. Returns true if plaintext and ciphertext are the same.
func sanityCheck(plaintext, ciphertext []byte) bool {
	if bytes.Compare(plaintext, ciphertext) == 0 {
		return true
	}
	return false
}

// Correct a rotation value if larger than 26.
func createRotationValue(r int) byte {
	return byte(r % 26)
}

func main() {

	// Defining
	rotationValue := flag.Int("r", 0, "Rotation value used to determine how many units to rotate alphabet.")
	textString := flag.String("s", "", "Plaintext string to be encrypted.")

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
