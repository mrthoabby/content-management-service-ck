package errorhandler

const (
	NotFoundErrorType = iota
	DomainErrorType
	ValidationErrorType
	ApiHandledError
)

type Commons interface {
	error
	GetType() uint
	SetTracerLauncherName(string)
	Error() string
}
