package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func VerifyHMAC(payload []byte, signature, secret string) bool {
	if signature == "" {
		return false
	}

	const prefix = "sha256="
	if !strings.HasPrefix(signature, prefix) {
		return false
	}

	sig := signature[len(prefix):]

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedMAC := mac.Sum(nil)

	expectedSig, err := hex.DecodeString(sig)
	if err != nil {
		return false
	}

	return hmac.Equal(expectedMAC, expectedSig)
}
