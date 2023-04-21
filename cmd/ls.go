/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all projects.",
	Long:  `Lists all directories in the $PROJECTS_HOME directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		// name, _ := os.UserHomeDir()
        viper.AddConfigPath("./configs")
        viper.SetConfigName("conf")
        viper.SetConfigType("yaml")
        viper.ReadInConfig()
        name := viper.Get("proj.dirname")
        iname := fmt.Sprint(name)        
     
        fmt.Println(name)
		dirs, err := os.ReadDir(iname)

		if err != nil {
			log.Fatal(err)
		}
        fmt.Println("Projects found:")
        fmt.Println("=========")
		for _, dir := range dirs {
			if dir.IsDir() {
				fmt.Println(dir.Name())
			}
		}
		bin, err := exec.LookPath("echo")
		if err != nil {
			log.Fatal(err)
		}
		argus := []string{"echo", "hello from golang"}
		env := os.Environ()
		execErr := syscall.Exec(bin, argus, env)
		if execErr != nil {
			log.Fatal(execErr)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
