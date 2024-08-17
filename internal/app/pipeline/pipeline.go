package pipeline

import (
	"context"
	"github.com/Pashgunt/pipeline/internal/app/fields"
	"github.com/Pashgunt/pipeline/internal/config"
	"github.com/Pashgunt/pipeline/internal/services/handlecarry"
	"github.com/Pashgunt/pipeline/internal/services/handleexception"
	"github.com/Pashgunt/pipeline/internal/services/send"
	"github.com/Pashgunt/pipeline/internal/services/then"
	"github.com/Pashgunt/pipeline/internal/services/thenreturn"
	"github.com/Pashgunt/pipeline/internal/services/through"
	"github.com/Pashgunt/pipeline/internal/services/via"
)

type Interface interface {
	Send(data config.SendDataType) *Pipeline
	Through(structures []interface{}) *Pipeline
	Via(method string) *Pipeline
	HandleException(callback func(ctx context.Context)) *Pipeline
	HandleCarry(callback func(ctx context.Context)) *Pipeline
	Pipes() []interface{}
	Then()
	ThenReturn() config.SendDataType
}

type Pipeline struct {
	fields *fields.PipelineFields
}

func NewPipeline() *Pipeline {
	return &Pipeline{fields: &fields.PipelineFields{}}
}

// Send Set the object being sent through the pipeline.
func (p *Pipeline) Send(data config.SendDataType) *Pipeline {
	send.Send(data, p.fields)

	return p
}

// Through Set the array of pipes.
func (p *Pipeline) Through(structures []interface{}) *Pipeline {
	through.Through(structures, p.fields)

	return p
}

// Via Set the method to call on the pipes.
func (p *Pipeline) Via(method string) *Pipeline {
	via.Via(method, p.fields)

	return p
}

// HandleException Handle the given exception.
func (p *Pipeline) HandleException(callback func(ctx context.Context)) *Pipeline {
	handleexception.HandleException(callback, p.fields)

	return p
}

// HandleCarry Handle the value returned from each pipe before passing it to the next.
func (p *Pipeline) HandleCarry(callback func(ctx context.Context)) *Pipeline {
	handlecarry.HandeCarry(callback, p.fields)

	return p
}

// Pipes Get the array of configured pipes.
func (p *Pipeline) Pipes() []interface{} {
	return p.fields.Structures()
}

// ThenReturn Run the pipeline and return the result.
func (p *Pipeline) ThenReturn() config.SendDataType {
	return thenreturn.ThenReturn(p.fields)
}

// Then Run the pipeline with a final destination callback.
func (p *Pipeline) Then() {
	then.Then(p.fields)
}
