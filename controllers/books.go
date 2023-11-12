package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/katerynamelnykova/fandom-cookbook-server/models"
	"github.com/katerynamelnykova/fandom-cookbook-server/mongo"
	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mh *mongo.MongoHandler

func Connect() error {
	config, configErr := models.LoadConfiguration()
	if configErr != nil {
		return configErr
	}
	mongoDbConnection := config.Database

	var err error
	mh, err = mongo.NewHandler(mongoDbConnection)

	return err
}

func ShortFandomsInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		projects, err := mh.GetShortFandomsInfo(bson.M{})

		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}
		renderOutput := render.New()
		renderOutput.JSON(w, 200, projects)
	}
}

func GetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		pid := chi.URLParam(r, "id")
		project := &models.Book{}
		objectId, idErr := primitive.ObjectIDFromHex(pid)

		if idErr != nil {
			http.Error(w, fmt.Sprintf("Invalid id"), 400)
			return
		}

		err := mh.GetOneBook(project, bson.M{"_id": objectId})

		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 200, project)
	}
}
