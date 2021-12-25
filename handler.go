package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, studentsList)
}

func getStudentById(c *gin.Context) {

	ID := c.Param("id")

	for _, s := range studentsList {
		if s.Id == ID {
			c.IndentedJSON(http.StatusOK, s)
			return
		}
	}

	c.IndentedJSON(
		http.StatusNotFound,
		gin.H{
			"message": "Student not found.",
		},
	)
}

func addStudent(c *gin.Context) {
	var student Student

	if err := c.BindJSON(&student); err == nil {
		studentsList = append(studentsList, student)
		c.IndentedJSON(http.StatusCreated, gin.H{
			"Student": student,
			"Status": "Student added successfully.",
		})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Not inserted."})
		return
	}
}

func removeStudentById(c *gin.Context) {
	ID := c.Param("id")

	for i, s := range studentsList {
		if s.Id == ID {
			studentsList = append(studentsList[:i], studentsList[i+1:]...)
			c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Student removed successfully."})
			return
		}
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Student not found."})
}

func help(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"/students": "Get a list of all the students.",
			"/student/add": "Add a student to the list.",
			"/student/get/:id": "Get an individual student with id.",
			"/student/remove/:id": "Remove a student from the list.",
		},
	)
}