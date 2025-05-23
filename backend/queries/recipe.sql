-- name: FilterRecipesByTagNamesAndParams :many
SELECT r.id
FROM recipes r
WHERE
  -- Min preparation time (optional)
  ($7 = 0 OR r.time >= $7)

  -- Max preparation time (optional)
  AND ($8 = 0 OR r.time <= $8)

  -- Min difficulty (optional)
  AND ($9 = 0 OR r.difficulty >= $9)

  -- Max difficulty (optional)
  AND ($10 = 0 OR r.difficulty <= $10)

  -- Type 1 (Dieta): OR within, AND across types
  AND ($1::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 1
      AND t.name = ANY($1)
  ))

  -- Type 2 (Region)
  AND ($2::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 2
      AND t.name = ANY($2)
  ))

  -- Type 3 (Rodzaj)
  AND ($3::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 3
      AND t.name = ANY($3)
  ))

  -- Type 4 (Alergie): must NOT include any of these
  AND ($4::text[] IS NULL OR NOT EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 4
      AND t.name = ANY($4)
  ))

  -- Type 5 (Składniki_odżywcze)
  AND ($5::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 5
      AND t.name = ANY($5)
  ))

  -- Type 6 (Inne)
  AND ($6::text[] IS NULL OR EXISTS (
    SELECT 1 FROM recipes_tags rt
    JOIN tags t ON t.id = rt.tag_id
    WHERE rt.recipe_id = r.id
      AND t.type_id = 6
      AND t.name = ANY($6)
  ))

ORDER BY r.id;
