package main

import (
	"time"
	"errors"
	"fmt"
)

type Todo struct {
	Title string 
	Completed bool 
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Done []Todo

func (d *Done) add(title string){
	todo := Todo{
		Title:			title,
		Completed: 		false,
		CreatedAt: 		time.Now(),
		CompletedAt:	nil ,
	}
	*d = append(*d, todo)
}

func (d *Done) validateIndex(index int) error{
	if index < 0 || index > len(*d) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (d *Done) Delete(index int) error{	
	if err := d.validateIndex(index) ; err != nil {
		return err
	}

	*d= append((*d)[:index],(*d)[index+1:]...)

	return nil
}

func (d *Done) toggle(index int) error{

	if err := d.validateIndex(index); err!= nil {
		return err
	}

	isCompleted := (*d)[index].Completed

	if !isCompleted{
		completionTime := time.Now()
		(*d)[index].CompletedAt = &completionTime
	}

	(*d)[index].Completed = !isCompleted

	return nil
}