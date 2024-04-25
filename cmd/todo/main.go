package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

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
	//complete := flag.Int("complete", 0, "mark a todo as completed")

	//del := flag.Int("delete", 0, "delete a todo")

	list := flag.Bool("list", false, "list all todos")

	// esta funcion analiza las banderas declaradas en el paquete flag
	flag.Parse()

	todoss := &todos.Todos{}
	// esto carga el json
	if err := todoss.Load(todoFile); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	// la utilidad a cada bandera
	switch {
	case *add:
		// dentro de esta logica va a estar nuestro agregar tarea
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		// añadiendo la tarea previamente leida por nuestra funcion getInput
		todoss.Add(task)
		// escribimos eb el archivo
		err = todoss.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

	case *list:
		// dentro logica de listar tareas
	}

}

// me permite captar el valor del titulo al agregar una tarea
// nosotros no sabemos si --> Tarea / Tarea de mati / Tarea de mati de matematica
func getInput(r io.Reader, args ...string) (string, error) {

	// explicar clase que viene
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)

	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) > 0 {
		return "", errors.New("todo vacio no es premitido")
	}

	return scanner.Text(), nil
}
