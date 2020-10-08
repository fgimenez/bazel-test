package pkg1

import (
	"fmt"

	"github.com/fgimenez/bazel-test/pkg/util"
)

func Pkg1() string {
	return fmt.Sprintf("return from Someutil: %q", util.SomeUtil())
}
