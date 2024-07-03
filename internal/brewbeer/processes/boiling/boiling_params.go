package boiling

type (
	Params struct {
		InitialDensity float64
		WortAmount     uint64 // WortAmount is the Quantity of wort in liters at the beginning of the process.
		HopAdditions   []Hop  // HopAdditions list of hops included in the process.
	}

	Hop struct {
		AlphaAcids float64
		Quantity   float64 // Quantity in grams.
		TimeOfWork uint64  // TimeOfWork in minutes.
	}
)
