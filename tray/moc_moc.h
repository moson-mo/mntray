/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.13.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <memory>
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.13.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_SystemTrayIcon2a00b6_t {
    QByteArrayData data[7];
    char stringdata0[62];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_SystemTrayIcon2a00b6_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_SystemTrayIcon2a00b6_t qt_meta_stringdata_SystemTrayIcon2a00b6 = {
    {
QT_MOC_LITERAL(0, 0, 20), // "SystemTrayIcon2a00b6"
QT_MOC_LITERAL(1, 21, 11), // "triggerSlot"
QT_MOC_LITERAL(2, 33, 0), // ""
QT_MOC_LITERAL(3, 34, 8), // "quintptr"
QT_MOC_LITERAL(4, 43, 1), // "a"
QT_MOC_LITERAL(5, 45, 1), // "b"
QT_MOC_LITERAL(6, 47, 14) // "connectionDead"

    },
    "SystemTrayIcon2a00b6\0triggerSlot\0\0"
    "quintptr\0a\0b\0connectionDead"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_SystemTrayIcon2a00b6[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       2,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    2,   24,    2, 0x0a /* Public */,
       6,    0,   29,    2, 0x0a /* Public */,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 3, QMetaType::Bool,    4,    5,
    QMetaType::Void,

       0        // eod
};

void SystemTrayIcon2a00b6::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<SystemTrayIcon2a00b6 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->triggerSlot((*reinterpret_cast< quintptr(*)>(_a[1])),(*reinterpret_cast< bool(*)>(_a[2]))); break;
        case 1: _t->connectionDead(); break;
        default: ;
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject SystemTrayIcon2a00b6::staticMetaObject = { {
    &QSystemTrayIcon::staticMetaObject,
    qt_meta_stringdata_SystemTrayIcon2a00b6.data,
    qt_meta_data_SystemTrayIcon2a00b6,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *SystemTrayIcon2a00b6::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *SystemTrayIcon2a00b6::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_SystemTrayIcon2a00b6.stringdata0))
        return static_cast<void*>(this);
    return QSystemTrayIcon::qt_metacast(_clname);
}

int SystemTrayIcon2a00b6::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QSystemTrayIcon::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 2)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 2)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 2;
    }
    return _id;
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
