package main

import (
	"context"
	"log"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
)

func main() {
	textreturn := detectText("fund.png")
	textreturn = strings.Replace(textreturn, "R$", "", -1)
	textreturn = strings.Split(textreturn, "\n")[0]
	log.Printf("Return: %v", textreturn)
}

func detectText(file string) string {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Printf("Error while executing NewImageAnnotatorClient")
	}

	f, err := os.Open(file)
	if err != nil {
		log.Printf("Error while executing Open")
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		log.Printf("Error while executing NewImageFromReader")
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Printf("Error while executing DetectTexts")
	}
	return string(annotations[0].Description)
}
