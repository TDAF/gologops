package gologops

var stringsForTesting = []string{
	"",
	"a",
	"aaa bbb ccc",
	"España y olé",
	`She said: "I know what it's like to be dead`,
	"{}}"}

var contextForTesting = C{"a": "A", "b": "BB", "España": "olé",
	"She said": `I know what it's like to be dead"`, "{": "}"}

type (
	contextLogFunc    func(l *Logger, context C, message string, params ...interface{})
	formatLogFunc     func(l *Logger, message string, params ...interface{})
	simpleLogFunction func(l *Logger, message string)
)
