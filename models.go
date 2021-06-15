package go_utils

import (
	"fmt"
	"time"
)

// PhoneOptIn is used to persist and manage phone communication whitelists
type PhoneOptIn struct {
	MSISDN  string `json:"msisdn" firestore:"msisdn"`
	OptedIn bool   `json:"optedIn" firestore:"optedIn"`
}

//IsEntity ...
func (p PhoneOptIn) IsEntity() {}

// USSDSessionLog is used to persist a log of USSD sessions
type USSDSessionLog struct {
	MSISDN    string `json:"msisdn" firestore:"msisdn"`
	SessionID string `json:"sessionID" firestore:"sessionID"`
}

//IsEntity ...
func (p USSDSessionLog) IsEntity() {}

// EmailOptIn is used to persist and manage email communication whitelists
type EmailOptIn struct {
	Email   string `json:"email" firestore:"optedIn"`
	OptedIn bool   `json:"optedIn" firestore:"optedIn"`
}

//IsEntity ...
func (e EmailOptIn) IsEntity() {}

// PIN is used to store a PIN (Personal Identifiation Number) associated
// to a phone number sign up to Firebase
type PIN struct {
	UID     string `json:"uid" firestore:"uid"`
	MSISDN  string `json:"msisdn,omitempty" firestore:"msisdn"`
	PIN     string `json:"pin,omitempty" firestore:"pin"`
	IsValid bool   `json:"isValid,omitempty" firestore:"isValid"`
}

//IsEntity ...
func (p PIN) IsEntity() {}

// Upload represents a file uploaded to cloud storage
type Upload struct {
	ID          string    `json:"id" firestore:"id"`
	URL         string    `json:"url" firestore:"url"`
	Size        int       `json:"size" firestore:"size"`
	Hash        string    `json:"hash" firestore:"hash"`
	Creation    time.Time `json:"creation" firestore:"creation"`
	Title       string    `json:"title" firestore:"title"`
	ContentType string    `json:"contentType" firestore:"contentType"`
	Language    string    `json:"language" firestore:"language"`
	Base64data  string    `json:"base64data" firestore:"base64data"`
}

// IsEntity marks upload as an apollo federation entity
func (u Upload) IsEntity() {}

// IsNode marks upload as a relay node
func (u Upload) IsNode() {}

// SetID marks upload as a relay node
func (u Upload) SetID(id string) {
	u.ID = id
}

// GetID marks upload as a relay node
func (u Upload) GetID() ID {
	return IDValue(u.ID)
}

// UploadInput is used to send data for new uploads
type UploadInput struct {
	Title       string `json:"title"`
	ContentType string `json:"contentType"`
	Language    string `json:"language"`
	Base64data  string `json:"base64data"`
	Filename    string `json:"filename"`
}

// CustomError represents a custom error struct
// Reference https://blog.golang.org/error-handling-and-go
type CustomError struct {
	Err     error  `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
