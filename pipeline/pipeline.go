package pipeline

import (
	"encoding/json"
	"github.com/ProtoML/ProtoML/types"
)

func NewPipeline() (pipeline *types.Pipeline) {
	pipeline = &types.Pipeline{}
	pipeline.Nodes = make([]types.PipelineNode,0,20)
	return
}

func LoadPipeline(PipelineFileBlob []byte) (pipeline *types.Pipeline, err error) {
	pipeline = &types.Pipeline{}
	err = json.Unmarshal(PipelineFileBlob, pipeline)
	return
}

func StorePipeline(pipeline *types.Pipeline) (pipelineBlob []byte, err error) {
	pipelineBlob, err = json.Marshal(pipeline)
	return
}
