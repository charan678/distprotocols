package neo

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type NeoConnector interface {
	Session()
	CloseSession()
	Close()
}

type neoConnector struct {
	driver  neo4j.Driver
	session neo4j.Session
}

func NewConnector(uri string, username string, password string) NeoConnector {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil
	}
	return &neoConnector{driver: driver}
}

func (this *neoConnector) Close() {
	this.driver.Close()
}

func (this *neoConnector) Session() {
	this.session = this.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
}

func (this *neoConnector) CloseSession() {
	this.session.Close()
}
