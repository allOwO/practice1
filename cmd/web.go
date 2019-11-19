package cmd

import (
	"PracticeItem/web"
	"github.com/spf13/cobra"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
			e := echo.New()
			e.Use(middleware.Recover())
			e.Static("/","./dist/index.html")
			e.Static("/static", "./dist/static")
			//e.GET("/",Index)
			e.POST("/senduser", web.CreateUser)
			e.POST("/changeuser", web.ChangeUser)
			e.GET("/checkuser",web.CheckUser)
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