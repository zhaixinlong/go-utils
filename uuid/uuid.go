package uuid

import (
	"github.com/teris-io/shortid"
)

func GenerateShortid() (string, error) {
	value, err := shortid.Generate()
	if err != nil {
		return "", err
	}
	// fmt.Println(value) // 例如：We8bXiCVg
	return value, nil
}
