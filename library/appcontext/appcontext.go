package appcontext

import (
	"context"
	"github.com/split-notes/pennant-flagger/configs"
	"github.com/split-notes/pennant-flagger/db"
)

type Context struct {
	DB db.Connection
	Config configs.Configuration
	GoContext context.Context
}