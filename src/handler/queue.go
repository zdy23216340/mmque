package handler

import (
	"../models"
	"../state"
	// "encoding/json"
	"fmt"
	"net"
)

func Push(conn net.Conn, iData interface{}) {

	data := iData.(*models.Task)

	queue := state.CreateQueue(data.Channal)
	result := queue.Push(data)

	res := &models.Res{
		Msg:    result,
		Status: "200",
	}

	models.Send(conn, res)
}

func Pull(conn net.Conn, iData interface{}) {

	data := iData.(*models.PullRequest)

	queue := state.Pool[data.Channal]

	res := &models.Res{}

	if queue == nil {

		res.Msg = "No this queue!"
		res.Status = "404"

		models.Send(conn, res)
		return
	}

	task := queue.Pull(data.Version)

	if task == nil {

		res.Msg = "No this queue!"
		res.Status = "404"

		models.Send(conn, res)
		return
	}

	res.Msg = task.Msg
	res.Status = "200"

	models.Send(conn, res)
}

func Ack(conn net.Conn, iData interface{}) {

	data := iData.(*models.PullRequest)

	queue := state.Pool[data.Channal]

	res := &models.Res{}

	if queue == nil {

		res.Msg = "No this queue!"
		res.Status = "404"

		models.Send(conn, res)
		return
	}

	done := queue.Ack(data.Version)

	fmt.Println(data)

	if done {
		res.Msg = "Done"
		res.Status = "200"

		models.Send(conn, res)
		return

	} else {
		res.Msg = "Not found this task."
		res.Status = "404"

		models.Send(conn, res)
		return

	}
}