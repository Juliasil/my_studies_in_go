package main

import (
  "bufio"
  "database/sql"
  "fmt"
  "log"
  "os"
  "strings"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql","usuario_exemplo:S3nh@_Ex@mpl3_2024@tcp(localhost:3306)/" ) 
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  fmt.Print("Enter database name: ")
  var dbName string
  fmt.Scanln(&dbName)

  if !checkDatabase(db, dbName) {
    queryDrop := "DROP DATABASE IF EXISTS " + dbName

    db.Exec(queryDrop)

    var query string = "CREATE DATABASE "
    fmt.Println("")
    fmt.Printf("Database %s does not exist. Creating...\"n", dbName)
    fmt.Println("")
    if _, err := db.Exec(query + dbName); err != nil {
      log.Fatal(err)
    }
    fmt.Println("")
    fmt.Println("==================")
    fmt.Println("Database created.")
    fmt.Println("==================")
    fmt.Println("")
  }
    db.Exec("USE " + dbName)
  


  createTable(db)

  for {
    fmt.Println("1. Create table")
    fmt.Println("2. Insert data")
    fmt.Println("3. Read data")
    fmt.Println("4. Update data")
    fmt.Println("5. Delete data")
    fmt.Println("6. Exit")
    fmt.Println("choose an option: ")
    var option int

    fmt.Scanln(&option)

    switch option {
    case 1:
      createTable(db)
    case 2:
      insertData(db)
    case 3:
      readData(db)
    case 4:
      // updateData(db)
    case 5:
      // deletData(db)
    case 6:
      fmt.Println("Existing program")
    default:
      fmt.Println("Invalid option.")
      os.Exit(0)
    }
  }
}



func checkDatabase(db *sql.DB, dbName string) bool {
  var temp string
  query := "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME" + dbName
  err := db.QueryRow(query).Scan(&temp)
  if err != nil {
    return false
  }

  return false
}

func createTable (db *sql.DB) {
  fmt.Printf("Enter table name:")
  var tableName string
  fmt.Scanln(&tableName)

  
  fmt.Println("Enter columns and types (e.g., id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255));")

  reader := bufio.NewReader(os.Stdin)
  columnDetails, _ := reader.ReadString('\n')
  columnDetails = strings.TrimSpace(columnDetails)

  
  query := fmt.Sprintf("CREATE TABLE %s (%s) ", tableName, columnDetails)
  fmt.Println(query)
  if _, err := db.Exec(query); err != nil {
    log.Fatal("err")
  }
  fmt.Println("Table created.")
}

func insertData(db *sql.DB) {
  fmt.Println("Enter table name")
  var tableName string
  fmt.Scanln((&tableName))

fmt.Println("Enter columns (e.g, id, name):")

  reader := bufio.NewReader(os.Stdin)
  columns, _ := reader.ReadString('\n')
  columns = strings.TrimSpace(columns)

  fmt.Println("Enter values (e.g., 1, 'Jhon Doe'):")
  values, _ := reader.ReadString('\n')
  values = strings.TrimSpace(values)

  query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columns, values)

  if _, err := db.Exec(query); err != nil {
    log.Fatal(err)
  }
  
  fmt.Println("Data inserted.")
}

func readData(db *sql.DB) {
  fmt.Println("Enter table name:")
  var tableName string
  fmt.Scanln(&tableName)

  query := fmt.Sprintf("SELECT * FROM %s WHERE id = 1 OR name = 'Jhon Doe'", tableName)
  rows, err  := db.Query(query)

  if err != nil {
    log.Fatal(err)
  }

  defer rows.Close()

  columns, err := rows.Columns()

  if err != nil {
    log.Fatal(err)
  }

  values := make([]sql.RawBytes, len(columns))
  scanArgs := make([]interface{}, len(values))

  for i := range values {
    scanArgs[i] = &values[i]
  }

  for rows.Next() {
    err = rows.Scan(scanArgs...)
    if err != nil {
      log.Fatal(err)
    }

    for i, column := range values {
      fmt.Println()
      fmt.Printf("%s: %s\n", columns[i], string(column))
      fmt.Println()
    }
  }
  
}



