package job

import (
	"blog_server/internal/svc"
	"blog_server/job/resetStar"
	"fmt"
	"github.com/robfig/cron/v3"
)

func CronRun(svc *svc.ServiceContext) {
	c := cron.New()
	//更新点赞状态
	_, err := c.AddFunc("59 23 * * *", func() {
		l := resetStar.NewResetStartLogic(svc)
		err := l.Start()
		if err != nil {
			fmt.Println(err)
		}
	})

	if err != nil {
		return
	}

	fmt.Println("定时任务运行中...")

	c.Start()

}
