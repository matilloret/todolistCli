package todos

import (
	"errors"
	"time"
)

// la funcionalidad que va a tener es representar a las tareas

type Item struct {
	Titutlo      string
	Done         bool
	Created_at   time.Time
	Completed_at time.Time
}

// El usuario tiene la capacidad de crear varias tareas lo que deberia hacer es crear un arreglo
type Todos []Item

func (t *Todos) Add(titutlo string) {
	todo := Item{
		Titulo:       titulo,
		Done:         false,
		Created_at:   time.Now(),
		Completed_at: time.Time{},
	}

	*t = append(*t, todo)
}

// el usuario pueda ser capaz de marcar como completada una tarea

func (t *Todos) Complete(index int) error {
	// crear una instancia de la estructura todos
	ls := *t
	// manejo de derror por si el usuario ingresa un numero  invalido
	if index <= 0 || index > len(ls) {
		return errors.New("Index invalido")
	}

	// Actualizar la tarea
	ls[index-1].Completed_at = time.Now()
	ls[index-1].Done = true

	return nil
}
