package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/v45h15h7/REWARD_API/pkg/models"
)

func GetAllUsers(c *gin.Context) { // Getting the web response and sending it to a different function for processing in Models module
	res := models.GetAllUsers() // Same for all other functions, you will find the same function name at both locations (mostly)
	c.IndentedJSON(http.StatusOK, res)
}

func GetUserById(c *gin.Context) {
	UserId := c.Param("UserId")
	res := models.GetUserById(UserId)
	c.IndentedJSON(http.StatusOK, res)
}

func CreateUser(c *gin.Context) {
	CreateUser := &models.User{}

	if err1 := c.BindJSON(&CreateUser); err1 != nil {
		return
	}

	res := models.CreateUser(*CreateUser)
	c.IndentedJSON(http.StatusCreated, res)
}

func GetUserStreak(c *gin.Context) {
	UserId := c.Param("UserId")
	res := models.GetUserStreak(UserId)
	c.IndentedJSON(http.StatusOK, res)
}

func PostUserStreak(c *gin.Context) {
	UserDetails := &models.User{}

	if err1 := c.BindJSON(&UserDetails); err1 != nil {
		return
	}

	CurrentStreak := models.GetUserStreak(UserDetails.UserId)
	// Here I am assuming that the data that is coming is fine, i.e. Streak either only increses by 1 (Streak++) or will be set to Zero (Streak = 0)
	if UserDetails.Streak != CurrentStreak { // If user streak has broken (=> Streak = 0) then Update in the system and grant 10*0 tokens
		models.UpdateUserStreak(UserDetails.UserId, UserDetails.Streak)
		models.UpdateUserRewards(UserDetails.UserId, UserDetails.Streak+"0") // Else Add 10*Streak Tokens into wallet
		res := models.GetUserById(UserDetails.UserId)
		c.IndentedJSON(http.StatusOK, res)
		return
	}
	res := models.GetUserById(UserDetails.UserId)
	c.IndentedJSON(http.StatusOK, res)
}

func GetUserLevel(c *gin.Context) {
	UserId := c.Param("UserId")
	res := models.GetUserLevel(UserId)
	c.IndentedJSON(http.StatusOK, res)
}

func PostUserLevel(c *gin.Context) {
	UserDetails := &models.User{}

	if err1 := c.BindJSON(&UserDetails); err1 != nil {
		return
	}

	CurrentLevel := models.GetUserLevel(UserDetails.UserId)
	// Here I am assuiming that the data is fine, i.e. either the Level stays same or increase by 1, based on this assumption I have implimented this function
	if CurrentLevel != UserDetails.Level { // If current level is not equal to the sent level means user has level up, send rewards
		models.UpdateUserLevel(UserDetails.UserId, UserDetails.Level)
		models.UpdateUserRewards(UserDetails.UserId, "10") // 10 Doss tokens on Leveling Up, I'm adding tokens in this funtion
		res := models.GetUserById(UserDetails.UserId)
		c.IndentedJSON(http.StatusOK, res)
		return
	}
	res := models.GetUserById(UserDetails.UserId)
	c.IndentedJSON(http.StatusOK, res)
}

func GetAllTasks(c *gin.Context) {
	res := models.GetAllTasks()
	c.IndentedJSON(http.StatusOK, res)
}

func CreateTask(c *gin.Context) {
	CreateTask := &models.AdminTasks{}

	if err1 := c.BindJSON(&CreateTask); err1 != nil {
		return
	}

	res := models.CreateTask(*CreateTask)
	c.IndentedJSON(http.StatusCreated, res)
}

func GetAllUserTaskStatus(c *gin.Context) {
	UserId := c.Param("UserId")
	res := models.GetAllUserTaskStatus(UserId)
	c.IndentedJSON(http.StatusAccepted, res)
}

func UpdateUserTaskStatus(c *gin.Context) {
	Details := &models.UserTasks{} // In Details, we will receive userId, TaskId and Completed Status in Post Request

	if err1 := c.BindJSON(&Details); err1 != nil { // Binding the JSON data to Details variable
		return
	}

	if Details.CompletedStatus == "1" { // If a task is completed, Update the status (If not already updated) and the rewards in Users Tab and in log(UserTasks)
		CurrentStatus := models.GetStatus(Details.UserId, Details.TaskId) // Get the status of Task and check if its already completed
		if CurrentStatus == "1" {
			c.IndentedJSON(http.StatusAlreadyReported, "Already Reported") // If a task is already completed then return,
			return
		}
		Reward := models.GetReward(Details.TaskId)                  // Get Rewards for this task, to update in User Tokens and Log
		models.UpdateStatus(Details.UserId, Details.TaskId, Reward) // Updating in the Log
		models.UpdateUserRewards(Details.UserId, Reward)            // Updating in User Wallet
		res := models.GetUserById(Details.UserId)                   // Retuning the User Details for Testing Purpose (Remove it)
		c.IndentedJSON(http.StatusOK, res)
		return
	}
	c.IndentedJSON(http.StatusOK, "NoTaskCompleted")
}
