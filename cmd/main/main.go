package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/v45h15h7/REWARD_API/pkg/controllers"
)

func main() {

	router := gin.Default()
	router.GET("/UserDetails/", controllers.GetAllUsers)                         // Get Details of all Users	(I have developed it for testing)
	router.GET("/UserDetails/:UserId", controllers.GetUserById)                  // Get details of user by userid	(I have developed it for testing)
	router.POST("/CreateUser/", controllers.CreateUser)                          // Create a new user by post request by sending data using JSON format (I have developed it for testing)
	router.GET("/UserStreak/:UserId", controllers.GetUserStreak)                 // Get user streak by sending userid
	router.POST("/UserStreak/", controllers.PostUserStreak)                      // We are sending user data in the json format, userid and userstreak
	router.GET("/UserLevel/:UserId", controllers.GetUserLevel)                   // get userlevel by sending userid
	router.POST("/UserLevel/", controllers.PostUserLevel)                        // We are sending user data in the json format, userid and userlevel
	router.POST("/CreateTask/", controllers.CreateTask)                          // Create a new task (Admin) by sending data using JSON
	router.GET("/GetAllTasks/", controllers.GetAllTasks)                         // Get all tasks
	router.GET("GetAllUserTaskStatus/:UserId", controllers.GetAllUserTaskStatus) // Get all User Tasks by User Id
	router.POST("/UpdateUserTaskStatus/", controllers.UpdateUserTaskStatus)      // Posting the User Id, Task Id and status in POST Request
	router.Run("localhost:9090")                                                 // Running in localhost:9090
}

/*

FORMATS FOR POST REQUESTS

How to send Post Request for /CreateUser/:

	1)	Make a POST REQUEST WITH URL :/CreateUser/
	2)	In the POST REQUEST include the following format

	{
		"id":"YOUR ID HERE",
		"level":"YOUR LEVEL HERE",
		"streak":"YOUR STREAK HERE",
		"tokens":"YOUR TOKENS HERE"
	}
	3) Send the request


How to send Post Request for /UserStreak/:

	1)	Make a POST REQUEST WITH URL :/UserStreak/
	2)	In the POST REQUEST include the following format

	{
		"id":"YOUR ID HERE",			//These both fields are Mandatory
		"streak":"YOUR STREAK HERE"
	}
	3) Send the request


How to send Post Request for /UserLevel/:

	1)	Make a POST REQUEST WITH URL :/UserLevel/
	2)	In the POST REQUEST include the following format

	{
		"id":"YOUR ID HERE",			//These both fields are Mandatory
		"level":"YOUR LEVEL HERE"
	}
	3) Send the request


How to send Post Request for /CreateTask/:

	1)	Make a POST REQUEST WITH URL :/CreateTask/
	2)	In the POST REQUEST include the following format

	{
		"taskid":"YOUR TASKID HERE",
		"taskname":"YOUR TASKNAME HERE",
		"actioncount":"YOUR COUNT HERE",
		"duration":"W/D FOR WEEKLY OR DAILY",
		"reward":"YOUR REWARD HERE"
	}
	3) Send the request


How to send Post Request for /UpdateUserTaskStatus/:

	1)	Make a POST REQUEST WITH URL :/UpdateUserTaskStatus/
	2)	In the POST REQUEST include the following format

	{
		"id":"YOUR ID HERE"
		"taskid":"YOUR TASKID HERE",
		"completedstatus":"YOUR COMPLETED STATUS HERE"
	}
	3) Send the request

*/

/*
FOR DATABASE I HAVE USED FOLLOWING TABLES IN MYSQL

CREATE TABLE WORK.USERS(
	UserId INT PRIMARY KEY,
    UserLevel INT DEFAULT 0,
    Streak INT,
    TOKENS INT
);

CREATE TABLE WORK.ADMINTASKS(
	TaskId INT PRIMARY KEY,
	TaskName VARCHAR(100),
    ActionCount INT,
    Duration VARCHAR(7),
    Reward INT
);

CREATE TABLE WORK.USERTASKS(
	UserId INT,
	TaskId INT,
	CompletedStatus INT,
    Rewards INT DEFAULT 0,
    CreatedAt time DEFAULT 0,
    UpdatedAt time DEFAULT 0,
    DeletedAt time DEFAULT NULL,
    PRIMARY KEY (UserId, TaskId)
);


*/
