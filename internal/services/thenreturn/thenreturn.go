package thenreturn

import (
	"github.com/Pashgunt/pipeline/internal/app/fields"
	"github.com/Pashgunt/pipeline/internal/config"
	"github.com/Pashgunt/pipeline/internal/services/then"
)

func ThenReturn(fields *fields.PipelineFields) config.SendDataType {
	then.Then(fields)

	return fields.SendData()
}
