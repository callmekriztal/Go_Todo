package main

func main(){
	done := Done{}
	
	storage := NewStorage[Done]("todos.json")
	storage.Load(&done)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&done)

	storage.Save(done)

}