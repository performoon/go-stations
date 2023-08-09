package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {

	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	stmt, err := s.db.PrepareContext(ctx, insert)
	if err != nil {
		fmt.Println("PrepareContext err")
		return nil, err
	}
	//defer stmt.Close()

	result, err := stmt.ExecContext(ctx, subject, description)
	//_, err = stmt.ExecContext(ctx, subject, description)
	if err != nil {
		fmt.Print("ExecContext err : ")
		fmt.Println(err)
		return nil, err
	}
	insertID, err := result.LastInsertId()

	fmt.Print("insert insertID : ")
	fmt.Println(insertID)
	fmt.Print("insertID type : ")
	fmt.Println(reflect.TypeOf(insertID))

	_, err = stmt.Exec("UPDATE todos SET id = ?", insertID)
	if err != nil {
		fmt.Print("ExecID err: ")
		fmt.Println(err)
	}

	stmt, err = s.db.PrepareContext(ctx, confirm)
	if err != nil {
		fmt.Println("PrepareContext err")
		return nil, err
	}
	//defer stmt.Close()
	fmt.Print("select PrepareContext : ")
	fmt.Println(stmt)

	_, err = stmt.ExecContext(ctx, insertID)
	//_, err = stmt.ExecContext(ctx, subject, description)
	if err != nil {
		fmt.Print("ExecContext err : ")
		fmt.Println(err)
		return nil, err
	}
	fmt.Print("select ExecContext : ")
	fmt.Println(result)

	todo := new(model.TODO)

	fmt.Print("insert insertID : ")
	fmt.Println(insertID)
	fmt.Print("insertID type : ")
	fmt.Println(reflect.TypeOf(insertID))

	todo.ID = insertID

	//fmt.Println(todo)

	row := stmt.QueryRowContext(ctx, insertID)

	err = row.Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		fmt.Print("row err : ")
		fmt.Println(err)
	}

	fmt.Print("todo : ")
	fmt.Println(todo)
	// row.Scan(todo)

	// s.db.QueryRowContext(ctx, confirm)

	//fmt.Println(todo)
	fmt.Println("finish")
	return todo, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	fmt.Print("prevID : ")
	fmt.Println(prevID)
	fmt.Print("size : ")
	fmt.Println(size)
	todos := []*model.TODO{}

	// if size < 5 {
	// 	size = 5
	// }

	if prevID == 0 {
		stmt, err := s.db.PrepareContext(ctx, read)
		if err != nil {
			fmt.Println("PrepareContext err")
			return nil, err
		}
		defer stmt.Close()

		// result, err := stmt.ExecContext(ctx, size)
		// //_, err = stmt.ExecContext(ctx, subject, description)
		// if err != nil {
		// 	fmt.Print("ExecContext err : ")
		// 	fmt.Println(err)
		// 	return nil, err
		// }

		rows, err := stmt.QueryContext(ctx, size)
		defer rows.Close()

		count := 0
		for rows.Next() {
			addTodo := &model.TODO{}
			if err := rows.Scan(&addTodo.ID, &addTodo.Subject, &addTodo.Description, &addTodo.CreatedAt, &addTodo.UpdatedAt); err != nil {
				log.Fatalf("getRows rows.Scan error err:%v", err)
			}
			todos = append(todos, addTodo)
			count++
		}
	} else {
		stmt, err := s.db.PrepareContext(ctx, readWithID)
		if err != nil {
			fmt.Println("PrepareContext err")
			return nil, err
		}
		defer stmt.Close()

		row := stmt.QueryRowContext(ctx, prevID, size)

		addTodo := &model.TODO{}
		err = row.Scan(&addTodo.ID, &addTodo.Subject, &addTodo.Description, &addTodo.CreatedAt, &addTodo.UpdatedAt)
		if err != nil {
			fmt.Print("row err : ")
			fmt.Println(err)
		}
		todos = append(todos, addTodo)
		fmt.Print(todos[0])
	}
	fmt.Print(todos)

	return todos, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	stmt, err := s.db.PrepareContext(ctx, update)
	if err != nil {
		fmt.Println("PrepareContext err")
		return nil, err
	}
	//defer stmt.Close()

	result, err := stmt.ExecContext(ctx, subject, description, id)
	//_, err = stmt.ExecContext(ctx, subject, description)
	if err != nil {
		fmt.Print("ExecContext err : ")
		fmt.Println(err)
		return nil, err
	}

	isRow, err := result.RowsAffected()
	if isRow == 0 {
		return nil, &model.ErrNotFound{}
	}
	insertID, err := result.LastInsertId()

	fmt.Print("insert insertID : ")
	fmt.Println(insertID)
	fmt.Print("insertID type : ")
	fmt.Println(reflect.TypeOf(insertID))

	// _, err = stmt.Exec("UPDATE todos SET id = ?", insertID)
	// if err != nil {
	// 	fmt.Print("ExecID err: ")
	// 	fmt.Println(err)
	// }

	stmt, err = s.db.PrepareContext(ctx, confirm)
	if err != nil {
		fmt.Println("PrepareContext err")
		return nil, err
	}
	//defer stmt.Close()
	fmt.Print("select PrepareContext : ")
	fmt.Println(stmt)

	_, err = stmt.ExecContext(ctx, insertID)
	//_, err = stmt.ExecContext(ctx, subject, description)
	if err != nil {
		fmt.Print("ExecContext err : ")
		fmt.Println(err)
		return nil, err
	}
	fmt.Print("select ExecContext : ")
	fmt.Println(result)

	todo := new(model.TODO)

	fmt.Print("insert insertID : ")
	fmt.Println(insertID)
	fmt.Print("insertID type : ")
	fmt.Println(reflect.TypeOf(insertID))

	todo.ID = insertID

	//fmt.Println(todo)

	row := stmt.QueryRowContext(ctx, insertID)

	err = row.Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		fmt.Print("row err : ")
		fmt.Println(err)
	}

	fmt.Print("todo : ")
	fmt.Println(todo)
	// row.Scan(todo)

	// s.db.QueryRowContext(ctx, confirm)

	//fmt.Println(todo)
	fmt.Println("finish")

	return todo, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
