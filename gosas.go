package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	SaS := createSharedAccessToken("<ServiceBusURL>", "<PolicyName>", "<SecretKey>")
	fmt.Println(SaS)
}

func createSharedAccessToken(uri string, saName string, saKey string) string {

	if len(uri) == 0 || len(saName) == 0 || len(saKey) == 0 {
		return "Missing required parameter"
	}

	encoded := url.PathEscape(uri)
	encoded = strings.Replace(encoded, ":", "%3A", -1)
	now := time.Now().Unix()
	week := 60 * 60 * 24 * 7
	ts := now + int64(week)
	signature := encoded + "\n" + strconv.Itoa(int(ts))
	key := []byte(saKey)
	hmac := hmac.New(sha256.New, key)
	hmac.Write([]byte(signature))
	hmacString := url.PathEscape(base64.StdEncoding.EncodeToString(hmac.Sum(nil)))
	hmacString = strings.Replace(hmacString, "=", "%3D", -1)

	result := "SharedAccessSignature sr=" + encoded + "&sig=" +
		hmacString + "&se=" + strconv.Itoa(int(ts)) + "&skn=" + saName
	return result
}
