package errors

// ServiceError : business error struct
// swagger:model
type ServiceError struct {
	// The error message
	Message string `json:"message"`
}
