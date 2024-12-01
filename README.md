# goa

A simple tool that generates commit messages from `git diff` using ChatGPT and automatically commits the changes for you. Perfect for developers looking to streamline their Git workflow while maintaining meaningful commit messages.

---

## **Features**
- Automatically generates commit messages based on `git diff`.
- Utilizes ChatGPT's API for natural language generation.
- Automatically commits changes with the generated message.

---

## **Requirements**
- A valid OpenAI API Key.
- Git installed and configured.

---

## **Installation**

1. Clone the repository:
   ```bash
   git clone github.com/manattan/goa
   cd goa
   ```

2. Install required dependencies:
   ```bash
   # Example for Go (if applicable):
   go mod tidy
   ```

3. Set your OpenAI API Key as an environment variable:
   ```bash
   export OPENAI_API_KEY="your-api-key-here"
   ```

---

## **Usage**

1. Stage your changes:
   ```bash
   git add <files>
   ```

2. Run the tool:
   ```bash
   go run main.go
   ```

3. The tool will:
   - Analyze the staged `git diff`.
   - Send the diff to ChatGPT for generating a meaningful commit message.
   - Automatically commit the changes with the generated message.

---

## **Environment Variables**

| Variable         | Description               |
|-------------------|---------------------------|
| `OPENAI_API_KEY` | Your OpenAI API Key (required). |
