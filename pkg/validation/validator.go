package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
)

func IsValidCPF(cpf string) bool {
	// gera o slug do CPF
	cpfSlug := slug.Make(cpf)

	// define o validador
	validate := validator.New()

	// adiciona o validador de CPF
	validate.RegisterValidation("cpf", func(fl validator.FieldLevel) bool {
		return cpfSlug == slug.Make(fl.Field().String())
	})

	// faz a validação
	err := validate.Var(cpf, "cpf")
	return err == nil
}
