package providers

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type DB struct {
	Config   *Config `inject:"config"`
	Client   *sql.DB
	adapter  string
	host     string
	username string
	password string
	dbname   string
	port     string
}

func NewDB() *DB {
	fmt.Println("NewDB")
	return &DB{}
}

func (this *DB)Init() {
	this.init()
	db, _ := sql.Open(this.adapter, this.getConn())
	this.Client = db
}

func (this *DB) init() {
	adapter, _ := this.Config.GetKey("database", "adapter")
	this.adapter = adapter.Value()

	host, _ := this.Config.GetKey("database", "host")
	this.host = host.Value()

	username, _ := this.Config.GetKey("database", "username")
	this.username = username.Value()

	password, _ := this.Config.GetKey("database", "password")
	this.password = password.Value()

	dbname, _ := this.Config.GetKey("database", "dbname")
	this.dbname = dbname.Value()

	port, _ := this.Config.GetKey("database", "port")
	this.port = port.Value()
}

func (this *DB)getConn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		this.username,
		this.password,
		this.host,
		this.port,
		this.dbname,
	)
}


