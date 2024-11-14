package main

import (
	"bufio"
	"eccea/internal/brewbeer/estimators/abv"
	"eccea/internal/brewbeer/estimators/beer"
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
	"unicode"
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
	fmt.Println(welcome)
	for {
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
	1: executeBeerUseCase,
}

/*
executeBeerUseCase es responsable de estimar IBU, ABV,
valor del color resultante, descripción y caracteristicas del color,
asi como caracteristicas del sabor, aroma y retrogusto;
para lograr ello consulta al usuario los valores que necesita y
carga los servicios y repositorio necesarios para el procesamiento
*/
func executeBeerUseCase() error {
	sqliteReader := repository.NewSqliteReader()
	boilingService := boiling.NewService(sqliteReader)
	macerationService := maceration.NewService(sqliteReader)
	maturationService := maturation.NewService()
	fermentationService := fermentation.NewService()
	estimator := beer.NewBeerImpl(macerationService, boilingService, fermentationService, maturationService)

	macerationParams, err := askForMacerationParams(true)
	if err != nil {
		return err
	}

	boilingParams, err := askForBoilingParams(false)
	if err != nil {
		return err
	}

	fermentationParams := &fermentation.Params{}
	fmt.Print("Ingrese la densidad final: ")
	finalDensity, FDErr := scanFloat()
	if FDErr != nil {
		fmt.Sprintf(buildParamError, "la densidad final")
		return FDErr
	}
	fermentationParams.FinalDensity = *finalDensity

	maturationParams, err := askForMaturationParams()
	if err != nil {
		return err
	}

	params := beer.Params{
		Maceration:   macerationParams,
		Boiling:      boilingParams,
		Fermentation: fermentationParams,
		Maturation:   maturationParams,
	}
	estimateResult, estimatorErr := estimator.Estimate(params)
	if estimatorErr != nil {
		fmt.Printf("No fue posible estimar las caracteristicas de la cerveza con los valores suministrados\n")
		time.Sleep(2 * time.Second)
		return estimatorErr
	}

	if len(estimateResult.FlavorCharacteristics) > 0 ||
		len(estimateResult.SmellCharacteristics) > 0 ||
		len(estimateResult.AfterTasteCharacteristics) > 0 {
		fmt.Printf("En base a los parámetros ingresados, se estima que la cerveza tendrá las siguientes caracteristicas: \n")
		fmt.Printf("%s", estimateResult.FlavorCharacteristics)
		fmt.Printf("%s", estimateResult.SmellCharacteristics)
		fmt.Printf("%s", estimateResult.AfterTasteCharacteristics)
		fmt.Printf("%s", estimateResult.ColorCharacteristics)

		fmt.Printf("Segun la cantidad de días de maduración, se estima que la cerveza tendrá:\n -IBU: %f \n -ABV: %f \n -COLOR: %d \n -Descripción del color: %s \n",
			estimateResult.IBU, estimateResult.ABV, estimateResult.ColorSRM, estimateResult.ColorDescription)
	}

	return err
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
	sqliteReader := repository.NewSqliteReader()
	boilingService := boiling.NewService(sqliteReader)
	estimator := ibu.NewIBUImpl(boilingService)

	boilingParams, err := askForBoilingParams(false)
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}

	ibuParams := ibu.Params{Boiling: *boilingParams}
	ibuEstimated, estimatorError := estimator.Estimate(ibuParams)
	if estimatorError != nil || ibuEstimated <= 0 {
		fmt.Printf("No fue posible calcula el IBU con los valores suministrados\n")
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
	sqliteReader := repository.NewSqliteReader()
	macerationService := maceration.NewService(sqliteReader)
	maturationService := maturation.NewService()
	estimator := color.NewColorImpl(macerationService, maturationService)

	macerationParams, err := askForMacerationParams(false)
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
	sqliteReader := repository.NewSqliteReader()
	macerationService := maceration.NewService(sqliteReader)
	boilingService := boiling.NewService(sqliteReader)
	maturationService := maturation.NewService()
	estimator := flavor.NewFlavorImpl(macerationService, boilingService, maturationService)

	macerationParams, err := askForMacerationParams(true)
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}

	boilingParams, err := askForBoilingParams(true)
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}

	maturationParams, err := askForMaturationParams()
	if err != nil {
		fmt.Println(buildParamsError)
		return err
	}

	flavorParams := flavor.NewParams(
		*macerationParams,
		*boilingParams,
		*maturationParams,
	)
	flavorEstimated, estimatorError := estimator.Estimate(*flavorParams)
	if flavorEstimated == nil ||
		(len(flavorEstimated.FlavorCharacteristics) == 0 &&
			len(flavorEstimated.SmellCharacteristics) == 0 &&
			len(flavorEstimated.AfterTasteCharacteristics) == 0) {
		fmt.Printf("No fue posible estimar las caracteristicas del sabor")

		if estimatorError != nil {
			fmt.Printf("- error: %s", estimatorError.Error())
		}
		time.Sleep(2 * time.Second)
		return estimatorError
	}

	if len(flavorEstimated.FlavorCharacteristics) > 0 ||
		len(flavorEstimated.SmellCharacteristics) > 0 ||
		len(flavorEstimated.AfterTasteCharacteristics) > 0 {
		fmt.Printf("En base a los parámetros ingresados, se estima que la cerveza tendrá las siguientes caracteristicas: \n")
		fmt.Printf("%s", flavorEstimated.FlavorCharacteristics)
		fmt.Printf("%s", flavorEstimated.SmellCharacteristics)
		fmt.Printf("%s", flavorEstimated.AfterTasteCharacteristics)

		fmt.Printf("Segun la cantidad de días de maduración, se estima que la cerveza tendrá %s \r\n",
			flavorEstimated.FlavorMaturation)
	}

	time.Sleep(3 * time.Second)
	return nil
}

/*
askForBoilingParams tiene como objetivo obtener del usuario
los valores de los parámetros del proceso de cocción
*/
func askForBoilingParams(onlyHops bool) (*boiling.Params, error) {
	boilingParams := boiling.Params{}

	if !onlyHops {
		fmt.Print("\r\nIngrese la densidad inicial: ")
		initialDensity, scanErr := scanFloat()
		if scanErr != nil || initialDensity == nil {
			fmt.Sprintf(buildParamError, " densidad inicial")
			return nil, scanErr
		}
		boilingParams.InitialDensity = *initialDensity

		fmt.Print("Ingrese la cantidad de litros del mosto para el proceso de cocción: ")
		wortAmount, scanErr := scanUint()
		if scanErr != nil || wortAmount == nil {
			fmt.Sprintf(buildParamError, " la cantidad de litros del mosto para la cocción")
			return nil, scanErr
		}
		boilingParams.WortAmount = *wortAmount
	}

	moreAdditions := yes
	for strings.Contains(yes, moreAdditions) {
		hopAddition := boiling.Hop{}

		fmt.Print("\r\nIngrese el nombre del lúpulo: ")
		nameID, scanErr := scanCharacter()
		if scanErr != nil || nameID == nil {
			fmt.Sprintf(buildParamError, "el nombre del lúpulo")
			return nil, scanErr
		}
		hopAddition.NameID = *nameID

		fmt.Print("Ingrese la cantidad del lúpulo en gramos: ")
		quantity, scanErr := scanFloat()
		if scanErr != nil || quantity == nil {
			fmt.Sprintf(buildParamError, "la cantidad del lúpulo en gramos")
			return nil, scanErr
		}
		hopAddition.Quantity = *quantity

		fmt.Print("Ingrese el tiempo de trabajo del lúpulo en minutos: ")
		timeOfWork, scanErr := scanUint()
		if scanErr != nil || timeOfWork == nil {
			fmt.Sprintf(buildParamError, "el tiempo de trabajo del lúpulo")
			return nil, scanErr
		}
		hopAddition.TimeOfWork = *timeOfWork

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
	numberOfDays, err := scanUint()
	if err != nil {
		fmt.Sprintf(buildParamError, "el tiempo de maduracion")
		return nil, err
	}
	params.NumberOfDays = *numberOfDays
	return params, nil
}

/*
askForMacerationParams tiene como objetivo consultar al usuario
los valores de los parámetros del proceso de maceración
*/
func askForMacerationParams(onlyMalts bool) (*maceration.Params, error) {
	var totalQuantity float64
	params := &maceration.Params{}

	if !onlyMalts {
		fmt.Print("\r\nIngrese la cantidad de litros inicial del mosto: ")
		wortAmount, scanErr := scanFloat()
		if scanErr != nil || wortAmount == nil {
			fmt.Sprintf(buildParamError, "la cantidad de litros del mosto")
			return nil, scanErr
		}
		params.WortAmount = *wortAmount
	}

	moreAdditions := yes
	for strings.Contains(yes, moreAdditions) {
		maltAddition := maceration.Malt{}

		fmt.Print("\r\nIngrese el nombre de la malta: ")
		nameID, err := scanCharacter()
		if err != nil {
			fmt.Sprintf(buildParamError, "el nombre de la malta")
			return nil, err
		}
		maltAddition.NameID = *nameID

		fmt.Print("Ingrese la cantidad en gramos de dicha malta: ")
		quantity, QErr := scanFloat()
		if QErr != nil {
			fmt.Sprintf(buildParamError, "la cantidad en gramos de la malta")
			return nil, err
		}
		maltAddition.Quantity = *quantity
		totalQuantity += maltAddition.Quantity
		params.MaltAdditions = append(params.MaltAdditions, maltAddition)
		fmt.Print("Desea ingresar otra malta? yes/no --> ")
		addMalt, addErr := scanCharacter()
		if addErr != nil {
			return nil, addErr
		}
		moreAdditions = *addMalt
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
	initialDensity, IDErr := scanFloat()
	if IDErr != nil {
		fmt.Sprintf(buildParamError, "la densidad inicial")
		return nil, IDErr
	}
	fermentationParams.InitialDensity = *initialDensity

	fmt.Print("Ingrese la densidad final: ")
	finalDensity, FDErr := scanFloat()
	if FDErr != nil {
		fmt.Sprintf(buildParamError, "la densidad final")
		return nil, FDErr
	}
	fermentationParams.FinalDensity = *finalDensity
	return &fermentationParams, nil
}

func scanUint() (*uint64, error) {
	var uintValue uint64
	_, err := fmt.Scanln(&uintValue)
	if err != nil {
		return nil, err
	}

	uintValueStr := strconv.FormatUint(uintValue, 10)
	if _, atoiErr := strconv.Atoi(uintValueStr); atoiErr != nil {
		fmt.Println("El valor ingresado no es un número")
		return nil, err
	}

	return &uintValue, nil
}

func scanFloat() (*float64, error) {
	var floatValueStr string
	_, err := fmt.Scanln(&floatValueStr)
	if err != nil {
		return nil, err
	}

	floatValue, parseErr := strconv.ParseFloat(floatValueStr, 64)
	if parseErr != nil {
		fmt.Println("El valor ingresodo no es un número flotante")
		return nil, parseErr
	}

	return &floatValue, nil
}

func scanCharacter() (*string, error) {
	var strValue string
	_, err := fmt.Scanln(&strValue)
	if err != nil {
		return nil, err
	}

	if len(strValue) == 0 || !unicode.IsLetter(rune(strValue[0])) {
		fmt.Println("El valor ingresado es incorrecto")
		return nil, err
	}
	return &strValue, nil
}
