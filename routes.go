package main

func intializeRoutes() {
	
	router.GET("/", help)

	router.GET("/students", getStudents)

	router.GET("/student/get/:id", getStudentById)
	router.POST("/student/add", addStudent)

	router.GET("/student/remove/:id", removeStudentById)
}