// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package nvcdi

import (
	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvml"
	"sync"
)

// Ensure, that nvmlUUIDerMock does implement nvmlUUIDer.
// If this is not the case, regenerate this file with moq.
var _ nvmlUUIDer = &nvmlUUIDerMock{}

// nvmlUUIDerMock is a mock implementation of nvmlUUIDer.
//
//	func TestSomethingThatUsesnvmlUUIDer(t *testing.T) {
//
//		// make and configure a mocked nvmlUUIDer
//		mockednvmlUUIDer := &nvmlUUIDerMock{
//			GetUUIDFunc: func() (string, nvml.Return) {
//				panic("mock out the GetUUID method")
//			},
//		}
//
//		// use mockednvmlUUIDer in code that requires nvmlUUIDer
//		// and then make assertions.
//
//	}
type nvmlUUIDerMock struct {
	// GetUUIDFunc mocks the GetUUID method.
	GetUUIDFunc func() (string, nvml.Return)

	// calls tracks calls to the methods.
	calls struct {
		// GetUUID holds details about calls to the GetUUID method.
		GetUUID []struct {
		}
	}
	lockGetUUID sync.RWMutex
}

// GetUUID calls GetUUIDFunc.
func (mock *nvmlUUIDerMock) GetUUID() (string, nvml.Return) {
	callInfo := struct {
	}{}
	mock.lockGetUUID.Lock()
	mock.calls.GetUUID = append(mock.calls.GetUUID, callInfo)
	mock.lockGetUUID.Unlock()
	if mock.GetUUIDFunc == nil {
		var (
			sOut      string
			returnOut nvml.Return
		)
		return sOut, returnOut
	}
	return mock.GetUUIDFunc()
}

// GetUUIDCalls gets all the calls that were made to GetUUID.
// Check the length with:
//
//	len(mockednvmlUUIDer.GetUUIDCalls())
func (mock *nvmlUUIDerMock) GetUUIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetUUID.RLock()
	calls = mock.calls.GetUUID
	mock.lockGetUUID.RUnlock()
	return calls
}
