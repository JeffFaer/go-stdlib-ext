package morecmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/slices"
)

func TestTrueFirst(t *testing.T) {
	l := []bool{
		true,
		false,
		true,
		false,
	}
	slices.SortFunc(l, TrueFirst())

	if diff := cmp.Diff([]bool{true, true, false, false}, l); diff != "" {
		t.Errorf("Diff (-want +got)n%s", diff)
	}
}

func TestFalseFirst(t *testing.T) {
	l := []bool{
		true,
		false,
		true,
		false,
	}
	slices.SortFunc(l, FalseFirst())

	if diff := cmp.Diff([]bool{false, false, true, true}, l); diff != "" {
		t.Errorf("Diff (-want +got)n%s", diff)
	}
}
