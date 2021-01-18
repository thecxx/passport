package errors

var (
	ErrorAccountNotFound      = NewServiceError("account not found")
	ErrorAccountAlreadyExists = NewServiceError("account already exists")
)

type ServiceError struct {
	message string
}

// NewServiceError returns an service error that formats as the given text.
func NewServiceError(text string) error {
	return &ServiceError{text}
}

func (s *ServiceError) Error() string {
	return s.message
}

type RuntimeError struct {
	message string
}

// NewRuntimeError returns an service error that formats as the given text.
func NewRuntimeError(text string) error {
	return &RuntimeError{text}
}

func (s *RuntimeError) Error() string {
	return s.message
}
