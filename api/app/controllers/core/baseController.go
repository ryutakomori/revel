package core

import (
	log "api/app/util/log"
	rand "api/app/util/rand"

	"github.com/revel/revel"
)

type BaseController struct {
	*revel.Controller
}

func (b *BaseController) before() revel.Result {
	log.Intialize()
	rand.Intialize()

	return nil
}

func (b *BaseController) after() revel.Result {
	return nil
}

func init() {
	revel.InterceptMethod((*BaseController).before, revel.BEFORE)
	revel.InterceptMethod((*BaseController).after, revel.AFTER)

}
