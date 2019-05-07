package lexer

type Lexer struct {
  input         string // source code
  position      int    // current already read character position, holding 'ch' byte
  ch            byte   // last read byte
  readPosition  int    // current position, unread 
}

func New (input string) *Lexer {
  // todo: 
  // 1. revisit long-hand assignment
  // 2. literal struct creation vs other options
  // 3. returning l vs returning & 
  l := &Lexer{input: input}
  return &l
}
