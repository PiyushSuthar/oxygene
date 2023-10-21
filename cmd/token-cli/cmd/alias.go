package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var aliasCmd = &cobra.Command{
	Use: "alias",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ErrMissingSubcommand
	},
}

type Alias struct {
	gorm.Model
	Alias   string `gorm:"unique"`
	Address string `gorm:"unique"`
}

var addAliasCmd = &cobra.Command{
	Use: "add",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("add alias lol <address>")
		}
		alias := args[0]
		address := args[1]

		db, err := gorm.Open(sqlite.Open("oxygene.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		db.AutoMigrate(&Alias{})

		db.Create(&Alias{Alias: alias, Address: address})

		fmt.Printf("Alias %s added for address %s\n", alias, address)

		return nil
	}}
