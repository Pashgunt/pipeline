package handleexception

import (
	"context"
	"pipelinePackage/internal/app/fields"
)

func HandleException(callback func(ctx context.Context), fields *fields.PipelineFields) {
	fields.SetHandleException(callback)
}
