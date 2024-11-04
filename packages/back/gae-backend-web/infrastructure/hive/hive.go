package hive

import (
	"context"
	"database/sql"
	"github.com/beltran/gohive"
	"log"
)

type Client interface {
	Query(raw string) (*sql.Rows, error)
	Modify(ctx context.Context, raw string)
}

func NewHiveClient(goHiveUrl string, goHivePort int, goHiveAuth string) (Client, error) {
	prestoConn := prestoConnect()
	hiveConn := goHiveConnect(goHiveUrl, goHivePort, goHiveAuth)
	return &hiveClient{readConn: prestoConn, modifyConn: hiveConn}, nil
}

func goHiveConnect(url string, port int, auth string) *gohive.Connection {
	//TODO 这里可能会报错
	conn, err := gohive.Connect(url, port, auth, nil)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	defer conn.Close()
	return conn
}

func prestoConnect() *sql.DB {
	conn, err := sql.Open("presto", "http://user@presto-server:8080/hive/default")
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return conn
}

type hiveClient struct {
	readConn   *sql.DB
	modifyConn *gohive.Connection
}

func (h *hiveClient) Query(raw string) (*sql.Rows, error) {
	query, err := h.readConn.Query(raw)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return query, nil
}

func (h *hiveClient) Modify(ctx context.Context, raw string) {
	h.modifyConn.Cursor().Exec(ctx, raw)
}
