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
//	func TestSomethingThatUsesdbusConn(t *testing.T) {
//
//		// make and configure a mocked dbusConn
//		mockeddbusConn := &dbusConnMock{
//			AddMatchSignalFunc: func(matchOptions ...dbus.MatchOption) error {
//				panic("mock out the AddMatchSignal method")
//			},
//			CloseFunc: func() error {
//				panic("mock out the Close method")
//			},
//			ObjectFunc: func(s string, objectPath dbus.ObjectPath) dbusBusObject {
//				panic("mock out the Object method")
//			},
//			SignalFunc: func(ch chan<- *dbus.Signal)  {
//				panic("mock out the Signal method")
//			},
//		}
//
//		// use mockeddbusConn in code that requires dbusConn
//		// and then make assertions.
//
//	}
type dbusConnMock struct {
	// AddMatchSignalFunc mocks the AddMatchSignal method.
	AddMatchSignalFunc func(matchOptions ...dbus.MatchOption) error

	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// ObjectFunc mocks the Object method.
	ObjectFunc func(s string, objectPath dbus.ObjectPath) dbusBusObject

	// SignalFunc mocks the Signal method.
	SignalFunc func(ch chan<- *dbus.Signal)

	// calls tracks calls to the methods.
	calls struct {
		// AddMatchSignal holds details about calls to the AddMatchSignal method.
		AddMatchSignal []struct {
			// MatchOptions is the matchOptions argument value.
			MatchOptions []dbus.MatchOption
		}
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Object holds details about calls to the Object method.
		Object []struct {
			// S is the s argument value.
			S string
			// ObjectPath is the objectPath argument value.
			ObjectPath dbus.ObjectPath
		}
		// Signal holds details about calls to the Signal method.
		Signal []struct {
			// Ch is the ch argument value.
			Ch chan<- *dbus.Signal
		}
	}
	lockAddMatchSignal sync.RWMutex
	lockClose          sync.RWMutex
	lockObject         sync.RWMutex
	lockSignal         sync.RWMutex
}

// AddMatchSignal calls AddMatchSignalFunc.
func (mock *dbusConnMock) AddMatchSignal(matchOptions ...dbus.MatchOption) error {
	if mock.AddMatchSignalFunc == nil {
		panic("dbusConnMock.AddMatchSignalFunc: method is nil but dbusConn.AddMatchSignal was just called")
	}
	callInfo := struct {
		MatchOptions []dbus.MatchOption
	}{
		MatchOptions: matchOptions,
	}
	mock.lockAddMatchSignal.Lock()
	mock.calls.AddMatchSignal = append(mock.calls.AddMatchSignal, callInfo)
	mock.lockAddMatchSignal.Unlock()
	return mock.AddMatchSignalFunc(matchOptions...)
}

// AddMatchSignalCalls gets all the calls that were made to AddMatchSignal.
// Check the length with:
//
//	len(mockeddbusConn.AddMatchSignalCalls())
func (mock *dbusConnMock) AddMatchSignalCalls() []struct {
	MatchOptions []dbus.MatchOption
} {
	var calls []struct {
		MatchOptions []dbus.MatchOption
	}
	mock.lockAddMatchSignal.RLock()
	calls = mock.calls.AddMatchSignal
	mock.lockAddMatchSignal.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *dbusConnMock) Close() error {
	if mock.CloseFunc == nil {
		panic("dbusConnMock.CloseFunc: method is nil but dbusConn.Close was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//
//	len(mockeddbusConn.CloseCalls())
func (mock *dbusConnMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// Object calls ObjectFunc.
func (mock *dbusConnMock) Object(s string, objectPath dbus.ObjectPath) dbusBusObject {
	if mock.ObjectFunc == nil {
		panic("dbusConnMock.ObjectFunc: method is nil but dbusConn.Object was just called")
	}
	callInfo := struct {
		S          string
		ObjectPath dbus.ObjectPath
	}{
		S:          s,
		ObjectPath: objectPath,
	}
	mock.lockObject.Lock()
	mock.calls.Object = append(mock.calls.Object, callInfo)
	mock.lockObject.Unlock()
	return mock.ObjectFunc(s, objectPath)
}

// ObjectCalls gets all the calls that were made to Object.
// Check the length with:
//
//	len(mockeddbusConn.ObjectCalls())
func (mock *dbusConnMock) ObjectCalls() []struct {
	S          string
	ObjectPath dbus.ObjectPath
} {
	var calls []struct {
		S          string
		ObjectPath dbus.ObjectPath
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
//
//	len(mockeddbusConn.SignalCalls())
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
