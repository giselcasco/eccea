package repository

import (
	"database/sql"
	"eccea/internal/brewbeer/ingredients"
	"eccea/internal/repository/dto"
	"os"
	"sync"

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
		data := os.Getenv("BREWBEER_DATA")
		newDb, err := sql.Open("sqlite3", data)
		if err != nil {
			panic(err)
		}
		dbInstance = newDb
	})

	return dbInstance
}

func (s *sqliteReader) GetMaltByName(maltName string) (*ingredients.Malt, error) {
	if len(maltName) == 0 {
		return nil, nil
	}

	dbConn := s.getDbConnection()
	queryMalt := `SELECT 
				ingredient_id, name, color, extract_fine_grind, extract_coarse_grind, diastatic_power 
				FROM malt
				WHERE name = ?`

	cursorMalt, err := dbConn.Query(queryMalt, maltName)
	if err != nil {
		return nil, err
	}

	defer cursorMalt.Close()

	if !cursorMalt.Next() {
		return nil, nil
	}

	item := dto.MaltResponse{}
	err = cursorMalt.Scan(
		&item.ID,
		&item.Name,
		&item.ColorSRM,
		&item.ExtractFineGrind,
		&item.ExtractCoarseGrind,
		&item.DiastaticPower)
	if err != nil {
		return nil, err
	}

	if charactList, err := s.getCharacteristics(dbConn, item.ID); err == nil {
		item.Characteristics = charactList
	}
	return item.ToDomain(), nil
}

func (s *sqliteReader) GetHopByName(hopName string) (*ingredients.Hop, error) {
	if len(hopName) == 0 {
		return nil, nil
	}

	dbConn := s.getDbConnection()
	queryHop := `SELECT p.ingredient_id, p.name, p.alpha_acids, p.beta_acids FROM hop p WHERE p.name = ?`

	cursorHop, err := dbConn.Query(queryHop, hopName)
	if err != nil {
		return nil, err
	}

	defer cursorHop.Close()

	if !cursorHop.Next() {
		return nil, nil
	}

	item := dto.HopResponse{}
	err = cursorHop.Scan(
		&item.ID,
		&item.Name,
		&item.AlphaAcids,
		&item.BetaAcids)
	if err != nil {
		return nil, err
	}

	if charactList, err := s.getCharacteristics(dbConn, item.ID); err == nil {
		item.Characteristics = charactList
	}
	return item.ToDomain(), nil
}

func (s *sqliteReader) GetYeastByName(yeastName string) (*ingredients.Yeast, error) {
	dbConn := s.getDbConnection()
	queryYeast := `SELECT 
				ingredient_id, name, temp_min, temp_max, time_min, time_recommended, time_max 
				FROM yeast
				WHERE name = ?`

	cursorYeast, err := dbConn.Query(queryYeast, yeastName)
	if err != nil {
		return nil, err
	}

	defer cursorYeast.Close()

	if !cursorYeast.Next() {
		return nil, nil
	}

	item := dto.YeastResponse{}
	err = cursorYeast.Scan(
		&item.ID,
		&item.Name,
		&item.TemperatureMin,
		&item.TemperatureMax,
		&item.TimeMin,
		&item.TimeRecommended,
		&item.TimeMax)
	if err != nil {
		return nil, err
	}

	return item.ToDomain(), nil
}

func (s *sqliteReader) getCharacteristics(dbConn *sql.DB, ingredientID string) ([]dto.CharacteristicResponse, error) {
	queryCharact := `SELECT c.id, c.description, c.adjetive_id, ic.type, ic.contribution FROM characteristic as c INNER JOIN ingredient_characteristic as ic ON c.id  =  ic.characteristic_id WHERE c.id  = ?`

	cursorCharact, err := dbConn.Query(queryCharact, ingredientID)
	if err != nil {
		return nil, err
	}

	defer cursorCharact.Close()

	charactList := []dto.CharacteristicResponse{}
	charactItem := dto.CharacteristicResponse{}
	for cursorCharact.Next() {
		err = cursorCharact.Scan(
			&charactItem.ID,
			&charactItem.Description,
			&charactItem.AdjetiveID,
			&charactItem.Type,
			&charactItem.Contribution)
		if err != nil {
			return nil, err
		}

		charactList = append(charactList, charactItem)
	}

	return charactList, nil
}
