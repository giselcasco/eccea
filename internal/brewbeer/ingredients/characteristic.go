package ingredients

type (
	// TypeAdjetives hace referencia al indice del mapa que contine
	// las listas de adjetivos que mejor acompañan a la descripcion del ingrediente.
	TypeAdjetives int

	// DegreeAdjetives hace referencia al indice de la lista de adjetivos
	// y se usa para identificar el adjetivo que mejor acompañana a una descripcion
	// teniendo en cuenta el grado o proporcion en que se encuentre presente el ingrediente.
	DegreeAdjetives int

	// AfterTaste es la estructura que contiene el adjetivo y descripcion de la caracteristica del sabor en boca
	AfterTaste struct {
		Type        TypeAdjetives
		Description string // Description descripcion del retrogusto (sabor en boca) que aporta un ingrediente.
	}

	// Smell es la estructura que contiene el adjetivo y descripcion de la caracteristica del aroma
	Smell struct {
		Type        TypeAdjetives
		Description string // SmellDescription descripcion del aroma que aporta un ingrediente.
	}

	// Flavor es la estructura que contiene el adjetivo y descripcion de la caracteristica del sabor
	Flavor struct {
		Type        TypeAdjetives
		Description string // FlavorDescription descripcion del sabor que aporta un ingrediente.
	}

	// Color es la estructura que contiene el adjetivo y descripcion de la caracteristica del color
	Color struct {
		Type        TypeAdjetives
		Description string // Description del color que aporta un ingrediente.
	}
)

const (
	// FlavorDescriptions
	CaramelFlavor = "caramelo"

	// ColorDescriptions
	GoldColor      = "tonalidades doradas"
	GoldAmberColor = "ambar dorado"

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
