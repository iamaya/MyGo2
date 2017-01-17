package main

func main() {
	mp := MessagePrinter{"foo", 5}
	mp.PrintMessage()

	emp := enhancedMessagePrinter{MessagePrinter{"bar", 10}}
	emp.PrintMessage()
	emp.EnhancedPrintMessage()
}

// MessagePrinter class has 2 members
type MessagePrinter struct {
	message string
	number  int
}

// PrintMessage is a function that prints a message
func (mp *MessagePrinter) PrintMessage() {
	println(mp.message, mp.number)
}

// enhancedMessagePrinter is a enhancement struct
type enhancedMessagePrinter struct {
	MessagePrinter
}

func (emp *enhancedMessagePrinter) EnhancedPrintMessage() {
	println("Enhanced -", emp.message, emp.number)
}
