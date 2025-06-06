-- Table: users
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(40) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    passwdhash VARCHAR(60) NOT NULL,
    email VARCHAR(50) NOT NULL,
    name VARCHAR(40) NOT NULL DEFAULT '',
    surname VARCHAR(40) NOT NULL DEFAULT '',
    phone_number VARCHAR(12) NOT NULL,
    age INTEGER NOT NULL,
    sex VARCHAR(13) NOT NULL,
    weight INTEGER NOT NULL DEFAULT 0,
    height INTEGER NOT NULL DEFAULT 0,
    BMI INTEGER NOT NULL DEFAULT 0
);

-- Table: recipes
CREATE TABLE IF NOT EXISTS recipes (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(100) NOT NULL,
    recipe VARCHAR(2500) NOT NULL,
    ingredients JSON NOT NULL,
    time INTEGER NOT NULL, -- Preparation time in minutes
    difficulty INTEGER NOT NULL
);

-- Table: reviews
CREATE TABLE IF NOT EXISTS reviews (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    recipe_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    review_score INTEGER NOT NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Table: ingredients
CREATE TABLE IF NOT EXISTS ingredients (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(30) NOT NULL
);

-- Table: recipes_ingredients
CREATE TABLE IF NOT EXISTS recipes_ingredients (
    recipe_id INTEGER NOT NULL,
    ingredient_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    unit INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id)
);

-- Table: tags_types
CREATE TABLE IF NOT EXISTS tags_types (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(30) NOT NULL
);

-- Table: tags
CREATE TABLE IF NOT EXISTS tags (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(30) NOT NULL,
    type_id INTEGER NOT NULL,
    FOREIGN KEY (type_id) REFERENCES tags_types(id)
);

-- Table: recipes_tags
CREATE TABLE IF NOT EXISTS recipes_tags (
    recipe_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);

-- Table: users_tags
CREATE TABLE IF NOT EXISTS users_tags (
    username VARCHAR(40) NOT NULL,
    tag_id INTEGER NOT NULL
);

-- Indexes
CREATE INDEX idx_reviews_recipe_id ON reviews (recipe_id);
CREATE INDEX idx_recipes_ingredients_recipe_id ON recipes_ingredients (recipe_id);
CREATE INDEX idx_tags_type_id ON tags (type_id);
CREATE INDEX idx_recipes_tags_recipe_id ON recipes_tags (recipe_id);