// var client *http.Client
// 	if tracing {
// 		client, err := createClientWithTracing()
// 		if err != nil {
// 			return err

// 		}
// 		log.Println(client)
// 	} else {
// 		client, err := createDefaultClient()
// 		if err != nil {
// 			return err
// 		}
// 		log.Println(client)
// 	} // Use client

package main

import "database/sql"

func main() {

}

// var db *sql.DB

// func init() {
// 	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
// 	d, err := sql.Open("mysql", dataSourceName)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	err = d.Ping()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	db = d
// }

func createClient(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
