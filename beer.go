// Code generated by jsonschema. DO NOT EDIT.
package beerjson

// ID: https://raw.githubusercontent.com/beerjson/beerjson/master/json/beer.json

// Root element of all beerjson documents.
type Beerjson struct {
	// A collection of steps providing process information for common mashing procedures.
	Mashes []MashProcedureType `json:"mashes,omitempty"`
	// Records containing a minimal collection of the description of ingredients, procedures and other required parameters necessary to recreate a batch of beer.
	Recipes []RecipeType `json:"recipes,omitempty"`
	// Records for adjuncts which do not contribute to the gravity of the beer.
	MiscellaneousIngredients []MiscellaneousType `json:"miscellaneous_ingredients,omitempty"`
	// Records detailing the characteristics of the beer styles for which judging guidelines have been established.
	Styles []StyleType `json:"styles,omitempty"`
	// A collection of steps providing process information for common fermentation procedures.
	Fermentations []FermentationProcedureType `json:"fermentations,omitempty"`
	// A collection of steps providing process information for common boil procedures.
	Boil []BoilProcedureType `json:"boil,omitempty"`
	// Explicitly encode beerjson version within list of records.
	Version VersionType `json:"version", validate:"required"`
	// Records for any ingredient that contributes to the gravity of the beer.
	Fermentables []FermentableType `json:"fermentables,omitempty"`
	// The timing object fully describes the timing of an addition with options for basis on time, gravity, or pH at any process step.
	TimingObject *TimingType `json:"timing_object,omitempty"`
	// Records detailing the wide array of unique cultures.
	Cultures []CultureInformation `json:"cultures,omitempty"`
	// Provides necessary information for brewing equipment.
	Equipments []EquipmentType `json:"equipments,omitempty"`
	// A collection of steps providing process information for common packaging procedures.
	Packaging []PackagingProcedureType `json:"packaging,omitempty"`
	// Records detailing the many properties of unique hop varieties.
	HopVarieties []VarietyInformation `json:"hop_varieties,omitempty"`
	// Records for water profiles used in brewing.
	Profiles []WaterBase `json:"profiles,omitempty"`
}
