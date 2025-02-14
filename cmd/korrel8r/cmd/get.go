/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"

	"github.com/korrel8r/korrel8r/internal/pkg/must"
	"github.com/korrel8r/korrel8r/pkg/korrel8r"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get DOMAIN QUERY",
	Short: "Execute QUERY in DOMAIN and print the results",
	Long: `
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		e := newEngine()
		d := must.Must1(e.DomainErr(args[0]))
		q := must.Must1(d.UnmarshalQuery([]byte(args[1])))
		s := must.Must1(e.StoreErr(d.String()))

		log.V(3).Info("get", "query", q, "class", korrel8r.ClassName(q.Class()))
		result := newPrinter(os.Stdout)
		must.Must(s.Get(context.Background(), q, result))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
