package gotestguide

import (
	"encoding/json"
	"fmt"
)

// Helper method to unmarshal raw test steps into the appropriate types.
func unmarshalRawTestStep(raw []json.RawMessage) ([]IAbstractTestStep, error) {
	ret := make([]IAbstractTestStep, len(raw))
	for i, step := range raw {
		// Get the type of the step
		var stepType struct {
			DType TestStepType `json:"dType"`
		}
		if err := json.Unmarshal(step, &stepType); err != nil {
			return nil, err
		}

		// Handle the different types
		switch stepType.DType {
		case TEST_STEP_TYPE_TEST_STEP:
			var ts TestStep
			if err := json.Unmarshal(step, &ts); err != nil {
				return nil, err
			}
			ret[i] = &ts
		case TEST_STEP_TYPE_TEST_STEP_FOLDER:
			var folder TestStepFolder
			if err := json.Unmarshal(step, &folder); err != nil {
				return nil, err
			}
			ret[i] = &folder
		default:
			return nil, fmt.Errorf("unknown type: %s", stepType.DType)
		}
	}
	return ret, nil
}
