package database

func Run() error {
	//数据库结构同步
	SyncDB()

	//数据库基础数据构建
	err := BuildDBData()
	if err != nil {
		panic(err)
	}

	//elasticsearch mappings 同步
	err = EsIndexSync()
	if err != nil {
		panic(err)
	}
	return nil
}
