package models

type RecipesFinderParams struct {
	Diet          []string
	Region        []string
	RecipeType    []string
	Allergies     []string
	Nutrients     []string
	Others        []string
	MinTime       int32
	MaxTime       int32
	MinDifficulty int32
	MaxDifficulty int32
	Limit         int32
	Offset        int32
}

type Ingredient struct {
	Name   string `json:"name"`
	Amount int32  `json:"amount"`
	Unit   string `json:"unit"`
}

type IngredientsJson struct {
	Ingredients []Ingredient `json:"ingredients"`
}

type TagGroup struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}
