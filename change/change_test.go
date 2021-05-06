package change

import (
	"testing"
)

func TestChange(t *testing.T) {
	change1 := RetrieveChange(55)

	if change1 != 2 {
		t.Errorf("Change is invalid")
	}

}
