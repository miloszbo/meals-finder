-- name: FilterRecipesByTagNamesAndParams :many
SELECT r.id, r.name, r.time, r.difficulty
FROM recipes r
WHERE
  -- User tags
  (NOT EXISTS (SELECT 1 FROM users_tags ut WHERE ut.username = @username::text) OR

  EXISTS (SELECT 1 FROM recipes_tags rt JOIN users_tags ut ON rt.tag_id = ut.tag_id WHERE
  ut.username = @username::text AND rt.recipe_id = r.id AND rt.tag_id IN (SELECT tag_id FROM users_tags ut WHERE ut.username = @username::text)))

  -- Min preparation time (optional)
  AND (@min_time::int = 0 OR r.time >= @min_time::int)

  -- Max preparation time (optional)
  AND (@max_time::int = 0 OR r.time <= @max_time::int)

  -- Min difficulty (optional)
  AND (@min_difficulty::int = 0 OR r.difficulty >= @min_difficulty::int)

  -- Max difficulty (optional)
  AND (@max_difficulty::int = 0 OR r.difficulty <= @max_difficulty::int)

  -- Type 1 (Dieta): OR within, AND across types
  AND (@diet::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 1
      AND t.name = ANY(@diet::text[])
  ))

  -- Type 2 (Region)
  AND (@region::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 2
      AND t.name = ANY(@region::text[])
  ))

  -- Type 3 (Rodzaj)
  AND (@recipe_type::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 3
      AND t.name = ANY(@recipe_type::text[])
  ))

  -- Type 4 (Alergie): must NOT include any of these
  AND (@allergies::text[] IS NULL OR NOT EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 4
      AND t.name = ANY(@allergies::text[])
  ))

  -- Type 5 (Składniki odżywcze)
  AND (@nutrients::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 5
      AND t.name = ANY(@nutrients::text[])
  ))

  -- Type 6 (Inne)
  AND (@others::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 6
      AND t.name = ANY(@others::text[])
  ))

ORDER BY r.id LIMIT @recipes_limit::int OFFSET @recipes_offset::int;

-- name: GetRecipeWithId :one
SELECT * FROM recipes WHERE id = $1;

-- name: GetAllTags :many
SELECT tt.name AS type_name, t.name AS tag_name
FROM tags t
JOIN tags_types tt ON t.type_id = tt.id
ORDER BY tt.id, t.name;

-- name: CreateRecipe :one
INSERT INTO recipes (name,recipe,ingredients,time,difficulty) VALUES 
(
  @name::text,
  @recipe::text,
  @ingredients,
  @time::int,
  @difficulty::int
) RETURNING id;

-- name: AddTagsForRecipe :exec
INSERT INTO recipes_tags (recipe_id, tag_id) VALUES
(@recipe_id::int, @tag_id::int);

-- name: GetTagId :one
SELECT t.id AS tag_id
FROM tags t WHERE t.name = @name::text;

-- name: InsertReview :exec
INSERT INTO reviews (recipe_id, username, review_score)
VALUES (
  @recipe_id::int,
  @username::text,
  @review_score::int
) ON CONFLICT (recipe_id, username) DO NOTHING;

-- name: GetReview :one
SELECT * FROM reviews r WHERE r.username = @username::text AND r.recipe_id = @recipe_id::int;

-- name: UpdateReview :exec
UPDATE reviews SET review_score = @new_review_score WHERE username = @username::text AND recipe_id = @recipe_id::int;

-- name: GetRatings :one
SELECT json_object_agg(review_score::text, count) AS ratings
FROM (
  SELECT review_score, COUNT(*) AS count
  FROM reviews
  WHERE recipe_id = $1
  GROUP BY review_score
) AS rating_counts;