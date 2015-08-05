package passbook

// Standard Field Dictionary Keys: Information about a field.
// These keys are used for all dictionaries that define a field.
type Field struct {
	Key               string         `json:"key"`                        // The key must be unique within the scope of the entire pass.
	Label             string         `json:"label,omitempty"`            // Label text for the field.
	Value             interface{}    `json:"value"`                      // Value of the field.
	AttributedValue   interface{}    `json:"attributedValue,omitempty"`  // Attributed value of the field.
	TextAlignment     TextAlignment  `json:"textAlignment,omitempty"`    // Alignment for the field’s contents.
	ChangeMessage     string         `json:"changeMessage,omitempty"`    // Format string for the alert text that is displayed when the pass is updated. The format string must contain the escape %@, which is replaced with the field’s new value. For example, “Gate changed to %@.”
	DataDetectorTypes []DataDetector `json:"dataDetectorTypes,omitempty` // Data dectors that are applied to the field’s value.
	// Date Style Keys: Information about how a date should be displayed in a field.
	// If any of these keys is present, the value of the field is treated as a date. Either specify both a date style and a time style, or neither.
	DateStyle       DateTimeStyle `json:"dateStyle,omitempty"`       // Style of date to display.
	TimeStyle       DateTimeStyle `json:"timeStyle,omitempty"`       // Style of time to display.
	IgnoresTimeZone bool          `json:"ignoresTimeZone,omitempty"` // Always display the time and date in the given time zone, not in the user’s current time zone.
	IsRelative      bool          `json:"isRelative,omitempty"`      // If true, the label’s value is displayed as a relative date; otherwise, it is displayed as an absolute date.
	// Number Style Keys: Information about how a number should be displayed in a field.
	// These keys are optional if the field’s value is a number; otherwise they are not allowed. Only one of these keys is allowed per field.
	CurrencyCode string      `json:"currencyCode,omitempty"` // ISO 4217 currency code for the field’s value.
	NumberStyle  NumberStyle `json:"numberStyle,omitempty"`  // Style of number to display.
}
