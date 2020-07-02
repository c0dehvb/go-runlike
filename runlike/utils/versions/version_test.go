package versions

import (
	"fmt"
	"testing"
)

func TestParseTolerant(t *testing.T) {
	version, err := ParseTolerant("v1.17.4-rancher1-3")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(version.Pre[0].IsNumeric())
	fmt.Println(version.String())
}
