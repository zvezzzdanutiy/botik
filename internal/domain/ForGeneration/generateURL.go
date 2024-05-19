package ForGeneration

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateURL(category string) string {
	const apiKey = "c2edc5973589c4a161d8dfbf5640ecab248e6b418bf81fd635d5a092d990335d"
	x := fmt.Sprintf("%d", time.Now().Unix())
	str := "pid=v26kt07850p19fb2278v&method=getRandItem&category=" + category + "&uts=" + x
	hash := GetMD5Hash(str + apiKey)
	return "http://anecdotica.ru/api?" + str + "&hash=" + hash
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
