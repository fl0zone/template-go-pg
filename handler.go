package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func getAllTodoHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := "SELECT id, title, completed FROM todo"
		rows, err := db.Query(q)
		if err != nil {
			defer rows.Close()
			c.JSON(400, err.Error())
			return
		}

		var todos []Todo = []Todo{}
		for rows.Next() {
			var id int
			var title string
			var completed bool
			rows.Scan(&id, &title, &completed)

			todo := Todo{
				Id:        id,
				Title:     title,
				Completed: completed,
			}
			todos = append(todos, todo)
		}

		defer rows.Close()

		c.JSON(200, todos)
	}
}

func createTodoHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.ParseForm()
		title := c.Request.Form.Get("title")
		q := "INSERT INTO todo (title, completed) VALUES ($1, $2)"
		_, err := db.Exec(q, title, false)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(400, err.Error())
			return
		}

		c.JSON(200, gin.H{"message": "Successfully created todo"})
	}
}

func getTodoHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		q := "SELECT id, title, completed FROM todo WHERE id=$1"
		row := db.QueryRow(q, id)
		var todo Todo
		err := row.Scan(&todo.Id, &todo.Title, &todo.Completed)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(400, err.Error())
			return
		}

		c.JSON(200, todo)
	}
}

func toggleTodoHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.Request.ParseForm()

		q := "SELECT completed FROM todo WHERE id=$1"
		row := db.QueryRow(q, id)
		var completed bool
		err := row.Scan(&completed)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(400, err.Error())
			return
		}

		q = "UPDATE todo SET completed=$1 WHERE id=$2"
		_, err = db.Exec(q, !completed, id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(400, err.Error())
			return
		}

		c.JSON(200, gin.H{"message": "Successfully updated todo"})
	}
}

func deleteTodoHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		q := "DELETE FROM todo WHERE id=$1"
		_, err := db.Exec(q, id)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(400, err.Error())
			return
		}

		c.JSON(200, gin.H{"message": "Successfully deleted todo"})
	}
}
