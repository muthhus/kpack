// Code generated by counterfeiter. DO NOT EDIT.
package buildfakes

import (
	"sync"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
	"github.com/pivotal/kpack/pkg/cnb"
	"github.com/pivotal/kpack/pkg/reconciler/v1alpha1/build"
)

type FakeMetadataRetriever struct {
	GetBuiltImageStub        func(*v1alpha1.Build) (cnb.BuiltImage, error)
	getBuiltImageMutex       sync.RWMutex
	getBuiltImageArgsForCall []struct {
		arg1 *v1alpha1.Build
	}
	getBuiltImageReturns struct {
		result1 cnb.BuiltImage
		result2 error
	}
	getBuiltImageReturnsOnCall map[int]struct {
		result1 cnb.BuiltImage
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMetadataRetriever) GetBuiltImage(arg1 *v1alpha1.Build) (cnb.BuiltImage, error) {
	fake.getBuiltImageMutex.Lock()
	ret, specificReturn := fake.getBuiltImageReturnsOnCall[len(fake.getBuiltImageArgsForCall)]
	fake.getBuiltImageArgsForCall = append(fake.getBuiltImageArgsForCall, struct {
		arg1 *v1alpha1.Build
	}{arg1})
	fake.recordInvocation("GetBuiltImage", []interface{}{arg1})
	fake.getBuiltImageMutex.Unlock()
	if fake.GetBuiltImageStub != nil {
		return fake.GetBuiltImageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBuiltImageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMetadataRetriever) GetBuiltImageCallCount() int {
	fake.getBuiltImageMutex.RLock()
	defer fake.getBuiltImageMutex.RUnlock()
	return len(fake.getBuiltImageArgsForCall)
}

func (fake *FakeMetadataRetriever) GetBuiltImageCalls(stub func(*v1alpha1.Build) (cnb.BuiltImage, error)) {
	fake.getBuiltImageMutex.Lock()
	defer fake.getBuiltImageMutex.Unlock()
	fake.GetBuiltImageStub = stub
}

func (fake *FakeMetadataRetriever) GetBuiltImageArgsForCall(i int) *v1alpha1.Build {
	fake.getBuiltImageMutex.RLock()
	defer fake.getBuiltImageMutex.RUnlock()
	argsForCall := fake.getBuiltImageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetadataRetriever) GetBuiltImageReturns(result1 cnb.BuiltImage, result2 error) {
	fake.getBuiltImageMutex.Lock()
	defer fake.getBuiltImageMutex.Unlock()
	fake.GetBuiltImageStub = nil
	fake.getBuiltImageReturns = struct {
		result1 cnb.BuiltImage
		result2 error
	}{result1, result2}
}

func (fake *FakeMetadataRetriever) GetBuiltImageReturnsOnCall(i int, result1 cnb.BuiltImage, result2 error) {
	fake.getBuiltImageMutex.Lock()
	defer fake.getBuiltImageMutex.Unlock()
	fake.GetBuiltImageStub = nil
	if fake.getBuiltImageReturnsOnCall == nil {
		fake.getBuiltImageReturnsOnCall = make(map[int]struct {
			result1 cnb.BuiltImage
			result2 error
		})
	}
	fake.getBuiltImageReturnsOnCall[i] = struct {
		result1 cnb.BuiltImage
		result2 error
	}{result1, result2}
}

func (fake *FakeMetadataRetriever) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBuiltImageMutex.RLock()
	defer fake.getBuiltImageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMetadataRetriever) recordInvocation(key string, args []interface{}) {
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

var _ build.MetadataRetriever = new(FakeMetadataRetriever)
