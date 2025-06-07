package common

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"math/rand/v2"
)

var City_Names = []string{
	"Chicago",
	"Morgantown",
	"Reston",
	"New York",
}

var PeopleNames = []string{
	"John",
	"Adam",
	"Josh",
	"Person",
	"Yeah",
	"I",
	"am",
	"bad",
	"at",
	"names",
}

const DEFAULT_URL = "postgres://myuser:mypassword@localhost:5432/mydb"

const TABLE_ONE_CREATE = "CREATE TABLE IF NOT EXISTS users (" +
	"userId UUID PRIMARY KEY," +
	"name TEXT NOT NULL" +
	")"
const TABLE_TWO_CREATE = "CREATE TABLE IF NOT EXISTS trips (" +
	"tripId UUID PRIMARY KEY," +
	"userId UUID NOT NULL," +
	"distanceTraveled NUMERIC NOT NULL," +
	"cityName TEXT NOT NULL" +
	")"

const INSERT_USER = "INSERT INTO users(userId, name) VALUES($1, $2)"
const INSERT_TABLE = "INSERT INTO trips(tripId, userId, distanceTraveled, cityName) VALUES($1, $2,$3, $4)"

type PostgresHelper struct {
	conn *pgx.Conn
}

func GeneratePostgresConnection(ctx context.Context, url string) (*PostgresHelper, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	return &PostgresHelper{
		conn: conn,
	}, nil
}

func (helper *PostgresHelper) CreateTables(ctx context.Context) error {
	_, err := helper.conn.Exec(ctx, TABLE_ONE_CREATE)
	if err != nil {
		return err
	}
	_, err = helper.conn.Exec(ctx, TABLE_TWO_CREATE)
	return err
}

func (helper *PostgresHelper) GenerateDefaultDemoData(ctx context.Context) error {
	return helper.GenerateDemoData(ctx, 300, []int{1, 5}, []int{10, 100})
}

func (helper *PostgresHelper) GenerateDemoData(ctx context.Context, userAmount int, amountOfTripsRange []int, distanceRange []int) error {
	for i := 0; i < userAmount; i++ {
		fmt.Println("hello world")
		batch := &pgx.Batch{}
		userId := uuid.New()
		nameIndex := rand.IntN(len(PeopleNames))
		name := PeopleNames[nameIndex] + fmt.Sprintf(" the %d", i)
		batch.Queue(INSERT_USER, userId, name)
		for i := 0; i < generateRandomRange(amountOfTripsRange[0], amountOfTripsRange[1]); i++ {
			tripId := uuid.New()
			cityIndex := rand.IntN(len(City_Names))
			batch.Queue(INSERT_TABLE, tripId, userId, generateRandomRange(distanceRange[0], distanceRange[1]), City_Names[cityIndex])
		}
		br := helper.conn.SendBatch(ctx, batch)
		err := br.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func generateRandomRange(min int, max int) int {
	val := min + rand.IntN(max-min+1)
	return val
}
