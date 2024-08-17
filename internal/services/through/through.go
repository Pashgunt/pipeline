package through

import (
	"github.com/Pashgunt/pipeline/internal/app/fields"
)

func Through(data []interface{}, fields *fields.PipelineFields) {
	fields.SetStructures(data)
}
