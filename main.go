package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	//変数でflagを定義します
	var (
		f = flag.String("f", "", "input git diff file path")
	)
	//ここで解析されます
	flag.Parse()

	if f == nil {
		panic("f is nil")
	}

	if *f == "" {
		log.Fatalf("please input the file path, file is empty string")
		return
	}

	ctx := context.Background()

	llm, err := openai.New()
	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(*f)
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
		return
	}

	prompt := `
以下に示す git diff から、commit メッセージを生成してください。必ず、以下の条件に従ってください。
- 変更差分を以下の3種類のどれに当てはまるかを理解し、あてはまったprefixをつけてください。
    - feat: 新機能の追加、機能の修正
	- fix: バグの改善
	- docs: ドキュメントのみの変更
- 100文字程度の英文で1つのcommitメッセージを記述してください。短くなりそうな場合は、適切に省略してください。1つのメッセージに複数の具体内容が含まれても構いません。
- 中身を理解し、適切なcommit messageを生成してください。
    - ダメな例: 'Add .env to .gitignore, create diff.sh script' のような、ファイルの差分をそのまま記述したもの
	- 良い例: 'fix: add .env to .gitignore' のような、変更内容を明確に示したもの
	`

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, prompt),
		llms.TextParts(llms.ChatMessageTypeHuman, string(file)),
	}

	response, err := llm.GenerateContent(ctx, content)

	if err != nil {
		log.Fatalf("cannot generate content: %v", err)
		return
	}

	fmt.Println(response.Choices[0].Content)
}
