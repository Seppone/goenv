package mock_test

import (
	"fmt"
	"testing"

	"github.com/seppone/goenv"
)

func SortedCompare(t *testing.T, a, b []goenv.Version) {
	t.Helper()
	errMessage := fmt.Sprintf("%v != %v", a, b)
	if len(a) != len(b) {
		t.Fatalf(errMessage)
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Fatalf(errMessage)
		}
	}
}
