package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	ctx := context.Background()

	llm, err := openai.New()
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter the text: ")

	scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
	scanner.Scan()
	text := scanner.Text()

	if text == "" {
		_ = fmt.Errorf("Please enter the text")
		return
	}

	completion, err := llm.Call(ctx, text, llms.WithMaxLength(100))
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
	}

	if completion == "" {
		fmt.Print("No completion")
		return
	} else {
		fmt.Print(completion)
	}
}
