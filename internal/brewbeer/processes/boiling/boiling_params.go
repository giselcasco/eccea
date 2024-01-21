package boiling

type (
	Params struct {
		TotalTime      uint32 // TotalTime is the total time in minutes of boiling process.
		InitialDensity uint32
		WortAmount     uint32         // WortAmount is the Quantity of wort in liters at the beginning of the process..
		HopAdditions   []HopAdditions // HopAdditions list of hops included in the process.
	}

	HopAdditions struct {
		ID         string
		Quantity   float32 // Quantity in grams.
		TimeOfWork uint32  // TimeOfWork in minutes.
	}
)
