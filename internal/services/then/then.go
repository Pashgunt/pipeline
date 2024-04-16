package then

import (
	"context"
	"pipelinePackage/internal/app/fields"
	"pipelinePackage/internal/config"
	reflectpipeline "pipelinePackage/internal/lib/reflect"
	"reflect"
)

func Then(fields *fields.PipelineFields) {
	if fields.HandleException() != nil {
		defer func() {
			if err := recover(); err != nil {
				ctxError := context.Background()
				ctxError = context.WithValue(ctxError, "err", err)
				fields.HandleException()(ctxError)
			}
		}()
	}

	for _, structure := range fields.Structures() {
		makeUpdateDataFields(fields, resolveKindCallType(structure, fields))
	}
}

func prepareHandleCarryStart(fields *fields.PipelineFields) {
	if fields.HandleCarry() == nil {
		return
	}

	ctxCarry := context.Background()
	ctxCarry = context.WithValue(ctxCarry, "data", fields.SendData())
	fields.HandleCarry()(ctxCarry)
}

func makeUpdateDataFields(fields *fields.PipelineFields, dataInterface []reflect.Value) {
	if data, ok := dataInterface[0].Interface().(config.SendDataType); ok {
		fields.SetSendData(data)

		prepareHandleCarryStart(fields)
	}
}

func resolveKindCallType(structure interface{}, fields *fields.PipelineFields) []reflect.Value {
	if reflect.TypeOf(structure).Kind() != reflect.Func {
		return reflectpipeline.CallMethodFromStructureWithParams(
			structure,
			fields.Via(),
			fields.SendData(),
		)
	}

	return reflectpipeline.CallFunctionWithParams(structure, fields.SendData())
}
