/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.13.1)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <memory>
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.13.1. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_SystemTrayIcond8a073_t {
    QByteArrayData data[8];
    char stringdata0[75];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_SystemTrayIcond8a073_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_SystemTrayIcond8a073_t qt_meta_stringdata_SystemTrayIcond8a073 = {
    {
QT_MOC_LITERAL(0, 0, 20), // "SystemTrayIcond8a073"
QT_MOC_LITERAL(1, 21, 14), // "newArticleSlot"
QT_MOC_LITERAL(2, 36, 0), // ""
QT_MOC_LITERAL(3, 37, 8), // "quintptr"
QT_MOC_LITERAL(4, 46, 2), // "v0"
QT_MOC_LITERAL(5, 49, 2), // "v1"
QT_MOC_LITERAL(6, 52, 9), // "errorSlot"
QT_MOC_LITERAL(7, 62, 12) // "hideIconSlot"

    },
    "SystemTrayIcond8a073\0newArticleSlot\0"
    "\0quintptr\0v0\0v1\0errorSlot\0hideIconSlot"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_SystemTrayIcond8a073[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       3,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    2,   29,    2, 0x0a /* Public */,
       6,    1,   34,    2, 0x0a /* Public */,
       7,    0,   37,    2, 0x0a /* Public */,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 3, QMetaType::Bool,    4,    5,
    QMetaType::Void, QMetaType::QString,    4,
    QMetaType::Void,

       0        // eod
};

void SystemTrayIcond8a073::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<SystemTrayIcond8a073 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->newArticleSlot((*reinterpret_cast< quintptr(*)>(_a[1])),(*reinterpret_cast< bool(*)>(_a[2]))); break;
        case 1: _t->errorSlot((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->hideIconSlot(); break;
        default: ;
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject SystemTrayIcond8a073::staticMetaObject = { {
    &QSystemTrayIcon::staticMetaObject,
    qt_meta_stringdata_SystemTrayIcond8a073.data,
    qt_meta_data_SystemTrayIcond8a073,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *SystemTrayIcond8a073::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *SystemTrayIcond8a073::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_SystemTrayIcond8a073.stringdata0))
        return static_cast<void*>(this);
    return QSystemTrayIcon::qt_metacast(_clname);
}

int SystemTrayIcond8a073::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QSystemTrayIcon::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 3)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 3)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 3;
    }
    return _id;
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
