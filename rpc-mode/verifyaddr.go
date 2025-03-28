package main

import (
	"fmt"
	"github.com/gagliardetto/solana-go"
)

func isValidPublicKey(pubkeyStr string) bool {
	_, err := solana.PublicKeyFromBase58(pubkeyStr)
	return err == nil
}

func main() {
	// Valid public key
	validKey := "Eyyr6FB3rwVc5ApfWjHknCasXu7bp3rPhGKhScUHLTGg"
	fmt.Println("Valid key:", isValidPublicKey(validKey)) // true

	// Invalid public keys
	invalidKey1 := "abc"                                          // Too short, invalid base58
	invalidKey2 := "65n1PdB5Q2G8cDK6v3N3eJ6D1DqnVDS7eNWFa7DbyNpv" // Too long
	invalidKey3 := "12345"                                        // Invalid base58 characters
	fmt.Println("Invalid key 1:", isValidPublicKey(invalidKey1))  // false
	fmt.Println("Invalid key 2:", isValidPublicKey(invalidKey2))  // false
	fmt.Println("Invalid key 3:", isValidPublicKey(invalidKey3))  // false
}
