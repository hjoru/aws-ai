package main

import (
	"aws-ai/model"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/bedrockruntime"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	region := "us-east-1"

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Fatal(err)
	}
	client := bedrockruntime.New(sess)
	prompt := viper.GetString("system_description")
	dbsql, _ := os.ReadFile("db.sql")
	prompt += string(dbsql) + "\n"
	queryExplainResult, _ := os.ReadFile("explain_result.csv")
	prompt += string(queryExplainResult)
	fmt.Println(model.InvokeJurassic2(client, prompt))
}
