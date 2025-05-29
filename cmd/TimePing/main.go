package main

import "timeping/internal/engine"
	
func main() {
	if err:=engine.Initial_engine();err!=nil{
		return
	}
	engine.Run();
}