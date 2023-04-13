package business

import (
	"github.com/erbaner/be-core/open_im_sdk_callback"
	"github.com/erbaner/be-core/pkg/log"
)

func (w *Business) SetListener(callback open_im_sdk_callback.OnCustomBusinessListener) {
	if callback == nil {
		log.NewError("", "callback is null")
		return
	}
	log.NewDebug("", "callback set success")
	w.listener = callback
}
