package model

import (
	"context"
	"glintfed/ent"
	"glintfed/ent/appregister"
	"glintfed/internal/data/client"
	"time"
)

type AppRegister struct {
	*ent.AppRegisterClient
}

func NewAppRegister(client *client.Database) *AppRegister {
	return &AppRegister{
		AppRegisterClient: client.Ent.AppRegister,
	}
}

// VerifyCodeExists
//
//	SELECT EXISTS(
//	  SELECT 1 FROM app_registers
//	  WHERE email = ? AND verify_code = ? AND created_at > ?
//	)
func (m *AppRegister) VerifyCodeExists(ctx context.Context, email, code string) (bool, error) {
	return m.Query().
		Where(
			appregister.Email(email),
			appregister.VerifyCode(code),
			appregister.CreatedAtGT(time.Now().AddDate(0, 0, -90)),
		).
		Exist(ctx)
}

// DeleteByEmail
//
//	DELETE FROM app_registers WHERE email = ?
func (m *AppRegister) DeleteByEmail(ctx context.Context, email string) error {
	_, err := m.Delete().
		Where(appregister.Email(email)).
		Exec(ctx)
	return err
}
