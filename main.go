package main

func main(){
	done := Done{}
	
	path,err := todoFilePath()
	if err!= nil {
		panic(err)
	}
	storage := NewStorage[Done](path)
	storage.Load(&done)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&done)

	storage.Save(done)

}