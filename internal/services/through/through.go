package through

import (
	"pipelinePackage/internal/app/fields"
)

func Through(data []interface{}, fields *fields.PipelineFields) {
	fields.SetStructures(data)
}
