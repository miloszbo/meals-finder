-- Table: users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(40) UNIQUE,
    created_at TIMESTAMP,
    passwd VARCHAR(40),
    email VARCHAR(40),
    name VARCHAR(40),
    surname VARCHAR(40),
    phone INTEGER,
    age INTEGER,
    sex BOOLEAN,
    weight INTEGER,
    height INTEGER,
    BMI INTEGER
);

-- Table: recipes
CREATE TABLE recipes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    synopsis VARCHAR(100),
    time INTEGER, -- Preparation time in minutes
    difficulty INTEGER
);

-- Table: reviews
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    recipe_id INTEGER,
    user_id INTEGER,
    review INTEGER,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Table: ingredients
CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30)
);

-- Table: recipes_ingredients
CREATE TABLE recipes_ingredients (
    recipe_id INTEGER,
    ingredient_id INTEGER,
    amount INTEGER,
    unit INTEGER DEFAULT 0,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id)
);

-- Table: tags_types
CREATE TABLE tags_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30)
);

-- Table: tags
CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30),
    type_id INTEGER,
    FOREIGN KEY (type_id) REFERENCES tags_types(id)
);

-- Table: recipes_tags
CREATE TABLE recipes_tags (
    recipe_id INTEGER,
    tag_id INTEGER,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);

-- Indexes
CREATE INDEX idx_reviews_recipe_id ON reviews (recipe_id);
CREATE INDEX idx_recipes_ingredients_recipe_id ON recipes_ingredients (recipe_id);
CREATE INDEX idx_tags_type_id ON tags (type_id);
CREATE INDEX idx_recipes_tags_recipe_id ON recipes_tags (recipe_id);
