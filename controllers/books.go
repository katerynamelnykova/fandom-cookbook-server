package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/katerynamelnykova/fandom-cookbook-server/models"
	"github.com/katerynamelnykova/fandom-cookbook-server/mongo"
	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/bson"
)

var mh *mongo.MongoHandler

func Connect() error {
	mongoDbConnection := os.Getenv("MONGO_DB_CONNECTION")
	if mongoDbConnection == "" {
		var err error
		return err
	}

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

		fandom := chi.URLParam(r, "fandom")
		project := &models.Book{}
		objectName := fmt.Sprintf("%v", fandom)

		err := mh.GetOneBook(project, bson.M{"fandom": objectName})

		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 200, project)
	}
}
