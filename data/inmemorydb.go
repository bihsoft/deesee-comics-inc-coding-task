package data

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"task/deeSeeComics/domain/models"
	"task/deeSeeComics/utils"
)

var superHeroes = []models.SuperheroDto{}

func SuperHeroes(filter string) []models.SuperheroDto {
	if len(strings.TrimSpace(filter)) == 0 {
		return superHeroes
	} else {
		var results []models.SuperheroDto
		for _, v := range superHeroes {
			if containsString(filter, v.SuperPowers) {
				results = append(results, v)
			}
		}
		return results
	}
}

func containsString(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Migrate() {
	path := "superheroes.json"
	data := readFile(path)
	if err := json.Unmarshal([]byte(data), &superHeroes); err != nil {
		panic(err)
	}
}

func readFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	utils.CheckErr(err)
	return data
}
