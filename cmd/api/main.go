package main

import (
	"bufio"
	"eccea/internal/brewbeer/estimators/abv"
	"eccea/internal/brewbeer/estimators/color"
	"eccea/internal/brewbeer/estimators/flavor"
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
	4: executeABVUseCase,
	3: executeColorUseCase,
	2: executeFlavorUseCase,
}

/*
executeIBUUseCase es responsable de calcular el IBU,
para ello consulta al usuario los valores que necesita y
carga los servicios y repositorio necesarios para el procesamiento,
entre ellos el servicio de estimacion de IBU,
el servicio del proceso de cocción
y el repositorio de ingredientes
*/
func executeIBUUseCase() error {
	repository := repository.NewSqliteReader()
	boilingService := boiling.NewService(repository)
	estimator := ibu.NewIBUImpl(boilingService)

	boilingParams, err := askForBoilingParams()
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}

	ibuParams := ibu.Params{Boiling: *boilingParams}
	ibuEstimated, estimatorError := estimator.Estimate(ibuParams)
	if estimatorError != nil || ibuEstimated <= 0 {
		fmt.Printf("No fue posible calcula el IBU con los valores suministrados\r\n")
		time.Sleep(2 * time.Second)
		return estimatorError
	}

	ibuValue := fmt.Sprintf("%.1f", ibuEstimated)
	fmt.Printf("El IBU estimado es %s: \r\n\n", ibuValue)
	time.Sleep(2 * time.Second)
	return nil
}

/*
executeABVUseCase es responsable de calcular el volumen de alcohol en la cerveza,
para ello consulta al usuario los valores que necesita y
carga los servicios necesarios para dicho calculo,
entre ellos el servicio de estimacion de ABV y
el servicio del proceso de fermentación
*/
func executeABVUseCase() error {
	fermentationService := fermentation.NewService()
	estimator := abv.NewABVImpl(fermentationService)

	fermentationParams, err := askForFermentationParams()
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}

	abvParams := abv.Params{Fermentation: *fermentationParams}
	abvEstimated, estimatorError := estimator.Estimate(abvParams)
	if estimatorError != nil || abvEstimated <= 0 {
		fmt.Printf("No fue posible calcula el ABV con los valores suministrados\r\n")
		time.Sleep(2 * time.Second)
		return estimatorError
	}

	abvValue := fmt.Sprintf("%.1f", abvEstimated)
	fmt.Printf("El ABV estimado es %s: \r\n", abvValue)
	time.Sleep(3 * time.Second)
	return nil
}

/*
executeColorUseCase es responsable de calcular
el valor del color resultante y las caracteristicas del color asociadas,
para ello consulta al usuario los valores que necesita y
carga los servicios y repositorio necesarios para el procesamiento,
entre ellos el servicio de estimacion de color,
el servicio del proceso de maceración,
el servicio del proceso de maduración
y el repositorio de ingredientes
*/
func executeColorUseCase() error {
	repository := repository.NewSqliteReader()
	macerationService := maceration.NewService(repository)
	maturationService := maturation.NewService()
	estimator := color.NewColorImpl(macerationService, maturationService)

	macerationParams, err := askForMacerationParams()
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}

	maturationParams, err := askForMaturationParams()
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}
	colorParams := color.Params{Maceration: *macerationParams, Maturation: *maturationParams}
	colorEstimated, estimatorError := estimator.Estimate(colorParams)
	if estimatorError != nil || colorEstimated == nil {
		fmt.Printf("No fue posible estimar el color con los valores suministrados\r\n")
		time.Sleep(2 * time.Second)
		return estimatorError
	}

	fmt.Printf("En base a los parámetros ingresados, se estima que la cerveza tendrá: \r\n"+
		"COLOR SRM %d - %s \r\n Caracteristicas del color: %s\r\n",
		colorEstimated.ColorSRM,
		colorEstimated.ColorDescription,
		colorEstimated.ColorCharacteristics)
	fmt.Printf("Segun la cantidad de días de maduración, se estima que la cerveza tendrá %s \r\n",
		colorEstimated.ColorIntensityDescription)

	time.Sleep(3 * time.Second)
	return nil
}

/*
executeFlavorUseCase es responsable de estimar las caracteristicas del sabor
asociadas a las maltas y lúpulos que intervienen en la elaboración,
para ello consulta al usuario los valores que necesita y
carga los servicios y repositorio necesarios para el procesamiento,
entre ellos el servicio de estimacion de sabor,
el servicio del proceso de maceración,
el servicio del proceso de cocción,
el servicio del proceso de maduración
y el repositorio de ingredientes
*/
func executeFlavorUseCase() error {
	repository := repository.NewSqliteReader()
	macerationService := maceration.NewService(repository)
	boilingService := boiling.NewService(repository)
	maturationService := maturation.NewService()
	estimator := flavor.NewFlavorImpl(macerationService, boilingService, maturationService)

	macerationParams, err := askForMacerationParams()
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}

	boilingParams, err := askForBoilingParams()
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}

	maturationParams, err := askForMaturationParams()
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}
	flavorParams := flavor.Params{
		Maceration: *macerationParams,
		Boiling:    *boilingParams,
		Maturation: *maturationParams,
	}
	flavorEstimated, estimatorError := estimator.Estimate(flavorParams)
	if estimatorError != nil || flavorEstimated == nil {
		fmt.Printf("No fue posible estimar las caracteristicas del sabor con los valores suministrados\r\n")
		time.Sleep(2 * time.Second)
		return estimatorError
	}

	fmt.Printf("En base a los parámetros ingresados, se estima que la cerveza tendrá las siguientes caracteristicas referidas a sabor: \r\n: ")
	if len(flavorEstimated.FlavorCharacteristics) > 0 {
		fmt.Printf("%s\r\n", flavorEstimated.FlavorCharacteristics)
	}
	if len(flavorEstimated.SmellCharacteristics) > 0 {
		fmt.Printf("%s\r\n", flavorEstimated.SmellCharacteristics)
	}
	if len(flavorEstimated.AfterTasteCharacteristics) > 0 {
		fmt.Printf("%s\r\n", flavorEstimated.AfterTasteCharacteristics)
	}

	time.Sleep(3 * time.Second)
	return nil
}

/*
askForBoilingParams tiene como objetivo obtener del usuario
los valores de los parámetros del proceso de cocción
*/
func askForBoilingParams() (*boiling.Params, error) {
	boilingParams := boiling.Params{}

	fmt.Print("\r\nIngrese la densidad inicial: ")
	okID, errID := fmt.Scanln(&boilingParams.InitialDensity)
	if errID != nil || okID == 0 || boilingParams.InitialDensity == 0 {
		fmt.Sprintf(buildParamError, " densidad inicial")
		return nil, errID
	}

	fmt.Print("Ingrese la cantidad de litros inicial del mosto: ")
	okWA, errWA := fmt.Scanln(&boilingParams.WortAmount)
	if errWA != nil || okWA == 0 || boilingParams.WortAmount == 0 {
		fmt.Sprintf(buildParamError, " la cantidad inicial del mosto")
		return nil, errWA
	}

	moreAdditions := yes
	for strings.Contains(yes, moreAdditions) {
		hopAddition := boiling.Hop{}

		fmt.Print("Ingrese el porcentaje de Alfa-acidos del lúpulo: ")
		okIDH, errAA := fmt.Scanln(&hopAddition.AlphaAcids)
		if errAA != nil || okIDH == 0 || hopAddition.AlphaAcids == 0 {
			fmt.Sprintf(buildParamError, "el porcentaje de Alfa-acidos")
			return nil, errAA
		}

		fmt.Print("Ingrese la cantidad del lúpulo en gramos: ")
		okQ, errQ := fmt.Scanln(&hopAddition.Quantity)
		if errQ != nil || okQ == 0 || hopAddition.Quantity == 0 {
			fmt.Sprintf(buildParamError, "la cantidad del lúpulo en gramos")
		}

		fmt.Print("Ingrese el tiempo de trabajo del lúpulo en minutos: ")
		okTOW, errTOW := fmt.Scanln(&hopAddition.TimeOfWork)
		if errTOW != nil || okTOW == 0 || hopAddition.TimeOfWork == 0 {
			fmt.Sprintf(buildParamError, "el tiempo de trabajo del lúpulo")
			return nil, errTOW
		}

		boilingParams.HopAdditions = append(boilingParams.HopAdditions, hopAddition)
		fmt.Print("Desea ingresar otro lúpulo? yes/no --> ")
		fmt.Scanln(&moreAdditions)
	}

	return &boilingParams, nil
}

/*
askForMaturationParams tiene como objetivo consultar al usuario
los valores de los parámetros del proceso de maduración
*/
func askForMaturationParams() (*maturation.Params, error) {
	params := &maturation.Params{}

	fmt.Print("\r\nIngrese el tiempo de maduración en días: ")
	ok, err := fmt.Scanln(&params.NumberOfDays)
	if err != nil || ok == 0 || params.NumberOfDays == 0 {
		fmt.Sprintf(buildParamError, "el tiempo de maduracion")
		return nil, err
	}

	return params, nil
}

/*
askForMacerationParams tiene como objetivo consultar al usuario
los valores de los parámetros del proceso de maceración
*/
func askForMacerationParams() (*maceration.Params, error) {
	var totalQuantity float64
	params := &maceration.Params{}

	fmt.Print("\r\nIngrese la cantidad de litros del mosto: ")
	ok, err := fmt.Scanln(&params.WortAmount)
	if err != nil || ok == 0 || params.WortAmount == 0 {
		fmt.Sprintf(buildParamError, "la cantidad de litros del mosto")
		return nil, err
	}

	moreAdditions := yes
	for strings.Contains(yes, moreAdditions) {
		maltAddition := maceration.Malt{}

		fmt.Print("Ingrese el nombre de la malta: ")
		ok, err := fmt.Scanln(&maltAddition.NameID)
		if err != nil || ok == 0 || len(maltAddition.NameID) == 0 {
			fmt.Sprintf(buildParamError, "el nombre de la malta")
			return nil, err
		}

		fmt.Print("Ingrese la cantidad en gramos de dicha malta: ")
		okQ, errQ := fmt.Scanln(&maltAddition.Quantity)
		if errQ != nil || okQ == 0 || maltAddition.Quantity == 0 {
			fmt.Sprintf(buildParamError, "la cantidad en gramos de la malta")
		}

		totalQuantity += maltAddition.Quantity
		params.MaltAdditions = append(params.MaltAdditions, maltAddition)
		fmt.Print("Desea ingresar otra malta? yes/no --> ")
		fmt.Scanln(&moreAdditions)
	}

	params.TotalQuantity = totalQuantity
	return params, nil
}

/*
askForFermentationParams tiene como objetivo consultar al usuario
los valores de los parámetros del proceso de fermentación
*/
func askForFermentationParams() (*fermentation.Params, error) {
	fermentationParams := fermentation.Params{}

	fmt.Print("\r\nIngrese la densidad inicial: ")
	okID, errID := fmt.Scanln(&fermentationParams.InitialDensity)
	if errID != nil || okID == 0 || fermentationParams.InitialDensity == 0 {
		fmt.Sprintf(buildParamError, "la densidad inicial")
		return nil, errID
	}

	fmt.Print("Ingrese la densidad final: ")
	okFD, errFD := fmt.Scanln(&fermentationParams.FinalDensity)
	if errFD != nil || okFD == 0 || fermentationParams.FinalDensity == 0 {
		fmt.Sprintf(buildParamError, "la densidad final")
		return nil, errFD
	}

	return &fermentationParams, nil
}
