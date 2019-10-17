

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


class SystemTrayIcond8a073: public QSystemTrayIcon
{
Q_OBJECT
public:
	SystemTrayIcond8a073(QObject *parent = Q_NULLPTR) : QSystemTrayIcon(parent) {qRegisterMetaType<quintptr>("quintptr");SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaType();SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaTypes();callbackSystemTrayIcond8a073_Constructor(this);};
	SystemTrayIcond8a073(const QIcon &icon, QObject *parent = Q_NULLPTR) : QSystemTrayIcon(icon, parent) {qRegisterMetaType<quintptr>("quintptr");SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaType();SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaTypes();callbackSystemTrayIcond8a073_Constructor(this);};
	 ~SystemTrayIcond8a073() { callbackSystemTrayIcond8a073_DestroySystemTrayIcon(this); };
	void Signal_Activated(QSystemTrayIcon::ActivationReason reason) { callbackSystemTrayIcond8a073_Activated(this, reason); };
	bool event(QEvent * e) { return callbackSystemTrayIcond8a073_Event(this, e) != 0; };
	void hide() { callbackSystemTrayIcond8a073_Hide(this); };
	void Signal_MessageClicked() { callbackSystemTrayIcond8a073_MessageClicked(this); };
	void setVisible(bool visible) { callbackSystemTrayIcond8a073_SetVisible(this, visible); };
	void show() { callbackSystemTrayIcond8a073_Show(this); };
	void showMessage(const QString & title, const QString & message, QSystemTrayIcon::MessageIcon icon, int millisecondsTimeoutHint) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };QByteArray t6f9b9a = message.toUtf8(); Moc_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackSystemTrayIcond8a073_ShowMessage(this, titlePacked, messagePacked, icon, millisecondsTimeoutHint); };
	void showMessage(const QString & title, const QString & message, const QIcon & icon, int millisecondsTimeoutHint) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };QByteArray t6f9b9a = message.toUtf8(); Moc_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackSystemTrayIcond8a073_ShowMessage2(this, titlePacked, messagePacked, const_cast<QIcon*>(&icon), millisecondsTimeoutHint); };
	void childEvent(QChildEvent * event) { callbackSystemTrayIcond8a073_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackSystemTrayIcond8a073_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackSystemTrayIcond8a073_CustomEvent(this, event); };
	void deleteLater() { callbackSystemTrayIcond8a073_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackSystemTrayIcond8a073_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackSystemTrayIcond8a073_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackSystemTrayIcond8a073_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackSystemTrayIcond8a073_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackSystemTrayIcond8a073_TimerEvent(this, event); };
signals:
public slots:
	void triggerSlot(quintptr a, bool b) { callbackSystemTrayIcond8a073_TriggerSlot(this, a, b); };
	void connectionDead(QString err) { QByteArray teb35c3 = err.toUtf8(); Moc_PackedString errPacked = { const_cast<char*>(teb35c3.prepend("WHITESPACE").constData()+10), teb35c3.size()-10 };callbackSystemTrayIcond8a073_ConnectionDead(this, errPacked); };
	void hideIcon() { callbackSystemTrayIcond8a073_HideIcon(this); };
private:
};

Q_DECLARE_METATYPE(SystemTrayIcond8a073*)


void SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaTypes() {
}

void SystemTrayIcond8a073_TriggerSlot(void* ptr, uintptr_t a, char b)
{
	QMetaObject::invokeMethod(static_cast<SystemTrayIcond8a073*>(ptr), "triggerSlot", Q_ARG(quintptr, a), Q_ARG(bool, b != 0));
}

void SystemTrayIcond8a073_ConnectionDead(void* ptr, struct Moc_PackedString err)
{
	QMetaObject::invokeMethod(static_cast<SystemTrayIcond8a073*>(ptr), "connectionDead", Q_ARG(QString, QString::fromUtf8(err.data, err.len)));
}

void SystemTrayIcond8a073_HideIcon(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<SystemTrayIcond8a073*>(ptr), "hideIcon");
}

int SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaType()
{
	return qRegisterMetaType<SystemTrayIcond8a073*>();
}

int SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<SystemTrayIcond8a073*>(const_cast<const char*>(typeName));
}

int SystemTrayIcond8a073_SystemTrayIcond8a073_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<SystemTrayIcond8a073>();
#else
	return 0;
#endif
}

int SystemTrayIcond8a073_SystemTrayIcond8a073_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<SystemTrayIcond8a073>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* SystemTrayIcond8a073___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SystemTrayIcond8a073___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SystemTrayIcond8a073___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* SystemTrayIcond8a073___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void SystemTrayIcond8a073___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* SystemTrayIcond8a073___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* SystemTrayIcond8a073___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SystemTrayIcond8a073___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SystemTrayIcond8a073___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SystemTrayIcond8a073___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SystemTrayIcond8a073___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SystemTrayIcond8a073___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SystemTrayIcond8a073___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SystemTrayIcond8a073___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SystemTrayIcond8a073___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SystemTrayIcond8a073_NewSystemTrayIcon(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(static_cast<QWindow*>(parent));
	} else {
		return new SystemTrayIcond8a073(static_cast<QObject*>(parent));
	}
}

void* SystemTrayIcond8a073_NewSystemTrayIcon2(void* icon, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(*static_cast<QIcon*>(icon), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(*static_cast<QIcon*>(icon), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(*static_cast<QIcon*>(icon), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(*static_cast<QIcon*>(icon), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(*static_cast<QIcon*>(icon), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(*static_cast<QIcon*>(icon), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(*static_cast<QIcon*>(icon), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new SystemTrayIcond8a073(*static_cast<QIcon*>(icon), static_cast<QWindow*>(parent));
	} else {
		return new SystemTrayIcond8a073(*static_cast<QIcon*>(icon), static_cast<QObject*>(parent));
	}
}

void SystemTrayIcond8a073_DestroySystemTrayIcon(void* ptr)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->~SystemTrayIcond8a073();
}

void SystemTrayIcond8a073_DestroySystemTrayIconDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char SystemTrayIcond8a073_EventDefault(void* ptr, void* e)
{
	return static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::event(static_cast<QEvent*>(e));
}

void SystemTrayIcond8a073_HideDefault(void* ptr)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::hide();
}

void SystemTrayIcond8a073_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::setVisible(visible != 0);
}

void SystemTrayIcond8a073_ShowDefault(void* ptr)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::show();
}

void SystemTrayIcond8a073_ShowMessageDefault(void* ptr, struct Moc_PackedString title, struct Moc_PackedString message, long long icon, int millisecondsTimeoutHint)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::showMessage(QString::fromUtf8(title.data, title.len), QString::fromUtf8(message.data, message.len), static_cast<QSystemTrayIcon::MessageIcon>(icon), millisecondsTimeoutHint);
}

void SystemTrayIcond8a073_ShowMessage2Default(void* ptr, struct Moc_PackedString title, struct Moc_PackedString message, void* icon, int millisecondsTimeoutHint)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::showMessage(QString::fromUtf8(title.data, title.len), QString::fromUtf8(message.data, message.len), *static_cast<QIcon*>(icon), millisecondsTimeoutHint);
}

void SystemTrayIcond8a073_ChildEventDefault(void* ptr, void* event)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::childEvent(static_cast<QChildEvent*>(event));
}

void SystemTrayIcond8a073_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void SystemTrayIcond8a073_CustomEventDefault(void* ptr, void* event)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::customEvent(static_cast<QEvent*>(event));
}

void SystemTrayIcond8a073_DeleteLaterDefault(void* ptr)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::deleteLater();
}

void SystemTrayIcond8a073_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char SystemTrayIcond8a073_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void SystemTrayIcond8a073_TimerEventDefault(void* ptr, void* event)
{
	static_cast<SystemTrayIcond8a073*>(ptr)->QSystemTrayIcon::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
