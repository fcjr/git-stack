package version

import (
	"fmt"
)

var Version = "source"

func String() string {
	return fmt.Sprintf("git-patch version: %s\n", Version)
}
