package lexer

import "monkey/token"

type Lexer struct {
  input         string // source code
  readPosition  int    // next position for peaking
  position      int    // current character, ch
  ch            byte   // last read byte
}

func New (input string) *Lexer {
  // todo: 
  // 1. revisit long-hand assignment for structs
  // 2. literal struct creation vs other options
  // 3. returning l vs returning & 
  l := &Lexer{input: input}
  l.ReadChar()

  return l
}

func (l *Lexer) ReadChar() {
  // l.ch is an ascii byte, but could be a rune if we support unicode
  // would have to adjust code to handle multi-byte characters
  if l.readPosition >= len(l.input) {
    l.ch = 0
  } else {
    l.ch = l.input[l.readPosition]
  }
  l.position = l.readPosition
  l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {

  var tok token.Token

  switch l.ch {
    case '=':
      tok = newToken(token.ASSIGN, l.ch)
    case ';':
      tok = newToken(token.SEMICOLON, l.ch)
    case '(':
      tok = newToken(token.LPAREN, l.ch)
    case ')':
      tok = newToken(token.RPAREN, l.ch)
    case ',':
      tok = newToken(token.COMMA, l.ch)
    case '+':
      tok = newToken(token.PLUS, l.ch)
    case '{':
      tok = newToken(token.LBRACE, l.ch)
    case '}':
      tok = newToken(token.RBRACE, l.ch)
    case 0:
      tok.Literal = ""
      tok.Type = token.EOF
  }

  l.ReadChar()
  return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
  return token.Token{Type: tokenType, Literal: string(ch)}
}

