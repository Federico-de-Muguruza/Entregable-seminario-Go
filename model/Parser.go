package model

import (
	"errors"
	"strconv"
	"strings"
)

// type ParserInterface interface {
// 	Parser(string) (*Result, bool)
// }

func Parser(entity string) (Result, error) {

	// Saco todos los espacios
	entity = strings.TrimSpace(entity)

	entityType := string(entity[0]) + string(entity[1])
	entityLength := string(entity[2])
	entityLength = verifyChar(entityLength)
	entityLength += string(entity[3])

	entityValue, entityValueCount := remainingValues(entity)

	entityLengthInteger, err := strconv.Atoi(entityLength)

	if err != nil {
		return newEntity(entityType, entityLengthInteger, entityValue), errors.New("la longitud tiene al menos un caracter y debe ser un número")
	}

	if entityLengthInteger == entityValueCount {
		return newEntity(entityType, entityLengthInteger, entityValue), nil
	}

	return newEntity(entityType, entityLengthInteger, entityValue), errors.New("la longitud declarada no coincide con la longitud real")
}

// Si el primer valor de Length es 0 lo elimino
func verifyChar(entityLength string) string {
	if string(entityLength) == "0" {
		entityLength = strings.Replace(string(entityLength), "0", "", 1)
	}
	return entityLength
}

// Devuelvo todos los caracteres a partir de la posición 4
// y devuelvo la cantidad de caracteres desde la posición 4
func remainingValues(entity string) (string, int) {
	return entity[4:], len(entity[4:])
}

func newEntity(t string, l int, v string) Result {
	return Result{t, l, v}
}

type Result struct {
	Type   string
	Length int
	Value  string
}
