# Caesar Cipher in Go

The Caesar Cipher(also called the shift cipher) is arguably the most basic form of encryption. Using this cipher involves mapping characters within an alphabet to other characters by shifting the locations of each letter by a fixed number of positions. For example, in the English alphabet of 26 letters, shifting left by 3 letters would map 'A' to 'D', 'B' to 'E', 'C' to 'F', and so forth. It's fairly obvious that the simplicity of this encryption scheme makes it unsuitable for any kind of modern crypto application, but it is useful as a tool for illustrating crypto concepts.

## Implementation in Go

This program takes two arguments: a plaintext string value and an integer value for shifting the alphabet. It outputs a comparison of the plaintext against a cyphertext created using the rotation value.

The implementation of this cipher in Go takes advantage of the language's ability to easily convert between strings and byte arrays. Since byte arrays contain a character's ASCII value, shifting the value of letters requires only basic arithmetic and bounds checking(if not in range, subtract by 26). This program accounts for both uppercase and lowercase ASCII values, ranged 65-90 and 97-122 respectively. All other ASCII characters remain unchanged.

For the English alphabet there are 25 distinct rotation positions that encrypt the plaintext(rotating by 26 is equivalent to 0 rotations and gives the original plaintext.) The modulus operator is used to correct any rotation value that is more than the range of letters(a rotation value of 27 is equivalent to a rotation value of 1).