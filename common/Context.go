package common

import (
	"gopkg.in/mgo.v2"
)

// Context used for maintaining HTTP Request Context
type Context struct {
	MongoSession *mgo.Session
	User         string
}

// Close mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

// DbCollection returns mgo.collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(mgoConfig.Database).C(name)
}

// NewContext creates a new Context object for each HTTP request
func NewContext() *Context {
	session := GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}
