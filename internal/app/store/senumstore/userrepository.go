package senumstore

import (
    "database/sql"

    "github.com/Gorynychdo/twilisip/internal/app/model"
)

type UserRepository struct {
    store *SenumStore
}

func (r *UserRepository) Create(u *model.User) error {
    _, err := r.store.db.Query(
        `INSERT INTO users (id, token)
        VALUES ($1, $2)`,
        u.ID,
        u.Token,
    )

    return err
}
