-- name: FilterRecipesByTagNamesAndParams :many
SELECT r.name, r.synopsis, r.time, r.difficulty
FROM recipes r
WHERE
  -- Min preparation time (optional)
  (@min_time::int = 0 OR r.time >= @min_time::int)

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

  -- Type 5 (Składniki_odżywcze)
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

ORDER BY r.id;
