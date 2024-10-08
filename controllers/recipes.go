package controllers

import (
	"encoding/json"
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

		id := chi.URLParam(r, "fandom")
		objectName := fmt.Sprintf("%v", id)

		rid := chi.URLParam(r, "recipe")
		recipeId, idErr := primitive.ObjectIDFromHex(rid)
		if idErr != nil {
			http.Error(w, fmt.Sprint("Invalid id"), 400)
			return
		}

		book := &models.FullBook{}
		err := mh.GetOneFullBook(book, bson.M{"fandom": objectName})
		if err != nil {
			http.Error(w, fmt.Sprintf("Not found"), 404)
			return
		}

		foundRecipe, err := findRecipe(recipeId, book.Recipes)

		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 200, foundRecipe)
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

func PostDB() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		var database []interface{}
		json.NewDecoder(r.Body).Decode(&database)

		_, err := mh.AddBooks(database)

		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		w.WriteHeader(201)
	}
}
