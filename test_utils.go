package go_utils

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/imroc/req"
	"github.com/stretchr/testify/assert"
)

const (
	anonymousUserUID  = "AgkGYKUsRifO2O9fTLDuVCMr2hb2" // This is an anonymous user
	verifyPhone       = "testing/verify_phone"
	createUserByPhone = "testing/create_user_by_phone"
	loginByPhone      = "testing/login_by_phone"
	removeUserByPhone = "testing/remove_user"
	addAdmin          = "testing/add_admin_permissions"
	updateBioData     = "testing/update_user_profile"
)

// ContextKey is used as a type for the UID key for the Firebase *auth.Token on context.Context.
// It is a custom type in order to minimize context key collissions on the context
// (.and to shut up golint).
type ContextKey string

// GetDefaultHeaders returns headers used in inter service communication acceptance tests
func GetDefaultHeaders(t *testing.T, rootDomain string, serviceName string) map[string]string {
	return req.Header{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": GetInterserviceBearerTokenHeader(t, rootDomain, serviceName),
	}
}

// GetInterserviceBearerTokenHeader returns a valid isc bearer token header
func GetInterserviceBearerTokenHeader(t *testing.T, rootDomain string, serviceName string) string {
	isc := GetInterserviceClient(t, rootDomain, serviceName)
	authToken, err := isc.CreateAuthToken()
	assert.Nil(t, err)
	assert.NotZero(t, authToken)
	bearerHeader := fmt.Sprintf("Bearer %s", authToken)
	return bearerHeader
}

// GetInterserviceClient returns an isc client used in acceptance testing
func GetInterserviceClient(t *testing.T, rootDomain string, serviceName string) *InterServiceClient {
	service := ISCService{
		Name:       serviceName,
		RootDomain: rootDomain,
	}
	isc, err := NewInterserviceClient(service)
	assert.Nil(t, err)
	assert.NotNil(t, isc)
	return isc
}

// RemoveTestPhoneNumberUser removes the records created by the
// test phonenumber user
func RemoveTestPhoneNumberUser(
	t *testing.T,
	onboardingClient *InterServiceClient,
) error {
	if onboardingClient == nil {
		return fmt.Errorf("nil ISC client")
	}

	payload := map[string]interface{}{
		"phoneNumber": TestUserPhoneNumber,
	}
	resp, err := onboardingClient.MakeRequest(
		http.MethodPost,
		removeUserByPhone,
		payload,
	)
	if err != nil {
		return fmt.Errorf("unable to make a request to remove test user: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil // This is a test utility. Do not block if the user is not found
	}

	return nil
}
