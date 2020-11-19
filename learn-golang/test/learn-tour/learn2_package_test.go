package tour

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_package(t *testing.T) {
	fmt.Println("My favorite number is", rand.Intn(10))
}
