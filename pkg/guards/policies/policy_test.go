package policies_test

import (
	"errors"
	"testing"

	"gowork/pkg/guards/policies"
	"gowork/pkg/guards/rules"
)

// MockRule é uma implementação de mock da interface Rule para fins de teste.
type MockRule[T any, R any] struct {
	ApplyFunc func(T) (R, error)
}

func (mr MockRule[T, R]) Apply(target T) (R, error) {
	return mr.ApplyFunc(target)
}

func (mr MockRule[T, R]) Combine(otherRules ...rules.Rule[T, R]) rules.Rule[T, R] {
	return MockRule[T, R]{
		ApplyFunc: func(target T) (R, error) {
			var lastResult R
			var lastError error
			for _, rule := range append([]rules.Rule[T, R]{mr}, otherRules...) {
				lastResult, lastError = rule.Apply(target)
				if lastError != nil {
					return *new(R), lastError
				}
			}
			return lastResult, nil
		},
	}
}

func (mr MockRule[T, R]) BatchApply(targets []T) ([]R, []error) {
	var results []R
	var errs []error
	for _, target := range targets {
		result, err := mr.Apply(target)
		results = append(results, result)
		errs = append(errs, err)
	}
	return results, errs
}

func TestPolicy_ApplyRules(t *testing.T) {
	// Definindo a tabela de testes.
	tests := []struct {
		name        string
		policyRules []rules.Rule[int, string]
		target      int
		wantResult  string
		wantErr     error
	}{
		{
			name:        "No rules",
			policyRules: []rules.Rule[int, string]{},
			target:      1,
			wantErr:     errors.New("no rules to apply"),
		},
		{
			name: "Single rule, success",
			policyRules: []rules.Rule[int, string]{
				MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "success", nil
					},
				},
			},
			target:     1,
			wantResult: "success",
			wantErr:    nil,
		},
		{
			name: "Single rule, fail",
			policyRules: []rules.Rule[int, string]{
				MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "", errors.New("specification not satisfied by 1")
					},
				},
			},
			target:  1,
			wantErr: errors.New("specification not satisfied by 1"),
		},
		{
			name: "Multiple rules, all pass",
			policyRules: []rules.Rule[int, string]{
				MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "success", nil
					},
				},
				MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "success", nil
					},
				},
			},
			target:     1,
			wantResult: "success",
			wantErr:    nil,
		},
		{
			name: "Multiple rules, one fails",
			policyRules: []rules.Rule[int, string]{
				MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "success", nil
					},
				},
				MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "", errors.New("specification not satisfied by 1")
					},
				},
				MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "success", nil
					},
				},
			},
			target:  1,
			wantErr: errors.New("specification not satisfied by 1"),
		},
	}

	// Executando os testes.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := policies.NewPolicy(tt.policyRules...)
			gotResult, gotErr := p.ApplyRules(tt.target)

			if (gotErr == nil) != (tt.wantErr == nil) || (gotErr != nil && gotErr.Error() != tt.wantErr.Error()) {
				t.Errorf("Policy.ApplyRules() error = %v, wantErr %v", gotErr, tt.wantErr)
			}
			if gotResult != tt.wantResult {
				t.Errorf("Policy.ApplyRules() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}