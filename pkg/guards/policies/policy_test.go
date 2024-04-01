package policies_test

import (
	"errors"
	"testing"

	"github.com/mateusmacedo/gowork/pkg/guards/policies"
	"github.com/mateusmacedo/gowork/pkg/guards/rules"
)

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
	tests := []struct {
		name        string
		setupPolicy func(*policies.Policy[int, string])
		target      int
		wantResult  string
		wantErr     error
	}{
		{
			name:        "No rules",
			setupPolicy: func(p *policies.Policy[int, string]) {},
			target:      1,
			wantErr:     errors.New("no rules to apply"),
		},
		{
			name: "Add rule after creation",
			setupPolicy: func(p *policies.Policy[int, string]) {
				additionalRule := MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "additional", nil
					},
				}
				p.AddRule(additionalRule)
			},
			target:     1,
			wantResult: "additional",
			wantErr:    nil,
		},
		{
			name: "Batch apply rules with mixed results",
			setupPolicy: func(p *policies.Policy[int, string]) {
				rule := MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						if i%2 == 0 {
							return "even", nil
						}
						return "", errors.New("odd number")
					},
				}
				p.AddRule(rule)
			},
			target:     1,
			wantResult: "",
			wantErr:    errors.New("odd number"),
		},
		{
			name: "Error case when applying rule",
			setupPolicy: func(p *policies.Policy[int, string]) {
				errorRule := MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "", errors.New("error applying rule")
					},
				}
				p.AddRule(errorRule)
			},
			target:  1,
			wantErr: errors.New("error applying rule"),
		},
		{
			name: "Multiple rules with one error",
			setupPolicy: func(p *policies.Policy[int, string]) {
				successRule := MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "success", nil
					},
				}
				errorRule := MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "", errors.New("error applying rule")
					},
				}
				p.AddRule(successRule)
				p.AddRule(errorRule)
			},
			target:  1,
			wantErr: errors.New("error applying rule"),
		},
		{
			name: "Combining multiple rules",
			setupPolicy: func(p *policies.Policy[int, string]) {
				p.AddRule(MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "rule1", nil
					},
				})
				p.AddRule(MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "rule2", nil
					},
				})
			},
			target:     1,
			wantResult: "rule2",
			wantErr:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := policies.NewPolicy[int, string]()
			tt.setupPolicy(p)
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

func TestPolicy_BatchApplyRules(t *testing.T) {
	tests := []struct {
		name        string
		setupPolicy func(*policies.Policy[int, string])
		targets     []int
		wantResults []string
		wantErrors  []error
	}{
		{
			name: "Batch apply with mixed results",
			setupPolicy: func(p *policies.Policy[int, string]) {
				rule := MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						if i%2 == 0 {
							return "even", nil
						}
						return "", errors.New("odd number")
					},
				}
				p.AddRule(rule)
			},
			targets:     []int{2, 4, 6, 7},
			wantResults: []string{"even", "even", "even", ""},
			wantErrors:  []error{nil, nil, nil, errors.New("odd number")},
		},
		{
			name: "No rules to apply for batch",
			setupPolicy: func(p *policies.Policy[int, string]) {
			},
			targets:     []int{1, 2, 3},
			wantResults: nil,
			wantErrors:  []error{errors.New("no rules to apply")},
		},
		{
			name: "Combine rules for batch apply",
			setupPolicy: func(p *policies.Policy[int, string]) {
				rule1 := MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "rule1", nil
					},
				}
				rule2 := MockRule[int, string]{
					ApplyFunc: func(i int) (string, error) {
						return "rule2", nil
					},
				}
				p.AddRule(rule1)
				p.AddRule(rule2)
			},
			targets:     []int{1, 2, 3},
			wantResults: []string{"rule2", "rule2", "rule2"},
			wantErrors:  []error{nil, nil, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := policies.NewPolicy[int, string]()
			tt.setupPolicy(p)
			gotResults, gotErrors := p.BatchApplyRules(tt.targets)

			if len(gotResults) != len(tt.wantResults) {
				t.Fatalf("Policy.BatchApplyRules() got %d results, want %d", len(gotResults), len(tt.wantResults))
			}

			if len(gotErrors) != len(tt.wantErrors) {
				t.Fatalf("Policy.BatchApplyRules() got %d errors, want %d", len(gotErrors), len(tt.wantErrors))
			}

			for i, gotResult := range gotResults {
				if gotResult != tt.wantResults[i] {
					t.Errorf("Policy.BatchApplyRules() gotResult[%d] = %v, want %v", i, gotResult, tt.wantResults[i])
				}
			}

			for i, gotErr := range gotErrors {
				if (gotErr == nil) != (tt.wantErrors[i] == nil) || (gotErr != nil && gotErr.Error() != tt.wantErrors[i].Error()) {
					t.Errorf("Policy.BatchApplyRules() gotError[%d] = %v, want %v", i, gotErr, tt.wantErrors[i])
				}
			}
		})
	}
}
