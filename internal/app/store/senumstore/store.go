package senumstore

import (
    "database/sql"
    _ "github.com/lib/pq"
)

type SenumStore struct {
    db             *sql.DB
    userRepository *UserRepository
}

func New(db *sql.DB) *SenumStore {
    return &SenumStore{
        db: db,
    }
}

func (s *SenumStore) User() *UserRepository {
    if s.userRepository == nil {
        s.userRepository = &UserRepository{
            store: s,
        }
    }

    return s.userRepository
}
