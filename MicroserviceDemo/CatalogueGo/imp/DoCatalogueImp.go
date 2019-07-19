package imp

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DoCatalogueImp struct {
}
type Sock struct {
	ID          string   `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	ImageURL    []string `json:"imageUrl" db:"-"`
	Price       float32  `json:"price" db:"price"`
	Count       int      `json:"count" db:"count"`
	Tags        []string `json:"tag" db:"-"`
}
type Count struct {
	Size int `json:"size"`
}

type Tags struct {
	Tags []string `json:"Tags"`
}

// ErrNotFound is returned when there is no sock for a given ID.
var ErrNotFound = errors.New("not found")

func (imp *DoCatalogueImp) EchoHello(name string, greeting *string) (int32, error) {
	*greeting = "hello " + name
	return 0, nil
}

func (imp *DoCatalogueImp) Catalogue() (string, error) {
	db, dbErr := sqlx.Open("mysql", "root@tcp(localhost:3306)/socksdb")
	if dbErr != nil {
		fmt.Printf("connect mysql fail ! [%s]", dbErr)
	} else {
		fmt.Println("connect to mysql success")
	}
	defer db.Close()

	rows, queryErr := db.Query("SELECT * FROM `sock`")
	if queryErr != nil {
		fmt.Println("select fail [%s]", queryErr)
	} else {
		fmt.Println("select success")
	}
	var socks []Sock
	for rows.Next() {
		var sock_id string
		var name string
		var description string
		var price float32
		var count int
		var image_url_1 string
		var image_url_2 string
		//var tag[] string
		rows.Columns()
		err := rows.Scan(&sock_id, &name, &description, &price, &count, &image_url_1, &image_url_2)
		if err != nil {
			fmt.Printf("get user info error [%s]", err)
		}
		var tags []string
		//tags = selectTags(&db, sock_id)
		tags = selectTags(db, sock_id)
		fmt.Println(tags)
		sock := Sock{sock_id, name, description, []string{image_url_1, image_url_2}, price, count, tags}
		socks = append(socks, sock)
	}

	sockjson, sockjsonerr := json.Marshal(socks)
	if (sockjsonerr != nil) {
		fmt.Sprintln(sockjsonerr)
	}
	fmt.Println(string(sockjson))
	return string(sockjson), nil
}
func selectTags(db *sqlx.DB, sock_id string) []string {
	query := "SELECT `name` FROM tag WHERE tag_id IN (SELECT DISTINCT tag_id FROM `sock_tag` WHERE sock_id =\"" + sock_id + "\")"
	rows, queryErr := db.Query(query)
	if queryErr != nil {
		fmt.Println("select fail [%s]", queryErr)
	} else {
		fmt.Println("select success")
	}
	var tags [] string
	for rows.Next() {
		var name string
		rows.Columns()
		err := rows.Scan(&name)
		if err != nil {
			fmt.Printf("get user info error [%s]", err)
		}
		fmt.Println(name)
		tags = append(tags, name)
	}
	return tags
}

func (imp *DoCatalogueImp) CatalogueId(sock_id string)(string, error) {
	db, dbErr := sqlx.Open("mysql", "root@tcp(localhost:3306)/socksdb")
	if dbErr != nil {
		fmt.Printf("connect mysql fail ! [%s]", dbErr)
	} else {
		fmt.Println("connect to mysql success")
	}
	defer db.Close()

	query := "SELECT * FROM `sock` WHERE sock_id = \"" + sock_id + "\""
	rows, queryErr := db.Query(query)
	if queryErr != nil {
		fmt.Println("select fail [%s]", queryErr)
	} else {
		fmt.Println("select success")
	}
	var sock Sock
	for rows.Next() {
		var sock_id string
		var name string
		var description string
		var price float32
		var count int
		var image_url_1 string
		var image_url_2 string
		//var tag[] string
		rows.Columns()
		err := rows.Scan(&sock_id, &name, &description, &price, &count, &image_url_1, &image_url_2)
		if err != nil {
			fmt.Printf("get user info error [%s]", err)
		}
		var tags []string
		//tags = selectTags(&db, sock_id)
		tags = selectTags(db, sock_id)
		fmt.Println(tags)
		sock = Sock{sock_id, name, description, []string{image_url_1, image_url_2}, price, count, tags}
	}

	sockjson, sockjsonerr := json.Marshal(sock)
	if (sockjsonerr != nil) {
		fmt.Sprintln(sockjsonerr)
	}
	fmt.Println(string(sockjson))
	return string(sockjson), nil
}

func (imp *DoCatalogueImp) Count()(string, error) {
	db, dbErr := sqlx.Open("mysql", "root@tcp(localhost:3306)/socksdb")
	if dbErr != nil {
		fmt.Printf("connect mysql fail ! [%s]", dbErr)
	} else {
		fmt.Println("connect to mysql success")
	}
	defer db.Close()

	query := "SELECT * FROM `sock`"
	rows, queryErr := db.Query(query)
	if queryErr != nil {
		fmt.Println("select fail [%s]", queryErr)
	} else {
		fmt.Println("select success")
	}

	var count int
	count = 0
	for rows.Next() {
		count += 1
	}
	var size Count
	size = Count{count}
	//fmt.Println(string(size))
	sizejson, sizejsonerr := json.Marshal(size)
	if (sizejsonerr != nil) {
		fmt.Sprintln(sizejsonerr)
	}
	fmt.Println(string(sizejson))
	return string(sizejson), nil
}

func (imp *DoCatalogueImp) Tags()(string, error) {
	db, dbErr := sqlx.Open("mysql", "root@tcp(localhost:3306)/socksdb")
	if dbErr != nil {
		fmt.Printf("connect mysql fail ! [%s]", dbErr)
	} else {
		fmt.Println("connect to mysql success")
	}
	defer db.Close()

	query := "SELECT DISTINCT NAME FROM `tag`"
	rows, queryErr := db.Query(query)
	if queryErr != nil {
		fmt.Println("select fail [%s]", queryErr)
	} else {
		fmt.Println("select success")
	}

	var tags []string
	for rows.Next() {
		var tag string
		//var tag[] string
		rows.Columns()
		err := rows.Scan(&tag)
		if err != nil {
			fmt.Printf("get user info error [%s]", err)
		}
		tags = append(tags, tag)
	}
	rettags := Tags{tags}

	//fmt.Println(string(size))
	rettagsjson, rettagsjsonerr := json.Marshal(rettags)
	if (rettagsjsonerr != nil) {
		fmt.Sprintln(rettagsjsonerr)
	}
	fmt.Println(string(rettagsjson))
	return string(rettagsjson), nil
}

