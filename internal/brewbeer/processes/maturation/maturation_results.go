package maturation

type (
	Results struct {
		flavorMaturation string
		colorMaturation  string
	}
)

func (r *Results) FlavorMaturation() string {
	return r.flavorMaturation
}

func (r *Results) ColorMaturation() string {
	return r.colorMaturation
}

func NewResults(colorMaturation, flavorMaturation string) Results {
	return Results{
		flavorMaturation: flavorMaturation,
		colorMaturation:  colorMaturation,
	}
}
