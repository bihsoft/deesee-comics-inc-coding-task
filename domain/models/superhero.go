package models

import (
	"reflect"
	"strings"
	"task/deeSeeComics/config"
	"task/deeSeeComics/cryptography"
	"task/deeSeeComics/utils"
)

type SuperheroDto struct {
	Name        string   `json:"name"`
	Identity    Identity `json:"identity"`
	Birthday    string   `json:"birthday"`
	SuperPowers []string `json:"superpowers"`
}

type Superhero struct {
	Name        string   `json:"name"`
	Identity    string   `json:"identity"`
	Birthday    string   `json:"birthday"`
	SuperPowers []string `json:"superpowers"`
}

func (s *SuperheroDto) MapToModel(encrypt bool, useFields ...string) Superhero {
	fs := utils.FieldSet(useFields...)
	rt, rv := reflect.TypeOf(*s), reflect.ValueOf(*s)
	dtoMap := make(map[string]interface{}, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		jsonKey := field.Tag.Get("json")
		if fs[jsonKey] {
			dtoMap[jsonKey] = rv.Field(i).Interface()
		}
	}
	var hero = Superhero{}
	for key, value := range dtoMap {
		if s, ok := value.(string); ok {
			switch key {
			case "name":
				hero.Name = s
			case "birthday":
				hero.Birthday = s
			}
		}
		if s, ok := value.([]string); ok {
			hero.SuperPowers = s
		}
		if s, ok := value.(Identity); ok {
			if encrypt {
				hero.Identity = strings.Map(func(r rune) rune {
					return cryptography.DeeSeeChiffre(r, config.DeeSeeChiffreSeed)
				}, s.FirstName) + " " + strings.Map(func(r rune) rune {
					return cryptography.DeeSeeChiffre(r, config.DeeSeeChiffreSeed)
				}, s.LastName)
			} else {
				hero.Identity = s.FirstName + " " + s.LastName
			}
		}
	}
	return hero
}
