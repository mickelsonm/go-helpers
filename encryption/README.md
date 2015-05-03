Encryption Helper
==========

This helper is setup to handle various forms of encryption and decryption mechanisms out there.

AES Encryption Example:
```go
	package main

	import (
		"fmt"
		"github.com/mickelsonm/go-helpers/encryption"
	)

	func main() {
		//Key needs to have a length of 16, 24, or 32 bytes to select
		//AES-128, AES-192, or AES-256 (per Golang documentation).
		CIPHER_KEY := []byte("DecryptMagic2015")
		msg := "This is a top secret message."

		encrypted, err := encryption.AES_Encrypt(CIPHER_KEY, msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Encrypted: %s\n", encrypted)

		decrypted, err := encryption.AES_Decrypt(CIPHER_KEY, encrypted)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Decrypted: %s\n", decrypted)
	}
```
