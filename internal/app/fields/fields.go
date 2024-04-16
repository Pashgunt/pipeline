package fields

import (
	"context"
	"pipelinePackage/internal/config"
)

type PipelineFields struct {
	sendData        config.SendDataType
	structures      []interface{}
	via             string
	mode            string
	handleException func(ctx context.Context)
	handleCarry     func(ctx context.Context)
}

func (p *PipelineFields) HandleException() func(ctx context.Context) {
	return p.handleException
}

func (p *PipelineFields) SetHandleException(handleException func(ctx context.Context)) *PipelineFields {
	p.handleException = handleException

	return p
}

func (p *PipelineFields) HandleCarry() func(ctx context.Context) {
	return p.handleCarry
}

func (p *PipelineFields) SetHandleCarry(handleCarry func(ctx context.Context)) *PipelineFields {
	p.handleCarry = handleCarry

	return p
}

func (p *PipelineFields) SendData() config.SendDataType {
	return p.sendData
}

func (p *PipelineFields) SetSendData(sendData config.SendDataType) *PipelineFields {
	p.sendData = sendData

	return p
}

func (p *PipelineFields) Structures() []interface{} {
	return p.structures
}

func (p *PipelineFields) SetStructures(structures []interface{}) *PipelineFields {
	p.structures = structures

	return p
}

func (p *PipelineFields) Via() string {
	if p.via != "" {
		return p.via
	}

	return config.ViaDefaultMethod
}

func (p *PipelineFields) SetVia(via string) *PipelineFields {
	p.via = via

	return p
}

func (p *PipelineFields) Mode() string {
	return p.mode
}

func (p *PipelineFields) SetMode(mode string) *PipelineFields {
	p.mode = mode

	return p
}
