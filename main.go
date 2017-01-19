package main

func main() {
	mp := MessagePrinter{"foo", 5}
	mp.PrintMessage()

	emp := enhancedMessagePrinter{MessagePrinter{"bar", 10}, 999}
	emp.PrintMessage()
	emp.EnhancedPrintMessage()

	emp2 := enhancedMessagePrinter{}
	emp2.message = "buz"
	emp2.number = 100
	emp2.enhancednumber = 888
	emp2.EnhancedPrintMessage()
}

// MessagePrinter class has 2 members
type MessagePrinter struct {
	message string
	number  int
}

// PrintMessage is a function that prints a message
func (mpx *MessagePrinter) PrintMessage() {
	println(mpx.message, mpx.number)
}

// enhancedMessagePrinter is a enhancement struct
type enhancedMessagePrinter struct {
	MessagePrinter
	enhancednumber int
}

func (emp *enhancedMessagePrinter) EnhancedPrintMessage() {
	println("Enhanced -", emp.message, emp.number, emp.enhancednumber)
}
