// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mpris

import (
	"github.com/godbus/dbus/v5"
	"sync"
)

// Ensure, that dbusConnMock does implement dbusConn.
// If this is not the case, regenerate this file with moq.
var _ dbusConn = &dbusConnMock{}

// dbusConnMock is a mock implementation of dbusConn.
//
// 	func TestSomethingThatUsesdbusConn(t *testing.T) {
//
// 		// make and configure a mocked dbusConn
// 		mockeddbusConn := &dbusConnMock{
// 			ObjectFunc: func(dest string, path dbus.ObjectPath) dbusBusObject {
// 				panic("mock out the Object method")
// 			},
// 		}
//
// 		// use mockeddbusConn in code that requires dbusConn
// 		// and then make assertions.
//
// 	}
type dbusConnMock struct {
	// ObjectFunc mocks the Object method.
	ObjectFunc func(dest string, path dbus.ObjectPath) dbusBusObject

	// calls tracks calls to the methods.
	calls struct {
		// Object holds details about calls to the Object method.
		Object []struct {
			// Dest is the dest argument value.
			Dest string
			// Path is the path argument value.
			Path dbus.ObjectPath
		}
	}
	lockObject sync.RWMutex
}

// Object calls ObjectFunc.
func (mock *dbusConnMock) Object(dest string, path dbus.ObjectPath) dbusBusObject {
	if mock.ObjectFunc == nil {
		panic("dbusConnMock.ObjectFunc: method is nil but dbusConn.Object was just called")
	}
	callInfo := struct {
		Dest string
		Path dbus.ObjectPath
	}{
		Dest: dest,
		Path: path,
	}
	mock.lockObject.Lock()
	mock.calls.Object = append(mock.calls.Object, callInfo)
	mock.lockObject.Unlock()
	return mock.ObjectFunc(dest, path)
}

// ObjectCalls gets all the calls that were made to Object.
// Check the length with:
//     len(mockeddbusConn.ObjectCalls())
func (mock *dbusConnMock) ObjectCalls() []struct {
	Dest string
	Path dbus.ObjectPath
} {
	var calls []struct {
		Dest string
		Path dbus.ObjectPath
	}
	mock.lockObject.RLock()
	calls = mock.calls.Object
	mock.lockObject.RUnlock()
	return calls
}
