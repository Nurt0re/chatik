package repository

import (
	"fmt"

	"github.com/Nurt0re/chatik"

)


func (r *AuthPostgres) UpdateUser(id int, input chatik.User) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, username=$2 WHERE id=$3", usersTable)
	_, err := r.db.Exec(query, input.Name, input.Username, id)
	return err
}

func (r *AuthPostgres) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	_, err := r.db.Exec(query, id)
	return err
}
