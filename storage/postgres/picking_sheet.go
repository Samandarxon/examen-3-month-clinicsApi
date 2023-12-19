package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Samandarxon/examen_3-month/clinics/models"
	"github.com/Samandarxon/examen_3-month/clinics/pkg/helpers"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PickingSheetRepo struct {
	db *pgxpool.Pool
}

func NewPickingSheetRepo(db *pgxpool.Pool) *PickingSheetRepo {
	return &PickingSheetRepo{
		db: db,
	}
}

func (r *PickingSheetRepo) Create(ctx context.Context, req models.CreatePickingSheet) (*models.PickingSheet, error) {
	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@", req)
	var (
		incrementIdNext, _ = helpers.NewIncrementId(r.db, "picking_sheet", "CD", 7)
		picking_sheetId    = uuid.New().String()
		query              = `
			INSERT INTO "picking_sheet"(
				"id",
				"increment_id",
				"product_id",
				"coming_id",
				"price",
				"quantity",
				"total",
				"updated_at"
			) VALUES ($1,$2,$3,$4,$5,$6,$7,NOW())`
	)
	_, err := r.db.Exec(ctx, query,
		picking_sheetId,
		incrementIdNext(),
		req.ProductID,
		req.ComingTableID,
		req.Price,
		req.Quantity,
		req.Total,
	)
	fmt.Println(query)
	// defer r.db.Close()
	if err != nil {
		return nil, err
	}

	// fmt.Println("CREATED")
	return r.GetById(ctx, models.PickingSheetPrimaryKey{Id: picking_sheetId})
}

func (c *PickingSheetRepo) GetById(ctx context.Context, req models.PickingSheetPrimaryKey) (*models.PickingSheet, error) {

	var (
		picking_sheet = models.PickingSheet{}
		query         = `
		SELECT 
				"id",
				"increment_id",
				"product_id",
				"coming_id",
				"price",
				"quantity",
				"total",
				"created_at",
				"updated_at"
		FROM "picking_sheet" WHERE id=$1`
	)

	var (
		Id            sql.NullString
		IncrementID   sql.NullString
		ProductID     sql.NullString
		ComingTableID sql.NullString
		Price         sql.NullFloat64
		Quantity      sql.NullInt64
		Total         sql.NullFloat64
		CreatedAt     sql.NullString
		UpdatedAt     sql.NullString
	)
	// fmt.Println(query)
	resp := c.db.QueryRow(ctx, query, req.Id)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&IncrementID,
		&ProductID,
		&ComingTableID,
		&Price,
		&Quantity,
		&Total,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	picking_sheet = models.PickingSheet{
		Id:            Id.String,
		IncrementID:   IncrementID.String,
		ProductID:     ProductID.String,
		ComingTableID: ComingTableID.String,
		Price:         Price.Float64,
		Quantity:      Quantity.Int64,
		Total:         Total.Float64,
		CreatedAt:     CreatedAt.String,
		UpdatedAt:     UpdatedAt.String,
	}
	return &picking_sheet, nil
}

func (r *PickingSheetRepo) GetList(ctx context.Context, req models.GetListPickingSheetRequest) (*models.GetListPickingSheetResponse, error) {
	var (
		respons  models.GetListPickingSheetResponse
		where    = " WHERE TRUE"
		offset   = " OFFSET 0"
		limit    = " LIMIT 10"
		sort     = " ORDER BY created_at DESC"
		querySql string
		query    = `
						SELECT 
							COUNT(*) OVER(),
							"id",
							"increment_id",
							"product_id",
							"coming_id",
							"price",
							"quantity",
							"total",
							"created_at",
							"updated_at"
						FROM "picking_sheet"`
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if len(req.Search) > 0 {
		where += " AND title ILIKE" + " '%" + req.Search + "%'"
	}

	if len(req.Query) > 0 {
		querySql = fmt.Sprintf(" AND %s", req.Query)
	}

	query += where + querySql + sort + offset + limit
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			Id            sql.NullString
			IncrementID   sql.NullString
			ProductID     sql.NullString
			ComingTableID sql.NullString
			Price         sql.NullFloat64
			Quantity      sql.NullInt64
			Total         sql.NullFloat64
			CreatedAt     sql.NullString
			UpdatedAt     sql.NullString
		)

		err := rows.Scan(
			&respons.Count,
			&Id,
			&IncrementID,
			&ProductID,
			&ComingTableID,
			&Price,
			&Quantity,
			&Total,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		respons.PickingSheets = append(respons.PickingSheets, &models.PickingSheet{
			Id:            Id.String,
			IncrementID:   IncrementID.String,
			ProductID:     ProductID.String,
			ComingTableID: ComingTableID.String,
			Price:         Price.Float64,
			Quantity:      Quantity.Int64,
			Total:         Total.Float64,
			CreatedAt:     CreatedAt.String,
			UpdatedAt:     UpdatedAt.String,
		})
	}

	return &respons, nil

}

func (r *PickingSheetRepo) Update(ctx context.Context, req models.UpdatePickingSheet) (*models.PickingSheet, error) {
	query := `
						UPDATE "picking_sheet" SET  
							"product_id" = $2,
							"coming_id" = $3,
							"price" = $4,
							"quantity" = $5,
							"total" = $6,
							"updated_at" = NOW()
						WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query,
		req.Id,
		req.ProductID,
		req.ComingTableID,
		req.Price,
		req.Quantity,
		req.Total,
	)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, models.PickingSheetPrimaryKey{Id: req.Id})

}

func (r *PickingSheetRepo) Delete(ctx context.Context, req models.PickingSheetPrimaryKey) error {
	query := `DELETE FROM "picking_sheet" WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil

}
