// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"sync"
)

// Ensure, that KVQueueMock does implement utils.KVQueue.
// If this is not the case, regenerate this file with moq.
var _ utils.KVQueue = &KVQueueMock{}

// KVQueueMock is a mock implementation of utils.KVQueue.
//
// 	func TestSomethingThatUsesKVQueue(t *testing.T) {
//
// 		// make and configure a mocked utils.KVQueue
// 		mockedKVQueue := &KVQueueMock{
// 			DequeueFunc: func(value codec.ProtoMarshaler) bool {
// 				panic("mock out the Dequeue method")
// 			},
// 			EnqueueFunc: func(key utils.Keyer, value codec.ProtoMarshaler)  {
// 				panic("mock out the Enqueue method")
// 			},
// 		}
//
// 		// use mockedKVQueue in code that requires utils.KVQueue
// 		// and then make assertions.
//
// 	}
type KVQueueMock struct {
	// DequeueFunc mocks the Dequeue method.
	DequeueFunc func(value codec.ProtoMarshaler) bool

	// EnqueueFunc mocks the Enqueue method.
	EnqueueFunc func(key utils.Keyer, value codec.ProtoMarshaler)

	// calls tracks calls to the methods.
	calls struct {
		// Dequeue holds details about calls to the Dequeue method.
		Dequeue []struct {
			// Value is the value argument value.
			Value codec.ProtoMarshaler
		}
		// Enqueue holds details about calls to the Enqueue method.
		Enqueue []struct {
			// Key is the key argument value.
			Key utils.Keyer
			// Value is the value argument value.
			Value codec.ProtoMarshaler
		}
	}
	lockDequeue sync.RWMutex
	lockEnqueue sync.RWMutex
}

// Dequeue calls DequeueFunc.
func (mock *KVQueueMock) Dequeue(value codec.ProtoMarshaler) bool {
	if mock.DequeueFunc == nil {
		panic("KVQueueMock.DequeueFunc: method is nil but KVQueue.Dequeue was just called")
	}
	callInfo := struct {
		Value codec.ProtoMarshaler
	}{
		Value: value,
	}
	mock.lockDequeue.Lock()
	mock.calls.Dequeue = append(mock.calls.Dequeue, callInfo)
	mock.lockDequeue.Unlock()
	return mock.DequeueFunc(value)
}

// DequeueCalls gets all the calls that were made to Dequeue.
// Check the length with:
//     len(mockedKVQueue.DequeueCalls())
func (mock *KVQueueMock) DequeueCalls() []struct {
	Value codec.ProtoMarshaler
} {
	var calls []struct {
		Value codec.ProtoMarshaler
	}
	mock.lockDequeue.RLock()
	calls = mock.calls.Dequeue
	mock.lockDequeue.RUnlock()
	return calls
}

// Enqueue calls EnqueueFunc.
func (mock *KVQueueMock) Enqueue(key utils.Keyer, value codec.ProtoMarshaler) {
	if mock.EnqueueFunc == nil {
		panic("KVQueueMock.EnqueueFunc: method is nil but KVQueue.Enqueue was just called")
	}
	callInfo := struct {
		Key   utils.Keyer
		Value codec.ProtoMarshaler
	}{
		Key:   key,
		Value: value,
	}
	mock.lockEnqueue.Lock()
	mock.calls.Enqueue = append(mock.calls.Enqueue, callInfo)
	mock.lockEnqueue.Unlock()
	mock.EnqueueFunc(key, value)
}

// EnqueueCalls gets all the calls that were made to Enqueue.
// Check the length with:
//     len(mockedKVQueue.EnqueueCalls())
func (mock *KVQueueMock) EnqueueCalls() []struct {
	Key   utils.Keyer
	Value codec.ProtoMarshaler
} {
	var calls []struct {
		Key   utils.Keyer
		Value codec.ProtoMarshaler
	}
	mock.lockEnqueue.RLock()
	calls = mock.calls.Enqueue
	mock.lockEnqueue.RUnlock()
	return calls
}