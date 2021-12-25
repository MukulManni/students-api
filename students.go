package main

type Student struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Department string `json:"department"`
}

var studentsList = []Student{
	Student{"20BCS9559", "Mukul", 19, "Cse"},
	Student{"20BCS9558", "Anonymous", 19, "Cse"},
}