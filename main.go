package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type Jurassic2Request struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"maxTokens,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
}

type Jurassic2Response struct {
	Completions []Completion `json:"completions"`
}
type Completion struct {
	Data Data `json:"data"`
}
type Data struct {
	Text string `json:"text"`
}

type InvokeModelWrapper struct {
	BedrockRuntimeClient *bedrockruntime.Client
}

// Invokes AI21 Labs Jurassic-2 on Amazon Bedrock to run an inference using the input
// provided in the request body.
func (wrapper InvokeModelWrapper) InvokeJurassic2(prompt string) (string, error) {
	modelId := viper.GetString("model_id")

	body, err := json.Marshal(Jurassic2Request{
		Prompt:      prompt,
		MaxTokens:   viper.GetInt("max_tokens"),
		Temperature: viper.GetFloat64("temperature"),
	})

	if err != nil {
		log.Fatal("failed to marshal", err)
	}

	output, err := wrapper.BedrockRuntimeClient.InvokeModel(context.TODO(), &bedrockruntime.InvokeModelInput{
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

func main() {
	region := "us-east-1"
	sdkConfig, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	client := bedrockruntime.NewFromConfig(sdkConfig)
	wrapper := InvokeModelWrapper{client}
	prompt := viper.GetString("system_description")
	dbsql, _ := os.ReadFile("db.sql")
	prompt += string(dbsql) + "\n"
	queryExplainResult, _ := os.ReadFile("explain_result.csv")
	prompt += string(queryExplainResult)
	fmt.Println(wrapper.InvokeJurassic2(prompt))
}
