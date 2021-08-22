package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	UsersStorage interface {
		Create(chat Chat) error
		Exists(chat Chat) (bool, error)
	}

	Users struct {
		*sqlx.DB
	}

	User struct {
		CreatedAt time.Time `sq:"created_at,omitempty"`
		ID        int64     `sq:"chat_id,omitempty"`
	}

	Chat interface {
		Recipient() string
	}
)

func (db *Users) Create(chat Chat) error {
	const q = `INSERT INTO users (id) VALUES (?)`
	_, err := db.Exec(q, chat.Recipient())
	return err
}

func (db *Users) Exists(chat Chat) (has bool, _ error) {
	const q = `SELECT EXISTS(SELECT 1 FROM users WHERE id=?)`
	return has, db.Get(&has, q, chat.Recipient())
}
