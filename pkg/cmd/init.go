package cmd

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/vmmgr/imacon/pkg/api/core/storage"
	"github.com/vmmgr/imacon/pkg/api/core/tool/config"
	"log"
	"strconv"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize",
	Long: `initialize command. For example:

database init: init database
`,
}
var initDBCmd = &cobra.Command{
	Use:   "store",
	Short: "store init",
	Long:  "store init cmd",
	RunE: func(cmd *cobra.Command, args []string) error {
		confPath, err := cmd.Flags().GetString("config")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return err
		}
		if config.GetConfig(confPath) != nil {
			log.Fatalf("error config process |%v", err)
			return err
		}

		db, err := gorm.Open("mysql", config.Conf.DB.User+":"+config.Conf.DB.Pass+"@"+
			"tcp("+config.Conf.DB.IP+":"+strconv.Itoa(config.Conf.DB.Port)+")"+"/"+config.Conf.DB.DBName+"?parseTime=true")
		if err != nil {
			return err
		}
		result := db.AutoMigrate(&storage.Storage{})
		log.Println(result.Error)
		log.Println("end")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.AddCommand(initDBCmd)
}
