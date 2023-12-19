package postgres

import (
	"context"
	"fmt"

	"github.com/Samandarxon/examen_3-month/clinics/config"
	"github.com/Samandarxon/examen_3-month/clinics/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db           *pgxpool.Pool
	branch       storage.BranchRepoI
	client       storage.ClientRepoI
	coming       storage.ComingTableRepoI
	picking      storage.PickingSheetRepoI
	product      storage.ProductRepoI
	remainder    storage.RemainderRepoI
	sale_product storage.SaleProductRepoI
	sale         storage.SaleRepoI
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	),
	)
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pgxpool,
	}, nil
}

func (s *Store) Branch() storage.BranchRepoI {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}
	return s.branch
}

func (s *Store) Client() storage.ClientRepoI {
	if s.client == nil {
		s.client = NewClientRepo(s.db)
	}
	return s.client
}

func (s *Store) ComingTable() storage.ComingTableRepoI {
	if s.coming == nil {
		s.coming = NewComingTableRepo(s.db)
	}
	return s.coming
}

func (s *Store) PickingSheet() storage.PickingSheetRepoI {
	if s.picking == nil {
		s.picking = NewPickingSheetRepo(s.db)
	}
	return s.picking
}

func (s *Store) Product() storage.ProductRepoI {
	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}
	return s.product
}

func (s *Store) Remainder() storage.RemainderRepoI {
	if s.remainder == nil {
		s.remainder = NewRemainderRepo(s.db)
	}
	return s.remainder
}

func (s *Store) SaleProduct() storage.SaleProductRepoI {
	if s.sale_product == nil {
		s.sale_product = NewSaleProductRepo(s.db)
	}
	return s.sale_product
}

func (s *Store) Sale() storage.SaleRepoI {
	if s.sale == nil {
		s.sale = NewSaleRepo(s.db)
	}
	return s.sale
}
