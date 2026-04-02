package main

import "fmt"

// =====================================================================
// STRATEGY DESIGN PATTERN
//
// Intuition:
// Imagine you are traveling to the airport. You can choose different "strategies" to get there:
// a Taxi, a Bus, or riding a Bicycle. The destination is the same, but the method (strategy) changes
// depending on your needs. The Strategy pattern lets you define all connecting algorithms
// separately and swap them dynamically at runtime without changing your core "Database" struct.
//
// Pros:
// + Swap algorithms (behaviors) at runtime dynamically.
// + Clean code: eliminates massive if/else or switch statements.
// + Open/Closed Principle: easily add new databases without changing existing context code.
//
// Cons:
// - Can overcomplicate the code if you only have a couple of algorithms that rarely change.
// - Callers must be aware of the different strategies to select the right one.
// =====================================================================

type startConnection interface {
	connect()
}

type Database struct {
	connectthis startConnection
}

func (db *Database) ConnectTODB() {
	db.connectthis.connect()
}

func (db *Database) switchModel(val startConnection) {
	db.connectthis = val
}

type MySql struct {
	dsn string
}

func (c *MySql) connect() {
	fmt.Println("connected to MySql with dsn string ", c.dsn)
}

type PostgreSql struct {
	dsn string
}

func (c *PostgreSql) connect() {
	fmt.Println("connected to postgresql with dsn string", c.dsn)
}

func main() {
	mysqldb := &MySql{
		dsn: "hello",
	}
	db := &Database{
		connectthis: mysqldb,
	}
	db.ConnectTODB()
	postgresqldb := &PostgreSql{
		dsn: "lets go",
	}
	db.switchModel(postgresqldb)
	db.ConnectTODB()
}
