package terminal
import "fmt"

const (
	YES = "YES"
	NO = "NO"
)

// AskForConfirmation asks a question for the user to confirm
func AskForConfirmation(question string) string  {
	answer := "2"
	fmt.Println(question)
	fmt.Println("1) Yes")
	fmt.Println("2) No")
	fmt.Print("> ")
	fmt.Scanln(&answer)
	switch answer {
	case "1":
		return YES
	case "2":
		return NO
	}

	return NO
}
