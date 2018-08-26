// Code generated by counterfeiter. DO NOT EDIT.
package gistfakes

import (
	"context"
	"sync"

	"github.com/google/go-github/github"
	"github.com/jutkko/copy-pasta/store/gist"
)

type FakeGistClient struct {
	GetStub        func(context.Context, string) (*github.Gist, *github.Response, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getReturns struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}
	getReturnsOnCall map[int]struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}
	CreateStub        func(context.Context, *github.Gist) (*github.Gist, *github.Response, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 *github.Gist
	}
	createReturns struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}
	createReturnsOnCall map[int]struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}
	EditStub        func(context.Context, string, *github.Gist) (*github.Gist, *github.Response, error)
	editMutex       sync.RWMutex
	editArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 *github.Gist
	}
	editReturns struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}
	editReturnsOnCall map[int]struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGistClient) Get(arg1 context.Context, arg2 string) (*github.Gist, *github.Response, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getReturns.result1, fake.getReturns.result2, fake.getReturns.result3
}

func (fake *FakeGistClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeGistClient) GetArgsForCall(i int) (context.Context, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2
}

func (fake *FakeGistClient) GetReturns(result1 *github.Gist, result2 *github.Response, result3 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGistClient) GetReturnsOnCall(i int, result1 *github.Gist, result2 *github.Response, result3 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *github.Gist
			result2 *github.Response
			result3 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGistClient) Create(arg1 context.Context, arg2 *github.Gist) (*github.Gist, *github.Response, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 *github.Gist
	}{arg1, arg2})
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createReturns.result1, fake.createReturns.result2, fake.createReturns.result3
}

func (fake *FakeGistClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeGistClient) CreateArgsForCall(i int) (context.Context, *github.Gist) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1, fake.createArgsForCall[i].arg2
}

func (fake *FakeGistClient) CreateReturns(result1 *github.Gist, result2 *github.Response, result3 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGistClient) CreateReturnsOnCall(i int, result1 *github.Gist, result2 *github.Response, result3 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *github.Gist
			result2 *github.Response
			result3 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGistClient) Edit(arg1 context.Context, arg2 string, arg3 *github.Gist) (*github.Gist, *github.Response, error) {
	fake.editMutex.Lock()
	ret, specificReturn := fake.editReturnsOnCall[len(fake.editArgsForCall)]
	fake.editArgsForCall = append(fake.editArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 *github.Gist
	}{arg1, arg2, arg3})
	fake.recordInvocation("Edit", []interface{}{arg1, arg2, arg3})
	fake.editMutex.Unlock()
	if fake.EditStub != nil {
		return fake.EditStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.editReturns.result1, fake.editReturns.result2, fake.editReturns.result3
}

func (fake *FakeGistClient) EditCallCount() int {
	fake.editMutex.RLock()
	defer fake.editMutex.RUnlock()
	return len(fake.editArgsForCall)
}

func (fake *FakeGistClient) EditArgsForCall(i int) (context.Context, string, *github.Gist) {
	fake.editMutex.RLock()
	defer fake.editMutex.RUnlock()
	return fake.editArgsForCall[i].arg1, fake.editArgsForCall[i].arg2, fake.editArgsForCall[i].arg3
}

func (fake *FakeGistClient) EditReturns(result1 *github.Gist, result2 *github.Response, result3 error) {
	fake.EditStub = nil
	fake.editReturns = struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGistClient) EditReturnsOnCall(i int, result1 *github.Gist, result2 *github.Response, result3 error) {
	fake.EditStub = nil
	if fake.editReturnsOnCall == nil {
		fake.editReturnsOnCall = make(map[int]struct {
			result1 *github.Gist
			result2 *github.Response
			result3 error
		})
	}
	fake.editReturnsOnCall[i] = struct {
		result1 *github.Gist
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGistClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.editMutex.RLock()
	defer fake.editMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGistClient) recordInvocation(key string, args []interface{}) {
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

var _ gist.GistClient = new(FakeGistClient)
