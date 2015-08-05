package passbook

import (
	"encoding/json"
	"errors"
)

// Fields Pass Structure Dictionary: Keys that define the structure of the pass.
// These keys are used for all pass styles and partition the fields into the various parts of the pass.
type Fields struct {
	TransitType TransitType `json:"transitType,omitempty"`     // Type of transit. Required for boarding passes; otherwise not allowed.
	Primary     FieldsData  `json:"primaryFields,omitempty"`   // Fields to be displayed prominently on the front of the pass.
	Secondary   FieldsData  `json:"secondaryFields,omitempty"` // Fields to be displayed on the front of the pass.
	Auxiliary   FieldsData  `json:"auxiliaryFields,omitempty"` // Additional fields to be displayed on the front of the pass.
	Back        FieldsData  `json:"backFields,omitempty"`      // Fields to be on the back of the pass.
	Header      FieldsData  `json:"headerFields,omitempty"`    // Fields to be displayed in the header on the front of the pass.
}

// FieldsData describe array of field dictionaries
type FieldsData map[string]Field

func (fd FieldsData) MarshalJSON() ([]byte, error) {
	var fields = make([]Field, 0, len(fd))
	for key, value := range fd {
		if key == "" {
			return nil, errors.New("The key of field must be not empty")
		}
		if value.Value == "" {
			return nil, errors.New("The value of field must be not empty")
		}
		value.Key = key
		fields = append(fields, value)
	}
	return json.Marshal(fields)
}

func (fd *FieldsData) UnmarshalJSON(data []byte) error {
	var fields []Field
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var fmap = make(FieldsData)
	for _, field := range fields {
		fmap[field.Key] = field
	}
	*fd = fmap
	return nil
}
