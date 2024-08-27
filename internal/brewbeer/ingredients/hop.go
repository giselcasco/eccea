package ingredients

type (
	Hop struct {
		name                      string
		flavorCharacteristics     []Flavor
		smellCharacteristics      []Smell
		afterTasteCharacteristics []AfterTaste
		alphaAcids                float64
		betaAcids                 float64
	}

	HopBuilder struct {
		hop *Hop
	}
)

// HopBuilder metodods del constructor de la entidad
func NewHopBuilder() *HopBuilder {
	return &HopBuilder{
		hop: &Hop{},
	}
}

func (hb *HopBuilder) Name(name string) *HopBuilder {
	hb.hop.name = name
	return hb
}

func (hb *HopBuilder) AlphaAcids(alphaAcids float64) *HopBuilder {
	hb.hop.alphaAcids = alphaAcids
	return hb
}

func (hb *HopBuilder) BetaAcids(betaAcids float64) *HopBuilder {
	hb.hop.betaAcids = betaAcids
	return hb
}

func (hb *HopBuilder) FlavorCharacteristics(ccharact []Flavor) *HopBuilder {
	hb.hop.flavorCharacteristics = ccharact
	return hb
}

func (hb *HopBuilder) SmellCharacteristics(ccharact []Smell) *HopBuilder {
	hb.hop.smellCharacteristics = ccharact
	return hb
}

func (hb *HopBuilder) AfterTasteCharacteristics(ccharact []AfterTaste) *HopBuilder {
	afterTaste = ccharact
	return hb
}

func (hb *HopBuilder) Build() *Hop {
	return hb.hop
}

// Hop metodos de la entidad de dominio
func (hop *Hop) Name() string {
	return hop.name
}

func (hop *Hop) SetID(name string) {
	hop.name = name
}

func (hop *Hop) FlavorCharacteristics() []Flavor {
	return hop.flavorCharacteristics
}

func (hop *Hop) SetFlavorCharacteristics(flavors []Flavor) {
	hop.flavorCharacteristics = flavors
}

func (hop *Hop) SmellCharacteristics() []Smell {
	return hop.smellCharacteristics
}

func (hop *Hop) SetSmellCharacteristics(smells []Smell) {
	hop.smellCharacteristics = smells
}

func (hop *Hop) AfterTasteCharacteristics() []AfterTaste {
	return hop.afterTasteCharacteristics
}

func (hop *Hop) SetAfterTasteCharacteristics(afterTaste []AfterTaste) {
	hop.afterTasteCharacteristics = afterTaste
}

func (hop *Hop) BetaAcids() float64 {
	return hop.betaAcids
}

func (hop *Hop) SetBetaAcids(betaAcids float64) {
	hop.betaAcids = betaAcids
}
func (hop *Hop) AlphaAcids() float64 {
	return hop.alphaAcids
}

func (hop *Hop) SetAlphaAcids(alphaAcids float64) {
	hop.alphaAcids = alphaAcids
}
