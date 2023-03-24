package requests

import (
	"github.com/gofrs/uuid"
	"log"
	"sort"
)

func InArray(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func LetterArr(startLetter int32, endLetter int32) (strArr []string) {

	for i := startLetter; i <= endLetter; i++ {
		strArr = append(strArr, string(i))
	}
	return
}

func GetUUID() string {
	u2, err := uuid.NewV6()
	if err != nil {
		log.Println("")
	}
	return u2.String()
}
