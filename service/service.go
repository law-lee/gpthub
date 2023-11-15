package service

import (
	"github.com/law-lee/gpthub/api"
	"github.com/law-lee/gpthub/pkg"
	"github.com/law-lee/gpthub/types"
)

func NewModel(modelType types.GPTModel) (api.GPT, error) {
	cfg := pkg.LoadConfig()
	switch modelType {
	case types.OpenAIModel:
		return &api.OpenAI{Model: cfg.Openapi}, nil
	case types.SparkModel:
		return &api.Spark{Model: cfg.Spark}, nil
	default:
		return nil, api.ErrNotSupportModel
	}

}
