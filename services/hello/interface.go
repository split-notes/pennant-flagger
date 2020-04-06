package hello

import (
	"github.com/split-notes/pennant-flagger/db/models"
	"github.com/split-notes/pennant-flagger/library/appcontext"
)

type Helper struct {
	AppCtx appcontext.Context
}
type Helpable struct{}

func(h Helpable) NewHelper(appCtx appcontext.Context) (interface{}, error) {
	return &Helper{
		AppCtx: appCtx,
	}, nil
}

func (h Helpable) ServiceName() string {
	return "HelloSvc"
}

type Service interface {
	Create(greetingModel models.Greetings) (*models.Greetings, error)
	Get() ([]models.Greetings, error)
}
