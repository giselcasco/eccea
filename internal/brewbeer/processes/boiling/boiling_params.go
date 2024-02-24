package boiling

type (
	Params struct {
		TotalTime      uint64 // TotalTime is the total time in minutes of boiling process.
		InitialDensity uint64
		WortAmount     uint64         // WortAmount is the Quantity of wort in liters at the beginning of the process..
		HopAdditions   []HopAdditions // HopAdditions list of hops included in the process.
	}

	HopAdditions struct {
		ID         string
		Quantity   float64 // Quantity in grams.
		TimeOfWork uint64  // TimeOfWork in minutes.
	}
)
