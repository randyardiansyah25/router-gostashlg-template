package databasefactory

func newMockImpl() Database {
	return &mockImpl{}
}

type mockImpl struct{}

func (m mockImpl) Connect() (er error) {
	return
}
func (m mockImpl) Ping() (er error) {
	return
}
func (m mockImpl) GetConnection() interface{} {
	return nil
}
func (m mockImpl) GetDriverName() string {
	return DRIVER_MOCK
}
func (m mockImpl) SetEnvironmentVariablePrefix(string) {

}
func (m mockImpl) Close() {

}
