package hash

import (
    "crypto/sha1"
    "encoding/hex"
)

func SHA1Hash(text string) string {
    hash := sha1.New()
    hash.Write([]byte(text))
    return hex.EncodeToString(hash.Sum(nil))
}
