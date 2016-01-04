package encryption

/*
 * Generates a private/public key pair in PEM format (not Certificate)
 *
 * If we were to generate a key it would be as follows:
 *
 * > openssl genrsa -out server.key 2048
 *
 * The generated private key can be parsed with openssl as follows:
 * > openssl rsa -in key.pem -text
 *
 * The generated public key can be parsed as follows:
 * > openssl rsa -pubin -in pub.pem -text
 */

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
)

// RSAKeyPair ...
type RSAKeyPair struct {
	PrivateKey string
	PublicKey  string
}

// GenerateRSAOptions ...
type GenerateRSAOptions struct {
	Bits       int
	Encryption *EncryptOptions
}

// EncryptOptions ...
type EncryptOptions struct {
	Password  string
	PEMCipher x509.PEMCipher
}

// GenerateRSAKeyPair ... generates an RSA private and public key values
func GenerateRSAKeyPair(opts GenerateRSAOptions) (*RSAKeyPair, error) {
	//creates the private key
	privateKey, err := rsa.GenerateKey(rand.Reader, opts.Bits)
	if err != nil {
		return nil, fmt.Errorf("error generating private key: %s\n", err)
	}

	//validates the private key
	err = privateKey.Validate()
	if err != nil {
		return nil, fmt.Errorf("error validating private key: %s\n", err)
	}

	// sets up the PEM block for private key
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(privateKey),
	}

	//check to see if we are applying encryption to this key
	if opts.Encryption != nil {
		//check to make sure we have a password specified
		pass := strings.TrimSpace(opts.Encryption.Password)
		if pass == "" {
			return nil, fmt.Errorf("%s", "need a password!")
		}
		//check to make sure we're using a supported PEMCipher
		encCipher := opts.Encryption.PEMCipher
		if encCipher != x509.PEMCipherDES &&
			encCipher != x509.PEMCipher3DES &&
			encCipher != x509.PEMCipherAES128 &&
			encCipher != x509.PEMCipherAES192 &&
			encCipher != x509.PEMCipherAES256 {
			return nil, fmt.Errorf("%s", "invalid PEMCipher")
		}
		//encrypt the private key block
		encBlock, err := x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", privateKeyBlock.Bytes, []byte(pass), encCipher)
		if err != nil {
			return nil, fmt.Errorf("error encrypting pirvate key: %s\n", err)
		}
		//replaces the starting one with the one we encrypted
		privateKeyBlock = *encBlock
	}

	// serializes the public key in a DER-encoded PKIX format (see docs for more)
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("error setting up public key: %s\n", err)
	}

	// sets up the PEM block for public key
	publicKeyBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   publicKeyBytes,
	}

	//returns the created key pair
	return &RSAKeyPair{
		PrivateKey: string(pem.EncodeToMemory(&privateKeyBlock)),
		PublicKey:  string(pem.EncodeToMemory(&publicKeyBlock)),
	}, nil
}

// GenerateCertificate ...
func GenerateCertificate(keypair RSAKeyPair) {
	fmt.Println("create the cert in here")
}

/**
func main() {
	kp, err := GenerateRSAKeyPair(GenerateRSAOptions{
		Bits: 2048,
		Encryption: &EncryptOptions{
			Password:  "gopher",
			PEMCipher: x509.PEMCipherAES256,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(kp.PrivateKey)
	fmt.Println(kp.PublicKey)
}
**/
