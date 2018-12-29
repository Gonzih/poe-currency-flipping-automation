package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "poe-flipper",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var uiCmd = &cobra.Command{
	Use: "ui",
	Run: func(cmd *cobra.Command, args []string) {
		renderUI()
	},
}

var scanCmd = &cobra.Command{
	Use: "scan",
	Run: func(cmd *cobra.Command, args []string) {
		pairs := getPairs()

		for i := 0; i < len(pairs); i++ {
			fmt.Println(pairs[i].String())
		}
	},
}

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		cur := make([]string, 0)
		for name := range currencyNames {
			cur = append(cur, name)
		}

		fmt.Printf("Available currencies:\n\"%s\"\n", strings.Join(cur, "\"\n\""))
	},
}

var (
	currenciesToScan []string
	onlineSearch     bool
	leagueToSearch   string
)

func init() {
	rootCmd.PersistentFlags().StringArrayVar(&currenciesToScan, "currencies", []string{}, "Currencies to scan")
	rootCmd.PersistentFlags().BoolVar(&onlineSearch, "online", true, "Perform online only search")
	rootCmd.PersistentFlags().StringVar(&leagueToSearch, "league", "Betrayal", "Which league to use for search")

	rootCmd.AddCommand(scanCmd, listCmd, uiCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
