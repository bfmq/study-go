package main

type singleton struct{
	name string
}

type Singleton struct {
	*singleton
}

var single *Singleton


func GetInstance() *Singleton  {
	return single
}

func main() {

}
