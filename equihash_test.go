package equihash

import (
	"encoding/hex"
	"testing"
)

func TestVerifySolution(t *testing.T) {
	input := "892025ec077f658d3a5d3db377dfa909590a3980071c43ef74fd1f9b2fe823b6"
	nonce := "9c7ed1f011b1b077"
	output := "02acdd8f8617549de9421bedad4adafce80507e53f793f0eb6a70526dcbcbad52678ed5b41095d1bb3fc246230acb24ad7db294507d68c1fe153282f54fafcbbc00dc45c1a39364906fa7e5c4a861ceb3db9482e99b41fab4ee3513bcd9b4f60eea465400bbd8220"
	inputBytes, _ := hex.DecodeString(input)
	nonceBytes, _ := hex.DecodeString(nonce)
	outputBytes, _ := hex.DecodeString(output)

	valid := VerifySolution(inputBytes, nonceBytes, outputBytes)
	if !valid {
		t.Fatal("Invalid")
	}

	nonce = "9c7ed1f011b1b078" // edit the last character to 8. others remain unchange.
	nonceBytes, _ = hex.DecodeString(nonce)

	valid = VerifySolution(inputBytes, nonceBytes, outputBytes)
	if valid {
		t.Fatal("Invalid")
	}
}
