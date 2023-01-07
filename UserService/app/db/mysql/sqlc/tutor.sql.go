// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: tutor.sql

package db

import (
	"context"
	"database/sql"
)

const createTutor = `-- name: CreateTutor :execresult
INSERT INTO Tutor (fullname, phone, gender, validate, adminId, datecreated) VALUES(?, ?, ?, false, 0, NOW())
`

type CreateTutorParams struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
}

func (q *Queries) CreateTutor(ctx context.Context, arg CreateTutorParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTutor, arg.Fullname, arg.Phone, arg.Gender)
}

const createTutorPassword = `-- name: CreateTutorPassword :execresult
INSERT INTO Tutor_Password (tutor_id, password) VALUES(?, ?)
`

type CreateTutorPasswordParams struct {
	TutorID  int32  `json:"tutor_id"`
	Password string `json:"password"`
}

func (q *Queries) CreateTutorPassword(ctx context.Context, arg CreateTutorPasswordParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTutorPassword, arg.TutorID, arg.Password)
}

const deleteTutor = `-- name: DeleteTutor :exec
DELETE FROM Tutor WHERE id = ?
`

func (q *Queries) DeleteTutor(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTutor, id)
	return err
}

const getTutor = `-- name: GetTutor :one
SELECT id, fullname, gender, phone, validate, adminid, datecreated, dateupdated FROM Tutor WHERE id = ? AND blocked = 0 LIMIT 1
`

func (q *Queries) GetTutor(ctx context.Context, id int32) (Tutor, error) {
	row := q.db.QueryRowContext(ctx, getTutor, id)
	var i Tutor
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Gender,
		&i.Phone,
		&i.Validate,
		&i.Adminid,
		&i.Datecreated,
		&i.Dateupdated,
	)
	return i, err
}

const getTutorByPhone = `-- name: GetTutorByPhone :one
SELECT id, fullname, gender, phone, validate, adminid, datecreated, dateupdated FROM Tutor WHERE phone = ? AND blocked = 0 LIMIT 1
`

func (q *Queries) GetTutorByPhone(ctx context.Context, phone string) (Tutor, error) {
	row := q.db.QueryRowContext(ctx, getTutorByPhone, phone)
	var i Tutor
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Gender,
		&i.Phone,
		&i.Validate,
		&i.Adminid,
		&i.Datecreated,
		&i.Dateupdated,
	)
	return i, err
}

const getTutorPassword = `-- name: GetTutorPassword :one
SELECT id, tutor_id, password FROM Tutor_Password WHERE tutor_id = ? LIMIT 1
`

func (q *Queries) GetTutorPassword(ctx context.Context, tutorID int32) (TutorPassword, error) {
	row := q.db.QueryRowContext(ctx, getTutorPassword, tutorID)
	var i TutorPassword
	err := row.Scan(&i.ID, &i.TutorID, &i.Password)
	return i, err
}

const updateTutorInfo = `-- name: UpdateTutorInfo :execresult
UPDATE Tutor SET fullname = ?, phone = ?, gender = ? WHERE id = ? AND blocked = 0
`

type UpdateTutorInfoParams struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	ID       int32  `json:"id"`
}

func (q *Queries) UpdateTutorInfo(ctx context.Context, arg UpdateTutorInfoParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTutorInfo,
		arg.Fullname,
		arg.Phone,
		arg.Gender,
		arg.ID,
	)
}

const updateTutorPassword = `-- name: UpdateTutorPassword :exec
UPDATE Tutor_Password SET password = ? WHERE tutor_id = ?
`

type UpdateTutorPasswordParams struct {
	Password string `json:"password"`
	TutorID  int32  `json:"tutor_id"`
}

func (q *Queries) UpdateTutorPassword(ctx context.Context, arg UpdateTutorPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateTutorPassword, arg.Password, arg.TutorID)
	return err
}
