package model

import (
    "log"    
	"../helpers"
)

type Task struct {
    Id string
	Name string
	Description string
}

var (
    id string
    name string
    description string
)

func (givenTask *Task) SaveTask() string{
    db := helper.OpenDb()

    message := "Success"

    _, err := db.Exec("insert into tasks(name, description) values(?, ?)", givenTask.Name, givenTask.Description)
    if err != nil{
        message = "Server error, unable to create record"
    }

    helper.CloseDb(db)

    return message
}

func (givenTask *Task) ListTasks() []Task{
    db := helper.OpenDb();

    rows, err := db.Query("select * from tasks")
    if err != nil {
        log.Fatal(err)
    }

    taskList := []Task{}

    for rows.Next() {
        err := rows.Scan(&id, &name, &description)
        if err != nil {
            log.Fatal(err)
        }

        taskList = append(taskList, Task{id, name, description})
    }

    err = rows.Err()
    if err != nil {
        log.Fatal(err)    
    }

    return taskList
}

func (givenTask *Task) UpdateTask() string{
    db := helper.OpenDb()

    message := "Success"

    _, err := db.Exec("update tasks set name=?,description=? where id=?", givenTask.Name, givenTask.Description, givenTask.Id)
    if err != nil{
        message = "Server error, unable to create record"
    }

    helper.CloseDb(db)

    return message
}

func (givenTask *Task) DeleteTask() string{
    db := helper.OpenDb()

    message := "Success"

    _, err := db.Exec("delete from tasks where id=?", givenTask.Id)
    if err != nil{
        message = "Server error, unable to create record"
    }

    helper.CloseDb(db)

    return message
}