// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package service

type Service interface {
	Name() string
}

type ServiceMap map[string]Service

type Generator func() (Service, error)

type Generators map[string]Generator

var services = make(Generators)

func Register(k string, gen Generator) {
	if _, ok := services[k]; ok {
		panic("service has been registered")
	}
	services[k] = gen
}

func GetServices() ServiceMap {
	var (
		serviceMap = make(ServiceMap)
		err error
	)
	for k, gen := range services {
		if serviceMap[k], err = gen(); err != nil {
			panic("service initialize fail")
		}
	}
	return serviceMap
}

func (serviceMap ServiceMap) Get(k string) Service {
	if v, ok := serviceMap[k]; ok {
		return v
	}
	panic("service not found")
}

func (serviceMap ServiceMap) GetOrNot(k string) (Service, bool) {
	v, ok := serviceMap[k]
	return v, ok
}

func (serviceMap ServiceMap) Add(k string, service Service) {
	if _, ok := serviceMap[k]; ok {
		panic("service exist")
	}
	serviceMap[k] = service
}
