package hash

import (
    "crypto/sha256"
    "encoding/hex"
)

func SHA256Hash(text string) string {
    hash := sha256.New()
    hash.Write([]byte(text))
    return hex.EncodeToString(hash.Sum(nil))
}
