/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Applies shenspace conf on project and opens it.",
	Long:  `Applies shenspace conf on project and opens it.`,
	Run: func(cmd *cobra.Command, args []string) {

		bin, err := exec.LookPath("tmux")
		if err != nil {
			log.Fatal(err)
		}
		argus := []string{"tmux", "new", "-s", args[0]}
		env := os.Environ()
		execErr := syscall.Exec(bin, argus, env)
		if execErr != nil {
			log.Fatal(execErr)
		}
	},
}

func init() {
	rootCmd.AddCommand(onCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
