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
// 			AddMatchSignalFunc: func(options ...dbus.MatchOption) error {
// 				panic("mock out the AddMatchSignal method")
// 			},
// 			ObjectFunc: func(dest string, path dbus.ObjectPath) dbusBusObject {
// 				panic("mock out the Object method")
// 			},
// 			SignalFunc: func(ch chan<- *dbus.Signal)  {
// 				panic("mock out the Signal method")
// 			},
// 		}
//
// 		// use mockeddbusConn in code that requires dbusConn
// 		// and then make assertions.
//
// 	}
type dbusConnMock struct {
	// AddMatchSignalFunc mocks the AddMatchSignal method.
	AddMatchSignalFunc func(options ...dbus.MatchOption) error

	// ObjectFunc mocks the Object method.
	ObjectFunc func(dest string, path dbus.ObjectPath) dbusBusObject

	// SignalFunc mocks the Signal method.
	SignalFunc func(ch chan<- *dbus.Signal)

	// calls tracks calls to the methods.
	calls struct {
		// AddMatchSignal holds details about calls to the AddMatchSignal method.
		AddMatchSignal []struct {
			// Options is the options argument value.
			Options []dbus.MatchOption
		}
		// Object holds details about calls to the Object method.
		Object []struct {
			// Dest is the dest argument value.
			Dest string
			// Path is the path argument value.
			Path dbus.ObjectPath
		}
		// Signal holds details about calls to the Signal method.
		Signal []struct {
			// Ch is the ch argument value.
			Ch chan<- *dbus.Signal
		}
	}
	lockAddMatchSignal sync.RWMutex
	lockObject         sync.RWMutex
	lockSignal         sync.RWMutex
}

// AddMatchSignal calls AddMatchSignalFunc.
func (mock *dbusConnMock) AddMatchSignal(options ...dbus.MatchOption) error {
	if mock.AddMatchSignalFunc == nil {
		panic("dbusConnMock.AddMatchSignalFunc: method is nil but dbusConn.AddMatchSignal was just called")
	}
	callInfo := struct {
		Options []dbus.MatchOption
	}{
		Options: options,
	}
	mock.lockAddMatchSignal.Lock()
	mock.calls.AddMatchSignal = append(mock.calls.AddMatchSignal, callInfo)
	mock.lockAddMatchSignal.Unlock()
	return mock.AddMatchSignalFunc(options...)
}

// AddMatchSignalCalls gets all the calls that were made to AddMatchSignal.
// Check the length with:
//     len(mockeddbusConn.AddMatchSignalCalls())
func (mock *dbusConnMock) AddMatchSignalCalls() []struct {
	Options []dbus.MatchOption
} {
	var calls []struct {
		Options []dbus.MatchOption
	}
	mock.lockAddMatchSignal.RLock()
	calls = mock.calls.AddMatchSignal
	mock.lockAddMatchSignal.RUnlock()
	return calls
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

// Signal calls SignalFunc.
func (mock *dbusConnMock) Signal(ch chan<- *dbus.Signal) {
	if mock.SignalFunc == nil {
		panic("dbusConnMock.SignalFunc: method is nil but dbusConn.Signal was just called")
	}
	callInfo := struct {
		Ch chan<- *dbus.Signal
	}{
		Ch: ch,
	}
	mock.lockSignal.Lock()
	mock.calls.Signal = append(mock.calls.Signal, callInfo)
	mock.lockSignal.Unlock()
	mock.SignalFunc(ch)
}

// SignalCalls gets all the calls that were made to Signal.
// Check the length with:
//     len(mockeddbusConn.SignalCalls())
func (mock *dbusConnMock) SignalCalls() []struct {
	Ch chan<- *dbus.Signal
} {
	var calls []struct {
		Ch chan<- *dbus.Signal
	}
	mock.lockSignal.RLock()
	calls = mock.calls.Signal
	mock.lockSignal.RUnlock()
	return calls
}
