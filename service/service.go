package service

import "fmt"

type Service struct {
	Name string
	Size int32
}

func (s *Service) CreateApp(name string) string {
	fmt.Println("create app (version-001) : ", name)
	return "app:" + name
}
