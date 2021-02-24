package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type flagIds []string

func (ids flagIds) String() string {
	return strings.Join(ids, ",")
}

func (ids *flagIds) Set(id string) error {
	*ids = append(*ids, id)
	return nil
}

type person struct {
	name string
	born time.Time
}

func (p person) String() string {
	return fmt.Sprintf("Name: %s, and I am %s", p.name, p.born)
}

func (p *person) Set(name string) error {
	p.name = name
	p.born = time.Now()
	return nil
}

func main() {
	var ids flagIds
	var p person

	flag.Var(&ids, "id", "Id to append to list")
	flag.Var(&p, "name", "Name of user")
	flag.Parse()
	fmt.Println(ids)
	fmt.Println(p)
}
