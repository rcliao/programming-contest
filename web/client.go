package web

import (
	"database/sql"

	contest "github.com/rcliao/programming-contest"
)

type Client struct {
	contestService ContestService

	Path string
	db   *DB
}

type DB struct {
	db *sql.DB
}

func NewClient() *Client {
	c := &Client{}
	c.contestService.client = c
	return c
}

func (c *Client) Connect() error {
	defer database.Close()
	database, err = sql.Open("sqlite3", "./db/task.db")
	c.db = &DB{database}
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Disconnect() {
	if c.db != nil && c.db.db != nil {
		c.db.db.Close
	}
}

func (c *Client) ContestService() contest.ContestService {
	return &c.contestService
}
