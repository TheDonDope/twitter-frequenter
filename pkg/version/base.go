package version

import (
	"strings"
)

var (
	major      = "1"
	minor      = "0"
	patch      = "0"
	gitVersion = "v" + strings.Join([]string{major, minor, patch}, ".")
	gitCommit  = "ea349782fed12be36b90bd0475c0349e55372039"
)
