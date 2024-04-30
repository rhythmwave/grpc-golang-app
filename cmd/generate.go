package cmd

import (
	"go-bponline/m/internal/database"

	"github.com/spf13/cobra"
	"gorm.io/gen"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Models",
	Run: func(cmd *cobra.Command, args []string) {
		database.ConnectWithDatabase()
		g := gen.NewGenerator(gen.Config{
			OutPath: "gen/query",
			Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		})
		g.UseDB(database.DB)
		g.ApplyBasic(
			// Generate structs from all tables of current database
			g.GenerateAllTable()...,
		)
		g.Execute()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
