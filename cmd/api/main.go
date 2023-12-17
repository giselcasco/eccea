package main

import (
	"bufio"
	"eccea/internal/brewbeer/estimators/ibu"
	"eccea/internal/brewbeer/processes/boiling"
	"eccea/internal/repository"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const (
		welcome = "¡Bienvenido aventurero/a! Gracias por elegir a CervezaDor como asistente en tu aventura `@-@´"
		options = "A continuación se listan las acciones que CervezaDor puede realizar: \n" +
			"1 - Estimar caracteristicas resultantes. \n " +
			"2 - Estimar caracteristicas del sabor. \n" +
			"3 - Estimar caracteristicas del color. \n " +
			"4 - Calcular ABV. \n" +
			"5 - Calcular IBU. \n" +
			"0 - Salir. "
		goodBeer              = "¡Buena Birra!"
		inputError            = "Error al leer la entrada "
		readOptionError       = "Error al leer el número de la opción seleccionada "
		parameterLoadingError = "Error al cargar parametros "
		runOptionError        = "Error al ejecutar la opción seleccionada "
	)

	for {
		fmt.Println(welcome)
		fmt.Println(options)

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(inputError, err)
			return
		}

		input = strings.TrimSpace(input)
		if strings.EqualFold(input, "0") {
			fmt.Println(goodBeer)
			break
		}

		option, err := strconv.ParseInt(input, 64, 10)
		if err != nil {
			fmt.Println(readOptionError, input, err)
			return
		}

		if askInoutParams, ok := inputOptions[option]; ok {
			params, errInput := askInoutParams()
			if errInput != nil {
				fmt.Println(parameterLoadingError, err)
				return
			}

			exeError := exeOptions[option](params)
			if exeError != nil {
				fmt.Println(runOptionError, err)
				return
			}
		}
	}
}

type askInput func() (interface{}, error)

var inputOptions = map[int64]askInput{
	0: exit,
	1: askForBeerParams,
	2: askForFlavorParams,
	3: askForColorParams,
	4: askForABVParams,
	5: askForIBUParams,
}

func exit() (interface{}, error) {
	return nil, nil
}

func askForBeerParams() (interface{}, error) {
	return nil, nil
}

func askForFlavorParams() (interface{}, error) {
	return nil, nil
}

func askForColorParams() (interface{}, error) {
	return nil, nil
}

func askForIBUParams() (interface{}, error) {
	return nil, nil
}

func askForABVParams() (interface{}, error) {
	return nil, nil
}

type executor func(interface{}) error

var exeOptions = map[int64]executor{
	1: executeNilUseCase,
	2: executeNilUseCase,
	3: executeNilUseCase,
	4: executeNilUseCase,
	5: executeIBUUseCase,
}

func executeIBUUseCase(interface{}) error {
	boilingService := boiling.NewService(repository.NewIngredientsRepository())
	estimator := ibu.NewIBUImpl(boilingService)

	// TODO conver interface to params
	ibuEstimated, estimatorError := estimator.Estimate(ibu.Params{})
	if estimatorError != nil {
		return estimatorError
	}

	fmt.Printf("El IBU estimado es %f: \n", ibuEstimated)
	return nil
}

func executeNilUseCase(interface{}) error {
	return nil
}
