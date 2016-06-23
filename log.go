package main

import ( "fmt"
	 gc "github.com/daviddengcn/go-colortext"
)

func logError(msg string)  {
  gc.Foreground(gc.Red, false)
  fmt.Printf(msg + "\n")
  gc.ResetColor()
}

func logSuccess(msg string) {
  gc.Foreground(gc.Green, false)
  fmt.Printf(msg + "\n")
  gc.ResetColor()
}

func logVerbose(msg string, verbose bool) {
  if verbose {
    fmt.Printf(msg + "\n")
  }
}
