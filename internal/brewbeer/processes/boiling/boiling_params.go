package boiling

type (
	Params struct {
		InitialDensity float64
		WortAmount     float64 // WortAmount is the Quantity of wort in liters at the beginning of the process.
		HopAdditions   []Hop   // HopAdditions list of hops included in the process.
	}

	Hop struct {
		NameID string // NameID es el nombre con que se conoce al lupulo.
		// AlphaAcids float64 // AlphaAcids is the value of AlphaAcids the hop
		Quantity   float64 // Quantity in grams.
		TimeOfWork uint64  // TimeOfWork in minutes.
	}
)
