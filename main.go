package main

import (
  "fmt"
  "strings"

  "github.com/spf13/cobra"
)

func main() {
  compatibility := &cobra.Command{
    Use:   "compatibility",
    Short: "Check compatibility validation against the GraphQL standard reference implementation.",
    Long: `This command allows you to compare a GraphQL implementation against the official reference implementation.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Print: " + strings.Join(args, " "))
    },
  }

  var rootCmd = &cobra.Command{Use: "app"}
  rootCmd.AddCommand(compatibility)
  rootCmd.Execute()
}
