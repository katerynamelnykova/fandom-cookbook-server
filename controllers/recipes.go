package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/katerynamelnykova/fandom-cookbook-server/models"
	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRecipe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		id := chi.URLParam(r, "id")
		bookId, idErr := primitive.ObjectIDFromHex(id)
		if idErr != nil {
			http.Error(w, fmt.Sprint("Invalid id"), 400)
			return
		}

		rid := chi.URLParam(r, "rid")
		recipeId, idErr := primitive.ObjectIDFromHex(rid)
		if idErr != nil {
			http.Error(w, fmt.Sprint("Invalid id"), 400)
			return
		}

		book := &models.FullBook{}
		err := mh.GetOneFullBook(book, bson.M{"_id": bookId})
		if err != nil {
			http.Error(w, fmt.Sprintf("Not found"), 404)
			return
		}

		foundRecipe, err := findRecipe(recipeId, book.Recipes)

		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		fandom := map[string]interface{}{
			"bgColor":    book.BgColor,
			"frontColor": book.FrontColor,
			"textColor":  book.TextColor,
		}

		response := map[string]interface{}{
			"recipe": foundRecipe,
			"fandom": fandom,
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 200, response)
	}
}

func findRecipe(id primitive.ObjectID, list []models.Recipe) (models.Recipe, error) {
	var foundRecipe models.Recipe
	for _, recipe := range list {
		if recipe.ID == id {
			foundRecipe = recipe
			break
		}
	}
	if foundRecipe.ID.IsZero() {
		return foundRecipe, fmt.Errorf("Missing id")
	}
	return foundRecipe, nil
}
