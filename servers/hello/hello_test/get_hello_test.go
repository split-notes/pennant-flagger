package hello_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/split-notes/pennant-flagger/db/models"
	"github.com/split-notes/pennant-flagger/library/appcontext"
	"github.com/split-notes/pennant-flagger/mocks/services_mocks"
	"github.com/split-notes/pennant-flagger/servers"
	"github.com/split-notes/pennant-flagger/servers/hello"
	"github.com/split-notes/pennant-flagger/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHelloServer(t *testing.T) {
	var GetHelloUnitTestData = []GetHelloTestData{
		{
			BaseTestData: BaseTestData{
				Description: "happy path",
				Request:     hello.GetHelloRequest{},
				Response:    hello.GetHelloReply{
					Greetings: []string{"one", "two"},
				},
			},
			MockGetHello: &MockGetHello{
				OutGreetings: []models.Greetings{
					{ Value: "one" },
					{ Value: "two" },
				},
				OutError:     nil,
			},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, data := range GetHelloUnitTestData {
		t.Run(data.Description, func(t *testing.T) {
			serviceBundle := MockGetHelloRequiredServices(mockCtrl, data)

			server := hello.Server{
				ServerContext: servers.ServerContext{
					AppCtx: appcontext.Context{},
					Bundle: *serviceBundle,
				},
			}

			requestTestData := data.Request.(hello.GetHelloRequest)
			responseTestData := data.Response.(hello.GetHelloReply)
			responseData, _ := server.GetHello(context.Background(), &requestTestData)

			assert.Equal(t, responseTestData, *responseData)
		})
	}
}

func MockGetHelloRequiredServices(mockCtrl *gomock.Controller, data GetHelloTestData) *services.Bundle {
	helloMock := services_mocks.NewMock_hello(mockCtrl)
	helloExpect := helloMock.EXPECT()

	if data.MockGetHello != nil {
		helloExpect.Get().Return(
				data.MockGetHello.OutGreetings,
				data.MockGetHello.OutError)
	}

	serviceBundle := services.Bundle{
		HelloSvc: helloMock,
	}

	return &serviceBundle
}
