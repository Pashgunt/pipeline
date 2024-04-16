package send

import (
	"pipelinePackage/internal/app/fields"
	"pipelinePackage/internal/config"
)

func Send(data config.SendDataType, fields *fields.PipelineFields) {
	fields.SetSendData(data)
}
