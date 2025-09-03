package main

type WireDb struct {
	Name string
	Port int
}

type WireConfig struct {
	Db      *WireDb
	Address string
}

func CreateName() string {
	return "use name"
}

func CreateWireDb(name string) *WireDb {
	return &WireDb{
		Name: name,
		Port: 3306,
	}
}

func CreateWireConfig(wb *WireDb) *WireConfig {
	return &WireConfig{
		Db:      wb,
		Address: "shenzhen",
	}
}
