package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"jlowell000.github.io/budgeting/internal/io"
	"jlowell000.github.io/budgeting/internal/model/data"
	"jlowell000.github.io/budgeting/internal/service"
	"jlowell000.github.io/budgeting/internal/service/accountservice"
	"jlowell000.github.io/budgeting/internal/service/dataservice"
	"jlowell000.github.io/budgeting/internal/service/periodicflowservice"

	"jlowell000.github.io/budgeting/internal/views"
	"jlowell000.github.io/budgeting/internal/views/accountlist"
	"jlowell000.github.io/budgeting/internal/views/accountview"
	"jlowell000.github.io/budgeting/internal/views/flowlist"
	"jlowell000.github.io/budgeting/internal/views/mainview"
)

const (
	CMD_CREATE         = "create"
	CMD_READ           = "read"
	CMD_QUIT           = "quit"
	FLG_ALL            = "all"
	VAR_CONTENT        = "content"
	ENTRYLIST_FILENAME = "./data.json"
)

var (
	d  *data.DataModel
	ds service.DataServiceInterface = &dataservice.DataService{
		Filename:    ENTRYLIST_FILENAME,
		GetDataJSON: io.ReadFromFile,
		PutDataJSON: io.WriteToFile,
	}
	flowService service.PeriodicFlowServiceInterface = &periodicflowservice.PeriodicFlowService{
		Dataservice: ds,
		GetTime:     time.Now,
		GetId:       uuid.New,
	}
	accountService service.AccountServiceInterface = &accountservice.AccountService{
		Dataservice: ds,
		GetTime:     time.Now,
		GetId:       uuid.New,
	}
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() views.AppModel {
	d = ds.GetData()

	return views.AppModel{
		Main: mainview.MainModel{
			Choice:         1,
			Selected:       make(map[int]struct{}),
			AccountService: accountService,
			FlowService:    flowService,
		},
		FlowList: flowlist.FlowListModel{
			Selected:       make(map[int]struct{}),
			AccountService: accountService,
			FlowService:    flowService,
		},
		AccountList: accountlist.AccountListModel{
			Selected:       make(map[int]struct{}),
			AccountService: accountService,
			FlowService:    flowService,
		},
		Account: accountview.AccountModel{
			AccountService: accountService,
		},
		SavaDataFunc: func() { d = ds.SaveData(d) },
	}
}
