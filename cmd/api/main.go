package main

import (
	"bufio"
	"eccea/internal/brewbeer/estimators/ibu"
	"eccea/internal/brewbeer/processes/boiling"
	"eccea/internal/brewbeer/processes/fermentation"
	"eccea/internal/brewbeer/processes/maceration"
	"eccea/internal/brewbeer/processes/maturation"
	"eccea/internal/repository"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	welcome = "¡Bienvenido aventurero/a! Gracias por elegir a CervezaDor como asistente en tu aventura `@-@´"
	options = "A continuación se listan las acciones que CervezaDor puede realizar: \n" +
		"1 - Estimar caracteristicas resultantes. \n " +
		"2 - Estimar caracteristicas del sabor. \n" +
		"3 - Estimar caracteristicas del color. \n " +
		"4 - Calcular ABV. \n" +
		"5 - Calcular IBU. \n" +
		"0 - Salir. "
	goodBeer         = "¡Buena Birra!"
	inputError       = "Error al leer la entrada "
	readOptionError  = "Error al leer el número de la opción seleccionada %s %s"
	runOptionError   = "Error al ejecutar la opción seleccionada "
	buildParamsError = "Error al cargar los parametros de ejecución"
	yes              = "yes"
)

func main() {

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

		option, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Printf(readOptionError, input, err)
			return
		}

		if exeOption, ok := exeOptions[option]; ok {
			exeError := exeOption()
			if exeError != nil {
				fmt.Println(runOptionError, err)
				return
			}
		}
	}
}

type executor func() error

var exeOptions = map[int64]executor{
	1: executeNilUseCase,
	2: executeNilUseCase,
	3: executeNilUseCase,
	4: executeNilUseCase,
	5: executeIBUUseCase,
}

func executeIBUUseCase() error {
	repository := repository.NewIngredientsRepository()
	boilingService := boiling.NewService(repository)
	estimator := ibu.NewIBUImpl(boilingService)

	boilingParams, err := askForBoilingParams()
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}
	ibuParams := ibu.Params{Boiling: *boilingParams}
	ibuEstimated, estimatorError := estimator.Estimate(ibuParams)
	if estimatorError != nil {
		return estimatorError
	}

	fmt.Printf("El IBU estimado es %f: \n", ibuEstimated)
	return nil
}

func executeNilUseCase() error {
	return nil
}

func askForMaturationParams() (*maturation.Params, error) {
	return nil, nil
}

func askForMacerationParams() (*maceration.Params, error) {
	return nil, nil
}

func askForFermentationParams() (*fermentation.Params, error) {
	return nil, nil
}

func askForBoilingParams() (*boiling.Params, error) {
	boilingParams := boiling.Params{}

	// Solicitar al usuario los valores para cada campo
	fmt.Print("Ingrese el tiempo total de hervor en minutos: ")
	fmt.Scanln(&boilingParams.TotalTime)

	fmt.Print("Ingrese la cantidad de litros inicial del mosto: ")
	fmt.Scanln(&boilingParams.WortAmount)

	fmt.Print("Ingrese la densidad inicial: ")
	fmt.Scanln(&boilingParams.InitialDensity)

	fmt.Print("Ingrese el volumen del mosto: ")
	fmt.Scanln(&boilingParams.Volume)

	hopAdditions := yes
	for strings.EqualFold(hopAdditions, yes) {
		hopAddition := boiling.HopAdditions{}

		fmt.Print("Ingrese el ID del lúpulo: ")
		fmt.Scanln(&hopAddition.ID)

		fmt.Print("Ingrese la cantidad del lúpulo en gramos: ")
		fmt.Scanln(&hopAddition.Quantity)

		fmt.Print("Ingrese el tiempo de trabajo del lúpulo en minutos: ")
		fmt.Scanln(&hopAddition.TimeOfWork)

		fmt.Print("Desea ingresar otro lúpulo? yes/no --> ")
		fmt.Scanln(&hopAdditions)
	}

	return &boilingParams, nil
}
