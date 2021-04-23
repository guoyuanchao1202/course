package config

type TotalConf struct {
	MySQL MySQLConfig `json:"mysql" yaml:"mysql"`
	DataPath PathConfig `json:"path" yaml:"datapath"`
}

type MySQLConfig struct {
	DBName     string `json:"DBName" yaml:"DBName"`
	DBHost     string `json:"DBHost" yaml:"DBHost"`
	DBPort     string `json:"DBPort" yaml:"DBPort"`
	DBUser     string `json:"DBUser" yaml:"DBUser"`
	DBPassWord string `json:"DBPassWord" yaml:"DBPassWord"`
}

type PathConfig struct {
	PathName string `json:"PathName" yaml:"PathName"`
}
