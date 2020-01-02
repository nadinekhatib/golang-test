package PackageNadine

import (
	"errors"
	"testing"
)

func TestMyName(t *testing.T) {
	if GetMyName() != "nadine" {
		t.Fatal(errors.New("GetMyName didn't reply correctly"))
	}
}
