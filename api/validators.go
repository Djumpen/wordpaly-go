package api

import (
	validator "gopkg.in/go-playground/validator.v8"
)

type ValidationFailers []validationFailer

type validationFailer struct {
	NameSpace string
	Field     string
	Rule      string
	RuleValue string
	Value     interface{}
}

func getValidationFailers(err error) (vf ValidationFailers, ok bool) {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, v := range ve {
			vf = append(vf, validationFailer{
				NameSpace: v.NameNamespace,
				Field:     v.Field,
				Rule:      v.Tag,
				RuleValue: v.Param,
				Value:     v.Value,
			})
			// fmt.Printf("KEY %#v\n", k)
			// fmt.Printf("VALUE %#v\n", *v)
		}
		return vf, true
	}
	return nil, false
}
