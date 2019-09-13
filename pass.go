package passbook

import (
	"encoding/json"
	"errors"
	"strings"
)

//
type Pass struct {
	// Standard Keys: Information that is required for all passes.
	FormatVersion      int    `json:"formatVersion"`      // Version of the file format. The value must be 1.
	PassTypeIdentifier string `json:"passTypeIdentifier"` // Pass type identifier, as issued by Apple. The value must correspond with your signing certificate.
	SerialNumber       string `json:"serialNumber"`       // Serial number that uniquely identifies the pass. No two passes with the same pass type identifier may have the same serial number.
	TeamIdentifier     string `json:"teamIdentifier"`     // Team identifier of the organization that originated and signed the pass, as issued by Apple.
	OrganizationName   string `json:"organizationName"`   // Display name of the organization that originated and signed the pass.
	Description        string `json:"description"`        // Brief description of the pass, used by the iOS accessibility technologies.
	// Associated App Keys: Information about an app that is associated with a pass.
	AppLaunchURL               string `json:"appLaunchURL,omitempty"`               // A URL to be passed to the associated app when launching it.
	AssociatedStoreIdentifiers []int  `json:"associatedStoreIdentifiers,omitempty"` // A list of iTunes Store item identifiers for the associated apps.
	// Companion App Keys: Custom information about a pass provided for a companion app to use.
	UserInfo map[string]interface{} `json:"userInfo,omitempty"` // Custom information for companion apps. This data is not displayed to the user.
	// Expiration Keys: Information about when a pass expires and whether it is still valid.
	// A pass is marked as expired if the current date is after the pass’s expiration date, or if the pass has been explicitly marked as voided.
	ExpirationDate *W3Time `json:"expirationDate,omitempty"` // Date and time when the pass expires.
	Voided         bool    `json:"voided,omitempty"`         // Indicates that the pass is void—for example, a one time use coupon that has been redeemed.
	// Relevance Keys: Information about where and when a pass is relevant.
	Beacons      []Beacon   `json:"beacons,omitempty"`      // Beacons marking locations where the pass is relevant.
	Locations    []Location `json:"locations,omitempty"`    // Locations where the pass is relevant.
	MaxDistance  uint       `json:"maxDistance,omitempty"`  // Maximum distance in meters from a relevant latitude and longitude that the pass is relevant.
	RelevantDate *W3Time    `json:"relevantDate,omitempty"` // Date and time when the pass becomes relevant.
	// Visual Appearance Keys: Visual styling and appearance of the pass.
	Barcode            *Barcode `json:"barcode,omitempty"`            // Information specific to barcodes.
	BackgroundColor    *Color   `json:"backgroundColor,omitempty"`    // Background color of the pass, specified as an CSS-style RGB triple.
	ForegroundColor    *Color   `json:"foregroundColor,omitempty"`    // Foreground color of the pass, specified as a CSS-style RGB triple.
	LabelColor         *Color   `json:"labelColor,omitempty"`         // Color of the label text, specified as a CSS-style RGB triple.
	LogoText           string   `json:"logoText,omitempty"`           // Text displayed next to the logo on the pass.
	GroupingIdentifier string   `json:"groupingIdentifier,omitempty"` // Optional for event tickets and boarding passes; otherwise not allowed. Identifier used to group related passes. If a grouping identifier is specified, passes with the same style, pass type identifier, and grouping identifier are displayed as a group. Otherwise, passes are grouped automatically.
	// Style Keys: Specifies the pass style.
	// Provide exactly one key—the key that corresponds with the pass’s type.
	Generic      *Fields `json:"generic,omitempty"`      // Information specific to a generic pass.
	BoardingPass *Fields `json:"boardingPass,omitempty"` // Information specific to a boarding pass.
	Coupon       *Fields `json:"coupon,omitempty"`       // Information specific to a coupon.
	EventTicket  *Fields `json:"eventTicket,omitempty"`  // Information specific to an event ticket.
	StoreCard    *Fields `json:"storeCard,omitempty"`    // Information specific to a store card.
	// Web Service Keys: Information used to update passes using the web service.
	AuthenticationToken string `json:"authenticationToken,omitempty"` // The authentication token to use with the web service. The token must be 16 characters or longer.
	WebServiceURL       string `json:"webServiceURL,omitempty"`       // The URL of a web service that conforms to the API described in Passbook Web Service Reference.
}

func (p Pass) Marshal() ([]byte, error) {
	if p.Description == "" {
		return nil, errors.New("Empty Description")
	}
	if p.FormatVersion != 1 {
		p.FormatVersion = 1
	}
	if p.OrganizationName == "" {
		return nil, errors.New("Empty Organization Name")
	}
	if p.PassTypeIdentifier == "" {
		return nil, errors.New("Empty Pass Type Identifier")
	}
	if p.SerialNumber == "" {
		return nil, errors.New("Empty Serial Number")
	}
	if p.TeamIdentifier == "" {
		return nil, errors.New("Empty Team Identifier")
	}
	if p.AppLaunchURL != "" && len(p.AssociatedStoreIdentifiers) == 0 {
		return nil, errors.New("Associated Store Identifiers is not defined")
	}
	if len(p.AuthenticationToken) < 16 {
		return nil, errors.New("The Authentication Token must be 16 characters or longer")
	}
	if p.WebServiceURL != "" && !strings.HasPrefix(p.WebServiceURL, "https://") {
		return nil, errors.New("The Web Service URL must use the HTTPS protocol")
	}
	return json.Marshal(p)
}
