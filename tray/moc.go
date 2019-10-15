package tray

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
	std_gui "github.com/therecipe/qt/gui"
	std_widgets "github.com/therecipe/qt/widgets"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type SystemTrayIcon_ITF interface {
	std_widgets.QSystemTrayIcon_ITF
	SystemTrayIcon_PTR() *SystemTrayIcon
}

func (ptr *SystemTrayIcon) SystemTrayIcon_PTR() *SystemTrayIcon {
	return ptr
}

func (ptr *SystemTrayIcon) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSystemTrayIcon_PTR().Pointer()
	}
	return nil
}

func (ptr *SystemTrayIcon) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSystemTrayIcon_PTR().SetPointer(p)
	}
}

func PointerFromSystemTrayIcon(ptr SystemTrayIcon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SystemTrayIcon_PTR().Pointer()
	}
	return nil
}

func NewSystemTrayIconFromPointer(ptr unsafe.Pointer) (n *SystemTrayIcon) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(SystemTrayIcon)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *SystemTrayIcon:
			n = deduced

		case *std_widgets.QSystemTrayIcon:
			n = &SystemTrayIcon{QSystemTrayIcon: *deduced}

		default:
			n = new(SystemTrayIcon)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackSystemTrayIcon2a00b6_Constructor
func callbackSystemTrayIcon2a00b6_Constructor(ptr unsafe.Pointer) {
	this := NewSystemTrayIconFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackSystemTrayIcon2a00b6_TriggerSlot
func callbackSystemTrayIcon2a00b6_TriggerSlot(ptr unsafe.Pointer, a C.uintptr_t, b C.char) {
	var aD article
	if aI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(a))); ok {
		qt.UnregisterTemp(unsafe.Pointer(uintptr(a)))
		aD = (*(*article)(aI))
	}
	if signal := qt.GetSignal(ptr, "triggerSlot"); signal != nil {
		(*(*func(article, bool))(signal))(aD, int8(b) != 0)
	}

}

func (ptr *SystemTrayIcon) ConnectTriggerSlot(f func(a article, b bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "triggerSlot"); signal != nil {
			f := func(a article, b bool) {
				(*(*func(article, bool))(signal))(a, b)
				f(a, b)
			}
			qt.ConnectSignal(ptr.Pointer(), "triggerSlot", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "triggerSlot", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SystemTrayIcon) DisconnectTriggerSlot() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "triggerSlot")
	}
}

func (ptr *SystemTrayIcon) TriggerSlot(a article, b bool) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(&a), unsafe.Pointer(&a))
		C.SystemTrayIcon2a00b6_TriggerSlot(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(&a))), C.char(int8(qt.GoBoolToInt(b))))
	}
}

//export callbackSystemTrayIcon2a00b6_ConnectionDead
func callbackSystemTrayIcon2a00b6_ConnectionDead(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectionDead"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *SystemTrayIcon) ConnectConnectionDead(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "connectionDead"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "connectionDead", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectionDead", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SystemTrayIcon) DisconnectConnectionDead() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "connectionDead")
	}
}

func (ptr *SystemTrayIcon) ConnectionDead() {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_ConnectionDead(ptr.Pointer())
	}
}

func SystemTrayIcon_QRegisterMetaType() int {
	return int(int32(C.SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaType()))
}

func (ptr *SystemTrayIcon) QRegisterMetaType() int {
	return int(int32(C.SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaType()))
}

func SystemTrayIcon_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaType2(typeNameC)))
}

func (ptr *SystemTrayIcon) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaType2(typeNameC)))
}

func SystemTrayIcon_QmlRegisterType() int {
	return int(int32(C.SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QmlRegisterType()))
}

func (ptr *SystemTrayIcon) QmlRegisterType() int {
	return int(int32(C.SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QmlRegisterType()))
}

func SystemTrayIcon_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *SystemTrayIcon) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *SystemTrayIcon) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SystemTrayIcon2a00b6___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SystemTrayIcon) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SystemTrayIcon) __children_newList() unsafe.Pointer {
	return C.SystemTrayIcon2a00b6___children_newList(ptr.Pointer())
}

func (ptr *SystemTrayIcon) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.SystemTrayIcon2a00b6___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *SystemTrayIcon) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *SystemTrayIcon) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.SystemTrayIcon2a00b6___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *SystemTrayIcon) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SystemTrayIcon2a00b6___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SystemTrayIcon) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SystemTrayIcon) __findChildren_newList() unsafe.Pointer {
	return C.SystemTrayIcon2a00b6___findChildren_newList(ptr.Pointer())
}

func (ptr *SystemTrayIcon) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SystemTrayIcon2a00b6___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SystemTrayIcon) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SystemTrayIcon) __findChildren_newList3() unsafe.Pointer {
	return C.SystemTrayIcon2a00b6___findChildren_newList3(ptr.Pointer())
}

func (ptr *SystemTrayIcon) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SystemTrayIcon2a00b6___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SystemTrayIcon) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SystemTrayIcon) __qFindChildren_newList2() unsafe.Pointer {
	return C.SystemTrayIcon2a00b6___qFindChildren_newList2(ptr.Pointer())
}

func NewSystemTrayIcon(parent std_core.QObject_ITF) *SystemTrayIcon {
	tmpValue := NewSystemTrayIconFromPointer(C.SystemTrayIcon2a00b6_NewSystemTrayIcon(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewSystemTrayIcon2(icon std_gui.QIcon_ITF, parent std_core.QObject_ITF) *SystemTrayIcon {
	tmpValue := NewSystemTrayIconFromPointer(C.SystemTrayIcon2a00b6_NewSystemTrayIcon2(std_gui.PointerFromQIcon(icon), std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackSystemTrayIcon2a00b6_DestroySystemTrayIcon
func callbackSystemTrayIcon2a00b6_DestroySystemTrayIcon(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~SystemTrayIcon"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSystemTrayIconFromPointer(ptr).DestroySystemTrayIconDefault()
	}
}

func (ptr *SystemTrayIcon) ConnectDestroySystemTrayIcon(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~SystemTrayIcon"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~SystemTrayIcon", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~SystemTrayIcon", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SystemTrayIcon) DisconnectDestroySystemTrayIcon() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~SystemTrayIcon")
	}
}

func (ptr *SystemTrayIcon) DestroySystemTrayIcon() {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_DestroySystemTrayIcon(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *SystemTrayIcon) DestroySystemTrayIconDefault() {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_DestroySystemTrayIconDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSystemTrayIcon2a00b6_Activated
func callbackSystemTrayIcon2a00b6_Activated(ptr unsafe.Pointer, reason C.longlong) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		(*(*func(std_widgets.QSystemTrayIcon__ActivationReason))(signal))(std_widgets.QSystemTrayIcon__ActivationReason(reason))
	}

}

//export callbackSystemTrayIcon2a00b6_Event
func callbackSystemTrayIcon2a00b6_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSystemTrayIconFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *SystemTrayIcon) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SystemTrayIcon2a00b6_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackSystemTrayIcon2a00b6_Hide
func callbackSystemTrayIcon2a00b6_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSystemTrayIconFromPointer(ptr).HideDefault()
	}
}

func (ptr *SystemTrayIcon) HideDefault() {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_HideDefault(ptr.Pointer())
	}
}

//export callbackSystemTrayIcon2a00b6_MessageClicked
func callbackSystemTrayIcon2a00b6_MessageClicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "messageClicked"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackSystemTrayIcon2a00b6_SetVisible
func callbackSystemTrayIcon2a00b6_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewSystemTrayIconFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *SystemTrayIcon) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackSystemTrayIcon2a00b6_Show
func callbackSystemTrayIcon2a00b6_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSystemTrayIconFromPointer(ptr).ShowDefault()
	}
}

func (ptr *SystemTrayIcon) ShowDefault() {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_ShowDefault(ptr.Pointer())
	}
}

//export callbackSystemTrayIcon2a00b6_ShowMessage
func callbackSystemTrayIcon2a00b6_ShowMessage(ptr unsafe.Pointer, title C.struct_Moc_PackedString, message C.struct_Moc_PackedString, icon C.longlong, millisecondsTimeoutHint C.int) {
	if signal := qt.GetSignal(ptr, "showMessage"); signal != nil {
		(*(*func(string, string, std_widgets.QSystemTrayIcon__MessageIcon, int))(signal))(cGoUnpackString(title), cGoUnpackString(message), std_widgets.QSystemTrayIcon__MessageIcon(icon), int(int32(millisecondsTimeoutHint)))
	} else {
		NewSystemTrayIconFromPointer(ptr).ShowMessageDefault(cGoUnpackString(title), cGoUnpackString(message), std_widgets.QSystemTrayIcon__MessageIcon(icon), int(int32(millisecondsTimeoutHint)))
	}
}

func (ptr *SystemTrayIcon) ShowMessageDefault(title string, message string, icon std_widgets.QSystemTrayIcon__MessageIcon, millisecondsTimeoutHint int) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.SystemTrayIcon2a00b6_ShowMessageDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))}, C.longlong(icon), C.int(int32(millisecondsTimeoutHint)))
	}
}

//export callbackSystemTrayIcon2a00b6_ShowMessage2
func callbackSystemTrayIcon2a00b6_ShowMessage2(ptr unsafe.Pointer, title C.struct_Moc_PackedString, message C.struct_Moc_PackedString, icon unsafe.Pointer, millisecondsTimeoutHint C.int) {
	if signal := qt.GetSignal(ptr, "showMessage2"); signal != nil {
		(*(*func(string, string, *std_gui.QIcon, int))(signal))(cGoUnpackString(title), cGoUnpackString(message), std_gui.NewQIconFromPointer(icon), int(int32(millisecondsTimeoutHint)))
	} else {
		NewSystemTrayIconFromPointer(ptr).ShowMessage2Default(cGoUnpackString(title), cGoUnpackString(message), std_gui.NewQIconFromPointer(icon), int(int32(millisecondsTimeoutHint)))
	}
}

func (ptr *SystemTrayIcon) ShowMessage2Default(title string, message string, icon std_gui.QIcon_ITF, millisecondsTimeoutHint int) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.SystemTrayIcon2a00b6_ShowMessage2Default(ptr.Pointer(), C.struct_Moc_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))}, std_gui.PointerFromQIcon(icon), C.int(int32(millisecondsTimeoutHint)))
	}
}

//export callbackSystemTrayIcon2a00b6_ChildEvent
func callbackSystemTrayIcon2a00b6_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewSystemTrayIconFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *SystemTrayIcon) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackSystemTrayIcon2a00b6_ConnectNotify
func callbackSystemTrayIcon2a00b6_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSystemTrayIconFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *SystemTrayIcon) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSystemTrayIcon2a00b6_CustomEvent
func callbackSystemTrayIcon2a00b6_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewSystemTrayIconFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *SystemTrayIcon) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackSystemTrayIcon2a00b6_DeleteLater
func callbackSystemTrayIcon2a00b6_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSystemTrayIconFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *SystemTrayIcon) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSystemTrayIcon2a00b6_Destroyed
func callbackSystemTrayIcon2a00b6_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackSystemTrayIcon2a00b6_DisconnectNotify
func callbackSystemTrayIcon2a00b6_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSystemTrayIconFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *SystemTrayIcon) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSystemTrayIcon2a00b6_EventFilter
func callbackSystemTrayIcon2a00b6_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSystemTrayIconFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *SystemTrayIcon) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SystemTrayIcon2a00b6_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackSystemTrayIcon2a00b6_ObjectNameChanged
func callbackSystemTrayIcon2a00b6_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackSystemTrayIcon2a00b6_TimerEvent
func callbackSystemTrayIcon2a00b6_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewSystemTrayIconFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *SystemTrayIcon) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcon2a00b6_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
