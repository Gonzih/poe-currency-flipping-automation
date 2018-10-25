package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "poe-flipper",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var scanCmd = &cobra.Command{
	Use: "scan",
	Run: func(cmd *cobra.Command, args []string) {
		online := true

		pairs := make([]Pair, 0)

		if len(currenciesToScan) == 0 {
			for name := range currencyNames {
				currenciesToScan = append(currenciesToScan, name)
			}
		}

		for _, name := range currenciesToScan {
			switch name {
			case "wisdom", "chaos":
				continue
			default:
			}

			offers1 := searchFor(online, "Delve", name, "chaos")
			offers2 := searchFor(online, "Delve", "chaos", name)

			for _, offer1 := range offers1 {
				p := Pair{of: offer1, ofs: make([]SearchOffer, 0)}

				for _, offer2 := range offers2 {
					if offer1.IsProfitable(offer2) {
						p.ofs = append(p.ofs, offer2)
					}
				}

				if len(p.ofs) > 0 {
					pairs = append(pairs, p)
				}
			}
		}

		sort.Slice(pairs, func(i, j int) bool { return pairs[i].MaxProfit() > pairs[j].MaxProfit() })

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

var currenciesToScan []string

func init() {
	rootCmd.PersistentFlags().StringArrayVar(&currenciesToScan, "currencies", []string{}, "Currencies to scan")

	rootCmd.AddCommand(scanCmd, listCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
