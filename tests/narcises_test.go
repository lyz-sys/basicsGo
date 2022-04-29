package tests

import (
	"demo/tools"
	"fmt"
	"testing"
)

func TestNarcises(t *testing.T) {
	res := tools.Narcises(100, 10000)
	fmt.Println(res)
}
