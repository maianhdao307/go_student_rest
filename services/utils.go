package services

import "github.com/lucasjones/reggen"

type Utils struct{}

type UtilsService interface {
	GenerateID(regex string, limit int) (string, error)
}

func (_self Utils) GenerateID(regex string, limit int) (string, error) {
	result, err := reggen.Generate(regex, limit)
	return result, err
}

