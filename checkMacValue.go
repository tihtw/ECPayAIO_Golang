package ecpayaio

import (
	"crypto/sha256"
	"fmt"
	// "log"
	"net/url"
	"sort"
	"strings"
)

var (
	EncryptTypeMD5    EncryptType = 0
	EncryptTypeSHA256 EncryptType = 1
)

type EncryptType int

func GenerateCheckMacValue(payload map[string]string, hashKey string, hashIV string, encryptType EncryptType) string {

	// log.Println(payload)

	keys := []string{}
	for id, _ := range payload {
		keys = append(keys, id)
	}
	sort.StringSlice(keys).Sort()

	message := "HashKey=" + hashKey

	for _, key := range keys {
		message += "&" + key + "=" + payload[key]
	}

	message += "&HashIV=" + hashIV

	escapedMessage := url.QueryEscape(message)

	lower := strings.ToLower(escapedMessage)
	if encryptType == EncryptTypeSHA256 {
		sum := sha256.Sum256([]byte(lower))
		ret := fmt.Sprintf("%X", sum)

		// log.Println(ret)

		return ret
	}
	return ""
}
