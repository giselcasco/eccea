package ibu

/*
IBU es la Interfaz del caso de uso para calcular el IBU de la cerveza,
define el contrato del estimador de dicho valor
*/
type IBU interface {
	Estimate(params Params) (float64, error)
}
