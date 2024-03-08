package rules_test

import (
	"errors"
	"testing"

	"gowork/pkg/guards/rules"
	specification "gowork/pkg/guards/specs"
)

type mockSpecification[T any] struct {
	isSatisfiedBy bool
}

func (ms mockSpecification[T]) IsSatisfiedBy(_ T) bool {
	return ms.isSatisfiedBy
}

func TestRule_Apply(t *testing.T) {
	type test[T, R any] struct {
		name          string
		specification specification.Specification[int]
		action        func(int) (string, error)
		target        int
		wantResult    string
		wantErr       bool
	}

	tests := []test[int, string]{
		{
			name:          "SatisfiedSpecification",
			specification: mockSpecification[int]{isSatisfiedBy: true},
			action: func(i int) (string, error) {
				return "ok", nil
			},
			target:     1,
			wantResult: "ok",
			wantErr:    false,
		},
		{
			name:          "UnsatisfiedSpecification",
			specification: mockSpecification[int]{isSatisfiedBy: false},
			action: func(i int) (string, error) {
				return "", nil
			},
			target:    2,
			wantErr:   true,
		},
		{
			name:          "ActionError",
			specification: mockSpecification[int]{isSatisfiedBy: true},
			action: func(i int) (string, error) {
				return "", errors.New("action error")
			},
			target:  3,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rules.NewRule(tt.specification, tt.action)
			gotResult, err := r.Apply(tt.target)

			if (err != nil) != tt.wantErr {
				t.Errorf("Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("Apply() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestRule_Combine(t *testing.T) {
	type test[T, R any] struct {
		name          string
		specification specification.Specification[int]
		action        func(int) (string, error)
		target        int
		rules         []rules.Rule[int, string]
		wantResult    string
		wantErr       bool
	}

	alwaysTrueSpec := mockSpecification[int]{isSatisfiedBy: true}
	alwaysFalseSpec := mockSpecification[int]{isSatisfiedBy: false}
	successAction := func(i int) (string, error) { return "ok", nil }
	failAction := func(i int) (string, error) { return "", errors.New("action error") }

	tests := []test[int, string]{
		{
			name:          "AllRulesSatisfied",
			specification: alwaysTrueSpec,
			action:        successAction,
			target:        1,
			rules: []rules.Rule[int, string]{
				rules.NewRule(alwaysTrueSpec, successAction),
				rules.NewRule(alwaysTrueSpec, successAction),
			},
			wantResult: "ok",
			wantErr:    false,
		},
		{
			name:          "FirstRuleUnsatisfied",
			specification: alwaysTrueSpec,
			action:        successAction,
			target:        2,
			rules: []rules.Rule[int, string]{
				rules.NewRule(alwaysFalseSpec, successAction),
				rules.NewRule(alwaysTrueSpec, successAction),
			},
			wantErr: true,
		},
		{
			name:          "SecondRuleUnsatisfied",
			specification: alwaysTrueSpec,
			action:        successAction,
			target:        3,
			rules: []rules.Rule[int, string]{
				rules.NewRule(alwaysTrueSpec, successAction),
				rules.NewRule(alwaysFalseSpec, successAction),
			},
			wantErr: true,
		},
		{
			name:          "ActionErrorInCombinedRule",
			specification: alwaysTrueSpec,
			action:        successAction,
			target:        4,
			rules: []rules.Rule[int, string]{
				rules.NewRule(alwaysTrueSpec, successAction),
				rules.NewRule(alwaysTrueSpec, failAction),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			r := rules.NewRule(tt.specification, tt.action).Combine(tt.rules...)
			gotResult, err := r.Apply(tt.target)

			if (err != nil) != tt.wantErr {
				t.Errorf("Combine().Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotResult != tt.wantResult {
				t.Errorf("Combine().Apply() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestRule_BatchApply(t *testing.T) {
	type test[T, R any] struct {
		name          string
		specification specification.Specification[int]
		action        func(int) (string, error)
		targets       []int
		wantResults   []string
		wantErrCount  int
	}

	alwaysTrueSpec := mockSpecification[int]{isSatisfiedBy: true}
	successAction := func(i int) (string, error) { return "ok", nil }
	failAction := func(i int) (string, error) { return "", errors.New("action error") }

	tests := []test[int, string]{
		{
			name:          "AllRulesSatisfied",
			specification: alwaysTrueSpec,
			action:        successAction,
			targets:       []int{1, 2, 3},
			wantResults:   []string{"ok", "ok", "ok"},
			wantErrCount:  0,
		},
		{
			name:          "ActionError",
			specification: alwaysTrueSpec,
			action:        failAction,
			targets:       []int{1, 2, 3},
			wantResults:   []string{},
			wantErrCount:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rules.NewRule(tt.specification, tt.action)
			gotResults, gotErrors := r.BatchApply(tt.targets)

			if len(gotErrors) != tt.wantErrCount {
				t.Errorf("BatchApply() got %v errors, want %v errors", len(gotErrors), tt.wantErrCount)
			}
			if len(gotResults) != len(tt.wantResults) {
				t.Errorf("BatchApply() got %v results, want %v results", len(gotResults), len(tt.wantResults))
			}
			for i, gotResult := range gotResults {
				if gotResult != tt.wantResults[i] {
					t.Errorf("BatchApply() gotResults[%d] = %v, want %v", i, gotResult, tt.wantResults[i])
				}
			}
		})
	}
}

func TestCombinedRule_Combine(t *testing.T) {
	type test[T, R any] struct {
		name         string
		initialRules []rules.Rule[int, string]
		newRules     []rules.Rule[int, string]
		target       int
		wantResult   string
		wantErr      bool
	}

	successAction := func(i int) (string, error) { return "ok", nil }
	failAction := func(i int) (string, error) { return "", errors.New("action error") }
	alwaysTrueSpec := mockSpecification[int]{isSatisfiedBy: true}
	alwaysFalseSpec := mockSpecification[int]{isSatisfiedBy: false}

	tests := []test[int, string]{
		{
			name:         "CombineWithNoNewRules",
			initialRules: []rules.Rule[int, string]{rules.NewRule(alwaysTrueSpec, successAction)},
			newRules:     []rules.Rule[int, string]{},
			target:       1,
			wantResult:   "ok",
			wantErr:      false,
		},
		{
			name:         "CombineWithNewRules",
			initialRules: []rules.Rule[int, string]{rules.NewRule(alwaysTrueSpec, successAction)},
			newRules:     []rules.Rule[int, string]{rules.NewRule(alwaysTrueSpec, successAction)},
			target:       1,
			wantResult:   "ok",
			wantErr:      false,
		},
		{
			name:         "CombineWithFailingNewRule",
			initialRules: []rules.Rule[int, string]{rules.NewRule(alwaysTrueSpec, successAction)},
			newRules:     []rules.Rule[int, string]{rules.NewRule(alwaysFalseSpec, successAction)},
			target:       1,
			wantErr:      true,
		},
		{
			name:         "CombineWithFailingAction",
			initialRules: []rules.Rule[int, string]{rules.NewRule(alwaysTrueSpec, successAction)},
			newRules:     []rules.Rule[int, string]{rules.NewRule(alwaysTrueSpec, failAction)},
			target:       1,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initialCombinedRule := rules.NewRule(alwaysTrueSpec, successAction).Combine(tt.initialRules...)
			combinedRule := initialCombinedRule.Combine(tt.newRules...)
			gotResult, err := combinedRule.Apply(tt.target)

			if (err != nil) != tt.wantErr {
				t.Errorf("Combine().Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotResult != tt.wantResult {
				t.Errorf("Combine().Apply() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestCombinedRule_BatchApply(t *testing.T) {
	type test[T, R any] struct {
		name         string
		combinedRule rules.Rule[int, string]
		targets      []int
		wantResults  []string
		wantErrCount int
	}

	successAction := func(i int) (string, error) { return "ok", nil }
	failAction := func(i int) (string, error) { return "", errors.New("action error") }
	alwaysTrueSpec := mockSpecification[int]{isSatisfiedBy: true}

	tests := []test[int, string]{
		{
			name:         "AllRulesSatisfied",
			combinedRule: rules.NewRule(alwaysTrueSpec, successAction).Combine(rules.NewRule(alwaysTrueSpec, successAction)),
			targets:      []int{1, 2, 3},
			wantResults:  []string{"ok", "ok", "ok"},
			wantErrCount: 0,
		},
		{
			name:         "ActionError",
			combinedRule: rules.NewRule(alwaysTrueSpec, failAction).Combine(rules.NewRule(alwaysTrueSpec, failAction)),
			targets:      []int{1, 2, 3},
			wantResults:  []string{},
			wantErrCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResults, gotErrors := tt.combinedRule.BatchApply(tt.targets)

			if len(gotErrors) != tt.wantErrCount {
				t.Errorf("BatchApply() got %v errors, want %v errors", len(gotErrors), tt.wantErrCount)
			}
			if len(gotResults) != len(tt.wantResults) {
				t.Errorf("BatchApply() got %v results, want %v results", len(gotResults), len(tt.wantResults))
			}
			for i, gotResult := range gotResults {
				if gotResult != tt.wantResults[i] {
					t.Errorf("BatchApply() gotResults[%d] = %v, want %v", i, gotResult, tt.wantResults[i])
				}
			}
		})
	}
}