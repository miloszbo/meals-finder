-- name: FindRecipe :many
SELECT re.name, re.synopsis, re.time, re.difficulty FROM recipes re
