package ctx

import (
	"context"
	"gorm.io/gorm"
)

type Context struct {
	DB  *gorm.DB
	Ctx context.Context
	//IsCenter bool
}

func NewContext(ctx context.Context, db *gorm.DB) *Context {
	//var api conf.CenterApi
	//if len(centerApis) > 0 {
	//	api = centerApis[0]
	//}

	return &Context{
		Ctx: ctx,
		DB:  db,
		//CenterApi: api,
		//IsCenter:  isCenter,
	}
}

// set db to Context
func (c *Context) SetDB(db *gorm.DB) {
	c.DB = db
}

// get context from Context
func (c *Context) GetContext() context.Context {
	return c.Ctx
}

// get db from Context
func (c *Context) GetDB() *gorm.DB {
	return c.DB
}
