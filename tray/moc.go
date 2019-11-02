package tray

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"errors"
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

//export callbackSystemTrayIcond8a073_Constructor
func callbackSystemTrayIcond8a073_Constructor(ptr unsafe.Pointer) {
	this := NewSystemTrayIconFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackSystemTrayIcond8a073_NewArticleSlot
func callbackSystemTrayIcond8a073_NewArticleSlot(ptr unsafe.Pointer, v0 C.uintptr_t, v1 C.char) {
	var v0D Article
	if v0I, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(v0))); ok {
		qt.UnregisterTemp(unsafe.Pointer(uintptr(v0)))
		v0D = (*(*Article)(v0I))
	}
	if signal := qt.GetSignal(ptr, "newArticleSlot"); signal != nil {
		(*(*func(Article, bool))(signal))(v0D, int8(v1) != 0)
	}

}

func (ptr *SystemTrayIcon) ConnectNewArticleSlot(f func(v0 Article, v1 bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "newArticleSlot"); signal != nil {
			f := func(v0 Article, v1 bool) {
				(*(*func(Article, bool))(signal))(v0, v1)
				f(v0, v1)
			}
			qt.ConnectSignal(ptr.Pointer(), "newArticleSlot", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newArticleSlot", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SystemTrayIcon) DisconnectNewArticleSlot() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "newArticleSlot")
	}
}

func (ptr *SystemTrayIcon) NewArticleSlot(v0 Article, v1 bool) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(&v0), unsafe.Pointer(&v0))
		C.SystemTrayIcond8a073_NewArticleSlot(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(&v0))), C.char(int8(qt.GoBoolToInt(v1))))
	}
}

//export callbackSystemTrayIcond8a073_ErrorSlot
func callbackSystemTrayIcond8a073_ErrorSlot(ptr unsafe.Pointer, v0 C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "errorSlot"); signal != nil {
		(*(*func(error))(signal))(errors.New(cGoUnpackString(v0)))
	}

}

func (ptr *SystemTrayIcon) ConnectErrorSlot(f func(v0 error)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorSlot"); signal != nil {
			f := func(v0 error) {
				(*(*func(error))(signal))(v0)
				f(v0)
			}
			qt.ConnectSignal(ptr.Pointer(), "errorSlot", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorSlot", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SystemTrayIcon) DisconnectErrorSlot() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorSlot")
	}
}

func (ptr *SystemTrayIcon) ErrorSlot(v0 error) {
	if ptr.Pointer() != nil {
		v0C := C.CString(func() string {
			tmp := v0
			if tmp != nil {
				return tmp.Error()
			}
			return ""
		}())
		defer C.free(unsafe.Pointer(v0C))
		C.SystemTrayIcond8a073_ErrorSlot(ptr.Pointer(), C.struct_Moc_PackedString{data: v0C, len: C.longlong(-1)})
	}
}

//export callbackSystemTrayIcond8a073_HideIconSlot
func callbackSystemTrayIcond8a073_HideIconSlot(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideIconSlot"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *SystemTrayIcon) ConnectHideIconSlot(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hideIconSlot"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hideIconSlot", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hideIconSlot", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SystemTrayIcon) DisconnectHideIconSlot() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hideIconSlot")
	}
}

func (ptr *SystemTrayIcon) HideIconSlot() {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073_HideIconSlot(ptr.Pointer())
	}
}

func SystemTrayIcon_QRegisterMetaType() int {
	return int(int32(C.SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaType()))
}

func (ptr *SystemTrayIcon) QRegisterMetaType() int {
	return int(int32(C.SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaType()))
}

func SystemTrayIcon_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaType2(typeNameC)))
}

func (ptr *SystemTrayIcon) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaType2(typeNameC)))
}

func SystemTrayIcon_QmlRegisterType() int {
	return int(int32(C.SystemTrayIcond8a073_SystemTrayIcond8a073_QmlRegisterType()))
}

func (ptr *SystemTrayIcon) QmlRegisterType() int {
	return int(int32(C.SystemTrayIcond8a073_SystemTrayIcond8a073_QmlRegisterType()))
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
	return int(int32(C.SystemTrayIcond8a073_SystemTrayIcond8a073_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.SystemTrayIcond8a073_SystemTrayIcond8a073_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *SystemTrayIcon) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SystemTrayIcond8a073___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SystemTrayIcon) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SystemTrayIcon) __children_newList() unsafe.Pointer {
	return C.SystemTrayIcond8a073___children_newList(ptr.Pointer())
}

func (ptr *SystemTrayIcon) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.SystemTrayIcond8a073___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *SystemTrayIcon) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *SystemTrayIcon) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.SystemTrayIcond8a073___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *SystemTrayIcon) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SystemTrayIcond8a073___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SystemTrayIcon) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SystemTrayIcon) __findChildren_newList() unsafe.Pointer {
	return C.SystemTrayIcond8a073___findChildren_newList(ptr.Pointer())
}

func (ptr *SystemTrayIcon) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SystemTrayIcond8a073___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SystemTrayIcon) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SystemTrayIcon) __findChildren_newList3() unsafe.Pointer {
	return C.SystemTrayIcond8a073___findChildren_newList3(ptr.Pointer())
}

func (ptr *SystemTrayIcon) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SystemTrayIcond8a073___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SystemTrayIcon) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SystemTrayIcon) __qFindChildren_newList2() unsafe.Pointer {
	return C.SystemTrayIcond8a073___qFindChildren_newList2(ptr.Pointer())
}

func NewSystemTrayIcon(parent std_core.QObject_ITF) *SystemTrayIcon {
	SystemTrayIcon_QRegisterMetaType()
	tmpValue := NewSystemTrayIconFromPointer(C.SystemTrayIcond8a073_NewSystemTrayIcon(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewSystemTrayIcon2(icon std_gui.QIcon_ITF, parent std_core.QObject_ITF) *SystemTrayIcon {
	SystemTrayIcon_QRegisterMetaType()
	tmpValue := NewSystemTrayIconFromPointer(C.SystemTrayIcond8a073_NewSystemTrayIcon2(std_gui.PointerFromQIcon(icon), std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackSystemTrayIcond8a073_DestroySystemTrayIcon
func callbackSystemTrayIcond8a073_DestroySystemTrayIcon(ptr unsafe.Pointer) {
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
		C.SystemTrayIcond8a073_DestroySystemTrayIcon(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *SystemTrayIcon) DestroySystemTrayIconDefault() {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073_DestroySystemTrayIconDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSystemTrayIcond8a073_Activated
func callbackSystemTrayIcond8a073_Activated(ptr unsafe.Pointer, reason C.longlong) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		(*(*func(std_widgets.QSystemTrayIcon__ActivationReason))(signal))(std_widgets.QSystemTrayIcon__ActivationReason(reason))
	}

}

//export callbackSystemTrayIcond8a073_Event
func callbackSystemTrayIcond8a073_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSystemTrayIconFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *SystemTrayIcon) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SystemTrayIcond8a073_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackSystemTrayIcond8a073_Hide
func callbackSystemTrayIcond8a073_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSystemTrayIconFromPointer(ptr).HideDefault()
	}
}

func (ptr *SystemTrayIcon) HideDefault() {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073_HideDefault(ptr.Pointer())
	}
}

//export callbackSystemTrayIcond8a073_MessageClicked
func callbackSystemTrayIcond8a073_MessageClicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "messageClicked"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackSystemTrayIcond8a073_SetVisible
func callbackSystemTrayIcond8a073_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewSystemTrayIconFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *SystemTrayIcon) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackSystemTrayIcond8a073_Show
func callbackSystemTrayIcond8a073_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSystemTrayIconFromPointer(ptr).ShowDefault()
	}
}

func (ptr *SystemTrayIcon) ShowDefault() {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073_ShowDefault(ptr.Pointer())
	}
}

//export callbackSystemTrayIcond8a073_ShowMessage
func callbackSystemTrayIcond8a073_ShowMessage(ptr unsafe.Pointer, title C.struct_Moc_PackedString, message C.struct_Moc_PackedString, icon C.longlong, millisecondsTimeoutHint C.int) {
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
		C.SystemTrayIcond8a073_ShowMessageDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))}, C.longlong(icon), C.int(int32(millisecondsTimeoutHint)))
	}
}

//export callbackSystemTrayIcond8a073_ShowMessage2
func callbackSystemTrayIcond8a073_ShowMessage2(ptr unsafe.Pointer, title C.struct_Moc_PackedString, message C.struct_Moc_PackedString, icon unsafe.Pointer, millisecondsTimeoutHint C.int) {
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
		C.SystemTrayIcond8a073_ShowMessage2Default(ptr.Pointer(), C.struct_Moc_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))}, std_gui.PointerFromQIcon(icon), C.int(int32(millisecondsTimeoutHint)))
	}
}

//export callbackSystemTrayIcond8a073_ChildEvent
func callbackSystemTrayIcond8a073_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewSystemTrayIconFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *SystemTrayIcon) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackSystemTrayIcond8a073_ConnectNotify
func callbackSystemTrayIcond8a073_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSystemTrayIconFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *SystemTrayIcon) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSystemTrayIcond8a073_CustomEvent
func callbackSystemTrayIcond8a073_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewSystemTrayIconFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *SystemTrayIcon) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackSystemTrayIcond8a073_DeleteLater
func callbackSystemTrayIcond8a073_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSystemTrayIconFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *SystemTrayIcon) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSystemTrayIcond8a073_Destroyed
func callbackSystemTrayIcond8a073_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackSystemTrayIcond8a073_DisconnectNotify
func callbackSystemTrayIcond8a073_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSystemTrayIconFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *SystemTrayIcon) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSystemTrayIcond8a073_EventFilter
func callbackSystemTrayIcond8a073_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSystemTrayIconFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *SystemTrayIcon) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SystemTrayIcond8a073_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackSystemTrayIcond8a073_ObjectNameChanged
func callbackSystemTrayIcond8a073_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackSystemTrayIcond8a073_TimerEvent
func callbackSystemTrayIcond8a073_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewSystemTrayIconFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *SystemTrayIcon) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SystemTrayIcond8a073_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
