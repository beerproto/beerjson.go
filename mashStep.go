// Code generated by jsonschema. DO NOT EDIT.
package beerjson

// ID: https://raw.githubusercontent.com/beerjson/beerjson/master/json/mash_step.json

// MashStepType - a per step representation occurring during the mash.
type MashStepType struct {
	StepTime TimeType `json:"step_time", validate:"required"`
	// The amount of time  that passes before this step begins. eg moving from a mash step (step 1) of 148F, to a new temperature step of 156F (step 2) may take 8 minutes to heat the mash. Step 2 would have a ramp time of 8 minutes.
	RampTime         *TimeType        `json:"ramp_time,omitempty"`
	EndTemperature   *TemperatureType `json:"end_temperature,omitempty"`
	StartPH          *AcidityType     `json:"start_pH,omitempty"`
	Name             string           `json:"name", validate:"required"`
	MashStepTypeType MashStepTypeType `json:"type", validate:"required"`
	Description      *string          `json:"description,omitempty"`
	// Also known as the mash thickness. eg 1.75 qt/lb or 3.65 L/kg.
	WaterGrainRatio *SpecificVolumeType `json:"water_grain_ratio,omitempty"`
	// Temperature of the water for an infusion step.
	InfuseTemperature *TemperatureType `json:"infuse_temperature,omitempty"`
	EndPH             *AcidityType     `json:"end_pH,omitempty"`
	Amount            *VolumeType      `json:"amount,omitempty"`
	StepTemperature   TemperatureType  `json:"step_temperature", validate:"required"`
}

type MashStepTypeType string

const (
	MashStepTypeType_Infusion     MashStepTypeType = "infusion"
	MashStepTypeType_Temperature  MashStepTypeType = "temperature"
	MashStepTypeType_Decoction    MashStepTypeType = "decoction"
	MashStepTypeType_SouringMash  MashStepTypeType = "souring mash"
	MashStepTypeType_SouringWort  MashStepTypeType = "souring wort"
	MashStepTypeType_DrainMashTun MashStepTypeType = "drain mash tun"
	MashStepTypeType_Sparge       MashStepTypeType = "sparge"
)
