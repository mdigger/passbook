package passbook

// Barcode format.
type BarcodeFormat string

// Supported Barcode formats.
const (
	PKBarcodeFormatQR     BarcodeFormat = "PKBarcodeFormatQR"
	PKBarcodeFormatPDF417               = "PKBarcodeFormatPDF417"
	PKBarcodeFormatAztec                = "PKBarcodeFormatAztec"
)

type DataDetector string

const (
	PKDataDetectorTypePhoneNumber   DataDetector = "PKDataDetectorTypePhoneNumber"
	PKDataDetectorTypeLink                       = "PKDataDetectorTypeLink"
	PKDataDetectorTypeAddress                    = "PKDataDetectorTypeAddress"
	PKDataDetectorTypeCalendarEvent              = "PKDataDetectorTypeCalendarEvent"
)

type TextAlignment string

const (
	PKTextAlignmentNatural TextAlignment = "PKTextAlignmentNatural"
	PKTextAlignmentLeft                  = "PKTextAlignmentLeft"
	PKTextAlignmentCenter                = "PKTextAlignmentCenter"
	PKTextAlignmentRight                 = "PKTextAlignmentRight"
)

type DateTimeStyle string

const (
	PKDateStyleNone   DateTimeStyle = "PKDateStyleNone"
	PKDateStyleShort                = "PKDateStyleShort"
	PKDateStyleMedium               = "PKDateStyleMedium"
	PKDateStyleLong                 = "PKDateStyleLong"
	PKDateStyleFull                 = "PKDateStyleFull"
)

// Type of transit
type TransitType string

// Supported types of transit
const (
	PKTransitTypeAir     TransitType = "PKTransitTypeAir"
	PKTransitTypeBoat                = "PKTransitTypeBoat"
	PKTransitTypeBus                 = "PKTransitTypeBus"
	PKTransitTypeGeneric             = "PKTransitTypeGeneric"
	PKTransitTypeTrain               = "PKTransitTypeTrain"
)

type NumberStyle string

const (
	PKNumberStyleDecimal    NumberStyle = "PKNumberStyleDecimal"
	PKNumberStylePercent                = "PKNumberStylePercent"
	PKNumberStyleScientific             = "PKNumberStyleScientific"
	PKNumberStyleSpellOut               = "PKNumberStyleSpellOut"
)
