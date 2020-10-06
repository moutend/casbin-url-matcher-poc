package role

import (
	"errors"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

// Adapter represents the readonly adapter for Casbin.
type Adapter struct {
	policies []string
}

// NewAdapter is the constructor for Adapter.
func NewAdapter(policies []string) *Adapter {
	return &Adapter{policies}
}

// LoadPolicy loads all policy rules from the storage.
func (a *Adapter) LoadPolicy(m model.Model) error {
	for _, policy := range a.policies {
		persist.LoadPolicyLine(policy, m)
	}

	return nil
}

// SavePolicy saves policy.
func (a *Adapter) SavePolicy(model model.Model) error {
	return errors.New("not implemented")
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

// AddPolicies adds policy rules to the storage.
func (a *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	return errors.New("not implemented")
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

// RemovePolicies removes policy rules from the storage.
func (a *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {
	return errors.New("not implemented")
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New("not implemented")
}
