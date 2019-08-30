// Code generated by counterfeiter. DO NOT EDIT.
package compositefakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/actor/v2action/composite"
)

type FakeGetServicePlanActor struct {
	GetServicePlanStub        func(string) (v2action.ServicePlan, v2action.Warnings, error)
	getServicePlanMutex       sync.RWMutex
	getServicePlanArgsForCall []struct {
		arg1 string
	}
	getServicePlanReturns struct {
		result1 v2action.ServicePlan
		result2 v2action.Warnings
		result3 error
	}
	getServicePlanReturnsOnCall map[int]struct {
		result1 v2action.ServicePlan
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGetServicePlanActor) GetServicePlan(arg1 string) (v2action.ServicePlan, v2action.Warnings, error) {
	fake.getServicePlanMutex.Lock()
	ret, specificReturn := fake.getServicePlanReturnsOnCall[len(fake.getServicePlanArgsForCall)]
	fake.getServicePlanArgsForCall = append(fake.getServicePlanArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServicePlan", []interface{}{arg1})
	fake.getServicePlanMutex.Unlock()
	if fake.GetServicePlanStub != nil {
		return fake.GetServicePlanStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getServicePlanReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeGetServicePlanActor) GetServicePlanCallCount() int {
	fake.getServicePlanMutex.RLock()
	defer fake.getServicePlanMutex.RUnlock()
	return len(fake.getServicePlanArgsForCall)
}

func (fake *FakeGetServicePlanActor) GetServicePlanCalls(stub func(string) (v2action.ServicePlan, v2action.Warnings, error)) {
	fake.getServicePlanMutex.Lock()
	defer fake.getServicePlanMutex.Unlock()
	fake.GetServicePlanStub = stub
}

func (fake *FakeGetServicePlanActor) GetServicePlanArgsForCall(i int) string {
	fake.getServicePlanMutex.RLock()
	defer fake.getServicePlanMutex.RUnlock()
	argsForCall := fake.getServicePlanArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGetServicePlanActor) GetServicePlanReturns(result1 v2action.ServicePlan, result2 v2action.Warnings, result3 error) {
	fake.getServicePlanMutex.Lock()
	defer fake.getServicePlanMutex.Unlock()
	fake.GetServicePlanStub = nil
	fake.getServicePlanReturns = struct {
		result1 v2action.ServicePlan
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGetServicePlanActor) GetServicePlanReturnsOnCall(i int, result1 v2action.ServicePlan, result2 v2action.Warnings, result3 error) {
	fake.getServicePlanMutex.Lock()
	defer fake.getServicePlanMutex.Unlock()
	fake.GetServicePlanStub = nil
	if fake.getServicePlanReturnsOnCall == nil {
		fake.getServicePlanReturnsOnCall = make(map[int]struct {
			result1 v2action.ServicePlan
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServicePlanReturnsOnCall[i] = struct {
		result1 v2action.ServicePlan
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGetServicePlanActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServicePlanMutex.RLock()
	defer fake.getServicePlanMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGetServicePlanActor) recordInvocation(key string, args []interface{}) {
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

var _ composite.GetServicePlanActor = new(FakeGetServicePlanActor)