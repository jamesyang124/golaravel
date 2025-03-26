package data

import (
	"time"

	upperDb "github.com/upper/db/v4"
)

type User struct {
	ID        int       `db:"id,omitempty"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Active    int       `db:"user_active"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	/**
	In Go struct tags with database annotations,
	db:"-" is a special directive that tells the ORM (Object-Relational Mapping)
	to ignore this field during database operations.
	*/
	Token Token `db:"-"`
}

func (u *User) Table() string {
	return "users"
}

func (u *User) GetAll(condition upperDb.Cond) ([]*User, error) {
	collection := upper.Collection(u.Table())

	var all []*User

	res := collection.Find(condition)
	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, nil
}

func (u *User) GetEmail(email string) (*User, error) {
	var user User
	collection := upper.Collection(u.Table())

	res := collection.Find(upperDb.Cond{"email": email})
	err := res.One(&user)
	if err != nil {
		return nil, err
	}

	var token Token
	collection = upper.Collection(token.Table())
	res = collection.Find(upperDb.Cond{"user_id": user.ID, "expiry <": time.Now()}).OrderBy("created_at_desc")
	err = res.One(&token)
	if err != nil {
		if err != upperDb.ErrNilRecord && err != upperDb.ErrNoMoreRows {
			return nil, err
		}
	}

	user.Token = token

	return &user, nil
}
