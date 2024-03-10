package conf

import (
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
	"strconv"
	"strings"
	"sync"
)

var once sync.Once

type SupportDbType string

type MailDb struct {
	Type       SupportDbType
	Connection string
}

const (
	Sqlite SupportDbType = "sqlite"
	PgSql  SupportDbType = "pgsql"
)

type Conf struct {
	AAGateway struct {
		Host string
	}
	Mail struct {
		Host     string
		Tls      bool
		Port     int
		User     string
		Password string
		Db       MailDb
	}
}

var conf *Conf

// Get 读取配置
// 优先使用环境变量，如果为空，则使用对应的appsettings.*.yaml
func Get() *Conf {
	once.Do(func() {
		if conf == nil {
			aaGatewayHost := os.Getenv("aagateway_host")
			mailHost := os.Getenv("mail__host")
			mailTls := os.Getenv("mail__tls")
			mailPortStr := os.Getenv("mail__port")
			var mailPort int64 = 995
			var err error
			if mailPort, err = strconv.ParseInt(mailPortStr, 0, 0); err != nil {
				mailPort = 995
			}
			mailUser := os.Getenv("mail__user")
			mailPassword := os.Getenv("mail__password")

			mailDbType := SupportDbType(os.Getenv("mail__db__type"))
			mailDbConnection := os.Getenv("mail__db__connection")

			filePath := getConfFilePath()
			confFile := getConfiguration(filePath)

			conf = &Conf{
				AAGateway: struct{ Host string }{
					Host: func() string {
						if aaGatewayHost == "" {
							return confFile.AAGateway.Host
						}
						return aaGatewayHost
					}(),
				},
				Mail: struct {
					Host     string
					Tls      bool
					Port     int
					User     string
					Password string
					Db       MailDb
				}{
					Host: func() string {
						if mailHost == "" {
							return confFile.Mail.Host
						}
						return mailHost
					}(),
					Tls: func() bool {
						if mailTls == "" {
							return confFile.Mail.Tls
						}
						return strings.EqualFold("true", mailTls)
					}(),
					Port: func() int {
						if mailPortStr == "" {
							return confFile.Mail.Port
						}
						return int(mailPort)
					}(),
					User: func() string {
						if mailUser == "" {
							return confFile.Mail.User
						}
						return mailUser
					}(),
					Password: func() string {
						if mailPassword == "" {
							return confFile.Mail.Password
						}
						return mailPassword
					}(),
					Db: func() MailDb {
						dbType := mailDbType
						dbConn := mailDbConnection
						if dbType == "" {
							dbType = confFile.Mail.Db.Type
						}
						if dbConn == "" {
							dbConn = confFile.Mail.Db.Connection
						}
						return MailDb{
							Type:       dbType,
							Connection: dbConn,
						}
					}(),
				},
			}
		}
	})
	return conf
}

// getConfiguration 读取配置
func getConfiguration(filePath *string) *Conf {
	if file, err := os.ReadFile(*filePath); err != nil {
		panic("conf lost")
	} else {
		c := Conf{}
		err := yaml.Unmarshal(file, &c)
		if err != nil {
			panic("conf lost")
		}
		return &c
	}
}
