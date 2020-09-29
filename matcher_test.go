package role

import (
	"net/http"
	"testing"

	"github.com/casbin/casbin"
	"github.com/stretchr/testify/require"
)

func TestMatchpath(t *testing.T) {
	enforcer := casbin.NewEnforcer("model.conf", "policy.csv")
	enforcer.AddFunction("match_path", MatchPath)

	tests := []struct {
		role      string
		path      string
		method    string
		isAllowed bool
	}{
		{
			role:      "member",
			path:      "/v1/api/tenants/xxxx/members/yyyy/facilities",
			method:    http.MethodGet,
			isAllowed: true,
		},
		{
			role:      "admin",
			path:      "/v1/api/tenants/xxxx/members/yyyy/facilities",
			method:    http.MethodGet,
			isAllowed: true,
		},
		{
			role:      "member",
			path:      "/v1/api/tenants/xxxx/members/yyyy/facilities",
			method:    http.MethodPut,
			isAllowed: false,
		},
		{
			role:      "admin",
			path:      "/v1/api/tenants/xxxx/members/yyyy/facilities",
			method:    http.MethodPut,
			isAllowed: true,
		},
		{
			role:      "member",
			path:      "/v1/api/teams/xxxx/members/yyyy/facilities",
			method:    http.MethodGet,
			isAllowed: false,
		},
		{
			role:      "admin",
			path:      "/v1/api/teams/xxxx/members/yyyy/facilities",
			method:    http.MethodGet,
			isAllowed: false,
		},
	}
	for _, test := range tests {
		t.Logf("%+v\n", test)

		require.Equal(t, test.isAllowed, enforcer.Enforce(test.role, test.path, test.method))
	}
}
