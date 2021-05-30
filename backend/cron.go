package startpage

import (
	"github.com/ecnepsnai/cron"
	"github.com/ecnepsnai/logtic"
)

func StartCron() {
	tab, err := cron.New([]cron.Job{
		{
			Pattern: "*/5 * * * *",
			Name:    "RefreshModules",
			Exec: func() {
				ModuleRefresh()
			},
		},
		{
			Pattern: "0 0 * * *",
			Name:    "RotateLogs",
			Exec: func() {
				logtic.Rotate()
			},
		},
	})
	if err != nil {
		panic(err)
	}

	go tab.Start()
}
