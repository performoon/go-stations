package service

import (
	"context"
	"database/sql"
	"fmt"
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

	todo.ID = int(insertID)

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

	return nil, nil
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

	error := &model.ErrNotFound{}
	isRow, err := result.RowsAffected()
	if isRow == 0 {
		return nil, error
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

	todo.ID = int(insertID)

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

	return nil, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
