package main

import "fmt"

func encryptByte(a, b byte) byte {
	a += b
	if a > 90 {
		a -= 26
	}
	return a
}

func encrypt(cleartext, key []byte) []byte {
	keyIndex := 0
	for i := range cleartext {
		cleartext[i] = encryptByte(cleartext[i], key[keyIndex])
		keyIndex++
		if keyIndex >= len(key) {
			keyIndex = 0
		}
	}
	return cleartext
}

func decryptByte(a, b byte) byte {
	a -= b
	if a < 65 {
		a += 26
	}
	return a
}

func decrypt(ciphertext, key []byte) []byte {
	keyIndex := 0
	for i := range ciphertext {
		ciphertext[i] = decryptByte(ciphertext[i], key[keyIndex])
		keyIndex++
		if keyIndex >= len(key) {
			keyIndex = 0
		}
	}
	return ciphertext
}

func main() {

	myKey := []byte("HELLO")
	fmt.Println(string(myKey))

	for i := range myKey {
		myKey[i] -= 64
	}

	// create a cleartext message:
	cleartext := []byte("THISISMYMESSAGE")

	ciphertext := encrypt(cleartext, myKey)
	fmt.Println(string(ciphertext))

	decrypted := decrypt(ciphertext, myKey)
	fmt.Println(string(decrypted))

}
