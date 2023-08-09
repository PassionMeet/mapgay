package usertoken

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func Gen() string {
	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}
