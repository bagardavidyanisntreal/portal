package portal

import (
	"fmt"
	"testing"
)

type subscribeTestHandler struct {
	marker  string
	storage *testStorage
}

func newSubscribeTestHandler(marker string, storage *testStorage) *subscribeTestHandler {
	return &subscribeTestHandler{
		marker:  marker,
		storage: storage,
	}
}

func (t subscribeTestHandler) Handle(msg any) {
	data := fmt.Sprintf(t.marker, msg)
	t.storage.Add(data)
}

func TestPortal_Subscribe_SubscriptionsAdded(t *testing.T) {
	t.Parallel()
	portal := New()

	storage := &testStorage{}
	portal.Subscribe(
		newSubscribeTestHandler("[handler-1]: '%v'", storage),
		newSubscribeTestHandler("[handler-2]: '%v'", storage),
	)

	portal.Send("I am test data")
	portal.Send("Some new data!")
	portal.Send("And more...")

	portal.Close()

	wantData := map[string]struct{}{
		"[handler-1]: 'I am test data'": {},
		"[handler-2]: 'I am test data'": {},
		"[handler-1]: 'Some new data!'": {},
		"[handler-2]: 'Some new data!'": {},
		"[handler-1]: 'And more...'":    {},
		"[handler-2]: 'And more...'":    {},
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

type uselessHandler struct{ marker string }

func (h uselessHandler) Handle(msg any) { fmt.Printf(h.marker, msg) }

func TestPortal_Subscribe_NoNegativeWGCounter(t *testing.T) {
	// this test checks nothing but stdout with sent messages on subscribed handlers
	// and that there is no panic on negative wg counter
	t.Parallel()

	gate := New()
	defer gate.Close()

	gate.Send("no one will see this")

	gate.Subscribe(
		&uselessHandler{marker: "[handler-1]: %v\n"},
		&uselessHandler{marker: "[handler-2]: %v\n"},
		&uselessHandler{marker: "[handler-3]: %v\n"},
	)
	gate.Send("first message")

	gate.Subscribe(
		&uselessHandler{marker: "[handler-4]: %v\n"},
		&uselessHandler{marker: "[handler-5]: %v\n"},
		&uselessHandler{marker: "[handler-6]: %v\n"},
	)
	gate.Send("second message")

	fmt.Println("[main] useless handlers sucks")
}
