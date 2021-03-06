// Code generated by counterfeiter. DO NOT EDIT.
package writersfakes

import (
	"sync"

	"github.com/jjmschofield/gocrawl/internal/pages"
	"github.com/jjmschofield/gocrawl/internal/writers"
)

type FakeWriter struct {
	StartStub        func(chan pages.Page)
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 chan pages.Page
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWriter) Start(arg1 chan pages.Page) {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 chan pages.Page
	}{arg1})
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		fake.StartStub(arg1)
	}
}

func (fake *FakeWriter) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeWriter) StartCalls(stub func(chan pages.Page)) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeWriter) StartArgsForCall(i int) chan pages.Page {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWriter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ writers.Writer = new(FakeWriter)
