// Code generated by jsonschema. DO NOT EDIT.
package beerjson

import "encoding/json"

// ID: https://raw.githubusercontent.com/beerjson/beerjson/master/json/culture.json

// CultureAdditionType collects the attributes of each culture ingredient for use in a recipe.
type CultureAdditionType struct {
	CultureBase *CultureBase `json:"CultureBase,omitempty"`
	// The expected, or measured apparent attenuation for a given culture in a given recipe. In comparison to attenuation range, this is a single value.
	Attenuation   *PercentType `json:"attenuation,omitempty"`
	TimesCultured *int32       `json:"times_cultured,omitempty"`
	// The timing object fully describes the timing of an addition with options for basis on time, gravity, or pH at any process step.
	Timing            *TimingType               `json:"timing,omitempty"`
	CellCountBillions *int32                    `json:"cell_count_billions,omitempty"`
	Amount            CultureAdditionTypeAmount `json:"amount,omitempty", validate:"oneof"`
}

func (s *CultureAdditionType) UnmarshalJSON(b []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil
	}

	cultureAdditionTypeAmount := func() CultureAdditionTypeAmount {
		raw, ok := m["amount"]
		if !ok {
			return nil
		}

		var volumeType VolumeType
		if err := json.Unmarshal(raw, &volumeType); err == nil {
			return &volumeType
		}

		var massType MassType
		if err := json.Unmarshal(raw, &massType); err == nil {
			return &massType
		}

		var unitType UnitType
		if err := json.Unmarshal(raw, &unitType); err == nil {
			return &unitType
		}

		return nil
	}
	type Alias CultureAdditionType
	aux := &struct {
		Amount CultureAdditionTypeAmount `json:"amount,omitempty", validate:"oneof"`
		*Alias
	}{
		Amount: cultureAdditionTypeAmount(),
		Alias:  (*Alias)(s),
	}

	s.Amount = aux.Amount

	return nil
}

// CultureAdditionTypeAmount
type CultureAdditionTypeAmount interface {
	CultureAdditionTypeamount()
}

// Provides unique properties to identify individual records of a culture.
type CultureBase struct {
	Name            string          `json:"name", validate:"required"`
	CultureBaseType CultureBaseType `json:"type", validate:"required"`
	CultureBaseForm CultureBaseForm `json:"form", validate:"required"`
	Producer        *string         `json:"producer,omitempty"`
	ProductId       *string         `json:"product_id,omitempty"`
}

type CultureBaseForm string

const (
	CultureBaseForm_Liquid  CultureBaseForm = "liquid"
	CultureBaseForm_Dry     CultureBaseForm = "dry"
	CultureBaseForm_Slant   CultureBaseForm = "slant"
	CultureBaseForm_Culture CultureBaseForm = "culture"
	CultureBaseForm_Dregs   CultureBaseForm = "dregs"
)

type CultureBaseType string

const (
	CultureBaseType_Ale          CultureBaseType = "ale"
	CultureBaseType_Bacteria     CultureBaseType = "bacteria"
	CultureBaseType_Brett        CultureBaseType = "brett"
	CultureBaseType_Champagne    CultureBaseType = "champagne"
	CultureBaseType_Kveik        CultureBaseType = "kveik"
	CultureBaseType_Lacto        CultureBaseType = "lacto"
	CultureBaseType_Lager        CultureBaseType = "lager"
	CultureBaseType_Malolactic   CultureBaseType = "malolactic"
	CultureBaseType_MixedCulture CultureBaseType = "mixed-culture"
	CultureBaseType_Other        CultureBaseType = "other"
	CultureBaseType_Pedio        CultureBaseType = "pedio"
	CultureBaseType_Spontaneous  CultureBaseType = "spontaneous"
	CultureBaseType_Wine         CultureBaseType = "wine"
)

// CultureInformation collects the attributes of a microbial culture.
type CultureInformation struct {
	Notes *string `json:"notes,omitempty"`
	// A glucoamylase positive culture is capable of producing glucoamylase, the enzyme produced through expression of the diastatic gene, which allows yeast to attenuate dextrins and starches leading to a very low FG. This is positive in some saison/brett yeasts as well as the new gulo hybrid by Omega yeast labs.
	Glucoamylase *bool                 `json:"glucoamylase,omitempty"`
	Inventory    *CultureInventoryType `json:"inventory,omitempty"`
	Zymocide     *Zymocide             `json:"zymocide,omitempty"`
	// Maximum number of times to reuse a culture before a new lab source is recommended.
	MaxReuse *int32 `json:"max_reuse,omitempty"`
	// Recommended styles for a particular culture.
	BestFor *string `json:"best_for,omitempty"`
	// Floculation refers to the ability of yeast to aggregate to form large flocs which drop out of suspension.
	Flocculation     *QualitativeRangeType `json:"flocculation,omitempty"`
	AttenuationRange *PercentRangeType     `json:"attenuation_range,omitempty"`
	// The recommended temperature range of fermentation by the culture producer.
	TemperatureRange *TemperatureRangeType `json:"temperature_range,omitempty"`
	// The recommended limit of abv by the culture producer before attenuation stops.
	AlcoholTolerance *PercentType `json:"alcohol_tolerance,omitempty"`
	CultureBase      *CultureBase `json:"CultureBase,omitempty"`
	// A POF+ culture is capable of producing phenols, which is a common distinctive property of saison, and brett yeasts.
	Pof *bool `json:"pof,omitempty"`
}

type CultureInventoryType struct {
	Liquid  *VolumeType `json:"liquid,omitempty"`
	Dry     *MassType   `json:"dry,omitempty"`
	Slant   *VolumeType `json:"slant,omitempty"`
	Culture *VolumeType `json:"culture,omitempty"`
}

// Zymocide, also known as killer yeast properties, is common among wine yeast. There are also some ale and brett yeasts that are immune to some zymocidic properties, these are known as killer neutral.
type Zymocide struct {
	No1     *bool `json:"1,omitempty"`
	No2     *bool `json:"2,omitempty"`
	No28    *bool `json:"28,omitempty"`
	Klus    *bool `json:"klus,omitempty"`
	Neutral *bool `json:"neutral,omitempty"`
}
