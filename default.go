package role

import "github.com/casbin/casbin/v2/model"

var (
	DefaultModel, _ = model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && match_path(p.obj, r.obj) && r.act == p.act
`)
	DefaultAdapter = NewAdapter([]string{
		"p, member, /v1/api/tenants/:tenant_id/members/:member_id/facilities, GET",
		"p, admin, /v1/api/tenants/:tenant_id/members/:member_id/facilities, GET",
		"p, admin, /v1/api/tenants/:tenant_id/members/:member_id/facilities, PUT",
	})
)
