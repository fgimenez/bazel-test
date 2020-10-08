package pkg1

import (
	"fmt"

	"github.com/fgimenez/test-bazel/pkg/util"
)

func Pkg1() string {
	return fmt.Sprintf("return from Someutil: %q", util.SomeUtil())
}
