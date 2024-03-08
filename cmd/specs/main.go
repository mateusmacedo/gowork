package main

import (
	"fmt"
	"gowork/pkg/guards/fixtures"
	specification "gowork/pkg/guards/specs"
)

func main() {
	type MyCandidate struct {
		Graduation bool
		Experience int
		Skills     []string
		Available  bool
	}
	graduationSpec := fixtures.NewDummySpecification(func(candidate any) bool {
		return candidate.(MyCandidate).Graduation
	})
	experienceSpec := fixtures.NewDummySpecification(func(candidate any) bool {
		candidateExperience := candidate.(MyCandidate).Experience
		return candidateExperience > 3
	})
	skillsSpec := fixtures.NewDummySpecification(func(candidate any) bool {
		skillList := []string{"Go", "Python", "SQL", "Java", "C++"}
		minimumRequiredSkills := 2
		matchingSkills := 0
		candidateSkills := candidate.(MyCandidate).Skills

		for _, skill := range candidateSkills {
			for _, requiredSkill := range skillList {
				if skill == requiredSkill {
					matchingSkills++
				}
			}
		}
		return matchingSkills >= minimumRequiredSkills
	})
	availabilitySpec := fixtures.NewDummySpecification(func(candidate any) bool {
		availability := candidate.(MyCandidate).Available
		return availability == true
	})

	// Criar um SpecificationBuilder e adicionar as especificações individualmente
	builder := specification.NewSpecificationBuilder[any]().
		WithSpecification(graduationSpec).
		And(skillsSpec).
		Or(experienceSpec).
		And(availabilitySpec)

	// Construir a especificação final
	finalSpecification := builder.Build()

	// Candidatos de exemplo
	candidates := []MyCandidate{
		{Graduation: true, Experience: 4, Skills: []string{"Go", "Python", "SQL"}, Available: false}, // Candidato 1
		{Graduation: false, Experience: 2, Skills: []string{"Java", "C++"}, Available: true},          // Candidato 2
		{Graduation: true, Experience: 5, Skills: []string{"Go", "Java"}, Available: true},           // Candidato 3
	}

	// Verificar se os candidatos satisfazem a especificação final
	var isSatisfied bool
	for i, candidate := range candidates {
		isSatisfied = finalSpecification.IsSatisfiedBy(candidate)
		if isSatisfied {
			fmt.Printf("Candidato %d atende aos critérios.\n", i+1)
		} else {
			fmt.Printf("Candidato %d não atende aos critérios.\n", i+1)
		}
	}
}
