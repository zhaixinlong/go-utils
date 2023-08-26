package main

import (
	"fmt"

	"github.com/zhaixinlong/go-utils/uuid"
)

func main() {
	value, err := uuid.GenerateShortid()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value) // 例如：We8bXiCVg
}
