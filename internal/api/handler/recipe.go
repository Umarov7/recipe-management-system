package handler

import (
	"errors"
	"recipe-management/internal/models"

	"github.com/gin-gonic/gin"
)

// Add godoc
// @Summary Adds a new recipe
// @Description Adds a new recipe
// @Tags recipes
// @Param data body models.NewRecipe true "New recipe data"
// @Success 201 {object} string "New recipe ID"
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router / [post]
func (h *Handler) Add(c *gin.Context) {
	var recipe models.NewRecipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		h.logger.Error(err.Error())
		return
	}

	id, err := h.storage.Create(c, &recipe)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		h.logger.Error(err.Error())
		return
	}

	c.JSON(201, gin.H{"New recipe ID": id})
}

// Get godoc
// @Summary Gets a recipe by ID
// @Description Gets a recipe by ID
// @Tags recipes
// @Param id path string true "Recipe ID"
// @Success 200 {object} models.Recipe
// @Failure 400 {object} string "Recipe ID is required"
// @Failure 500 {object} string "Server error while processing request"
// @Router /{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "Recipe ID is required"})
		h.logger.Error(errors.New("recipe ID not provided").Error())
		return
	}

	recipe, err := h.storage.Read(c, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		h.logger.Error(err.Error())
		return
	}

	c.JSON(200, gin.H{"Recipe": recipe})
}

// Update godoc
// @Summary Updates a recipe
// @Description Updates a recipe
// @Tags recipes
// @Param id path string true "Recipe ID"
// @Param data body models.NewRecipeDataNoID true "Recipe data"
// @Success 200 {object} string "Recipe updated"
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "Recipe ID is required"})
		h.logger.Error(errors.New("recipe ID not provided").Error())
		return
	}

	var recipe models.NewRecipeDataNoID
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		h.logger.Error(err.Error())
		return
	}

	err := h.storage.Update(c, &models.NewRecipeData{
		ID:           id,
		Category:     recipe.Category,
		Title:        recipe.Title,
		Description:  recipe.Description,
		Ingredients:  recipe.Ingredients,
		Instructions: recipe.Instructions,
		PrepTime:     recipe.PrepTime,
		CookTime:     recipe.CookTime,
		Servings:     recipe.Servings,
	})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		h.logger.Error(err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Recipe updated"})
}

// Delete godoc
// @Summary Deletes a recipe by ID
// @Description Deletes a recipe by ID
// @Tags recipes
// @Param id path string true "Recipe ID"
// @Success 200 {object} string "Recipe deleted"
// @Failure 400 {object} string "Recipe ID is required"
// @Failure 500 {object} string "Server error while processing request"
// @Router /{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "Recipe ID is required"})
		h.logger.Error(errors.New("recipe ID not provided").Error())
		return
	}

	if err := h.storage.Delete(c, id); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		h.logger.Error(err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Recipe deleted"})
}

// List godoc
// @Summary Lists all recipes
// @Description Lists all recipes, paginated
// @Tags recipes
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {array} models.RecipeInfo "Recipes"
// @Failure 400 {object} string "Invalid pagination parameters"
// @Failure 500 {object} string "Server error while processing request"
// @Router / [get]
func (h *Handler) List(c *gin.Context) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, err := parseIntQueryParam(pageStr)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid pagination parameters"})
		h.logger.Error(errors.New("invalid pagination parameters").Error())
		return
	}

	limit, err := parseIntQueryParam(limitStr)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid pagination parameters"})
		h.logger.Error(errors.New("invalid pagination parameters").Error())
		return
	}

	recipes, err := h.storage.Fetch(c, &models.Pagination{Limit: limit, Offset: (page - 1) * limit})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		h.logger.Error(err.Error())
		return
	}

	c.JSON(200, gin.H{"Recipes": recipes})
}
