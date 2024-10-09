package test_repositories

//func LoadConfig(configPath string) (*config.Config, error) {
//	file, err := os.Open(configPath)
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close()
//
//	var config config.Config
//	if err := json.NewDecoder(file).Decode(&config); err != nil {
//		return nil, err
//	}
//
//	return &config, nil
//}
//
//func SetupTestDatabase() (testcontainers.Container, DBInterface) {
//	var dbContainer testcontainers.Container
//	var db DBInterface
//
//	cfg, err := LoadConfig("config.json")
//	if err != nil {
//		return nil, nil
//	}
//
//	switch cfg.DatabaseType {
//	case "postgres":
//		dbContainer, db = postgres_init.SetupTestDatabase()
//	case "mongo":
//		dbContainer, db = mongo_init.SetupTestDatabase()
//	default:
//		log.Fatal("Unsupported database type")
//	}
//
//	return dbContainer, db
//}
