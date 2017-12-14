package main

import (
	"bliss"
	"fmt"
	"params"
	"sampler"
)

func main() {
	version := params.BLISS_B_0
	seed := []uint8{
		0, 1, 2, 3, 4, 5, 6, 7,
		10, 11, 12, 13, 14, 15, 16, 17,
		20, 21, 22, 23, 24, 25, 226, 27,
		30, 31, 32, 33, 34, 35, 36, 37,
		40, 41, 42, 43, 44, 45, 46, 47,
		50, 51, 52, 53, 54, 55, 56, 57,
		60, 61, 62, 63, 64, 65, 66, 67,
		70, 71, 72, 73, 74, 75, 76, 77,
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
		fmt.Printf("Private Key: %s\n", key.String())
	}

	pub := key.PublicKey()
	fmt.Printf("Public Key: %s\n", pub.String())

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
