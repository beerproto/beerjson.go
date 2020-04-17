// Code generated by jsonschema. DO NOT EDIT.
package beerjson

// ID: https://raw.githubusercontent.com/beerjson/beerjson/master/json/recipe.json

// The efficiencyType stores each efficiency component.
type EfficiencyType struct {
	// The percentage of sugar that makes it from the grain to the fermenter.
	Brewhouse PercentType `json:"brewhouse", validate:"required"`
	// The percentage of sugar from the grain yield that is extracted and converted during the mash.
	Conversion *PercentType `json:"conversion,omitempty"`
	// The percentage of sugar that makes it from the mash tun to the kettle.
	Lauter *PercentType `json:"lauter,omitempty"`
	// The percentage of sugar that makes it from the grain to the kettle.
	Mash *PercentType `json:"mash,omitempty"`
}

type IngredientsType struct {
	// miscellaneous_additions collects all the miscellaneous items for use in a recipe
	MiscellaneousAdditions []MiscellaneousAdditionType `json:"miscellaneous_additions,omitempty"`
	// culture_additions collects all the culture items for use in a recipe
	CultureAdditions []CultureAdditionType `json:"culture_additions,omitempty"`
	// water_additions collects all the water items for use in a recipe
	WaterAdditions []WaterAdditionType `json:"water_additions,omitempty"`
	// fermentable_additions collects all the fermentable ingredients for use in a recipe
	FermentableAdditions []FermentableAdditionType `json:"fermentable_additions", validate:"required"`
	// hop_additions collects all the hops for use in a recipe
	HopAdditions []HopAdditionType `json:"hop_additions,omitempty"`
}

// RecipeType composes the information stored in a beerjson recipe.
type RecipeType struct {
	AlcoholByVolume *PercentType `json:"alcohol_by_volume,omitempty"`
	// The final carbonation of the beer when packaged or served.
	Carbonation *float64 `json:"carbonation,omitempty"`
	// Used to store each efficiency component, including conversion, and brewhouse.
	Efficiency EfficiencyType `json:"efficiency", validate:"required"`
	// The gravity of wort when transffered to the fermenter.
	OriginalGravity *GravityType     `json:"original_gravity,omitempty"`
	Style           *RecipeStyleType `json:"style,omitempty"`
	// This defines the procedure for performing unique mashing processes.
	Mash *MashProcedureType `json:"mash,omitempty"`
	// Used to differentiate which IBU formula is being used in a recipe. If formula is modified in any way, eg to support whirlpool/flameout additions etc etc, please use `Other` for transparency.
	IbuEstimate *IBUEstimateType `json:"ibu_estimate,omitempty"`
	// The final beer pH at the end of fermentation.
	BeerPH *AcidityType `json:"beer_pH,omitempty"`
	// The total apparent attenuation of the finished beer after fermentation.
	ApparentAttenuation *PercentType `json:"apparent_attenuation,omitempty"`
	// Describes the procedure for packaging your beverage.
	Packaging *PackagingProcedureType `json:"packaging,omitempty"`
	Coauthor  *string                 `json:"coauthor,omitempty"`
	// The volume into the fermenter.
	BatchSize VolumeType `json:"batch_size", validate:"required"`
	Notes     *string    `json:"notes,omitempty"`
	// The gravity of beer at the end of fermentation.
	FinalGravity *GravityType `json:"final_gravity,omitempty"`
	// Defines the procedure for performing a boil. A boil procedure with no steps is the same as a standard single step boil.
	Boil            *BoilProcedureType `json:"boil,omitempty"`
	CaloriesPerPint *float64           `json:"calories_per_pint,omitempty"`
	Name            string             `json:"name", validate:"required"`
	RecipeTypeType  RecipeTypeType     `json:"type", validate:"required"`
	// A collection of all ingredients used for the recipe.
	Ingredients IngredientsType `json:"ingredients", validate:"required"`
	// The color of the finished beer, using SRM or EBC.
	ColorEstimate *ColorType `json:"color_estimate,omitempty"`
	// FermentationProcedureType defines the procedure for performing fermentation.
	Fermentation *FermentationProcedureType `json:"fermentation,omitempty"`
	// Used to store subjective tasting notes, and rating.
	Taste   *TasteType `json:"taste,omitempty"`
	Author  string     `json:"author", validate:"required"`
	Created *DateType  `json:"created,omitempty"`
}

type RecipeTypeType string

const (
	RecipeTypeType_Cider       RecipeTypeType = "cider"
	RecipeTypeType_Kombucha    RecipeTypeType = "kombucha"
	RecipeTypeType_Soda        RecipeTypeType = "soda"
	RecipeTypeType_Other       RecipeTypeType = "other"
	RecipeTypeType_Mead        RecipeTypeType = "mead"
	RecipeTypeType_Wine        RecipeTypeType = "wine"
	RecipeTypeType_Extract     RecipeTypeType = "extract"
	RecipeTypeType_PartialMash RecipeTypeType = "partial mash"
	RecipeTypeType_AllGrain    RecipeTypeType = "all grain"
)

type TasteType struct {
	Notes  string  `json:"notes", validate:"required"`
	Rating float64 `json:"rating", validate:"required"`
}
