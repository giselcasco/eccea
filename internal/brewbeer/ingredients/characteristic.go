package ingredients

type (
	// Characteristic es la estructura que contiene el adjetivo y descripcion de la caracteristica del sabor en boca
	Characteristic struct {
		CharacteristicType string  // CharacteristicType hace referencia al tipo de caracteristica, estas pueden ser "aroma", "sabor", "color", "retrogusto", etc.
		TypeAdjetives      int     // TypeAdjetives referencia al indice del mapa que contine los adjetivos que mejor acompañan a la descripcion del ingrediente.
		Description        string  // Description descripcion de la caracteristica que aporta un ingrediente.
		Contribution       float64 // grado en que contribuye la caracteristica en el ingrediente
	}
)

const (
	// FlavorDescriptions
	CaramelFlavor = "caramelo"

	// ColorDescriptions
	GoldColor      = "tonalidades doradas"
	GoldAmberColor = "ambar dorado"

	maxContribution = 5

	space = " "
)

var mapTypeAdjetives = map[int][]string{
	1: {"sutiles", "suaves", "marcadas", "intensas"},
	2: {"sutil", "suave", "definido", "intenso"},
	3: {"sutil", "suave", "definida", "intensa"},
	4: {"sutiles", "suaves", "marcados", "intensos"},
	5: {"dejo a", "notas de", "presencia de"},
	6: {"dejos", "notas", "presencias"},
}

// GetDescriptionByProportion devuelve la descripcion de la caracteristica teniendo en cuenta
// el grado en que dicha descripcion puede cumplirse
func (c *Characteristic) GetDescriptionByProportion(proportion float64) string {
	contributionPercentage := getContributionPercentage(c.Contribution, proportion)
	if adjectives, ok := mapTypeAdjetives[c.TypeAdjetives]; ok && len(adjectives) > 0 {
		partitionBase := 100 / len(adjectives)
		degree := partitionBase
		for _, adjective := range adjectives {
			if contributionPercentage <= float64(degree) {
				return adjective + space + c.Description
			}
			degree += partitionBase
		}
	}
	return space + string(c.Description)
}

func getContributionPercentage(contribution, proportion float64) float64 {
	if contribution == maxContribution {
		return proportion
	}

	return (contribution / maxContribution) * proportion
}
