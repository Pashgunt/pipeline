package handlecarry

import (
	"context"
	"pipelinePackage/internal/app/fields"
)

func HandeCarry(callback func(ctx context.Context), fields *fields.PipelineFields) {
	fields.SetHandleCarry(callback)
}
