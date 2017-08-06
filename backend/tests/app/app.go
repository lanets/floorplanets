// Package app provides testing facilities for the app package
package app

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/lanets/floorplanets/backend/app"
	"github.com/lanets/floorplanets/backend/models"
)

type TestApp struct {
	t       *testing.T
	tempDir string
	App     *app.App
}

// NewTestApp creates a TestApp with using a database in a temporary directory
func NewTestApp(t *testing.T) *TestApp {
	tempDir, err := ioutil.TempDir("", "floorplanets-tests")
	if err != nil {
		t.Fatal(err)
	}

	// Setup the database
	databaseFile := path.Join(tempDir, "testdb.sqlite")
	database, err := setupDatabase(databaseFile)
	if err != nil {
		t.Fatal(err)
	}

	// Setup the app
	application := app.NewApp(database)

	testApp := TestApp{
		t:       t,
		tempDir: tempDir,
		App:     application,
	}

	return &testApp
}

func setupDatabase(databaseFile string) (*gorm.DB, error) {
	database, err := gorm.Open("sqlite3", databaseFile)
	if err != nil {
		return nil, err
	}

	models.RunMigrations(database)

	return database, nil
}

func (testApp *TestApp) Close() {
	err := os.RemoveAll(testApp.tempDir)
	if err != nil {
		testApp.t.Fatal(err)
	}
}
