package go_utils

import (
	"fmt"
	"io"
	"strconv"
)

// Gender is a code system for administrative gender.
//
// See: https://www.hl7.org/fhir/valueset-administrative-gender.html
type Gender string

// gender constants
const (
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
	GenderOther   Gender = "other"
	GenderUnknown Gender = "unknown"
)

// AllGender is a list of known genders
var AllGender = []Gender{
	GenderMale,
	GenderFemale,
	GenderOther,
	GenderUnknown,
}

// IsValid returns True if the enum value is valid
func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale, GenderOther, GenderUnknown:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

// UnmarshalGQL translates from the supplied value to a valid enum value
func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

// MarshalGQL writes the enum value to the supplied writer
func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// ContentType defines accepted content types
type ContentType string

// Constants used to map to allowed MIME types
const (
	ContentTypePng ContentType = "PNG"
	ContentTypeJpg ContentType = "JPG"
	ContentTypePdf ContentType = "PDF"
)

// AllContentType is a list of all acceptable content types
var AllContentType = []ContentType{
	ContentTypePng,
	ContentTypeJpg,
	ContentTypePdf,
}

// IsValid ensures that the content type value is valid
func (e ContentType) IsValid() bool {
	switch e {
	case ContentTypePng, ContentTypeJpg, ContentTypePdf:
		return true
	}
	return false
}

func (e ContentType) String() string {
	return string(e)
}

// UnmarshalGQL turns the supplied value into a content type value
func (e *ContentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContentType", str)
	}
	return nil
}

// MarshalGQL writes the value of this enum to the supplied writer
func (e ContentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))

}

// Language defines allowed languages for uploads
type Language string

// Constants used to map to allowed languages
const (
	LanguageEn Language = "en"
	LanguageSw Language = "sw"
)

// LanguageCodingSystem is the FHIR language coding system
const LanguageCodingSystem = "urn:ietf:bcp:47"

// LanguageCodingVersion is the FHIR language value
const LanguageCodingVersion = ""

// LanguageNames is a map of language codes to language names
var LanguageNames = map[Language]string{
	LanguageEn: "English",
	LanguageSw: "Swahili",
}

// AllLanguage is a list of all allowed languages
var AllLanguage = []Language{
	LanguageEn,
	LanguageSw,
}

// IsValid ensures that the supplied language value is correct
func (e Language) IsValid() bool {
	switch e {
	case LanguageEn, LanguageSw:
		return true
	}
	return false
}

func (e Language) String() string {
	return string(e)
}

// UnmarshalGQL translates the input to a language type value
func (e *Language) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Language(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Language", str)
	}
	return nil
}

// MarshalGQL writes the value of this enum to the supplied writer
func (e Language) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))

}

// PractitionerSpecialty is a list of recognised health worker specialties.
//
// See: https://medicalboard.co.ke/resources_page/gazetted-specialties/
type PractitionerSpecialty string

// list of known practitioner specialties
const (
	PractitionerSpecialtyUnspecified                     PractitionerSpecialty = "UNSPECIFIED"
	PractitionerSpecialtyAnaesthesia                     PractitionerSpecialty = "ANAESTHESIA"
	PractitionerSpecialtyCardiothoracicSurgery           PractitionerSpecialty = "CARDIOTHORACIC_SURGERY"
	PractitionerSpecialtyClinicalMedicalGenetics         PractitionerSpecialty = "CLINICAL_MEDICAL_GENETICS"
	PractitionerSpecialtyClincicalPathology              PractitionerSpecialty = "CLINCICAL_PATHOLOGY"
	PractitionerSpecialtyGeneralPathology                PractitionerSpecialty = "GENERAL_PATHOLOGY"
	PractitionerSpecialtyAnatomicPathology               PractitionerSpecialty = "ANATOMIC_PATHOLOGY"
	PractitionerSpecialtyClinicalOncology                PractitionerSpecialty = "CLINICAL_ONCOLOGY"
	PractitionerSpecialtyDermatology                     PractitionerSpecialty = "DERMATOLOGY"
	PractitionerSpecialtyEarNoseAndThroat                PractitionerSpecialty = "EAR_NOSE_AND_THROAT"
	PractitionerSpecialtyEmergencyMedicine               PractitionerSpecialty = "EMERGENCY_MEDICINE"
	PractitionerSpecialtyFamilyMedicine                  PractitionerSpecialty = "FAMILY_MEDICINE"
	PractitionerSpecialtyGeneralSurgery                  PractitionerSpecialty = "GENERAL_SURGERY"
	PractitionerSpecialtyGeriatrics                      PractitionerSpecialty = "GERIATRICS"
	PractitionerSpecialtyImmunology                      PractitionerSpecialty = "IMMUNOLOGY"
	PractitionerSpecialtyInfectiousDisease               PractitionerSpecialty = "INFECTIOUS_DISEASE"
	PractitionerSpecialtyInternalMedicine                PractitionerSpecialty = "INTERNAL_MEDICINE"
	PractitionerSpecialtyMicrobiology                    PractitionerSpecialty = "MICROBIOLOGY"
	PractitionerSpecialtyNeurosurgery                    PractitionerSpecialty = "NEUROSURGERY"
	PractitionerSpecialtyObstetricsAndGynaecology        PractitionerSpecialty = "OBSTETRICS_AND_GYNAECOLOGY"
	PractitionerSpecialtyOccupationalMedicine            PractitionerSpecialty = "OCCUPATIONAL_MEDICINE"
	PractitionerSpecialtyOphthalmology                   PractitionerSpecialty = "OPGTHALMOLOGY"
	PractitionerSpecialtyOrthopaedicSurgery              PractitionerSpecialty = "ORTHOPAEDIC_SURGERY"
	PractitionerSpecialtyOncology                        PractitionerSpecialty = "ONCOLOGY"
	PractitionerSpecialtyOncologyRadiotherapy            PractitionerSpecialty = "ONCOLOGY_RADIOTHERAPY"
	PractitionerSpecialtyPaediatricsAndChildHealth       PractitionerSpecialty = "PAEDIATRICS_AND_CHILD_HEALTH"
	PractitionerSpecialtyPalliativeMedicine              PractitionerSpecialty = "PALLIATIVE_MEDICINE"
	PractitionerSpecialtyPlasticAndReconstructiveSurgery PractitionerSpecialty = "PLASTIC_AND_RECONSTRUCTIVE_SURGERY"
	PractitionerSpecialtyPsychiatry                      PractitionerSpecialty = "PSYCHIATRY"
	PractitionerSpecialtyPublicHealth                    PractitionerSpecialty = "PUBLIC_HEALTH"
	PractitionerSpecialtyRadiology                       PractitionerSpecialty = "RADIOLOGY"
	PractitionerSpecialtyUrology                         PractitionerSpecialty = "UROLOGY"
)

// AllPractitionerSpecialty is the set of known practitioner specialties
var AllPractitionerSpecialty = []PractitionerSpecialty{
	PractitionerSpecialtyUnspecified,
	PractitionerSpecialtyAnaesthesia,
	PractitionerSpecialtyCardiothoracicSurgery,
	PractitionerSpecialtyClinicalMedicalGenetics,
	PractitionerSpecialtyClincicalPathology,
	PractitionerSpecialtyGeneralPathology,
	PractitionerSpecialtyAnatomicPathology,
	PractitionerSpecialtyClinicalOncology,
	PractitionerSpecialtyDermatology,
	PractitionerSpecialtyEarNoseAndThroat,
	PractitionerSpecialtyEmergencyMedicine,
	PractitionerSpecialtyFamilyMedicine,
	PractitionerSpecialtyGeneralSurgery,
	PractitionerSpecialtyGeriatrics,
	PractitionerSpecialtyImmunology,
	PractitionerSpecialtyInfectiousDisease,
	PractitionerSpecialtyInternalMedicine,
	PractitionerSpecialtyMicrobiology,
	PractitionerSpecialtyNeurosurgery,
	PractitionerSpecialtyObstetricsAndGynaecology,
	PractitionerSpecialtyOccupationalMedicine,
	PractitionerSpecialtyOphthalmology,
	PractitionerSpecialtyOrthopaedicSurgery,
	PractitionerSpecialtyOncology,
	PractitionerSpecialtyOncologyRadiotherapy,
	PractitionerSpecialtyPaediatricsAndChildHealth,
	PractitionerSpecialtyPalliativeMedicine,
	PractitionerSpecialtyPlasticAndReconstructiveSurgery,
	PractitionerSpecialtyPsychiatry,
	PractitionerSpecialtyPublicHealth,
	PractitionerSpecialtyRadiology,
	PractitionerSpecialtyUrology,
}

// IsValid returns True if the practitioner specialty is valid
func (e PractitionerSpecialty) IsValid() bool {
	switch e {
	case PractitionerSpecialtyUnspecified, PractitionerSpecialtyAnaesthesia, PractitionerSpecialtyCardiothoracicSurgery, PractitionerSpecialtyClinicalMedicalGenetics, PractitionerSpecialtyClincicalPathology, PractitionerSpecialtyGeneralPathology, PractitionerSpecialtyAnatomicPathology, PractitionerSpecialtyClinicalOncology, PractitionerSpecialtyDermatology, PractitionerSpecialtyEarNoseAndThroat, PractitionerSpecialtyEmergencyMedicine, PractitionerSpecialtyFamilyMedicine, PractitionerSpecialtyGeneralSurgery, PractitionerSpecialtyGeriatrics, PractitionerSpecialtyImmunology, PractitionerSpecialtyInfectiousDisease, PractitionerSpecialtyInternalMedicine, PractitionerSpecialtyMicrobiology, PractitionerSpecialtyNeurosurgery, PractitionerSpecialtyObstetricsAndGynaecology, PractitionerSpecialtyOccupationalMedicine, PractitionerSpecialtyOphthalmology, PractitionerSpecialtyOrthopaedicSurgery, PractitionerSpecialtyOncology, PractitionerSpecialtyOncologyRadiotherapy, PractitionerSpecialtyPaediatricsAndChildHealth, PractitionerSpecialtyPalliativeMedicine, PractitionerSpecialtyPlasticAndReconstructiveSurgery, PractitionerSpecialtyPsychiatry, PractitionerSpecialtyPublicHealth, PractitionerSpecialtyRadiology, PractitionerSpecialtyUrology:
		return true
	}
	return false
}

func (e PractitionerSpecialty) String() string {
	return string(e)
}

// UnmarshalGQL converts the supplied value to a practitioner specialty
func (e *PractitionerSpecialty) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PractitionerSpecialty(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PractitionerSpecialty", str)
	}
	return nil
}

// MarshalGQL writes the practitioner specialty to the supplied writer
func (e PractitionerSpecialty) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))

}

// CalendarView is used to determine what view of a calendar to render
type CalendarView string

// calendar view constants
const (
	// CalendarViewDay ...
	CalendarViewDay CalendarView = "DAY"
	// CalendarViewWeek ...
	CalendarViewWeek CalendarView = "WEEK"
)

// AllCalendarView is a list of calendar views
var AllCalendarView = []CalendarView{
	CalendarViewDay,
	CalendarViewWeek,
}

// IsValid returns true if a calendar view is valid
func (e CalendarView) IsValid() bool {
	switch e {
	case CalendarViewDay, CalendarViewWeek:
		return true
	}
	return false
}

// String ...
func (e CalendarView) String() string {
	return string(e)
}

// UnmarshalGQL converts the input value into a calendar view
func (e *CalendarView) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CalendarView(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CalendarView", str)
	}
	return nil
}

// MarshalGQL writes the calendar view value to the supplied writer
func (e CalendarView) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// AddressType represents the types of addresses we have
type AddressType string

// AddressTypeHome is an example of an address type
const (
	AddressTypeHome AddressType = "HOME"
	AddressTypeWork AddressType = "WORK"
)

// AllAddressType contains a slice of all addresses types
var AllAddressType = []AddressType{
	AddressTypeHome,
	AddressTypeWork,
}

// IsValid checks if the address type is valid
func (e AddressType) IsValid() bool {
	switch e {
	case AddressTypeHome, AddressTypeWork:
		return true
	}
	return false
}

func (e AddressType) String() string {
	return string(e)
}

// UnmarshalGQL converts the input, if valid, into an address type value
func (e *AddressType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AddressType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AddressType", str)
	}
	return nil
}

// MarshalGQL converts address type into a valid JSON string
func (e AddressType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// IdentificationDocType defines the various supplier IdentificationDocTypes
type IdentificationDocType string

// IdentificationDocTypeNationalid is an example of a IdentificationDocType
const (
	IdentificationDocTypeNationalid IdentificationDocType = "NATIONALID"
	IdentificationDocTypePassport   IdentificationDocType = "PASSPORT"
	IdentificationDocTypeMilitary   IdentificationDocType = "MILITARY"
)

// AllIdentificationDocType contains a slice of all IdentificationDocTypes
var AllIdentificationDocType = []IdentificationDocType{
	IdentificationDocTypeNationalid,
	IdentificationDocTypePassport,
	IdentificationDocTypeMilitary,
}

// IsValid checks if the IdentificationDocType is valid
func (e IdentificationDocType) IsValid() bool {
	switch e {
	case IdentificationDocTypeNationalid, IdentificationDocTypePassport, IdentificationDocTypeMilitary:
		return true
	}
	return false
}

func (e IdentificationDocType) String() string {
	return string(e)
}

// UnmarshalGQL converts the input, if valid, into an IdentificationDocType value
func (e *IdentificationDocType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IdentificationDocType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IdentificationDocType", str)
	}
	return nil
}

// MarshalGQL converts IdentificationDocType into a valid JSON string
func (e IdentificationDocType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
