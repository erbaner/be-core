package wasm_wrapper

import (
	"syscall/js"

	"github.com/erbaner/be-core/open_im_sdk"
	"github.com/erbaner/be-core/pkg/utils"
	"github.com/erbaner/be-core/wasm/event_listener"
)

// ------------------------------------group---------------------------
type WrapperSignaling struct {
	*WrapperCommon
}

func NewWrapperSignaling(wrapperCommon *WrapperCommon) *WrapperSignaling {
	return &WrapperSignaling{WrapperCommon: wrapperCommon}
}

func (w *WrapperSignaling) SignalingInviteInGroup(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.SignalingInviteInGroup, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperSignaling) SignalingInvite(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.SignalingInvite, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperSignaling) SignalingAccept(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.SignalingAccept, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperSignaling) SignalingReject(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.SignalingReject, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperSignaling) SignalingCancel(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.SignalingCancel, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperSignaling) SignalingHungUp(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(open_im_sdk.SignalingHungUp, callback, &args).AsyncCallWithCallback()
}
