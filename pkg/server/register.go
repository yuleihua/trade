package server

var registers map[string]Service

func init() {
	registers = make(map[string]Service)
}

func Register(name string, server Service) {
	if _, ok := registers[name]; !ok {
		registers[name] = server
	}
}

func Deregister(name string) {
	if _, ok := registers[name]; ok {
		delete(registers, name)
	}
}

func GetService(name string) Service {
	if _, ok := registers[name]; ok {
		return registers[name]
	}
	return nil
}
