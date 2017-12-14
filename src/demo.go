package main

import (
	"bliss"
	"fmt"
	"params"
	"sampler"
	"encoding/hex"
)

func main() {
	version := params.BLISS_B_0
	seed := []uint8{
		0x11, 0x1a, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7,
		0x10, 0x1b, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x20, 0x21, 0x2c, 0x23, 0x2d, 0x25, 0x26, 0x27,
		0x30, 0x31, 0x32, 0x33, 0x3e, 0x35, 0x36, 0xa7,
		0x40, 0x41, 0x4a, 0x4f, 0x44, 0x45, 0xa6, 0x47,
		0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0xc6, 0x57,
		0x60, 0x61, 0x22, 0xa3, 0x64, 0xd5, 0x66, 0xc7,
		0x70, 0xa1, 0x72, 0xe3, 0xd4, 0x75, 0x76, 0x77,
	}
	msg := "Hello world"

	entropy, err := sampler.NewEntropy(seed)
	if err != nil {
		fmt.Printf("Error in creating entropy: %s\n", err.Error())
		return
	}

	key, err := bliss.GeneratePrivateKey(version, entropy)
	if err != nil {
		fmt.Printf("Error in generating private key: %s\n", err.Error())
		return
	} else {
		// fmt.Printf("Private Key: %s\n", key.String())
		fmt.Printf("Private Key: %s\n", hex.EncodeToString(key.Encode()))
	}

	pub := key.PublicKey()
	// fmt.Printf("Public Key: %s\n", pub.String())
	fmt.Printf("Public Key: %s\n", hex.EncodeToString(pub.Encode()))

	sig, err := key.Sign([]byte(msg), entropy)
	if err != nil {
		fmt.Printf("Error in signing: %s\n", err.Error())
		return
	} else {
		fmt.Printf("Signature: %s\n", sig.String())
	}

	res, err := pub.Verify([]byte(msg), sig)
	if res {
		fmt.Printf("Verified!\n")
	} else {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
