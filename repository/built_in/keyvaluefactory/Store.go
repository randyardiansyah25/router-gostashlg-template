package keyvaluefactory

var AppStore Store

type Store interface {
	Open() (er error)
	Echo() (er error)
	GetStore() interface{}
	GetDriverName() (driverName string)
	Close()
}
