### 1. Ticker
https://pkg.go.dev/time#Tick

## Important Points
1. Remember: if a program exits before its goroutines have completed, those goroutines will be killed silently. Which of the function calls probably shouldn't run in the background as a goroutine?