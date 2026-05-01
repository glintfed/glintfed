package model

import (
	"context"
	"errors"
	"time"

	"glintfed/ent"
	"glintfed/ent/user"
	"glintfed/internal/data/client"
	"glintfed/internal/lib/libstr"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	*ent.UserClient
}

func NewUser(client *client.Database) *User {
	return &User{
		UserClient: client.Ent.User,
	}
}

// CreateUserParams holds the fields required to create a new user.
type CreateUserParams struct {
	Name            string
	Username        string
	Email           string
	Password        string // plaintext; hashed before storing
	AppRegisterIP   string
	RegisterSource  string
	EmailVerifiedAt time.Time
}

// Create
//
//	INSERT INTO users (name, username, email, password, app_register_ip, register_source, email_verified_at)
//	VALUES (?, ?, ?, ?, ?, ?, ?)
func (m *User) Create(ctx context.Context, params CreateUserParams) (*ent.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return m.UserClient.Create().
		SetNillableName(libstr.ToPtr(params.Name)).
		SetUsername(params.Username).
		SetEmail(params.Email).
		SetPassword(string(hashed)).
		SetNillableAppRegisterIP(libstr.ToPtr(params.AppRegisterIP)).
		SetNillableRegisterSource(libstr.ToPtr(params.RegisterSource)).
		SetEmailVerifiedAt(params.EmailVerifiedAt).
		Save(ctx)
}

// CountAll
//
//	SELECT count(*) FROM users
func (m *User) CountAll(ctx context.Context) (int, error) {
	return m.Query().Count(ctx)
}

// CountActiveSince
//
//	SELECT count(*)
//	FROM users
//	WHERE updated_at > ? OR last_active_at > ?
func (m *User) CountActiveSince(ctx context.Context, since time.Time) (int, error) {
	return m.Query().
		Where(user.Or(
			user.UpdatedAtGT(since),
			user.LastActiveAtGT(since),
		)).
		Count(ctx)
}

// ErrInvalidCredentials is returned when username/password verification fails.
var ErrInvalidCredentials = errors.New("invalid username or password")

// Authenticate
//
//	SELECT * FROM users WHERE (username = ? OR email = ?) LIMIT 1
func (m *User) Authenticate(ctx context.Context, username, password string) (*ent.User, error) {
	u, err := m.UserClient.Query().
		Where(user.Or(
			user.Username(username),
			user.Email(username),
		)).
		First(ctx)
	if ent.IsNotFound(err) {
		return nil, ErrInvalidCredentials
	}
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}
	return u, nil
}
