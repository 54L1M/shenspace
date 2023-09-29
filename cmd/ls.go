/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/0xshen/shenspace/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all projects.",
	Long:  `Lists all directories in the $PROJECTS_HOME directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		// name, _ := os.UserHomeDir()
		homeDir, err := os.UserHomeDir()
		cobra.CheckErr(err)

		sf := viper.GetString("proj.dir")
		confPath := homeDir + "/" + sf
		fmt.Println(confPath)
		utils.CheckConf(sf)

		dirs, err := os.ReadDir(confPath)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Projects found:")
		fmt.Println("===============")
		for _, dir := range dirs {
			if dir.IsDir() {
				fmt.Println(dir.Name())
				fmt.Println("=========")

			}
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
