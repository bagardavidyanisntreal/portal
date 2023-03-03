package portal

import (
	"fmt"
	"testing"
)

type awaitTestHandler1 struct {
	storage *testStorage
}

func (t awaitTestHandler1) Handle(msg any) {
	data := fmt.Sprintf("[awaitTestHandler1]: '%v'", msg)
	t.storage.Add(data)
}

type awaitTestHandler2 struct {
	storage *testStorage
}

func (t awaitTestHandler2) Handle(msg any) {
	data := fmt.Sprintf("[awaitTestHandler2]: '%v'", msg)
	t.storage.Add(data)
}

func TestPortal_Await_SubscriptionsAdded(t *testing.T) {
	t.Parallel()
	portal := New()

	storage := &testStorage{}

	portal.Subscribe(
		&awaitTestHandler1{storage: storage},
		&awaitTestHandler2{storage: storage},
	)
	portal.Send("I am await function test data")

	portal.Close()

	wantData := map[string]struct{}{
		"[awaitTestHandler1]: 'I am await function test data'": {},
		"[awaitTestHandler2]: 'I am await function test data'": {},
	}

	gotData := storage.Data()
	/*if len(gotData) != len(wantData) {
		t.Errorf("gotData len %v not eq to wantData len %v", len(gotData), len(wantData))
	}*/

	for _, stored := range gotData {
		_, ok := wantData[stored]
		if !ok {
			t.Errorf("cannot find needed message in got: %s", stored)
		}
	}
}
