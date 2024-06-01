package model

import (
	//	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/bedrockruntime"
	"github.com/aws/aws-sdk-go/service/bedrockruntime/bedrockruntimeiface"
)

type Jurassic2Request struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"maxTokens,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
}

type Completion struct {
	Data Data `json:"data"`
}
type Data struct {
	Text string `json:"text"`
}

type Jurassic2Response struct {
	Completions []Completion `json:"completions"`
}

// Invokes AI21 Labs Jurassic-2 on Amazon Bedrock to run an inference using the input
// provided in the request body.
func InvokeJurassic2(client bedrockruntimeiface.BedrockRuntimeAPI, prompt string) (string, error) {
	modelId := viper.GetString("model_id")

	body, err := json.Marshal(Jurassic2Request{
		Prompt:      prompt,
		MaxTokens:   viper.GetInt("max_tokens"),
		Temperature: viper.GetFloat64("temperature"),
	})

	if err != nil {
		log.Fatal("failed to marshal", err)
	}

	output, err := client.InvokeModel(&bedrockruntime.InvokeModelInput{
		ModelId:     aws.String(modelId),
		ContentType: aws.String("application/json"),
		Body:        body,
	})

	if err != nil {
		fmt.Println(err, modelId)
	}

	var response Jurassic2Response
	if err := json.Unmarshal(output.Body, &response); err != nil {
		log.Fatal("failed to unmarshal", err)
	}

	return response.Completions[0].Data.Text, nil
}
