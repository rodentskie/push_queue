package mongo

type DatabaseInfo struct {
	Uri      string
	Database string
}

func InitDbConnection(uri string, db string) DatabaseInfo {
	return DatabaseInfo{
		Uri:      uri,
		Database: db,
	}
}
