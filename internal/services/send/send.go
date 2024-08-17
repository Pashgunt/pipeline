package send

import (
	"github.com/Pashgunt/pipeline/internal/app/fields"
	"github.com/Pashgunt/pipeline/internal/config"
)

func Send(data config.SendDataType, fields *fields.PipelineFields) {
	fields.SetSendData(data)
}
