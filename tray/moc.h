

#pragma once

#ifndef GO_MOC_2a00b6_H
#define GO_MOC_2a00b6_H

#include <stdint.h>

#ifdef __cplusplus
class SystemTrayIcon2a00b6;
void SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void SystemTrayIcon2a00b6_TriggerSlot(void* ptr, uintptr_t a, char b);
void SystemTrayIcon2a00b6_ConnectionDead(void* ptr);
int SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaType();
int SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QRegisterMetaType2(char* typeName);
int SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QmlRegisterType();
int SystemTrayIcon2a00b6_SystemTrayIcon2a00b6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* SystemTrayIcon2a00b6___children_atList(void* ptr, int i);
void SystemTrayIcon2a00b6___children_setList(void* ptr, void* i);
void* SystemTrayIcon2a00b6___children_newList(void* ptr);
void* SystemTrayIcon2a00b6___dynamicPropertyNames_atList(void* ptr, int i);
void SystemTrayIcon2a00b6___dynamicPropertyNames_setList(void* ptr, void* i);
void* SystemTrayIcon2a00b6___dynamicPropertyNames_newList(void* ptr);
void* SystemTrayIcon2a00b6___findChildren_atList(void* ptr, int i);
void SystemTrayIcon2a00b6___findChildren_setList(void* ptr, void* i);
void* SystemTrayIcon2a00b6___findChildren_newList(void* ptr);
void* SystemTrayIcon2a00b6___findChildren_atList3(void* ptr, int i);
void SystemTrayIcon2a00b6___findChildren_setList3(void* ptr, void* i);
void* SystemTrayIcon2a00b6___findChildren_newList3(void* ptr);
void* SystemTrayIcon2a00b6___qFindChildren_atList2(void* ptr, int i);
void SystemTrayIcon2a00b6___qFindChildren_setList2(void* ptr, void* i);
void* SystemTrayIcon2a00b6___qFindChildren_newList2(void* ptr);
void* SystemTrayIcon2a00b6_NewSystemTrayIcon(void* parent);
void* SystemTrayIcon2a00b6_NewSystemTrayIcon2(void* icon, void* parent);
void SystemTrayIcon2a00b6_DestroySystemTrayIcon(void* ptr);
void SystemTrayIcon2a00b6_DestroySystemTrayIconDefault(void* ptr);
char SystemTrayIcon2a00b6_EventDefault(void* ptr, void* e);
void SystemTrayIcon2a00b6_HideDefault(void* ptr);
void SystemTrayIcon2a00b6_SetVisibleDefault(void* ptr, char visible);
void SystemTrayIcon2a00b6_ShowDefault(void* ptr);
void SystemTrayIcon2a00b6_ShowMessageDefault(void* ptr, struct Moc_PackedString title, struct Moc_PackedString message, long long icon, int millisecondsTimeoutHint);
void SystemTrayIcon2a00b6_ShowMessage2Default(void* ptr, struct Moc_PackedString title, struct Moc_PackedString message, void* icon, int millisecondsTimeoutHint);
void SystemTrayIcon2a00b6_ChildEventDefault(void* ptr, void* event);
void SystemTrayIcon2a00b6_ConnectNotifyDefault(void* ptr, void* sign);
void SystemTrayIcon2a00b6_CustomEventDefault(void* ptr, void* event);
void SystemTrayIcon2a00b6_DeleteLaterDefault(void* ptr);
void SystemTrayIcon2a00b6_DisconnectNotifyDefault(void* ptr, void* sign);
char SystemTrayIcon2a00b6_EventFilterDefault(void* ptr, void* watched, void* event);
;
void SystemTrayIcon2a00b6_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif