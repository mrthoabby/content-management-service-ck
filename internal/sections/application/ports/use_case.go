package ports

import "context"

type UseCaseWithParam[T any] interface {
	Execute(context.Context, T)
}

type UseCaseWithReturn[R any] interface {
	Execute(context.Context) R
}

type UseCaseWithParamAndReturn[T any, R any] interface {
	Execute(context.Context, T) R
}

type UseCase interface {
	Execute(context.Context)
}
