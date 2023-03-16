package main

import (
	"dtmWebHook.com/m/v2/Config"
	"dtmWebHook.com/m/v2/Driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DtmReport struct {
	Gid        string `json:"gid"`
	Status     string `json:"status"`
	RetryCount int    `json:"retry_count"`
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.POST("/dtm-hook", func(c *gin.Context) {

		dtmReport := DtmReport{}

		c.ShouldBind(&dtmReport)

		configs := Config.InitConfig()

		message := fmt.Sprintf("Gid: %s; 重试状态: %s; 重试次数: %d", dtmReport.Gid, dtmReport.Status, dtmReport.RetryCount)

		for _, config := range configs.Configs {
			if config.Type == "dingTalk" {
				d := new(Driver.DingTalkMessageDriver)
				d.Send(config, message)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.Run(":9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
