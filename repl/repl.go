package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-i/evaluator"
	"monkey-i/lexer"
	"monkey-i/object"
	"monkey-i/parser"
	"monkey-i/token"
)

const PROMPT = ">> "

func StartLexer(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)

		if !scanner.Scan() {
			return
		}

		l := lexer.New(scanner.Text())

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

func StartParser(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		if !scanner.Scan() {
			return
		}

		l := lexer.New(scanner.Text())
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

	}
}

const MONKEY_FACE = `🙈🙈🙈`
const MONKEY_TAIL = `🐒🐒🐒`

func printParserErrors(out io.Writer, errors []error) {
	io.WriteString(out, MONKEY_FACE+"\n")
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, fmt.Sprintf("\t%s\n", msg.Error()))
	}
	io.WriteString(out, MONKEY_TAIL+"\n")
}

func StartEvaluator(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		if !scanner.Scan() {
			return
		}

		l := lexer.New(scanner.Text())
		p := parser.New(l)
		program := p.ParseProgram()
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}

}
