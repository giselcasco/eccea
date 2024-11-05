package beer

import (
	"eccea/internal/brewbeer/processes/boiling"
	"eccea/internal/brewbeer/processes/fermentation"
	"eccea/internal/brewbeer/processes/maceration"
	"eccea/internal/brewbeer/processes/maturation"
)

type Params struct {
	Maceration   *maceration.Params
	Boiling      *boiling.Params
	Fermentation *fermentation.Params
	Maturation   *maturation.Params
}
