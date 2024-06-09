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
	"time"
)

const (
	welcome = "¡Bienvenido aventurero/a! Gracias por elegir a CervezaDor como asistente en tu aventura `@-@´"
	options = "A continuación se listan las acciones que CervezaDor puede realizar: \n" +
		"1 - Estimar caracteristicas resultantes. \n" +
		"2 - Estimar caracteristicas del sabor. \n" +
		"3 - Estimar caracteristicas del color. \n" +
		"4 - Calcular ABV. \n" +
		"5 - Calcular IBU. \n" +
		"0 - Salir."
	goodBeer         = "¡Buena Birra!"
	inputError       = "Error al leer la entrada "
	readOptionError  = "Error al leer el número de la opción seleccionada %s %s"
	runOptionError   = "Error al ejecutar la opción seleccionada "
	buildParamsError = "Error al cargar los parametros de ejecución"
	buildParamError  = "Error al cargar %s"
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
				fmt.Sprintln(runOptionError, err)
				return
			}
		}
	}
}

type executor func() error

/*
 exeOptions es un mapa que asocia la opción ingresada por el usuario
 con el ejecutor de dicha funcionalidad
 */
var exeOptions = map[int64]executor{
	5: executeIBUUseCase,
}

/*
 executeIBUUseCase es responsable de calcular el IBU,
 para ello consulta al usuario los valores que necesita y 
 carga los servicios y repositorio necesarios para el procesamiento,
 entre ellos el servicio de estimacion de IBU,
 el servicio cocción y el repositorio de ingredientes
*/
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

	fmt.Printf("El IBU estimado es %f: \n", ibuEstimated.IBU)
	time.Sleep(2 * time.Second)
	return nil
}

/*
 askForBoilingParams tiene como objetivo obtener del usuario
 los valores de los parámetros del proceso de cocción
*/
func askForBoilingParams() (*boiling.Params, error) {
	boilingParams := boiling.Params{}

	fmt.Print("Ingrese el tiempo total de hervor en minutos: ")
	okTT, errTT := fmt.Scanln(&boilingParams.TotalTime)
	if errTT != nil || okTT == 0 || boilingParams.TotalTime == 0 {
		fmt.Sprintf(buildParamError, "tiempo total de hervor")
		return nil, errTT
	}

	fmt.Print("Ingrese la densidad inicial: ")
	okID, errID := fmt.Scanln(&boilingParams.InitialDensity)
	if errID != nil || okID == 0 || boilingParams.TotalTime == 0 {
		fmt.Sprintf(buildParamError, " densidad inicial")
		return nil, errID
	}

	fmt.Print("Ingrese la cantidad de litros inicial del mosto: ")
	okWA, errWA := fmt.Scanln(&boilingParams.WortAmount)
	if errWA != nil || okWA == 0 || boilingParams.WortAmount == 0 {
		fmt.Sprintf(buildParamError, " la cantidad inicial del mosto")
		return nil, errWA
	}

	hopAdditions := yes
	for strings.EqualFold(hopAdditions, yes) {
		hopAddition := boiling.HopAdditions{}

		fmt.Print("Ingrese el ID del lúpulo: ")
		okIDH, errIDH := fmt.Scanln(&hopAddition.ID)
		if errIDH != nil || okIDH == 0 || len(hopAddition.ID) == 0 {
			fmt.Sprintf(buildParamError, "el ID del lúpulo")
			return nil, errIDH
		}

		fmt.Print("Ingrese la cantidad del lúpulo en gramos: ")
		okQ, errQ := fmt.Scanln(&hopAddition.Quantity)
		if errQ != nil || okQ == 0 || len(hopAddition.ID) == 0 {
			fmt.Sprintf(buildParamError, "el ID del lúpulo")
			return nil, errQ
		}

		fmt.Print("Ingrese el tiempo de trabajo del lúpulo en minutos: ")
		okTOW, errTOW := fmt.Scanln(&hopAddition.TimeOfWork)
		if errTOW != nil || okTOW == 0 || len(hopAddition.ID) == 0 {
			fmt.Sprintf(buildParamError, "el ID del lúpulo")
			return nil, errTOW
		}

		boilingParams.HopAdditions = append(boilingParams.HopAdditions, hopAddition)
		fmt.Print("Desea ingresar otro lúpulo? yes/no --> ")
		fmt.Scanln(&hopAdditions)
	}

	return &boilingParams, nil
}

/*
 askForMaturationParams tiene como objetivo consultar al usuario
 los valores de los parámetros del proceso de maduración
*/
func askForMaturationParams() (*maturation.Params, error) {
	return nil, nil
}

/*
 askForMacerationParams tiene como objetivo consultar al usuario
 los valores de los parámetros del proceso de maceración
*/
func askForMacerationParams() (*maceration.Params, error) {
	return nil, nil
}

/*
 askForFermentationParams tiene como objetivo consultar al usuario
 los valores de los parámetros del proceso de fermentación
*/
func askForFermentationParams() (*fermentation.Params, error) {
	return nil, nil
}
