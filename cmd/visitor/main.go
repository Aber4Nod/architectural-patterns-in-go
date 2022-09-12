package main

import (
	"github.com/aber4nod/architectural-patterns-in-go/pkg/visitor/company"
	"github.com/aber4nod/architectural-patterns-in-go/pkg/visitor/visitor"
)

func main() {
	auditor := visitor.NewVisitor()
	for _, auditedCompany := range [...]company.Company{
		company.NewSteelMill(),
		company.NewChemicalFactory(),
		company.NewCarFactory(),
	} {
		auditedCompany.Accept(auditor)
	}
}
