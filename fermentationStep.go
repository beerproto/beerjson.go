// Code generated by jsonschema. DO NOT EDIT.
package beerjson

// ID: https://raw.githubusercontent.com/beerjson/beerjson/master/json/fermentation_step.json

// FermentationStepType - a per step representation of a fermentation action.
type FermentationStepType struct {
	StepTime *TimeType `json:"step_time,omitempty"`
	// Free rise is used to indicate a fermentation step where the exothermic fermentation is allowed to raise the temperature without restriction This is either True or false.
	FreeRise         *bool            `json:"free_rise,omitempty"`
	StartGravity     *GravityType     `json:"start_gravity,omitempty"`
	StartPh          *AcidityType     `json:"start_ph,omitempty"`
	EndPh            *AcidityType     `json:"end_ph,omitempty"`
	EndTemperature   *TemperatureType `json:"end_temperature,omitempty"`
	Description      *string          `json:"description,omitempty"`
	StartTemperature *TemperatureType `json:"start_temperature,omitempty"`
	EndGravity       *GravityType     `json:"end_gravity,omitempty"`
	Vessel           *string          `json:"vessel,omitempty"`
	Name             string           `json:"name", validate:"required"`
}
