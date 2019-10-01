package twiml

import (
	"encoding/xml"
	"fmt"
)

// Language and speaker options
const (
	SSMLNone    = "none"
	SSMLXWeak   = "x-weak"
	SSMLWeak    = "weak"
	SSMLMedium  = "medium"
	SSMLStrong  = "strong"
	SSMLXStrong = "x-strong"
)

// SSMLText
type SSMLText struct {
	XMLName xml.Name `xml:""`
}

// Validate returns an error if the TwiML is constructed improperly
func (c *SSMLText) Validate() error {
	return nil
}

// Type returns the XML name of the verb
func (c *SSMLText) Type() string {
	return "SSMLText"
}

// SSMLBreak
type SSMLBreak struct {
	XMLName  xml.Name `xml:"break"`
	Strength string   `xml:"strength,attr,omitempty"`
	Time     string   `xml:"time,attr,omitempty"`
}

// Validate returns an error if the TwiML is constructed improperly
func (c *SSMLBreak) Validate() error {
	ok := Validate(
		OneOfOpt(c.Strength, SSMLNone, SSMLXWeak, SSMLWeak, SSMLMedium, SSMLStrong, SSMLXStrong),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", s.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (c *SSMLBreak) Type() string {
	return "SSMLBreak"
}
