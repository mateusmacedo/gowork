package rules

import (
	"fmt"
	specification "gowork/pkg/guards/specs"
)

type Rule[T any, R any] interface {
	Apply(T) (R, error)
	Combine(...Rule[T, R]) Rule[T, R]
	BatchApply([]T) ([]R, []error)
}

type rule[T any, R any] struct {
	Specification specification.Specification[T]
	Action        func(target T) (R, error)
}

func NewRule[T any, R any](spec specification.Specification[T], action func(target T) (R, error)) Rule[T, R] {
	return &rule[T, R]{
		Specification: spec,
		Action:        action,
	}
}

func (r *rule[T, R]) Apply(target T) (R, error) {
	if !r.Specification.IsSatisfiedBy(target) {
		var zero R
		return zero, fmt.Errorf("specification not satisfied by %v", target)
	}

	result, err := r.Action(target)
	if err != nil {
		return result, fmt.Errorf("action failed: %v", err)
	}

	return result, nil
}

func (r *rule[T, R]) BatchApply(targets []T) ([]R, []error) {
	var results []R
	var errors []error

	for _, target := range targets {
		result, err := r.Apply(target)
		if err != nil {
			errors = append(errors, err)
		} else {
			results = append(results, result)
		}
	}

	return results, errors
}

func (r *rule[T, R]) Combine(rules ...Rule[T, R]) Rule[T, R] {
	return &combinedRule[T, R]{
		rules: append([]Rule[T, R]{r}, rules...),
	}
}

type combinedRule[T any, R any] struct {
	rules []Rule[T, R]
}

func (cr *combinedRule[T, R]) Apply(target T) (R, error) {
	var zero R
	for _, rule := range cr.rules {
		if _, err := rule.Apply(target); err != nil {
			return zero, err
		}
	}
	return cr.rules[len(cr.rules)-1].Apply(target)
}

func (cr *combinedRule[T, R]) Combine(rules ...Rule[T, R]) Rule[T, R] {
	return &combinedRule[T, R]{
		rules: append(cr.rules, rules...),
	}
}

func (cr *combinedRule[T, R]) BatchApply(targets []T) ([]R, []error) {
	var results []R
	var errors []error

	for _, target := range targets {
		result, err := cr.Apply(target)
		if err != nil {
			errors = append(errors, err)
		} else {
			results = append(results, result)
		}
	}

	return results, errors
}