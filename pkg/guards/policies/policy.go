package policies

import (
	"errors"

	"github.com/mateusmacedo/gowork/pkg/guards/rules"
)

type Policy[T any, R any] struct {
	rules []rules.Rule[T, R]
}

func NewPolicy[T any, R any](rules ...rules.Rule[T, R]) *Policy[T, R] {
	return &Policy[T, R]{rules: rules}
}

func (p *Policy[T, R]) AddRule(r rules.Rule[T, R]) {
	p.rules = append(p.rules, r)
}

func (p *Policy[T, R]) ApplyRules(target T) (R, error) {
	var lastResult R
	var err error
	var combinedRule rules.Rule[T, R]
	if len(p.rules) > 0 {
		combinedRule = p.rules[0]
		for _, r := range p.rules[1:] {
			combinedRule = combinedRule.Combine(r)
		}
		lastResult, err = combinedRule.Apply(target)
	} else {
		return *new(R), errors.New("no rules to apply")
	}

	if err != nil {
		return *new(R), err
	}

	return lastResult, nil
}

func (p *Policy[T, R]) BatchApplyRules(targets []T) ([]R, []error) {
	if len(p.rules) == 0 {
		return nil, []error{errors.New("no rules to apply")}
	}

	combinedRule := p.rules[0]
	for _, r := range p.rules[1:] {
		combinedRule = combinedRule.Combine(r)
	}

	return combinedRule.BatchApply(targets)
}
