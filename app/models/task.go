package model

import (
    "log"    
	"../helpers"
)

type Task struct {
    Id string
	Name string
}

var (
    id string
    name string
)

func (givenTask *Task) SaveTask() string{
    db := helper.OpenDb()

    message := "Success"

    _, err := db.Exec("insert into tasks(name) values(?)", givenTask.Name)
    if err != nil{
        message = "Server error, unable to create record"
    }

    helper.CloseDb(db)

    return message
}

func (givenTask *Task) ListTasks() []Task{
    db := helper.OpenDb();

    rows, err := db.Query("select * from tasks order by 1 desc")
    if err != nil {
        log.Fatal(err)
    }

    taskList := []Task{}

    for rows.Next() {
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }

        taskList = append(taskList, Task{id, name})
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

    _, err := db.Exec("update tasks set name=? where id=?", givenTask.Name, givenTask.Id)
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