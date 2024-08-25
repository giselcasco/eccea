package repository

import (
	"bytes"
	"context"
	"database/sql"
	"eccea/internal/brewbeer/characteristic"
	"eccea/internal/brewbeer/ingredients"
	"eccea/internal/repository/responses"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"os"

	_ "github.com/mattn/go-sqlite3"
)

var dbInstance *sql.DB

type sqliteReader struct {
	once sync.Once
}

func NewSqliteReader() Reader {
	return &sqliteReader{}
}

func (s *sqliteReader) getDbConnection() *sql.DB {
	s.once.Do(func() {
		newDb, err := sql.Open("sqlite3", "data.sqlite")
		if err != nil {
			panic(err)
		}
		dbInstance = newDb
	})

	return dbInstance
}

type ingredientRepo struct {
}

func NewIngredientsRepository() ingredients.Repository {
	return &ingredientRepo{}
}

func (s *sqliteReader) Get(ctx context.Context, id int) (*ingredients.Malt, error) {
	dbConn := s.getDbConnection()

	query := `SELECT 
				id, description, created_at 
				FROM something
				WHERE id = ?`

	cursor, err := dbConn.Query(query, id)
	if err != nil {
		return nil, err
	}

	defer cursor.Close()

	if cursor.Next() {
		item := MaltResponse{}
		err = cursor.Scan(&item.Name,
			&item.ColorSRM)

		if err != nil {
			return nil, err
		}

		return item.ToDomain(), nil
	}

	return nil, nil
}

func (i *ingredientRepo) GetMalt(idMalt string) (*ingredients.Malt, error) {

	// TODO getMalt from sqlite
	colorCharacts := []characteristic.Color{{
		Type:        1,
		Description: characteristic.GoldColor,
	}}
	maltBuilder := ingredients.NewMaltBuilder()
	maltBuilder.ColorSRM(10)
	maltBuilder.ColorCharacteristics(colorCharacts)
	return maltBuilder.Build(), nil
}

func (i *ingredientRepo) GetHop(hopID string) (*ingredients.Hop, error) {
	hop := &responses.Hop{}
	pathHopResources := fmt.Sprintf("internal/resources/hops/%s.json", hopID)
	if i.getResource(pathHopResources, hop); hop.ID == hopID {
		return hop.ToDomain(), nil
	}

	return nil, errors.New("el lupulo no se encuentra en nuestra base de datos")
}

func (i *ingredientRepo) GetYeast(idYeast string) (*ingredients.Yeast, error) {
	return nil, nil
}

func (i *ingredientRepo) getResource(path string, ingredient interface{}) {
	byteArray, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	buffer := new(bytes.Buffer)
	json.Compact(buffer, byteArray)

	if errUnmarshal := json.Unmarshal([]byte(buffer.String()), ingredient); err != nil {
		panic(errUnmarshal)
	}
}
