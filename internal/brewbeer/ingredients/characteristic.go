package ingredients

type (
	// TypeAdjetives hace referencia al indice del mapa que contine
	// las listas de adjetivos que mejor acompañan a la descripcion del ingrediente.
	TypeAdjetives int

	// DegreeAdjetives hace referencia al indice de la lista de adjetivos
	// y se usa para identificar el adjetivo que mejor acompañana a una descripcion
	// teniendo en cuenta el grado o proporcion en que se encuentre presente el ingrediente.
	DegreeAdjetives int

	// MouthfeelDescription es la descripcion del retrogusto (sabor en boca) que aporta un ingrediente.
	MouthfeelDescription string

	// Mouthfeel es la estructura que contiene el adjetivo y descripcion de la caracteristica del sabor en boca
	Mouthfeel struct {
		Type        TypeAdjetives
		Description MouthfeelDescription
	}

	// SmellDescription es la descripcion del aroma que aporta un ingrediente.
	SmellDescription string

	// Smell es la estructura que contiene el adjetivo y descripcion de la caracteristica del aroma
	Smell struct {
		Type        TypeAdjetives
		Description SmellDescription
	}

	// FlavorDescription es la descripcion del sabor que aporta un ingrediente.
	FlavorDescription string

	// Flavor es la estructura que contiene el adjetivo y descripcion de la caracteristica del sabor
	Flavor struct {
		Type        TypeAdjetives
		Description FlavorDescription
	}

	// ColorDescription es la descripcion del color que aporta un ingrediente.
	ColorDescription string

	// Color es la estructura que contiene el adjetivo y descripcion de la caracteristica del color
	Color struct {
		Type        TypeAdjetives
		Description ColorDescription
	}
)

const (
	// FlavorDescriptions
	CandyFlavor   FlavorDescription = "caramelo dulce"
	CaramelFlavor FlavorDescription = "caramelo"

	// ColorDescriptions
	GoldColor      ColorDescription = "tonalidades doradas"
	GoldAmberColor ColorDescription = "ambar dorado"

	space = " "
)

var mapTypeAdjetives = map[TypeAdjetives][]string{
	1: {"sutiles", "suaves", "marcadas", "intensas"},
	2: {"sutil", "suave", "definido", "intenso"},
}

// GetDescriptionColor devuelve la descripcion del color teniendo en cuenta
// el grado en que dicha descripcion puede cumplirse
func (c *Color) GetDescriptionColor(proportion float64) string {
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
