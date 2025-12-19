package keyvaluefactory

func newStoreMockImpl() Store {
	return &storeMockImpl{}
}

type storeMockImpl struct{}

func (s storeMockImpl) Open() (er error) {
	return
}
func (s storeMockImpl) Echo() (er error) {
	return
}
func (s storeMockImpl) GetStore() interface{} {
	return nil
}
func (s storeMockImpl) GetDriverName() (driverName string) {
	return DRIVER_MOCK
}
func (s storeMockImpl) Close() {}
