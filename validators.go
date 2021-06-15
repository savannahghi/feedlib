package go_utils

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/ttacon/libphonenumber"
)

// IsMSISDNValid uses regular expression to validate the a phone number
func IsMSISDNValid(msisdn string) bool {
	if len(msisdn) < 10 {
		return false
	}
	reKen := regexp.MustCompile(`^(?:254|\+254|0)?((7|1)(?:(?:[129][0-9])|(?:0[0-8])|(4[0-1]))[0-9]{6})$`)
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if !reKen.MatchString(msisdn) {
		return re.MatchString(msisdn)
	}
	return reKen.MatchString(msisdn)
}

// NormalizeMSISDN validates the input phone number.
// For valid phone numbers, it normalizes them to international format
// e.g +2547........
func NormalizeMSISDN(msisdn string) (*string, error) {
	if !IsMSISDNValid(msisdn) {
		return nil, fmt.Errorf("invalid phone number: %s", msisdn)
	}
	num, err := libphonenumber.Parse(msisdn, defaultRegion)
	if err != nil {
		return nil, err
	}
	formatted := libphonenumber.Format(num, libphonenumber.INTERNATIONAL)
	cleaned := strings.ReplaceAll(formatted, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	return &cleaned, nil
}

// MustNormalizeMSISDN validates the input phone number otherwise it panics
func MustNormalizeMSISDN(msisdn string) (*string, error) {
	if !IsMSISDNValid(msisdn) {
		return nil, fmt.Errorf("invalid phone number: %s", msisdn)
	}
	num, err := libphonenumber.Parse(msisdn, defaultRegion)
	if err != nil {
		log.Panic(fmt.Sprintf("invalid phone number: %s", msisdn))
	}
	formatted := libphonenumber.Format(num, libphonenumber.INTERNATIONAL)
	cleaned := strings.ReplaceAll(formatted, " ", "")
	return &cleaned, nil
}

// StringSliceContains tests if a string is contained in a slice of strings
func StringSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// IntSliceContains tests if a string is contained in a slice of strings
func IntSliceContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
