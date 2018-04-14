package request

type AuthenticateRequest struct {
	PhoneNumber   string
	DeviceID      string
	FirebaseToken string
}
