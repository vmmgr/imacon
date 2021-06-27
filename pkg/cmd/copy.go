package cmd

import (
	"github.com/spf13/cobra"
	controllerInt "github.com/vmmgr/imacon/pkg/api/core/controller"
	controller "github.com/vmmgr/imacon/pkg/api/core/controller/v2"
	storage "github.com/vmmgr/imacon/pkg/api/core/storage/v2"
	"github.com/vmmgr/imacon/pkg/api/core/tool/config"
	"log"
	"time"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "start client server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		confPath, err := cmd.Flags().GetString("config")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if config.GetConfig(confPath) != nil {
			log.Fatalf("error config process |%v", err)
		}

		uuid, err := cmd.Flags().GetString("uuid")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
			return
		}
		if uuid == "" {
			log.Fatalln("uuid is not select")
			return
		}
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if url == "" {
			log.Fatalln("url is not select")
			return
		}
		srcPath, err := cmd.Flags().GetString("src")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if srcPath == "" {
			log.Fatalln("srcPath is not select")
			return
		}
		dstPath, err := cmd.Flags().GetString("dst")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if dstPath == "" {
			log.Fatalln("dstPath is not select")
			return
		}
		dstAddr, err := cmd.Flags().GetString("addr")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if dstAddr == "" {
			log.Fatalln("dstAddr is not select")
			return
		}
		dstUser, err := cmd.Flags().GetString("user")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if dstUser == "" {
			log.Fatalln("dstUser is not select")
			return
		}

		err = storage.Copy(uuid, url, srcPath, dstPath, dstAddr, dstUser, config.Conf.PublicKeyPath)
		if err != nil {
			log.Println(err)
			go func() {
				for {
					<-time.NewTimer(10 * time.Second).C
					err = controller.SendController(url, controllerInt.Controller{UUID: uuid, Error: err.Error()})
					if err == nil {
						break
					}
				}
			}()
		}

		log.Println("end")
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.PersistentFlags().StringP("config", "c", "", "config")
	copyCmd.PersistentFlags().StringP("uuid", "u", "", "UUID")
	copyCmd.PersistentFlags().StringP("url", "l", "", "URL")
	copyCmd.PersistentFlags().StringP("src", "s", "", "src path")
	copyCmd.PersistentFlags().StringP("dst", "d", "", "dst path")
	copyCmd.PersistentFlags().StringP("addr", "a", "", "dst Addr")
	copyCmd.PersistentFlags().StringP("user", "r", "", "dst User")
}
