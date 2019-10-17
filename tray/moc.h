

#pragma once

#ifndef GO_MOC_d8a073_H
#define GO_MOC_d8a073_H

#include <stdint.h>

#ifdef __cplusplus
class SystemTrayIcond8a073;
void SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void SystemTrayIcond8a073_TriggerSlot(void* ptr, uintptr_t a, char b);
void SystemTrayIcond8a073_ConnectionDead(void* ptr, struct Moc_PackedString err);
void SystemTrayIcond8a073_HideIcon(void* ptr);
int SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaType();
int SystemTrayIcond8a073_SystemTrayIcond8a073_QRegisterMetaType2(char* typeName);
int SystemTrayIcond8a073_SystemTrayIcond8a073_QmlRegisterType();
int SystemTrayIcond8a073_SystemTrayIcond8a073_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* SystemTrayIcond8a073___children_atList(void* ptr, int i);
void SystemTrayIcond8a073___children_setList(void* ptr, void* i);
void* SystemTrayIcond8a073___children_newList(void* ptr);
void* SystemTrayIcond8a073___dynamicPropertyNames_atList(void* ptr, int i);
void SystemTrayIcond8a073___dynamicPropertyNames_setList(void* ptr, void* i);
void* SystemTrayIcond8a073___dynamicPropertyNames_newList(void* ptr);
void* SystemTrayIcond8a073___findChildren_atList(void* ptr, int i);
void SystemTrayIcond8a073___findChildren_setList(void* ptr, void* i);
void* SystemTrayIcond8a073___findChildren_newList(void* ptr);
void* SystemTrayIcond8a073___findChildren_atList3(void* ptr, int i);
void SystemTrayIcond8a073___findChildren_setList3(void* ptr, void* i);
void* SystemTrayIcond8a073___findChildren_newList3(void* ptr);
void* SystemTrayIcond8a073___qFindChildren_atList2(void* ptr, int i);
void SystemTrayIcond8a073___qFindChildren_setList2(void* ptr, void* i);
void* SystemTrayIcond8a073___qFindChildren_newList2(void* ptr);
void* SystemTrayIcond8a073_NewSystemTrayIcon(void* parent);
void* SystemTrayIcond8a073_NewSystemTrayIcon2(void* icon, void* parent);
void SystemTrayIcond8a073_DestroySystemTrayIcon(void* ptr);
void SystemTrayIcond8a073_DestroySystemTrayIconDefault(void* ptr);
char SystemTrayIcond8a073_EventDefault(void* ptr, void* e);
void SystemTrayIcond8a073_HideDefault(void* ptr);
void SystemTrayIcond8a073_SetVisibleDefault(void* ptr, char visible);
void SystemTrayIcond8a073_ShowDefault(void* ptr);
void SystemTrayIcond8a073_ShowMessageDefault(void* ptr, struct Moc_PackedString title, struct Moc_PackedString message, long long icon, int millisecondsTimeoutHint);
void SystemTrayIcond8a073_ShowMessage2Default(void* ptr, struct Moc_PackedString title, struct Moc_PackedString message, void* icon, int millisecondsTimeoutHint);
void SystemTrayIcond8a073_ChildEventDefault(void* ptr, void* event);
void SystemTrayIcond8a073_ConnectNotifyDefault(void* ptr, void* sign);
void SystemTrayIcond8a073_CustomEventDefault(void* ptr, void* event);
void SystemTrayIcond8a073_DeleteLaterDefault(void* ptr);
void SystemTrayIcond8a073_DisconnectNotifyDefault(void* ptr, void* sign);
char SystemTrayIcond8a073_EventFilterDefault(void* ptr, void* watched, void* event);
;
void SystemTrayIcond8a073_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif