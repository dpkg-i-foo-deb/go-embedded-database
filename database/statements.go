package database

import "database/sql"

var CreateTestTableStmt *sql.Stmt

func InitStatements() {
	CreateTestTableStmt = PrepareStatement(`CREATE TABLE 
											IF NOT EXISTS cat 
												(id INTEGER PRIMARY KEY AUTOINCREMENT,
												name VARCHAR(80) NOT NULL)`)
}
