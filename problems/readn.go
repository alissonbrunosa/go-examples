package problems

// Using a read7() method that returns 7 characters from a file, implement readN(n) which reads n characters.
//
// For example, given a file with the content “Hello world”, three read7() returns “Hello w”, “orld” and then “”.

type StringBuffer struct {
	position int
	buffer   []byte
}

func (b *StringBuffer) empty() bool { return len(b.buffer) <= b.position }

func (b *StringBuffer) ReadN(n int) string {
	if b.empty() {
		return ""
	}

	if n > len(b.buffer) || (b.position+n) > len(b.buffer) {
		n = len(b.buffer)
	}

	bytes := b.buffer[b.position:n]
	b.position += len(bytes)

	return string(bytes)
}

func NewStringBuffer(text string) *StringBuffer {
	return &StringBuffer{
		buffer: []byte(text),
	}
}
