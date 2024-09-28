package errorhandler

const (
	NotFoundErrorType = iota
	DomainErrorType
	ValidationErrorType
)

type Commons interface {
	GetType() uint
	SetTracerLauncherName(string)
	GetMessage() string
}
