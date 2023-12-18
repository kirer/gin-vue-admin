package config

import "fmt"

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) EmptyDsn() string {
	if m.Path == "" {
		m.Path = "127.0.0.1"
	}
	if m.Port == "" {
		m.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", m.Username, m.Password, m.Path, m.Port)
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
