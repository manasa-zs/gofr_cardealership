package engine

import (
	"cardealership/db"
	"cardealership/models"
	_ "database/sql"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
)

type engine struct {
	db []models.Engine
}

func New() *engine {
	return &engine{db.Engine}
}

const (
	getByID = "SELECT engine_id,displacement,no_of_cylinders,`range` FROM engine where engine_id = ?"
	post    = "INSERT INTO engine(engine_id,displacement,no_of_cylinders,`range`) VALUES(?,?,?,?)"
	put     = "UPDATE engine SET displacement=?,no_of_cylinders=?,`range`=? where engine_id=?"
	del     = "delete from engine where engine_id=?"
)

// EngineCreate is used to insert a row into the engine table
func (e *engine) EngineCreate(ctx *gofr.Context, engine *models.Engine) (*models.Engine, error) {
	_, err := ctx.DB().ExecContext(ctx, post, engine.ID, engine.Displacement, engine.NoOfCylinders, engine.Range)
	if err != nil {
		return &models.Engine{}, errors.DB{Err: err}
	}

	return engine, nil
}

// EngineGet method takes id as input and returns a row with the given id
func (e *engine) EngineGet(ctx *gofr.Context, id uuid.UUID) (*models.Engine, error) {
	var engine = models.Engine{}

	rows := ctx.DB().QueryRowContext(ctx, getByID, id.String())
	err := rows.Scan(&engine.ID, &engine.Displacement, &engine.NoOfCylinders, &engine.Range)
	if err != nil {
		return &models.Engine{}, errors.DB{Err: err}
	}

	return &engine, nil
}

// EngineUpdate method is used to update/modify a particular row
func (e *engine) EngineUpdate(ctx *gofr.Context, engine *models.Engine) (*models.Engine, error) {
	_, err := ctx.DB().Exec(put, engine.Displacement, engine.NoOfCylinders, engine.Range, engine.ID)
	if err != nil {
		return &models.Engine{}, errors.DB{Err: err}
	}

	return engine, nil
}

// EngineDelete is used to delete a row based on given id.It takes id as input.
func (e *engine) EngineDelete(ctx *gofr.Context, id uuid.UUID) error {
	_, err := ctx.DB().ExecContext(ctx, del, id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
