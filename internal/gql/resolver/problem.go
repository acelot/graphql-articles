package resolver

import "github.com/acelot/articles/internal/gql/model"

func NewInternalErrorProblem() model.InternalErrorProblem {
	return model.InternalErrorProblem{Message: "internal server error"}
}

func NewVersionMismatchProblem() model.VersionMismatchProblem {
	return model.VersionMismatchProblem{Message: "version mismatch"}
}
