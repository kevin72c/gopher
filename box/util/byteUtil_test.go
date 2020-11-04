package util

import (
	"fmt"
	"testing"
)

func TestInt32ToBytes(t *testing.T) {

	fmt.Println(11111111111)
	slice := make([]interface{}, 0)
	slice = append(slice, byte(1))
	slice = append(slice, 2)
	slice = append(slice, "xxx")
	slice = append(slice, uint64(32))
	slice = append(slice, 2)
	fmt.Printf("slice: %v len: %d, cap: %d\n", slice, len(slice), cap(slice))

}
