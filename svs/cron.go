package svs

import (
	"github.com/robfig/cron/v3"
	"log"
)

type cronService struct {
	*cron.Cron
}

func NewService(list ...Crontab) {
	log.Println("init cron")
	s := &cronService{
		Cron: cron.New(),
	}
	for _, crontab := range list {
		log.Println("add cron: ", crontab.Name())
		if _, err := s.Cron.AddFunc(crontab.Spec(), cronCommandHandler(crontab)); err != nil {
			log.Fatalln(err)
		}
	}
	go s.Cron.Start()
	log.Println("init cron finish")
}

type Crontab interface {
	Name() string
	Spec() string
	Command() func() (interface{}, error)
}

func cronCommandHandler(c Crontab) func() {
	return func() {
		data, err := c.Command()()
		if err != nil {
			log.Printf("cron error: %s (%s)\n", c.Name(), err)
		}
		log.Printf("cron success: %s (%s)\n", c.Name(), data)
	}
}
