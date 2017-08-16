package passbook

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
type FieldsData []Field
