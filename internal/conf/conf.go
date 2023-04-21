package conf

type Data struct {
	Database *Data_Database `json:"database,omitempty"`
}
type Data_Database struct {
	Driver string `json:"driver,omitempty"`
	Source string `json:"source,omitempty"`
}
type MySQLConfig struct {
	UserName     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"databasename"`
	Charset      string `json:"charset"`
	ParseTime    string `json:"parsetime"`
	Loc          string `json:"loc"`
}
type KafkaConfig struct {
	KafkaAddress string `json:"kafkaaddress"` // kafka地址
}
type Env struct {
	MySQLConfig *MySQLConfig `json:"mysqlconfig"`
	KafkaConfig *KafkaConfig `json:"kafkaconfig"`
	LogTopic    string       `json:"logtopic"`
	EnableKafka bool         `json:"enablekafka"`
	Data        *Data        `json:"data"`
}
