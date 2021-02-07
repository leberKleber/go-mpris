package mpris

import "github.com/godbus/dbus/v5"

type dbusConnWrapper struct {
	conn *dbus.Conn
}

func (w dbusConnWrapper) Object(dest string, path dbus.ObjectPath) dbusBusObject {
	return dbusBusObjectWrapper{
		obj: w.conn.Object(dest, path),
	}
}

type dbusBusObjectWrapper struct {
	obj dbus.BusObject
}

func (w dbusBusObjectWrapper) Call(method string, flags dbus.Flags, args ...interface{}) dbusCall {
	return dbusCallWrapper{
		call: w.obj.Call(method, flags, args),
	}
}

func (w dbusBusObjectWrapper) GetProperty(p string) (dbus.Variant, error) {
	return w.obj.GetProperty(p)
}

type dbusCallWrapper struct {
	call *dbus.Call
}

func (w dbusCallWrapper) Store(retvalues ...interface{}) error {
	return w.call.Store(retvalues)
}
