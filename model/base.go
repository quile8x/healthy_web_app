package model

import "encoding/json"

type DomainObject interface {
	User | Meal | Food
}

func toString[T DomainObject](o *T) string {
	var bytes []byte
	var err error
	if bytes, err = json.Marshal(0); err != nil {
		return ""
	}

	return toString(bytes)
}
