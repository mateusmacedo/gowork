# O Poder dos Padrões de Design na Engenharia de Software Moderna

## Introdução

Criar sistemas que sejam funcionais, tolerantes a mudanças e fáceis de manter e manter é um problema comum e persistente no campo do desenvolvimento de software. Os "padrões de projeto" nos ajudam, fornecendo soluções prontas, testadas e aprovadas para os problemas comuns com os quais todos os desenvolvedores lidam diariamente.
Realizar isso adequadamente e com eficiência ao utilizar o conhecimento atemporal que faz parte dos elementos da "Engenharia", independentemente de quão complicado isso possa parecer inicialmente para aqueles que não estão preparados para construir software robusto e confiável. Além disso, é uma máxima inquestionável que nada é mais complicado que escrever tudo do zero sem garantias.
Este estudo examina três padrões principais: `Specification`, `Builder` e o `Composite`. O objetivo é demonstrar como a qualidade, eficiência e reuso podem ser usadas para criar software `Enterprise`.

## O Padrão `Specification`

### Conceito

Como o próprio nome sugere, o padrão `Specification` foi introduzido pela primeira vez no campo de modelagem de `Rich Domain` no livro `Domain-Driven Design: Tackling Complexity in the Heart of Software`. Ele oferece uma maneira simples e declarativa de implementar lógica de negócios, permitindo que a composição construa de forma fácil e modular regras de negócios complexas. Este padrão define uma `Specification` como um objeto capaz de determinar se outro objeto satisfaz ou não um critério específico.

Ao separar a lógica de seleção (o "o quê") da lógica de execução (o "como"), o padrão `Specification` dá poder e torna claro como as regras de negócios podem mudar e se trocar. Além disso, permite que as regras de negócios anteriores sejam reutilizadas, atendendo diretamente ao princípio `Open–closed principle` do `SOLID`

### Benefícios

- **Reusabilidade**: Permite que a definição de regras de negócio atômicas e reutilizáveis que permitem serem combinadas para formar lógicas de validação complexas.
- **Composição**: Permite combinar especificações simples em especificações complexas sem que as implementações existentes tenham que ser alteradas, sendo assim aderente o princípio aberto/fechado S**O**LID.
- **Clareza**: Permite que o código seja mais legível e revelador de suas intenções, ao separar nitidamente as regras de negócio da lógica de aplicação.

## O Padrão `Builder`

### Conceito

Quando construir objetos complexos se torna um problema, o padrão de `Builder` é útil. Ele define um processo construtivo que permite que representações de objetos com a mesma estrutura sejam construídas com diferentes detalhes. Separar o processo de construção da maneira como o objeto final é representado permite a criação de objetos muito mais adaptáveis, o que permite aos desenvolvedores controlar a complexidade de alterar os objetos criados para requisitos específicos.

Além disso, encapsular a lógica de construção em um diretor ou grupo de construtores concretos facilita a manutenção do sistema, pois novas variações de objetos não comprometem o código já existente. A aplicação do `Builder`no contexto de `Specification` torna a configuração fluente de regras complexas mais fácil durante a construção de vários tipos de lógica de negócios.

### Benefícios

- **Encapsulamento**: Isola a lógica de construção e detalhes de implementação do usuário, simplificando o contrato de criação.
- **Flexibilidade**: Permite maior controle sobre o processo de instanciação do que quando com a inicialização direta, permitindo ajustes específicos e configurações complexas.
- **Fluidez**: Com a criação através contratos fluídos melhorando a legibilidade e a usabilidade do código de construção de `Specification`.

## O Padrão `Composite`

### Conceito

O padrão de projeto conhecido como `Composite`  organiza os objetos em uma estrutura que lembra uma árvore de hierarquias parte-todo. Permite que as relações entre objetos individuais ou suas composições sejam usadas de forma transparente. Isso é muito útil para a gestão de regras de negócios, onde o padrão `Composite`  torna a criação e implementação de sistemas complexos muito fáceis.
Um objeto composto é constituído por outros objetos individuais que podem ser outros objetos compostos. Uma grande vantagem sobre o padrão `Composite`  é que permite ao cliente tratar os objetos individuais e suas composições de maneira comum.

Diversos tipos de software usam `Composite`. Por exemplo, em um sistema de arquivos, como arquivos e diretórios; em interfaces de usuário, quando configurações de relacionamento pai-filho para itens, como botões e menus, processamento de documentos ou representação de redes de computadores são outras aplicações.

### Benefícios

- **Uniformidade**: Objetos simples e compostos são tratados de maneira uniforme, simplificando a interface para o desenvolvedor.
- **Flexibilidade de Estrutura**: Permite a criação dinâmica de estruturas complexas, que podem ser expandidas ou modificadas conforme novos cenários são necessários.
- **Simplicidade**: Diminue a complexidade do código, pois permiti que operações complexas sejam executadas em estruturas compostas como se fossem simples objetos.

Ao utilizar os padrões `Specification`, `Builder` e `Composite` na engenharia de software uma abordagem poderosa e obtida para desenvolver sistemas complexos. Permitir a definição clara, construção flexível e a composição hierárquica de componentes, favorecem criar sistemas que são, robustos, adaptáveis e fáceis de manter.

# Aplicando Padrões de Design em Sistemas de Especificações e Políticas

Já que exploramos a teoria dos padrões `Specification`, `Builder` e `Composite`. Avancemos em nossa compreensão das aplicações práticas deles em um sistema de especificações, regras e políticas, destacando como eles podem ser usados para lidar com desafios comuns de desenvolvimento e promover a criação de sistemas flexíveis e manutíveis. Para ilustrar, usamos a linguagem  `Golang`.

## Implementando o Padrão de Especificação

### Aplicação Prática

A implementação do padrão de Especificação começa com a definição de uma interface genérica para `Specification`. é o ponto inicial para a aplicação do padrão de Especificação. Este contrato especifica um método chamado `IsSatisfiedBy` para verificar se um objeto específico do candidato atende a uma especificação.

```go
package specification

type Candidate interface{}

type Specification[T Candidate] interface {
 IsSatisfiedBy(candidate T) bool
}
```

Especificações concretas são implementadas para que as regras de negócio específicas sejam encapsuladas, e que essas regras possam ser combinadas usando operações lógicas como `AND`, `OR` e `NOT`.

```go
package specification

type AndSpecification[T Candidate] struct {
 specs []Specification[T]
}

func NewAndSpecification[T Candidate](specs ...Specification[T]) *AndSpecification[T] {
 return &AndSpecification[T]{specs: specs}
}

func (s *AndSpecification[T]) IsSatisfiedBy(candidate T) bool {
 for _, spec := range s.specs {
  if !spec.IsSatisfiedBy(candidate) {
   return false
  }
 }
 return true
}

type OrSpecification[T Candidate] struct {
 specs []Specification[T]
}

func NewOrSpecification[T Candidate](specs ...Specification[T]) *OrSpecification[T] {
 return &OrSpecification[T]{specs: specs}
}

func (s *OrSpecification[T]) IsSatisfiedBy(candidate T) bool {
 for _, spec := range s.specs {
  if spec.IsSatisfiedBy(candidate) {
   return true
  }
 }
 return false
}

type NotSpecification[T Candidate] struct {
 spec Specification[T]
}

func NewNotSpecification[T Candidate](spec Specification[T]) *NotSpecification[T] {
 return &NotSpecification[T]{spec: spec}
}

func (s *NotSpecification[T]) IsSatisfiedBy(candidate T) bool {
 return !s.spec.IsSatisfiedBy(candidate)
}
```

### Desafios Superados

- **Modularidade**: Especificações que podem ser recombinadas assim formando novas regras de negócio sem alterar as especificações existentes, produzindo reusabilidade.
- **Manutenção**: Alterações em uma regra de negócio específica que exigem modificação pontual, minimiza o impacto no sistema evitando anomalias ou comportamentos inesperados.

## Aplicando o Padrão Builder

### Aplicação Prática

O uso do padrão `Builder` facilita criar especificações complexas. Métodos encadeados oferecidos pelo `SpecificationBuilder` permitem a construção fluente de especificações compostas. Ao contrário da criação manual de objetos compostos, essa abordagem melhora significativamente a legibilidade do código ao construir regras de negócios complexas.

```go
package specification

type SpecificationBuilder[T Candidate] interface {
 WithSpecification(spec Specification[T]) SpecificationBuilder[T]
 And(spec Specification[T]) SpecificationBuilder[T]
 Or(spec Specification[T]) SpecificationBuilder[T]
 Not() SpecificationBuilder[T]
 Build() Specification[T]
}

type BaseSpecificationBuilder[T Candidate] struct {
 spec Specification[T]
}

func NewSpecificationBuilder[T Candidate]() SpecificationBuilder[T] {
 return &BaseSpecificationBuilder[T]{}
}

func (b *BaseSpecificationBuilder[T]) WithSpecification(spec Specification[T]) SpecificationBuilder[T] {
 b.spec = spec
 return b
}

func (b *BaseSpecificationBuilder[T]) And(spec Specification[T]) SpecificationBuilder[T] {
 b.spec = NewAndSpecification(b.spec, spec)
 return b
}

func (b *BaseSpecificationBuilder[T]) Or(spec Specification[T]) SpecificationBuilder[T] {
 b.spec = NewOrSpecification(b.spec, spec)
 return b
}

func (b *BaseSpecificationBuilder[T]) Not() SpecificationBuilder[T] {
 b.spec = NewNotSpecification(b.spec)
 return b
}

func (b *BaseSpecificationBuilder[T]) Build() Specification[T] {
 return b.spec
}
```

Usando o `SpecificationBuilder`, os desenvolvedores podem combinar várias especificações com uma sintaxe clara e expressiva, como `builder.WithSpecification(spec1).And(spec2).Not().Build()`, para criar uma nova especificação que representa a lógica de negócio desejada.

### Desafios Superados

- **Complexidade de Criação**: Reduz a complexidade envolvida na construção de especificações compostas, encapsulando a lógica de composição.
- **Flexibilidade**: Permite ajustes finos e configurações personalizadas de especificações complexas, sem expor a complexidade interna.

## Implementando o Padrão Compositor

### Aplicação Prática

O padrão Compositor é aplicado ao agrupar múltiplas regras ou especificações em uma estrutura hierárquica que pode ser tratada como uma única entidade. Isso é particularmente útil para sistemas que precisam avaliar em conjunto conjuntos de regras complexas. É possível aplicar políticas compostas por várias regras aos objetivos de negócios, permitindo uma avaliação unificada que abstrai a complexidade das regras individuais.

No contexto de nosso sistema, uma `Policy` pode ser composta por várias `Rules`, onde cada `Rule` é uma aplicação de uma `Specification`. A `Policy` pode então aplicar todas as suas `Rules` a um objeto de negócio, simplificando o processo de validação.

```go
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
 results := make([]R, 0, len(targets))
 errors := make([]error, 0)

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
 newRules := make([]Rule[T, R], 0, len(rules)+1)
 newRules = append(newRules, r)
 newRules = append(newRules, rules...)
 return &combinedRule[T, R]{rules: newRules}
}

type combinedRule[T any, R any] struct {
 rules []Rule[T, R]
}

func (cr *combinedRule[T, R]) Apply(target T) (R, error) {
 var lastResult R
 for _, rule := range cr.rules {
  var err error
  lastResult, err = rule.Apply(target)
  if err != nil {
   return *new(R), err
  }
 }
 return lastResult, nil
}

func (cr *combinedRule[T, R]) Combine(rules ...Rule[T, R]) Rule[T, R] {
 newRules := make([]Rule[T, R], len(cr.rules), len(cr.rules)+len(rules))
 copy(newRules, cr.rules)
 newRules = append(newRules, rules...)
 return &combinedRule[T, R]{rules: newRules}
}

func (cr *combinedRule[T, R]) BatchApply(targets []T) ([]R, []error) {
 results := make([]R, 0, len(targets))
 errors := make([]error, 0)

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
```

### Desafios Superados

- **Uniformidade**: Trata composições complexas de regras da mesma forma que regras individuais, simplificando a interface de aplicação.
- **Manutenção e Expansão**: Novas regras podem ser adicionadas às políticas existentes sem alterar o código de aplicação, facilitando a manutenção e a expansão.

### Exemplo

Segue um exemplo de uso para `Specification`, e espero ver exemplos do uso combinado com `Rule` e `Policy`.

```go
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
```

## Conclusão

A aplicação dos padrões `Specification`, `Builder` e o `Composite` em um sistema de especificações, regras e políticas não apenas aborda desafios comuns de desenvolvimento, mas também estabelece uma fundação sólida para a criação de sistemas flexíveis, extensíveis e fáceis de manter. Ao separar claramente as regras de negócio da lógica de aplicação, promovendo a reutilização de código e simplificando a criação de regras complexas, esses padrões de design elevam a qualidade e a eficiência do desenvolvimento de software.

Estas implementações fornecem um exemplo prático de como conceitos teóricos e padrões de design podem ser aplicados com sucesso para resolver problemas reais de engenharia de software, ao mesmo tempo, em que fornece percepções úteis para desenvolvedores que buscam melhorar a arquitetura e a manutenibilidade de seus sistemas. Podemos enfrentar com confiança a crescente complexidade das demandas de negócios e tecnologia no desenvolvimento de software contemporâneo compreendendo e aplicando esses padrões.
