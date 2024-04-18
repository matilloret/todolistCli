package todos

import "time"

// la funcionalidad que va a tener es representar a las tareas

type Item struct {
	Titutlo      string
	Done         bool
	Created_at   time.Time
	Completed_at time.Time
}

// El usuario tiene la capacidad de crear varias tareas lo que deberia hacer es crear un arreglo
type Todos []Item

func (t *Todos) Add() {

}
