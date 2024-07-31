package tools

import (
	"fmt"
	"math/rand"
)

func Rand4Num() string {
	intn := rand.Intn(9999)
	for intn < 1000 {
		intn = rand.Intn(9999)
	}
	return fmt.Sprintf("%d", intn)
}
