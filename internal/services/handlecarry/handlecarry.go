package handlecarry

import (
	"context"
	"github.com/Pashgunt/pipeline/internal/app/fields"
)

func HandeCarry(callback func(ctx context.Context), fields *fields.PipelineFields) {
	fields.SetHandleCarry(callback)
}
