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
	Username      string
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

type RecipeTags struct {
	Name    string `json:"name"`
	TagType string `json:"type"`
}

type RecipeAdd struct {
	Name        string          `json:"name"`
	Recipe      string          `json:"recipe"`
	Ingredients IngredientsJson `json:"ingredients"`
	Time        int32           `json:"time"`
	Difficulty  int32           `json:"difficulty"`
	Tags        []RecipeTags    `json:"tags"`
}
