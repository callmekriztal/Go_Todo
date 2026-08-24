package main

import(
	"os"
	"errors"
)

func main(){
	done := Done{}
	
	path,err := todoFilePath()
	if err!= nil {
		panic(err)
	}
	storage := NewStorage[Done](path)

	if err:= storage.Load(&done); err!= nil{
		if !errors.Is(err,os.ErrNotExist){
			panic(err)
		}
	}
	
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&done)

	storage.Save(done)

}