package ingredients

type Repository interface {
	GetMalt(idMalt string) (*Malt, error)
	GetHop(idHop string) (*Hop, error)
	GetYeast(idYeast string) (*Yeast, error)
}
