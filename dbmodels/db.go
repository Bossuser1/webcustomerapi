package dbmodels

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/lib/pq"
)

var db *sql.DB

const (
    dbhost = "DBHOST"
    dbport = "DBPORT"
    dbuser = "DBUSER"
    dbpass = "DBPASS"
    dbname = "DBNAME"
)



func InitDb() {
    config := dbConfig()
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        config[dbhost], config[dbport],
        config[dbuser], config[dbpass], config[dbname])

    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
    conf := make(map[string]string)
    host, ok := os.LookupEnv(dbhost)
    if !ok {
        panic("DBHOST environment variable required but not set")
    }
    port, ok := os.LookupEnv(dbport)
    if !ok {
        panic("DBPORT environment variable required but not set")
    }
    user, ok := os.LookupEnv(dbuser)
    if !ok {
        panic("DBUSER environment variable required but not set")
    }
    password, ok := os.LookupEnv(dbpass)
    if !ok {
        panic("DBPASS environment variable required but not set")
    }
    name, ok := os.LookupEnv(dbname)
    if !ok {
        panic("DBNAME environment variable required but not set")
    }
    conf[dbhost] = host
    conf[dbport] = port
    conf[dbuser] = user
    conf[dbpass] = password
    conf[dbname] = name
    return conf
}

// Close closes the db connection (called via defer)
func Close() {
	db.Close()
}


// repository contains the details of a repository
type repositorySummary struct {
	ID         int
	Ip       string
	DateConnect      string
}

type Repositories struct {
	Repositories []repositorySummary
}

// repository contains the details of a repository
type CustomersSummary struct {
	ID         int
	CustomerName       string
	Phone      string
	Email      string
}

type Customers struct {
	Customers []CustomersSummary
}


// queryRepos first fetches the repositories data from the db
func QueryRepos(repos *Repositories) error {
	rows, err := db.Query(`SELECT connect_id,ip,date_connect FROM connection`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := repositorySummary{}
		err = rows.Scan(
			&repo.ID,
			&repo.Ip,
			&repo.DateConnect,
		)
		if err != nil {
			return err
		}
		repos.Repositories = append(repos.Repositories, repo)
	}
	err = rows.Err()
	if err != nil {
		return err
	}


	return nil
}

// queryRepos first fetches the repositories data from the db
func queryRepos2(repos *Customers) error {
	//rows, err := db.Query(`SELECT connect_id,ip,date_connect FROM connection`)
	rows, err := db.Query(`select cu.customer_id, cu.customer_name,phone,email from customers as cu,contacts as cs where  cu.customer_id=cs.customer_id`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		repo := CustomersSummary{}
		err = rows.Scan(&repo.ID,&repo.CustomerName,&repo.Phone,&repo.Email)
		if err != nil {
			return err
		}
		repos.Customers = append(repos.Customers, repo)
	}
	err = rows.Err()
	if err != nil {
		return err
	}


	return nil
}
