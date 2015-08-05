package passbook

import (
	"encoding/json"
	"errors"
)

// Barcode Dictionary: Information about a pass’s barcode.
type Barcode struct {
	Format          BarcodeFormat `json:"format"`            // Barcode format.
	Message         string        `json:"message"`           // Message or payload to be displayed as a barcode.
	MessageEncoding string        `json:"messageEncoding"`   // Textencodingthatisusedtoconvertthemessage from the string representation to a data representation to render the barcode.
	AltText         string        `json:"altText,omitempty"` // Text displayed near the barcode. For example, a human-readable version of the barcode data in case the barcode doesn’t scan.
}

func (b Barcode) MarshalJSON() ([]byte, error) {
	if b.Format != PKBarcodeFormatQR ||
		b.Format != PKBarcodeFormatPDF417 ||
		b.Format != PKBarcodeFormatAztec {
		return nil, errors.New("Barcode format must be one of the following values: " +
			"PKBarcodeFormatQR, PKBarcodeFormatPDF417, PKBarcodeFormatAztec")
	}
	if b.Message == "" {
		return nil, errors.New("Message of barcode must be set")
	}
	if b.MessageEncoding == "" {
		b.MessageEncoding = "iso-8859-1"
	}
	return json.Marshal(b)
}
