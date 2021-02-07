// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mpris

import (
	"github.com/godbus/dbus/v5"
	"sync"
)

var (
	lockdbusWrapperIMockCall sync.RWMutex
)

// Ensure, that dbusWrapperIMock does implement dbusWrapperI.
// If this is not the case, regenerate this file with moq.
var _ dbusWrapperI = &dbusWrapperIMock{}

// dbusWrapperIMock is a mock implementation of dbusWrapperI.
//
//     func TestSomethingThatUsesdbusWrapperI(t *testing.T) {
//
//         // make and configure a mocked dbusWrapperI
//         mockeddbusWrapperI := &dbusWrapperIMock{
//             CallFunc: func(dest string, path dbus.ObjectPath, method string, flags dbus.Flags, args ...interface{})  {
// 	               panic("mock out the Call method")
//             },
//         }
//
//         // use mockeddbusWrapperI in code that requires dbusWrapperI
//         // and then make assertions.
//
//     }
type dbusWrapperIMock struct {
	// CallFunc mocks the Call method.
	CallFunc func(dest string, path dbus.ObjectPath, method string, flags dbus.Flags, args ...interface{})

	// calls tracks calls to the methods.
	calls struct {
		// Call holds details about calls to the Call method.
		Call []struct {
			// Dest is the dest argument value.
			Dest string
			// Path is the path argument value.
			Path dbus.ObjectPath
			// Method is the method argument value.
			Method string
			// Flags is the flags argument value.
			Flags dbus.Flags
			// Args is the args argument value.
			Args []interface{}
		}
	}
}

// Call calls CallFunc.
func (mock *dbusWrapperIMock) Call(dest string, path dbus.ObjectPath, method string, flags dbus.Flags, args ...interface{}) {
	if mock.CallFunc == nil {
		panic("dbusWrapperIMock.CallFunc: method is nil but dbusWrapperI.Call was just called")
	}
	callInfo := struct {
		Dest   string
		Path   dbus.ObjectPath
		Method string
		Flags  dbus.Flags
		Args   []interface{}
	}{
		Dest:   dest,
		Path:   path,
		Method: method,
		Flags:  flags,
		Args:   args,
	}
	lockdbusWrapperIMockCall.Lock()
	mock.calls.Call = append(mock.calls.Call, callInfo)
	lockdbusWrapperIMockCall.Unlock()
	mock.CallFunc(dest, path, method, flags, args...)
}

// CallCalls gets all the calls that were made to Call.
// Check the length with:
//     len(mockeddbusWrapperI.CallCalls())
func (mock *dbusWrapperIMock) CallCalls() []struct {
	Dest   string
	Path   dbus.ObjectPath
	Method string
	Flags  dbus.Flags
	Args   []interface{}
} {
	var calls []struct {
		Dest   string
		Path   dbus.ObjectPath
		Method string
		Flags  dbus.Flags
		Args   []interface{}
	}
	lockdbusWrapperIMockCall.RLock()
	calls = mock.calls.Call
	lockdbusWrapperIMockCall.RUnlock()
	return calls
}
