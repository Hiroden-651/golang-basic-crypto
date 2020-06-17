# Vigenere Cipher

This cipher is similar to the Caesar Cipher except that instead of shifting all letters by a single value, the number of shifts per letter are determined by a specified key. This key is used for both encrypting and decrypting a plaintext message.

### Example
Let's say we want to encrypt the following plaintext message: 
```bash
HELLOWORLD
```
We'll choose a simple key called "KEY" which is 3 letters.

Because the key is shorter than the plaintext, we repeat the key's pattern until it's the same length as the plaintext:
```bash
HELLOWORLD
KEYKEYKEYK
```
Each letter in the plaintext will be rotated based on the numeric postion of it's corresponding key letter. H is rotated by 11 units, E by 5, L by 25, L by 11, and so forth.

Encrypting with this key gives the resulting ciphertext:
```bash
SJLWTWZWLO
```

To decrypt the cipher text back into plaintext we use the same key to rotate the letter back to their original positions.

```bash
SJLWTWZWLO
KEYKEYKEYK
```
...becomes...

```bash
HELLOWORLD
```

### Go Implementation