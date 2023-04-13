package wasm_wrapper

import (
	"syscall/js"

	"github.com/erbaner/be-core/open_im_sdk"
	"github.com/erbaner/be-core/pkg/utils"
	"github.com/erbaner/be-core/wasm/event_listener"
)

// ------------------------------------group---------------------------
type WrapperUser struct {
	*WrapperCommon
}

func NewWrapperUser(wrapperCommon *WrapperCommon) *WrapperUser {
	return &WrapperUser{WrapperCommon: wrapperCommon}
}

func (w *WrapperUser) GetSelfUserInfo(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.GetSelfUserInfo, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperUser) SetSelfInfo(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.SetSelfInfo, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperUser) GetUsersInfo(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.GetUsersInfo, callback, &args).AsyncCallWithCallback()

}
