package hash

import (
    "errors"
    "regexp"
)

func DetectHashType(hash string) (string, error) {
    if len(hash) == 32 && isHex(hash) {
        return "md5", nil
    }
    if len(hash) == 40 && isHex(hash) {
        return "sha1", nil
    }
    if len(hash) == 64 && isHex(hash) {
        return "sha256", nil
    }
    return "", errors.New("unknown hash type")
}

func isHex(s string) bool {
    match, _ := regexp.MatchString("^[a-fA-F0-9]+$", s)
    return match
}
