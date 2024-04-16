package thenreturn

import (
	"pipelinePackage/internal/app/fields"
	"pipelinePackage/internal/config"
	"pipelinePackage/internal/services/then"
)

func ThenReturn(fields *fields.PipelineFields) config.SendDataType {
	then.Then(fields)

	return fields.SendData()
}
