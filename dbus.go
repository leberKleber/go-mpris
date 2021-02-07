package mpris

import "github.com/godbus/dbus/v5"

type dbusWrapper struct {
	conn *dbus.Conn
}

func (w dbusWrapper) Call(dest string, path dbus.ObjectPath, method string, flags dbus.Flags, args ...interface{}) {
	w.conn.Object(dest, path).Call(method, flags, args)
}
