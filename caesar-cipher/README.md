# Caesar Cipher in Go

The Caesar Cipher(also called the shift cipher) is arguably the most basic form of encryption. Using this cipher involves mapping characters within an alphabet to other characters by shifting the locations of each letter by a fixed number of positions. For example, in the English alphabet of 26 letters, shifting left by 3 letters would map 'A' to 'D', 'B' to 'E', 'C' to 'F', and so forth. It's fairly obvious that the simplicity of this encryption scheme makes it unsuitable for any kind of modern crypto application, but it's useful as a tool for illustrating crypto concepts.

## Implementation in Go

This program takes two arguments: a plaintext string value and an integer value for shifting the alphabet. It outputs a comparison of the plaintext against a cyphertext created using the rotation value.

The implementation of this cipher in Go takes advantage of the language's ability to easily convert between strings and byte arrays. Since byte arrays contain a character's ASCII value, shifting the value of letters requires only basic arithmetic and bounds checking(if not in range, subtract by 26). This program accounts for both uppercase and lowercase ASCII values, ranged 65-90 and 97-122 respectively. All other ASCII characters remain unchanged.

For the English alphabet there are 25 distinct rotation positions that encrypt the plaintext(rotating by 26 is equivalent to 0 rotations and gives the original plaintext.) The modulus operator is used to correct any rotation value that is more than the range of letters(a rotation value of 27 is equivalent to a rotation value of 1).

## Usage

The simplist way to run this program in BASH is to use Go's run command which compiles and runs the program all at once.

An example of the program in action: applying a rotation value of 3 to the input cleartext "ABCDabcd".
```bash
go run caesar.go -s "ABCDabcd" -r 3
```

The output of this program is as follows:
```bash
Caeser Cipher with 3 rotation(s).
Plaintext:	ABCDabcd
Ciphertext:	DEFGdefg
```


Another example using a sentence and rotation of 9:

Input:
```bash
go run caesar.go -s "I am the greatest guitar player ever to have lived!" -r 9
```
Output:
```bash
Caeser Cipher with 9 rotation(s).
Plaintext:	I am the greatest guitar player ever to have lived!
Ciphertext:	R jv cqn panjcnbc pdrcja yujhna nena cx qjen urenm!
```
The above example shows that whitespace and punctuation characters are not affected by the cipher rotation.