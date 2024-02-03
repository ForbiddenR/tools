package safemap

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	InitDefaultPool()
	sn := "34499059959595"
	_ = Set(sn)
	value := IsExist(sn)
	fmt.Println("value", value)
	_ = Delete(sn)
	value = IsExist(sn)
	fmt.Println("value", value)
}
