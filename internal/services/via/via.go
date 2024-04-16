package via

import (
	"fmt"
	"pipelinePackage/internal/app/fields"
	"strings"
)

const invalidViaMethodError = "The construction cannot call a private method %s for correct operation it must begin with a capital letter"

func Via(via string, fields *fields.PipelineFields) {
	if via[0] != strings.ToTitle(via)[0] {
		panic(fmt.Sprintf(invalidViaMethodError, via))
	}

	fields.SetVia(via)
}
