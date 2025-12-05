package main

import (
	"aoc/helpers"
	"fmt"
	"sort"
	"strings"
)

type Range struct {
	Min int
	Max int
}

type Ingredient struct {
	Id int
}

func main() {
	data := helpers.ReadFile("input/2025/day05.txt")
	splitData := strings.Split(data, "\n\n")

	idRangeData := strings.Split(splitData[0], "\n")
	idRanges := make([]Range, len(idRangeData))
	for i, idRange := range idRangeData {
		splitId := strings.Split(idRange, "-")
		idRanges[i].Min = helpers.StringToInt(splitId[0])
		idRanges[i].Max = helpers.StringToInt(splitId[1])
	}

	ingredientData := strings.Split(splitData[1], "\n")
	ingredients := make([]Ingredient, len(ingredientData))
	for i, id := range ingredientData {
		ingredients[i].Id = helpers.StringToInt(id)
	}

	mergedIdRanges := mergeIdRanges(idRanges)

	freshIngredientMap := make(map[Ingredient]struct{})
	totalFreshIngredients := 0

	for _, idRange := range mergedIdRanges {
		for _, ingredient := range ingredients {
			if ingredient.Id >= idRange.Min && ingredient.Id <= idRange.Max {
				freshIngredientMap[ingredient] = struct{}{}
			}
		}

		totalFreshIngredients += idRange.Max - idRange.Min + 1
	}

	fmt.Printf("fresh ingredients: %s\n", fmt.Sprint(len(freshIngredientMap)))
	fmt.Printf("possible fresh ingredients: %s\n", fmt.Sprint(totalFreshIngredients))
}

func mergeIdRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})

	merged := []Range{}
	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]

		if next.Min <= current.Max+1 {
			current.Max = max(current.Max, next.Max)
		} else {
			merged = append(merged, current)
			current = next
		}
	}

	merged = append(merged, current)

	return merged
}
