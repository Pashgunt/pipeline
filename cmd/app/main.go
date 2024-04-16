package main

import (
	"pipelinePackage/internal/app/pipeline"
	"pipelinePackage/internal/config"
)

func main() {
	(pipeline.NewPipeline()).
		Send(
			config.SendDataType{
				"1": 1,
			},
		).
		Through(
			[]interface{}{
				&TestStructure1{},
				func(data config.SendDataType) config.SendDataType {
					data["3"] = "3"
					return data
				},
				&TestStructure2{},
			},
		).
		Via("Test").
		Then()
}

type TestStructure1 struct {
}

func (test1 *TestStructure1) Test(data config.SendDataType) config.SendDataType {
	data["2"] = "2"
	return data
}

type TestStructure2 struct {
}

func (test2 *TestStructure2) Test(data config.SendDataType) config.SendDataType {
	data["4"] = "4"
	return data
}
