package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/lanets/floorplanets/backend/api"
	"github.com/lanets/floorplanets/backend/app"
	"github.com/lanets/floorplanets/backend/models"
)

type ApiTest struct {
	router       *mux.Router
	app          *app.App
	databaseFile string
	t            *testing.T
}

func NewApiTest(t *testing.T) *ApiTest {
	// Setup the database
	databaseFile := "testdb.sqlite"
	database, err := setupDatabase(databaseFile)
	if err != nil {
		t.Fatal(err)
	}

	// Setup the app
	application := app.NewApp(database)

	// Create the ApiTest
	apiTest := ApiTest{
		router:       api.NewRouter(application),
		databaseFile: databaseFile,
		app:          application,
	}

	return &apiTest
}

func setupDatabase(databaseFile string) (*gorm.DB, error) {
	database, err := gorm.Open("sqlite3", databaseFile)
	if err != nil {
		return nil, err
	}

	models.RunMigrations(database)

	return database, nil
}

func (apitest *ApiTest) ServeHTTP(req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	apitest.router.ServeHTTP(res, req)
	return res
}

func (apitest *ApiTest) Close() {
	err := os.Remove(apitest.databaseFile)
	if err != nil {
		apitest.t.Fatal(err)
	}
}
