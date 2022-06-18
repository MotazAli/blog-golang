package utilities

import "github.com/go-playground/validator"

var Validate = validator.New()

//that is mean T is the generic type which is slice of Element where the Element is any type
func GRemoveElementByIndex[T ~[]Element , Element any] (s T, index int) T {
    return append(s[:index], s[index+1:]...)
}