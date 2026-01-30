package main

import (
	"fmt"
	"strconv"
	"flag"
	"os"
	"strings"
)

type CmdFlags struct {
	Add string 
	Del int
	Edit string 
	Toggle int
	List bool 
}


func NewCmdFlags() *CmdFlags{
	cf := CmdFlags{}

	flag.StringVar(&cf.Add,"add","","Add a new Task")
	flag.StringVar(&cf.Edit,"edit","","edit a todo title by index and specify a new title. id:new_title")
	flag.IntVar(&cf.Del,"del",-1,"delete a todo by index")
	flag.IntVar(&cf.Toggle,"toggle",-1,"mark as completed")
	flag.BoolVar(&cf.List,"list",false,"list all the Tasks")

	flag.Parse()

	return &cf
}


func (cf *CmdFlags) Execute(done *Done){
	switch {
	case cf.List:
		done.print()
	case cf.Add !="":
		done.add(cf.Add)
	case cf.Edit !="":
		parts := strings.SplitN(cf.Edit,":",2)
		if len(parts) != 2 {
			fmt.Println("Error,invalid format,use id:new_title")
			os.Exit(1)
		}
		
		index ,err := strconv.Atoi(parts[0])
		if err != nil || index<=0 {
			fmt.Println("Error: invalid index")
			os.Exit(1)
		}
		index--
		done.edit(index,parts[1])

	case cf.Toggle != -1:
		done.toggle(cf.Toggle-1)
	case cf.Del != -1:
		done.delete(cf.Del-1)

	default:
		fmt.Println("Invalid command")
	}
}
