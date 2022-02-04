package integers

import (
	"testing"
)

func TestIter(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// func ExampleAdd() {
// 	repeated :=
// 		fmt.Println(sum)
// 	// Output: 6
// }
