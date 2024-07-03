package maceration

type (
	Params struct {
		WortAmount    float64 // WortAmount es la cantidad en litros del mosto
		MaltAdditions []Malt  // MaltAdditions is a list of malts included in the process.
	}

	Malt struct {
		Quantity   float64 // quantity in grams.
		ColorSRM   float64
		TimeOfWork uint32 // timeOfWork in minutes.
	}
)
