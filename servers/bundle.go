package servers

import (
	"github.com/split-notes/pennant-flagger/library/appcontext"
	"github.com/split-notes/pennant-flagger/services"
	"google.golang.org/grpc"
)

type ServerContext struct {
	AppCtx appcontext.Context
	Bundle services.Bundle
}

type Includable func(*grpc.Server, ServerContext)

var Includables []Includable

func Include(includable Includable) {
	Includables = append(Includables, includable)
}

func BundleAll(server *grpc.Server, appCtx appcontext.Context, bundle services.Bundle) {
	sharedContext := ServerContext{
		AppCtx: appCtx,
		Bundle: bundle,
	}
	for _, includable := range Includables {
		includable(server, sharedContext)
	}
}