package handleexception

import (
	"context"
	"github.com/Pashgunt/pipeline/internal/app/fields"
)

func HandleException(callback func(ctx context.Context), fields *fields.PipelineFields) {
	fields.SetHandleException(callback)
}
