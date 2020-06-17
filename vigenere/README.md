# Vigenere Cipher
This cipher is similar to the Caesar Cipher except that instead of shifting all letters by a single value, the number of shifts per letter are determined by a specified key. This key is used for both encrypting and decrypting a cleartext message.

## Example
Let's say we want to encrypt the following cleartext message: 
```bash
HELLOWORLD
```
We'll choose a simple key called "KEY" which is 3 letters.

Because the key is shorter than the cleartext, we repeat the key's pattern until it's the same length as the cleartext:
```bash
HELLOWORLD
KEYKEYKEYK
```
Each letter in the cleartext will be rotated based on the numeric postion of it's corresponding key letter. H is rotated by 11 units, E by 5, L by 25, L by 11, and so forth.

Encrypting with this key gives the resulting ciphertext:
```bash
SJLWTWZWLO
```

To decrypt the cipher text back into cleartext we use the same key to rotate the letter back to their original positions.

```bash
SJLWTWZWLO
KEYKEYKEYK
```
...becomes...

```bash
HELLOWORLD
```

## Go Implementation
This Go program performs a simple version of the Vigenere Cipher using cleartexts and keys containing exclusively uppercase letters. Rotation for encryption and decryption is done based on byte values(ASCII values between 65 and 90).

To run the program in BASH, we use Go's run command to type the following:
```bash
go run vigenere.go -c CLEARTEXT -k KEY
```
This program uses two flags c and k for cleartext and key respectively. If either values are missing the program will exit.

When given proper parameters, the program produces the following output:
```bash
go run vigenere.go -c THISISASECRETMESSAGE -k TELLNOONE

Cipher key:		TELLNOONE
Cleartext message:	THISISASECRETMESSAGE
Ciphertext message:	MLTDVGOFIVVPEZSGFEZI
Decrypted message:	THISISASECRETMESSAGE

```