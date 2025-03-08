package messageformatters

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}
type Bold struct {
	message string
}
type Code struct {
	message string
}

func (c Bold) Format() string {
	return `**` + c.message + `**`
}
func (c PlainText) Format() string {
	return c.message
}
func (c Code) Format() string {
	return "`" + c.message + "`"
}

// func formattedMessage(f Formatter) string {
// 	return fmt.Println(f.Format())
// }

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
