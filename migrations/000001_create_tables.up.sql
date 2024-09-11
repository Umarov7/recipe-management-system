CREATE TABLE recipes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category VARCHAR(100) NOT NULL,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    instructions TEXT NOT NULL,
    prep_time INTEGER, -- Preparation time in minutes
    cook_time INTEGER, -- Cooking time in minutes
    servings INTEGER, -- Number of servings
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ingredients (
    id SERIAL PRIMARY KEY,
    recipe_id UUID REFERENCES recipes(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    quantity DECIMAL(10, 2), -- Quantity of the ingredient
    unit VARCHAR(20), -- Unit of measurement (e.g., grams, cups)
    UNIQUE (recipe_id, name)
);
