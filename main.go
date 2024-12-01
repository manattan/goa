package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	ctx := context.Background()

	cmd := exec.Command("bash", "-c", `git diff --cached | grep -v "^---" | grep -e '^+' -e '^-' | sed 's/^+++ b\//+++ \.\//'`)

	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cannot run command: %v", err)
		return
	}

	diff := buf.String()

	llm, err := openai.New()
	if err != nil {
		panic(err)
	}

	prompt := `
	以下に示す git diff から、commit メッセージを生成してください。必ず、以下の条件に従ってください。
	- 変更差分を以下の3種類のどれに当てはまるかを理解し、あてはまったprefixをつけてください。
		- feat: 新機能の追加、機能の修正
		- fix: バグの改善
		- docs: ドキュメントのみの変更
	- 最大100文字程度の英文で1つのcommitメッセージを記述してください。短くなりそうな場合は、適切に省略してください。
	- 複数の英文になっては絶対にいけません。
	- もし diff が存在しない場合は、'empty commit' と記述してください。
	- commit の description は不要です。
	- 中身を理解し、適切なcommit messageを生成してください。
		- ダメな例: 'Add .env to .gitignore, create diff.sh script' のような、ファイルの差分をそのまま記述したもの
		- 良い例: 'fix: add .env to .gitignore' のような、変更内容を明確に示したもの
	`

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, prompt),
		llms.TextParts(llms.ChatMessageTypeHuman, diff),
	}

	response, err := llm.GenerateContent(ctx, content)

	if err != nil {
		log.Fatalf("cannot generate content: %v", err)
		return
	}

	fmt.Println(response.Choices[0].Content)

	cmd = exec.Command("git", "commit", "-m", response.Choices[0].Content)

	err = cmd.Run()
	if err != nil {
		log.Fatalf("cannot run command: %v", err)
		return
	}
}
