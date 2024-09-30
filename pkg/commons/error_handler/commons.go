package errorhandler

const (
	NotFoundErrorType = iota
	DomainErrorType
	ValidationErrorType
	ApiHandledError
	ConflictErrorType
	InvalidFormatErrorType
)

type Commons interface {
	error
	GetType() uint
	SetTracerLauncherName(string)
	Error() string
}
