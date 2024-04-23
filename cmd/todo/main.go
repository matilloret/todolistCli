package main

import (
	"flag"
	"fmt"
	"os"

	todos "github.com/matilloret/todolistCli"
)

// inicializar mediante una constante el JSON
const (
	todoFile = ".todos.json"
)

// TODO LIST

// TODOs

// Que va a hacer el usuario?

func main() {
	add := flag.Bool("add", false, "add a new todo")
	// flag -> el usuario indica que tarea va a marcar como completada -> complete
	complete := flag.Int("complete", 0, "mark a todo as completed")

	del := flag.Int("delete", 0, "delete a todo")

	list := flag.Bool("list", false, "list all todos")

	// esta funcion analiza las banderas declaradas en el paquete flag
	flag.Parse()

	todoss := &todos.Todos{}
	// esto carga el json
	if err := todoss.Load(todoFile); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
