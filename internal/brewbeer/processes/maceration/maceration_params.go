package maceration

type (
	Params struct {
		WortAmount    float64 // WortAmount es la cantidad en litros del mosto.
		TotalQuantity float64 // TotalQuantity en gramos de las maltas utilizadas.
		MaltAdditions []Malt  // MaltAdditions is a list of malts included in the process.
	}

	Malt struct {
		NameID   string  // NameID es el nombre con que se conoce la malta.
		Quantity float64 // Quantity en gramos de la malta.
	}
)
