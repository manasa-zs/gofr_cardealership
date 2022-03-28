package engine

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"manasa/models"
)

type DB struct {
	DB *sql.DB
}

func New(db *sql.DB) DB {
	return DB{DB: db}
}

const (
	create  = "insert into engine (engine_id,displacement,no_of_cylinders,`range`)values (?,?,?,?)"
	getByID = "select engine_id,displacement,no_of_cylinders,`range` from engine where engine_id=?"
	update  = "update engine set displacement=?,no_of_cylinders=?,`range`=? where engine_id = ?"
	del     = "delete from engine where engine_id = ?"
)

// EngineCreate method used to insert rows into engine table.
func (d DB) EngineCreate(ctx context.Context, engine *models.Engine) (*models.Engine, error) {
	_, err := d.DB.ExecContext(ctx, create, engine.ID.String(), engine.Displacement, engine.NoOfCylinders, engine.Range)
	if err != nil {
		return &models.Engine{}, err
	}

	return engine, nil
}

// EngineGetByID  method used to get values from engine table.
func (d DB) EngineGetByID(ctx context.Context, id uuid.UUID) (*models.Engine, error) {
	var engine models.Engine

	row := d.DB.QueryRowContext(ctx, getByID, id.String())

	err := row.Scan(&engine.ID, &engine.Displacement, &engine.NoOfCylinders, &engine.Range)
	if err != nil {
		return &models.Engine{}, err
	}

	return &engine, nil
}

// EngineUpdate method used to update rows in engine table.
func (d DB) EngineUpdate(ctx context.Context, engine *models.Engine) (*models.Engine, error) {
	_, err := d.DB.ExecContext(ctx, update, engine.Displacement, engine.NoOfCylinders, engine.Range, engine.ID.String())
	if err != nil {
		return &models.Engine{}, err
	}

	return engine, nil
}

// EngineDelete method used to delete rows from engine table.
func (d DB) EngineDelete(ctx context.Context, id uuid.UUID) error {
	_, err := d.DB.ExecContext(ctx, del, id.String())
	if err != nil {
		return err
	}

	return nil
}
