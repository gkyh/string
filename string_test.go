package stringx

import (
	"testing"
)

func TestParse(t *testing.T) {

	str := Parse(123.22)
	t.Log(str)

	f64 := Float64(str)

	t.Log(f64)
}
