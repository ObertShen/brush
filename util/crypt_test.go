package util

import (
	"fmt"
	"testing"
)

func TestGetSHA256String(t *testing.T) {
	fmt.Println(GetSHA256String("123456"))
}
