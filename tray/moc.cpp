

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QIcon>
#include <QLayout>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QString>
#include <QSystemTrayIcon>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class SystemTrayIcon2a00b6: public QSystemTrayIcon
{
Q_OBJECT
public:
	SystemTrayIcon2a00b6(QObject *parent = Q_NULLPTR) : QSystemTrayIcon(parent) {qRegisterMetaType<quintptr>("quintptr");SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaType();SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaTypes();callbackSystemTrayIcon2a00b6_Constructor(this);};
	SystemTrayIcon2a00b6(const QIcon &icon, QObject *parent = Q_NULLPTR) : QSystemTrayIcon(icon, parent) {qRegisterMetaType<quintptr>("quintptr");SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaType();SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaTypes();callbackSystemTrayIcon2a00b6_Constructor(this);};
	 ~SystemTrayIcon2a00b6() { callbackSystemTrayIcon2a00b6_DestroySystemTrayIcon(this); };
	void Signal_Activated(QSystemTrayIcon::ActivationReason reason) { callbackSystemTrayIcon2a00b6_Activated(this, reason); };
	bool event(QEvent * e) { return callbackSystemTrayIcon2a00b6_Event(this, e) != 0; };
	void hide() { callbackSystemTrayIcon2a00b6_Hide(this); };
	void Signal_MessageClicked() { callbackSystemTrayIcon2a00b6_MessageClicked(this); };
	void setVisible(bool visible) { callbackSystemTrayIcon2a00b6_SetVisible(this, visible); };
	void show() { callbackSystemTrayIcon2a00b6_Show(this); };
	void showMessage(const QString & title, const QString & message, QSystemTrayIcon::MessageIcon icon, int millisecondsTimeoutHint) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };QByteArray t6f9b9a = message.toUtf8(); Moc_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackSystemTrayIcon2a00b6_ShowMessage(this, titlePacked, messagePacked, icon, millisecondsTimeoutHint); };
	void showMessage(const QString & title, const QString & message, const QIcon & icon, int millisecondsTimeoutHint) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };QByteArray t6f9b9a = message.toUtf8(); Moc_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackSystemTrayIcon2a00b6_ShowMessage2(this, titlePacked, messagePacked, const_cast<QIcon*>(&icon), millisecondsTimeoutHint); };
	void childEvent(QChildEvent * event) { callbackSystemTrayIcon2a00b6_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackSystemTrayIcon2a00b6_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackSystemTrayIcon2a00b6_CustomEvent(this, event); };
	void deleteLater() { callbackSystemTrayIcon2a00b6_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackSystemTrayIcon2a00b6_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackSystemTrayIcon2a00b6_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackSystemTrayIcon2a00b6_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackSystemTrayIcon2a00b6_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackSystemTrayIcon2a00b6_TimerEvent(this, event); };
signals:
public slots:
	void triggerSlot(quintptr a, bool b) { callbackSystemTrayIcon2a00b6_TriggerSlot(this, a, b); };
	void connectionDead() { callbackSystemTrayIcon2a00b6_ConnectionDead(this); };
private:
};

Q_DECLARE_METATYPE(SystemTrayIcon2a00b6*)


void SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaTypes() {
}

void SystemTrayIcon2a00b6_TriggerSlot(void* ptr, uintptr_t a, char b)
{
	QMetaObject::invokeMethod(static_cast<SystemTrayIcon2a00b6*>(ptr), "triggerSlot", Q_ARG(quintptr, a), Q_ARG(bool, b != 0));
}

void SystemTrayIcon2a00b6_ConnectionDead(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<SystemTrayIcon2a00b6*>(ptr), "connectionDead");
}

int SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaType()
{
	return qRegisterMetaType<SystemTrayIcon2a00b6*>();
}

int SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<SystemTrayIcon2a00b6*>(const_cast<const char*>(typeName));
}

int SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<SystemTrayIcon2a00b6>();
#else
	return 0;
#endif
}

int SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<SystemTrayIcon2a00b6>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* SystemTrayIcon2a00b6___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SystemTrayIcon2a00b6___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SystemTrayIcon2a00b6___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* SystemTrayIcon2a00b6___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void SystemTrayIcon2a00b6___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* SystemTrayIcon2a00b6___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* SystemTrayIcon2a00b6___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SystemTrayIcon2a00b6___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SystemTrayIcon2a00b6___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SystemTrayIcon2a00b6___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SystemTrayIcon2a00b6___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SystemTrayIcon2a00b6___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SystemTrayIcon2a00b6___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SystemTrayIcon2a00b6___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SystemTrayIcon2a00b6___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SystemTrayIcon2a00b6_NewSystemTrayIcon(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(static_cast<QWindow*>(parent));
	} else {
		return new SystemTrayIcon2a00b6(static_cast<QObject*>(parent));
	}
}

void* SystemTrayIcon2a00b6_NewSystemTrayIcon2(void* icon, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(*static_cast<QIcon*>(icon), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(*static_cast<QIcon*>(icon), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(*static_cast<QIcon*>(icon), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(*static_cast<QIcon*>(icon), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(*static_cast<QIcon*>(icon), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(*static_cast<QIcon*>(icon), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(*static_cast<QIcon*>(icon), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcon2a00b6(*static_cast<QIcon*>(icon), static_cast<QWindow*>(parent));
	} else {
		return new SystemTrayIcon2a00b6(*static_cast<QIcon*>(icon), static_cast<QObject*>(parent));
	}
}

void SystemTrayIcon2a00b6_DestroySystemTrayIcon(void* ptr)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->~SystemTrayIcon2a00b6();
}

void SystemTrayIcon2a00b6_DestroySystemTrayIconDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char SystemTrayIcon2a00b6_EventDefault(void* ptr, void* e)
{
	return static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::event(static_cast<QEvent*>(e));
}

void SystemTrayIcon2a00b6_HideDefault(void* ptr)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::hide();
}

void SystemTrayIcon2a00b6_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::setVisible(visible != 0);
}

void SystemTrayIcon2a00b6_ShowDefault(void* ptr)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::show();
}

void SystemTrayIcon2a00b6_ShowMessageDefault(void* ptr, struct Moc_PackedString title, struct Moc_PackedString message, long long icon, int millisecondsTimeoutHint)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::showMessage(QString::fromUtf8(title.data, title.len), QString::fromUtf8(message.data, message.len), static_cast<QSystemTrayIcon::MessageIcon>(icon), millisecondsTimeoutHint);
}

void SystemTrayIcon2a00b6_ShowMessage2Default(void* ptr, struct Moc_PackedString title, struct Moc_PackedString message, void* icon, int millisecondsTimeoutHint)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::showMessage(QString::fromUtf8(title.data, title.len), QString::fromUtf8(message.data, message.len), *static_cast<QIcon*>(icon), millisecondsTimeoutHint);
}

void SystemTrayIcon2a00b6_ChildEventDefault(void* ptr, void* event)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::childEvent(static_cast<QChildEvent*>(event));
}

void SystemTrayIcon2a00b6_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void SystemTrayIcon2a00b6_CustomEventDefault(void* ptr, void* event)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::customEvent(static_cast<QEvent*>(event));
}

void SystemTrayIcon2a00b6_DeleteLaterDefault(void* ptr)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::deleteLater();
}

void SystemTrayIcon2a00b6_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char SystemTrayIcon2a00b6_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void SystemTrayIcon2a00b6_TimerEventDefault(void* ptr, void* event)
{
	static_cast<SystemTrayIcon2a00b6*>(ptr)->QSystemTrayIcon::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
