package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// These are some structs for Easily Handling the data from HTTP requests
// This Module also contains various Functions which is helping in the controller Module.

type User struct {
	UserId string `json:"id"`
	Level  string `json:"level"`  // Current User level
	Streak string `json:"streak"` // Days of Streak
	Tokens string `json:"tokens"` // Number of token a user has
}

type AdminTasks struct {
	TaskId      string `json:"taskid"`
	TaskName    string `json:"taskname"`
	ActionCount string `json:"actioncount"`
	Duration    string `json:"duration"` // Duration type of a task, Daily or Weekly
	Reward      string `json:"reward"`   // Rewards associated with a task
}

type UserTasks struct {
	UserId          string    `json:"id"`
	TaskId          string    `json:"taskid"`
	CompletedStatus string    `json:"completedstatus"` // 1 to represent Completed, 0 to represent Not Completed
	Reward          string    // To keep track of the rewards sent to a user on completeting a task
	CreatedAt       time.Time // To keep track of time Created, Updated and Deleted (If)
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

var (
	db  *sql.DB
	err error
)

// Connecting with the DB, in this case MYSQL, verifying the Connection
func init() {
	db, err = sql.Open("mysql", "root:PASSWORD@tcp(localhost:3306)/WORK?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("erro validating sql.Open() arguments")
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping()")
		panic(err.Error())
	}

}

// Function to Get all Users from the database (Developed for testing)

func GetAllUsers() []User {
	var res *sql.Rows
	res, err = db.Query("SELECT * FROM `WORK`.`USERS`;") // Executing the Query and storing the result
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	names := make([]User, 0)
	for res.Next() {
		var name User
		if err = res.Scan(&name.UserId, &name.Level, &name.Streak, &name.Tokens); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}
	return names
}

func GetUserById(UserId string) User {
	var user User
	err = db.QueryRow("SELECT * FROM `WORK`.`USERS` WHERE `USERID`=?;", UserId).Scan(&user.UserId, &user.Level, &user.Streak, &user.Tokens)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	return user
}

func CreateUser(A User) User {
	_, err = db.Exec("INSERT INTO `WORK`.`USERS` (`USERID`, `USERLEVEL`, `STREAK`, `TOKENS`) VALUES (?,?,?,?);", A.UserId, A.Level, A.Streak, A.Tokens)
	if err != nil {
		panic(err.Error())
	}
	return A
}

func GetUserStreak(UserId string) string {
	var Streak string
	err = db.QueryRow("SELECT `STREAK` FROM `WORK`.`USERS` WHERE `USERID`=?;", UserId).Scan(&Streak)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	return Streak
}

func GetUserTokens(UserId string) string {
	var UserTokens string
	err = db.QueryRow("SELECT `TOKENS` FROM `WORK`.`USERS` WHERE `USERID`=?;", UserId).Scan(&UserTokens)
	if err != nil {
		panic(err.Error())
	}
	return UserTokens
}

func addStrings(A string, B string) string { // Function to add two strings as I am using strings for all Columns
	a, err := strconv.ParseInt(A, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	b, err := strconv.ParseInt(B, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	c := a + b
	res := fmt.Sprintf("%d", c)
	return res
}

func UpdateUserRewards(UserId string, Rewards string) {
	OldTokens := GetUserTokens(UserId)
	NewTokens := addStrings(OldTokens, Rewards)
	_, err = db.Exec("UPDATE `WORK`.`USERS` SET `TOKENS` = ? WHERE `USERID`=?;", NewTokens, UserId)
	if err != nil {
		panic(err.Error())
	}
}

func UpdateUserStreak(UserId string, Streak string) {
	_, err = db.Exec("UPDATE `WORK`.`USERS` SET `STREAK` = ? WHERE `USERID`=?;", Streak, UserId)
	if err != nil {
		panic(err.Error())
	}
}

func GetUserLevel(UserId string) string {
	var UserLevel string
	err = db.QueryRow("SELECT `USERLEVEL` FROM `WORK`.`USERS` WHERE `USERID`=?;", UserId).Scan(&UserLevel)
	if err != nil {
		panic(err.Error())
	}
	return UserLevel
}

func UpdateUserLevel(UserId string, Level string) {
	_, err = db.Exec("UPDATE `WORK`.`USERS` SET `USERLEVEL` = ? WHERE `USERID`=?;", Level, UserId)
	if err != nil {
		panic(err.Error())
	}
}

func CreateTask(A AdminTasks) AdminTasks {
	_, err = db.Exec("INSERT INTO `WORK`.`ADMINTASKS` (`TASKID`, `TASKNAME`, `ACTIONCOUNT`, `DURATION`, `REWARD`) VALUES (?,?,?,?,?);", A.TaskId, A.TaskName, A.ActionCount, A.Duration, A.Reward)
	if err != nil {
		panic(err.Error())
	}
	return A
}

func GetAllTasks() []AdminTasks {
	var res *sql.Rows
	res, err = db.Query("SELECT * FROM `WORK`.`ADMINTASKS`;")
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	names := make([]AdminTasks, 0)
	for res.Next() {
		var name AdminTasks
		if err = res.Scan(&name.TaskId, &name.TaskName, &name.ActionCount, &name.Duration, &name.Reward); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}
	return names
}

type UserTaskDetail struct {
	TaskId          string `json:"taskid"`
	TaskName        string `json:"taskname"`
	ActionCount     string `json:"actioncount"`
	Duration        string `json:"duration"`
	Reward          string `json:"reward"`
	CompletedStatus string `json:"completedstatus"`
}

func GetAllUserTaskStatus(UserId string) []UserTaskDetail {
	var res *sql.Rows
	res, err = db.Query("SELECT `B`.TaskId, `B`.`TaskName`, `B`.`ActionCount`, `B`.`Duration`, `B`.`Reward`, `C`.`CompletedStatus` FROM `AdminTasks` `B` JOIN `UserTasks` `C` ON `B`.`TaskId` = `C`.`TaskId` WHERE C.`UserId` = ?;", UserId)
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()

	names := make([]UserTaskDetail, 0)
	for res.Next() {
		var name UserTaskDetail
		if err = res.Scan(&name.TaskId, &name.TaskName, &name.ActionCount, &name.Duration, &name.Reward, &name.CompletedStatus); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}
	return names
}

func GetStatus(UserId string, TaskId string) string {
	var Status string
	err = db.QueryRow("SELECT `COMPLETEDSTATUS` FROM `WORK`.`USERTASKS` WHERE `UserId`=? AND `TASKID`=?;", UserId, TaskId).Scan(&Status)
	if err != nil {
		panic(err.Error())
	}
	return Status
}

func UpdateStatus(UserId string, TaskId string, Rewards string) {
	_, err = db.Exec("UPDATE `WORK`.`USERTASKS` SET `COMPLETEDSTATUS`=1, `UPDATEDAT`=?, `REWARDS`=? WHERE `USERID`=? AND `TASKID`=? ;", time.Now(), Rewards, UserId, TaskId)
	if err != nil {
		panic(err.Error())
	}
}

func GetReward(TaskId string) string {
	var Reward string
	err = db.QueryRow("SELECT `REWARD` FROM `WORK`.`ADMINTASKS` WHERE `TASKID`=?;", TaskId).Scan(&Reward)
	if err != nil {
		panic(err.Error())
	}
	return Reward
}
