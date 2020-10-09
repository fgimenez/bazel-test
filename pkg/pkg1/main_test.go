package pkg1_test

import (
	"fmt"
	"testing"

	"github.com/fgimenez/bazel-test/pkg/pkg1"
)

func TestPkg1(t *testing.T) {
	actual := pkg1.Pkg1()
	expected := fmt.Sprintf("return from Someutil: %q", "utl")

	if actual != expected {
		t.Errorf("expected: %q, actual %q", expected, actual)
	}
}
