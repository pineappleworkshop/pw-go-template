package stores

const (
	DB_NAME = "{{<db_name>}}"
)

var (
	MONGOHOSTS_WORKSTATION = []string{"localhost:27017"}
	MONGOHOSTS_CLUSTER     = []string{
		"{{<mongo_rs>}}-mongodb-replicaset-0",
		"{{<mongo_rs>}}-mongodb-replicaset-1",
		"{{<mongo_rs>}}-mongodb-replicaset-2",
	}
)
