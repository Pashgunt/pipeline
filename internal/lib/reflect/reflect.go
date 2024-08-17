package reflectpipeline

import (
	"fmt"
	"github.com/Pashgunt/pipeline/internal/config"
	"reflect"
)

const invalidCallMethodByName = "In structures method %s does not exists"

func CallMethodFromStructure(
	structure interface{},
	method string,
) []reflect.Value {
	valueOfMethod := reflect.ValueOf(structure).MethodByName(method)

	if !valueOfMethod.IsValid() {
		panic(fmt.Sprintf(invalidCallMethodByName, method))
	}

	return valueOfMethod.Call(nil)
}

func CallMethodFromStructureWithParams(
	structure interface{},
	method string,
	data config.SendDataType,
) []reflect.Value {
	valueOfMethod := reflect.ValueOf(structure).MethodByName(method)

	if !valueOfMethod.IsValid() {
		panic(fmt.Sprintf(invalidCallMethodByName, method))
	}

	return valueOfMethod.Call([]reflect.Value{reflect.ValueOf(data)})
}

func CallFunctionWithParams(
	function interface{},
	data config.SendDataType,
) []reflect.Value {
	return reflect.ValueOf(function).Call([]reflect.Value{reflect.ValueOf(data)})
}
