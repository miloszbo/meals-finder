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
