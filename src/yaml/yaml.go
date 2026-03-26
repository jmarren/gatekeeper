package yaml

import (
	"fmt"
	"io"
	"slices"
)

type Parser struct {
	content string
	pos     int
	length  int
	tokens  []string
	token   string
}

var emptyChars []rune = []rune{'\n', '\t', ' '}

func NewParser(content string) *Parser {
	return &Parser{
		content: content,
		pos:     0,
		length:  len(content),
	}
}

func (p *Parser) Curr() (rune, error) {
	if p.pos > p.length {
		return ' ', io.EOF
	}
	return rune(p.content[p.pos]), nil
}

func (p *Parser) IsEmpty() bool {
	curr, err := p.Curr()

	if err != nil {
		return true
	}
	return slices.Contains(emptyChars, curr)
}

func (p *Parser) fwd() error {
	p.pos++
	if p.pos >= p.length {
		return io.EOF
	}

	return nil
}

func (p *Parser) Tokens() {

	var err error

	for err != io.EOF {
		err = p.Next()
		if err != nil {
			break
		}
		err = p.CaptureToken()

		if err != nil {
			break
		}
		fmt.Printf("token = %s\n", p.token)
	}

	fmt.Printf("tokens = %v\n", p.tokens)

}

func (p *Parser) CaptureToken() error {
	for !p.IsEmpty() {
		curr, err := p.Curr()
		if err != nil {
			return err
		}
		p.token += string(curr)
		err = p.fwd()
		if err != nil {
			return err
		}
	}

	p.tokens = append(p.tokens, p.token)
	return nil
}

func (p *Parser) Next() error {
	p.token = ""
	for p.IsEmpty() {
		err := p.fwd()
		if err != nil {
			return err
		}
	}

	return nil
}
