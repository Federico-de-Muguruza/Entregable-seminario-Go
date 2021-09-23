package model

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Parser(entity string) (Result, error) {

	// Saco todos los espacios
	entity = strings.TrimSpace(entity)
	entityType := string(entity[0]) + string(entity[1])
	entityLength := verifyLength(string(entity[2]))
	entityLength += string(entity[3])
	entityValue, entityValueCount := remainingValues(entity)

	errType := verifyType(entityType, entityValue)
	entityLengthInteger, errLength := strconv.Atoi(entityLength)

	if errType == nil {
		if errLength != nil {
			return newEntity(entityType, entityLengthInteger, entityValue), errors.New("la longitud tiene al menos un caracter y debe ser un número")
		}

		if entityLengthInteger == entityValueCount {
			return newEntity(entityType, entityLengthInteger, entityValue), nil
		}

		return newEntity(entityType, entityLengthInteger, entityValue), errors.New("la longitud declarada no coincide con la longitud real")
	}
	return newEntity(entityType, entityLengthInteger, entityValue), errType
}

// Si el primer valor de Length es 0 lo elimino
func verifyLength(entityLength string) string {
	if string(entityLength) == "0" {
		entityLength = strings.Replace(string(entityLength), "0", "", 1)
	}
	return entityLength
}

// Verifico que el tipo sea TX o NN y si el valor de cada una tiene algún caracter mal
func verifyType(entityType string, entityValue string) error {
	if entityType == "TX" {
		return containsNumber(entityValue)
	}

	if entityType == "NN" {
		return containsChar(entityValue)
	}
	return errors.New("el tipo de valor no es TX ni NN")
}

func containsNumber(entityValue string) error {
	for _, c := range entityValue {
		if unicode.IsDigit(c) {
			return errors.New("el tipo es TX y el valor contiene al menos un número")
		}
	}
	return nil
}

func containsChar(entity string) error {
	_, err := strconv.Atoi(entity)
	if err != nil {
		return errors.New("el tipo es NN y el valor contiene al menos un caracter")
	}
	return nil
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
