package cmd

import (
	"PracticeItem"
	"PracticeItem/controllers"
	"PracticeItem/service"
	"github.com/fpay/foundation-go/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	clientCmd := &cobra.Command{
		Use:   "dashboard",
		Short: "Messenger web",
		Run: func(cmd *cobra.Command, args []string) {
			env:=new(PracticeItem.AppConfig)
			env.Load()
			db, err := database.NewDatabase(database.DatabaseOptions{Driver: "mysql",Dsn:env.Mysqldsn})
			if err!=nil{
				log.Fatalln("Mysql error:",err)
			}
			user:=service.NewDBservice(db)
			web:=controllers.NewWebController(user,user)

			e := echo.New()
			e.Use(middleware.Recover())
			e.Static("/","./dist/index.html")
			e.Static("/static", "./dist/static")
			e.POST("/senduser", web.CreateUser)
			e.POST("/changeuser", web.ChangeUser)
			e.GET("/checkuser", web.CheckUser)
			e.Logger.Fatal(e.Start(":8000"))

			//退出
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			<-quit
			// Finish after all clients disconnected
			log.Println("recv exited")
		},
	}
	rootCmd.AddCommand(clientCmd)
}