package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (repository Users) FindUsersByNameOrNick(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name like ? or nick like ?",
		nameOrNick,
		nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) FindUserById(userID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		userID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository Users) UpdateUser(userID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, userID); err != nil {
		return err
	}

	return nil
}

func (repository Users) DeleteUser(userID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from users where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID); err != nil {
		return err
	}

	return nil
}

func (repository Users) FindUserByEmail(email string) (models.User, error) {
	line, err := repository.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository Users) FollowUser(userID, userIDToFollow uint64) error {
	statement, err := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, userIDToFollow); err != nil {
		return err
	}

	return nil
}

func (repository Users) UnfollowUser(userID, userIDToUnfollow uint64) error {
	statement, err := repository.db.Prepare(
		"delete from followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, userIDToUnfollow); err != nil {
		return err
	}

	return nil
}