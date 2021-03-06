package main

import (
	"flag"
	"fmt"
	"os"
)

// Create a byte slice encryption/decryption key from a string argument.
func createKey(s string) []byte {
	myKey := []byte(s)
	for i := range myKey {
		myKey[i] -= 65
	}
	return myKey
}

// Return true if a byte value is an upper case letter(ASCII).
func isLetter(a byte) bool {
	if a >= 65 && a <= 95 {
		return true
	}
	return false
}

// Modify original byte value 'a' using a key byte value 'b'. Return 'a'.
func encryptByte(a, b byte) byte {
	a += b
	if a > 90 {
		a -= 26
	}
	return a
}

// Encrypt cleartext byte slice using key byte slice. Return the encrypted message.
func encrypt(cleartext, key []byte) []byte {
	keyIndex := 0
	for i := range cleartext {
		if isLetter(cleartext[i]) {
			cleartext[i] = encryptByte(cleartext[i], key[keyIndex])
			keyIndex++
			if keyIndex >= len(key) {
				keyIndex = 0
			}
		}
	}
	return cleartext
}

// Modify byte 'a' by value of byte 'b'. Return 'a'.
func decryptByte(a, b byte) byte {
	a -= b
	if a < 65 {
		a += 26
	}
	return a
}

// Decrypt ciphertext using the same key used for encrypting the cleartext. Return the deciphered text.
func decrypt(ciphertext, key []byte) []byte {
	keyIndex := 0
	for i := range ciphertext {
		if isLetter(ciphertext[i]) {
			ciphertext[i] = decryptByte(ciphertext[i], key[keyIndex])
			keyIndex++
			if keyIndex >= len(key) {
				keyIndex = 0
			}
		}
	}
	return ciphertext
}

func main() {

	// Declaring -c and -k flags for cleartext and key respectively.
	clearTextString := flag.String("c", "", "Cleartext string to encrypt.")
	keyString := flag.String("k", "", "String for encryption/decryption.")
	flag.Parse()

	// Exit program if either cleartext or key remain blank.
	if *clearTextString == "" {
		fmt.Println("Error! No cleartext provided!")
		os.Exit(1)
	}

	if *keyString == "" {
		fmt.Println("Error! No key provided!")
		os.Exit(1)
	}

	fmt.Printf("Cipher key:\t\t%s\n", *keyString)
	myKey := createKey(*keyString)

	// create a cleartext message:
	cleartext := []byte(*clearTextString)
	fmt.Printf("Cleartext message:\t%s\n", string(cleartext))

	ciphertext := encrypt(cleartext, myKey)
	fmt.Printf("Ciphertext message:\t%s\n", string(ciphertext))

	decrypted := decrypt(ciphertext, myKey)
	fmt.Printf("Decrypted message:\t%s\n", string(decrypted))

}
