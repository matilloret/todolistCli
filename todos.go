package todos

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

// la funcionalidad que va a tener es representar a las tareas

type Item struct {
	Titulo       string
	Done         bool
	Created_at   time.Time
	Completed_at time.Time
}

// El usuario tiene la capacidad de crear varias tareas lo que deberia hacer es crear un arreglo
type Todos []Item

func (t *Todos) Add(titulo string) {
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

// tarea: hacer la funcionalidad delete
func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("Index invalido")
	}

	//eliminar
	*t = append(ls[:index-1], ls[index-1:]...)

	return nil
}

// leer el json o cargar
func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	// comprobar si el archivo esta vacio
	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return errors.New("deserializacion invalida")
	}
	return nil

}

func (t *Todos) Store(filename string) error {
	data, err := json.MarshalIndent(t, "", "")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
