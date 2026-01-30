package main

import (
	"time"
	"errors"
	"fmt"
	"strconv"
	"os"
	"github.com/aquasecurity/table"
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
	if index < 0 || index >= len(*d) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (d *Done) delete(index int) error{	
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

	if !isCompleted { 
		completionTime := time.Now()
		(*d)[index].CompletedAt = &completionTime
		(*d)[index].Completed = !isCompleted
	}

	return nil
}

func (d *Done) edit(index int,title string) error{

	if err := d.validateIndex(index);err != nil {
		return nil 
	} 

	(*d)[index].Title = title

	return nil 
}


func (d *Done) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Task No","Title","Completed","Created At","Completed At")

	for index,t := range *d{
		completed := "🙅‍♀️" 
		completedAt := ""
		
		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(
			strconv.Itoa(index),
			t.Title,
			completed,
			t.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}
	table.Render()
}