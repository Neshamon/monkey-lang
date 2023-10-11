package lexer

type Lexer struct {
	input string
	position int  // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch byte  // current char under examination
}

/* position & readPosition access characters in input by using the characters
 as an index.

   Why there are two positions is so both the input of the current
   character and the subsequent one can be inspected. position describes the
   current character and readPosition describes the next. ch represents the byte
   of the current character in the input
*/

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
