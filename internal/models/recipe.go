package models

type NewRecipe struct {
	Category     string       `json:"category"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions string       `json:"instructions"`
	PrepTime     int          `json:"prep_time"`
	CookTime     int          `json:"cook_time"`
	Servings     int          `json:"servings"`
}

type NewRecipeData struct {
	ID           string       `json:"id"`
	Category     string       `json:"category"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions string       `json:"instructions"`
	PrepTime     int          `json:"prep_time"`
	CookTime     int          `json:"cook_time"`
	Servings     int          `json:"servings"`
}

type NewRecipeDataNoID struct {
	Category     string       `json:"category"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions string       `json:"instructions"`
	PrepTime     int          `json:"prep_time"`
	CookTime     int          `json:"cook_time"`
	Servings     int          `json:"servings"`
}

type Recipe struct {
	ID           string       `json:"id"`
	Category     string       `json:"category"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions string       `json:"instructions"`
	PrepTime     int          `json:"prep_time"`
	CookTime     int          `json:"cook_time"`
	Servings     int          `json:"servings"`
	CreatedAt    string       `json:"created_at"`
	UpdatedAt    string       `json:"updated_at"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Unit     string `json:"unit"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type RecipeInfo struct {
	ID          string `json:"id"`
	Category    string `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
