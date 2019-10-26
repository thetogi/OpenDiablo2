package UI

import (
	"github.com/essial/OpenDiablo2/Common"
)

// Widget defines an object that is a UI widget
type Widget interface {
	Common.Drawable
	GetEnabled() bool
	SetEnabled(enabled bool)
}