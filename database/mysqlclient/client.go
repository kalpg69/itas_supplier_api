package mysqlclient

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kalpg69/itas_supplier_api/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	queryCreate = "INSERT INTO suppliers (supplier_code, supplier_name, email) VALUES (?, ?, ?);"
)

func NewMySQLClient(dbUser string, dbPassword string, dbHost string, dbSchema string) domain.ISupplier {

	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		dbUser,
		dbPassword,
		dbHost,
		dbSchema,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		//return nil, errors.New("failed to open mysql connection")
		panic((err))
	}

	if err = db.Ping(); err != nil {
		//return nil, errors.New("failed to ping database")
		panic(err)
	}

	return &mysqlclient{db: db}

}

type mysqlclient struct {
	db *sql.DB
}

func (m *mysqlclient) Create(ctx context.Context, s *domain.Supplier) (*int64, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Error(codes.Canceled, "timeout hit")
	}

	stmt, err := m.db.PrepareContext(ctx, queryCreate)
	if err != nil {
		return nil, status.Error(codes.Internal, "prepare statement failed")
	}

	res, err := stmt.ExecContext(ctx, s.SupplierCode, s.SupplierName, s.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create supplier")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve created supplier id")
	}

	return &id, nil
}
