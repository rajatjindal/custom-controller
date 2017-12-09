package handlers

import (
	"github.com/rajatjindal/custom-controller/config"
)

// Handler is implemented by any handler.
// The Handle method is used to process event
type Handler interface {
	Init(c *config.Config) error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(oldObj, newObj interface{})
}

// Default handler implements Handler interface,
// print each event with JSON format
type Default struct {
}

// Init initializes handler configuration
// Do nothing for default handler
func (d *Default) Init(c *config.Config) error {
	return nil
}

// ObjectCreated Handle Object Creation
func (d *Default) ObjectCreated(obj interface{}) {

}

// ObjectDeleted Handle Object Deletion
func (d *Default) ObjectDeleted(obj interface{}) {

}

// ObjectUpdated Handle Object Updates
func (d *Default) ObjectUpdated(oldObj, newObj interface{}) {

}
