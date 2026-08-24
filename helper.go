package main

import (
	"errors"
	"github.com/aquasecurity/table"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var ErrInvalidIndex = errors.New("invalid index")

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Done []Todo

func (d *Done) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*d = append(*d, todo)
}

func (d *Done) validateIndex(index int) error {
	if index < 0 || index >= len(*d) {
		return ErrInvalidIndex
	}
	return nil
}

func (d *Done) delete(index int) error {
	if err := d.validateIndex(index); err != nil {
		return err
	}

	*d = append((*d)[:index], (*d)[index+1:]...)

	return nil
}

func (d *Done) toggle(index int) error {

	if err := d.validateIndex(index); err != nil {
		return err
	}

	if (*d)[index].Completed {
		(*d)[index].Completed = false
		(*d)[index].CompletedAt = nil
	} else {
		completionTime := time.Now()
		(*d)[index].CompletedAt = &completionTime
		(*d)[index].Completed = true
	}

	return nil
}

func (d *Done) edit(index int, title string) error {

	if err := d.validateIndex(index); err != nil {
		return err
	}

	(*d)[index].Title = title

	return nil
}

func (d *Done) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Task No", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *d {
		completed := "✖️"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(
			strconv.Itoa(index+1),
			t.Title,
			completed,
			t.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}
	table.Render()
}

func todoFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".todocli")

	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}

	return filepath.Join(dir, "todos.json"), nil
}

func (d *Done) clear() {
	*d = Done{}
}
