package services

import (
	"github.com/split-notes/pennant-flagger/library/appcontext"
	"github.com/split-notes/pennant-flagger/services/hello"
	"reflect"
)

type Bundle struct {
	HelloSvc hello.Service
}

var bundlables = []bundlable{
	hello.Helpable{},
}

type bundlable interface {
	NewHelper(appCtx appcontext.Context) (interface{}, error)
	ServiceName() string
}

func NewBundle(appCtx appcontext.Context) (*Bundle, error) {
	bundle := &Bundle{}

	for _, bundlable := range bundlables {
		helper, err := bundlable.NewHelper(appCtx)
		if err != nil {
			return nil, err
		}
		SetField(bundle, bundlable.ServiceName(), helper)
	}

	return bundle, nil
}

func SetField(bundler *Bundle, field string, value interface{}) {
	v := reflect.ValueOf(bundler).Elem().FieldByName(field)
	if v.IsValid() {
		v.Set(reflect.ValueOf(value))
	}
}