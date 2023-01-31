package portal

import (
	"context"
	"fmt"
	"testing"
)

type awaitTestHandler1 struct {
	storage *testStorage
}

func (t awaitTestHandler1) Support(_ Message) bool {
	return true
}

func (t awaitTestHandler1) Handle(msg Message) {
	data := fmt.Sprintf("[awaitTestHandler1]: '%v'", msg.Data())
	t.storage.Add(data)
}

type awaitTestHandler2 struct {
	storage *testStorage
}

func (t awaitTestHandler2) Support(_ Message) bool {
	return true
}

func (t awaitTestHandler2) Handle(msg Message) {
	data := fmt.Sprintf("[awaitTestHandler2]: '%v'", msg.Data())
	t.storage.Add(data)
}

type awaitTestMessage struct{}

func (a awaitTestMessage) Data() any {
	return "I am await function test data"
}

func TestPortal_Await_SubscriptionsAdded(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	portal := New(ctx)

	storage := &testStorage{}

	portal.Await(ctx,
		&awaitTestHandler1{storage: storage},
		&awaitTestHandler2{storage: storage},
	)
	portal.Send(awaitTestMessage{})

	portal.Close()

	wantData := map[string]struct{}{
		"[awaitTestHandler1]: 'I am await function test data'": {},
		"[awaitTestHandler2]: 'I am await function test data'": {},
	}

	gotData := storage.Data()
	for _, stored := range gotData {
		_, ok := wantData[stored]
		if !ok {
			t.Errorf("cannot find needed message in got: %s", stored)
		}
	}
}
