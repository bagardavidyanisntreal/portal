package portal

import (
	"fmt"
	"testing"
)

type subscribeTestHandler1 struct {
	storage *testStorage
}

func (t subscribeTestHandler1) Handle(msg any) {
	data := fmt.Sprintf("[subscribeTestHandler1]: '%v'", msg)
	t.storage.Add(data)
}

type subscribeTestHandler2 struct {
	storage *testStorage
}

func (t subscribeTestHandler2) Handle(msg any) {
	data := fmt.Sprintf("[subscribeTestHandler2]: '%v'", msg)
	t.storage.Add(data)
}

func TestPortal_Await_SubscriptionsAdded(t *testing.T) {
	t.Parallel()
	portal := New()

	storage := &testStorage{}

	portal.Subscribe(
		&subscribeTestHandler1{storage: storage},
		&subscribeTestHandler2{storage: storage},
	)
	portal.Send("I am test data")
	portal.Send("Some new data!")
	portal.Send("And more...")

	portal.Close()

	wantData := map[string]struct{}{
		"[subscribeTestHandler1]: 'I am test data'": {},
		"[subscribeTestHandler2]: 'I am test data'": {},
		"[subscribeTestHandler1]: 'Some new data!'": {},
		"[subscribeTestHandler2]: 'Some new data!'": {},
		"[subscribeTestHandler1]: 'And more...'":    {},
		"[subscribeTestHandler2]: 'And more...'":    {},
	}

	gotData := storage.Data()
	if len(gotData) != len(wantData) {
		t.Errorf("gotData len %v not eq to wantData len %v", len(gotData), len(wantData))
	}

	for _, stored := range gotData {
		_, ok := wantData[stored]
		if !ok {
			t.Errorf("cannot find needed message in got: %s", stored)
		}
	}
}
