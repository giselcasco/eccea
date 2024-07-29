package characteristic

type (
	// ColorDescription es la descripcion del color que aporta un ingrediente.
	ColorDescription string

	// TypeAdjetives hace referencia al indice del mapa que contine
	// las listas de adjetivos que mejor acompañan a la descripcion del color.
	TypeAdjetives int

	// DegreeAdjetives hace referencia al indice de la lista de adjetivos
	// y se usa para identificar el adjetivo que mejor acompañana a una descripcion de color
	// teniendo en cuenta el grado o proporcion en que se encuentre presente el ingrediente.
	DegreeAdjetives int

	Color struct {
		Type        TypeAdjetives
		Description ColorDescription
	}
)

const (
	space = " "

	GoldColor      ColorDescription = "tonalidades doradas"
	GoldAmberColor ColorDescription = "ambar dorado"
)

var mapTypeAdjetives = map[TypeAdjetives][]string{
	1: {"suaves", "marcadas", "intensas"},
}

// GetDescriptionMaltColor devuelve la descripcion del color teniendo en cuenta
// el grado en que dicha descripcion puede cumplirse
func (c *Color) GetDescriptionMaltColor(proportion float64) string {
	if adjetives, ok := mapTypeAdjetives[c.Type]; ok && len(adjetives) > 0 {
		partitionBase := 100 / len(adjetives)
		degree := partitionBase
		for _, adjetive := range adjetives {
			if proportion <= float64(degree) {
				return adjetive + space + string(c.Description)
			}
			degree += partitionBase
		}
	}
	return space + string(c.Description)
}
