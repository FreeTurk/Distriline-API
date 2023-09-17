package main

import (
	"crypto/sha256"
	"encoding/json"
	"os"
)

func CheckUserIntegrity(input map[string]interface{}) (bool, [32]byte) {
	checksumSecurityKey := os.Getenv("CHECKSUM_SEC_KEY")

	prevChecksum := input["Checksum"]

	delete(input, "Checksum")

	keyedInput, err := json.Marshal(input)

	keyedInput = append(keyedInput, []byte(checksumSecurityKey)...)

	if err != nil {
		panic(err)
	}

	var checksum [32]byte = sha256.Sum256(keyedInput)

	return checksum == prevChecksum, checksum
	// TODO
}
