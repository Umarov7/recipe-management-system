package postgres

import (
	"context"
	"database/sql"
	"log"
	"recipe-management/internal/models"
	"recipe-management/internal/storage"

	"github.com/pkg/errors"
)

type RecipeRepo struct {
	db *sql.DB
}

func NewRecipeRepo(db *sql.DB) storage.IRecipeStorage {
	return &RecipeRepo{db: db}
}

func (r *RecipeRepo) Create(ctx context.Context, in *models.NewRecipe) (string, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to begin transaction")
	}

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("failed to rollback transaction: %v\n", rollbackErr)
			}
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				log.Printf("failed to commit transaction: %v\n", commitErr)
			}
		}
	}()

	query1 := `
	insert into recipes
		(category, title, description, instructions, prep_time, cook_time, servings)
	values
		($1, $2, $3, $4, $5, $6, $7, $8)
	returning id
	`

	query2 := `
	insert into ingredients
		(recipe_id, name, quantity, unit)
	values
		($1, $2, $3, $4)
	`

	var id string
	err = tx.QueryRowContext(ctx, query1, in.Category, in.Title, in.Description,
		in.Instructions, in.PrepTime, in.CookTime, in.Servings).Scan(&id)
	if err != nil {
		return "", errors.Wrap(err, "failed to add recipe")
	}

	for _, i := range in.Ingredients {
		_, err = tx.ExecContext(ctx, query2, id, i.Name, i.Quantity, i.Unit)
		if err != nil {
			return "", errors.Wrap(err, "failed to add ingredients")
		}
	}

	return id, nil
}

func (r *RecipeRepo) Read(ctx context.Context, id string) (*models.Recipe, error) {
	query1 := `
	select
		category, title, description, instructions,
		prep_time, cook_time, servings, created_at, updated_at
	from
		recipes
	where
		id = $1
	`

	query2 := `
	select
		name, quantity, unit
	from
		ingredients
	where
		recipe_id = $1
	`

	rec := models.Recipe{ID: id}
	err := r.db.QueryRowContext(ctx, query1, id).Scan(&rec.Category, &rec.Title,
		&rec.Description, &rec.Instructions, &rec.PrepTime, &rec.CookTime,
		&rec.Servings, &rec.CreatedAt, &rec.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(err, "recipe not found")
		}
		return nil, errors.Wrap(err, "failed to read recipe")
	}

	rows, err := r.db.QueryContext(ctx, query2, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read ingredients")
	}

	for rows.Next() {
		var ing models.Ingredient
		err = rows.Scan(&ing.Name, &ing.Quantity, &ing.Unit)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read ingredients")
		}
		rec.Ingredients = append(rec.Ingredients, ing)
	}
	defer rows.Close()

	return &rec, nil
}

func (r *RecipeRepo) Update(ctx context.Context, in *models.NewRecipeData) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("failed to rollback transaction: %v\n", rollbackErr)
			}
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				log.Printf("failed to commit transaction: %v\n", commitErr)
			}
		}
	}()

	query1 := `
	update
		recipes
	set
		category = $1, title = $2, description = $3, instructions = $4,
		prep_time = $5, cook_time = $6, servings = $7, updated_at = now()
	where
		id = $8
	`

	query2 := `
	insert into ingredients
		(recipe_id, name, quantity, unit)
	values
		($1, $2, $3, $4)
	on conflict
		(recipe_id, name)
	do nothing
	`

	res, err := tx.ExecContext(ctx, query1, in.Category, in.Title, in.Description, in.Instructions,
		in.PrepTime, in.CookTime, in.Servings, in.ID)
	if err != nil {
		return errors.Wrap(err, "failed to update recipe")
	}

	n, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to get rows affected")
	}
	if n < 1 {
		return errors.New("recipe not found")
	}

	for _, i := range in.Ingredients {
		_, err = tx.ExecContext(ctx, query2, in.ID, i.Name, i.Quantity, i.Unit)
		if err != nil {
			return errors.Wrap(err, "failed to update ingredients")
		}
	}

	return nil
}

func (r *RecipeRepo) Delete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx, "delete from recipes where id = $1", id)
	if err != nil {
		return errors.Wrap(err, "failed to delete recipe")
	}

	n, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to get rows affected")
	}
	if n < 1 {
		return errors.New("recipe not found")
	}

	return nil
}

func (r *RecipeRepo) Fetch(ctx context.Context, p *models.Pagination) ([]*models.RecipeInfo, error) {
	query := `
	select
		id, category, title, description
	from
		recipes
	limit $1 offset $2
	`

	rows, err := r.db.QueryContext(ctx, query, p.Limit, p.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read recipes")
	}

	var recipes []*models.RecipeInfo
	for rows.Next() {
		var rec models.RecipeInfo
		err = rows.Scan(&rec.ID, &rec.Category, &rec.Title, &rec.Description)
		if err != nil {
			return nil, errors.Wrap(err, "failed to fetch recipe")
		}

		recipes = append(recipes, &rec)
	}
	defer rows.Close()

	return recipes, nil
}
