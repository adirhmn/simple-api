package cmd

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"simpleapi/app"
	"simpleapi/handler"
	"simpleapi/tools/dbpool"
)

var rootCmd = &cobra.Command{
	Use:   "simpleapi",
	Short: "simpleapi service",
	Long:  "simpleapi service",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := &app.Config{}
		viper.AddConfigPath(".")
		viper.SetConfigName(".env")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Panic(err)
		}

		viper.AutomaticEnv()

		err = viper.Unmarshal(&cfg)

		dbConfig := &dbpool.Config{}
		err = copier.Copy(dbConfig, cfg)
		if err != nil {
			log.Panic(err)
		}

		db, err := dbpool.New(dbConfig, []*dbpool.Config{dbConfig}, &dbpool.ConnectionConfig{
			SetMaxIdleConns: 100,
			SetMaxOpenConns: 100,
		})
		if err != nil {
			log.Panic(err)
		}

		cont, err := app.NewServiceContainers(cfg, db)
		if err != nil {
			log.Panic(err)
		}
		defer cont.Close()

		e := echo.New()
		handler.RegisterRoutes(cont, e)
		e.Logger.Fatal(e.Start(":1323"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
