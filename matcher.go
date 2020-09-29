package role

import (
	"strings"
)

// MatchPath returns true if two paths of URL are same.
//
// The pattern can include `:`, which indicates routing parameter.
//
// e.g. MatchPath("/v1/api/users/:user_id", "/v1/api/users/12345") returns true.
func MatchPath(args ...interface{}) (interface{}, error) {
	if len(args) < 2 {
		return false, nil
	}

	pattern, ok := args[0].(string)

	if !ok {
		return false, nil
	}

	path, ok := args[1].(string)

	if !ok {
		return false, nil
	}

	xs := strings.Split(pattern, "/")
	ys := strings.Split(path, "/")

	if len(xs) != len(ys) {
		return false, nil
	}
	for i, x := range xs {
		if !strings.HasPrefix(x, ":") && x != ys[i] {
			return false, nil
		}
	}

	return true, nil
}
