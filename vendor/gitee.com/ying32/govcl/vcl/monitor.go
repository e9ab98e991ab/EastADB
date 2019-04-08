
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TMonitor struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewMonitor
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMonitor() *TMonitor {
    m := new(TMonitor)
    m.instance = Monitor_Create()
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MonitorFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func MonitorFromInst(inst uintptr) *TMonitor {
    m := new(TMonitor)
    m.instance = inst
    m.ptr = unsafe.Pointer(inst)
    return m
}

// MonitorFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func MonitorFromObj(obj IObject) *TMonitor {
    m := new(TMonitor)
    m.instance = CheckPtr(obj)
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MonitorFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func MonitorFromUnsafePointer(ptr unsafe.Pointer) *TMonitor {
    m := new(TMonitor)
    m.instance = uintptr(ptr)
    m.ptr = ptr
    return m
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (m *TMonitor) Free() {
    if m.instance != 0 {
        Monitor_Free(m.instance)
        m.instance = 0
        m.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMonitor) Instance() uintptr {
    return m.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMonitor) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMonitor) IsValid() bool {
    return m.instance != 0
}

// TMonitorClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMonitorClass() TClass {
    return Monitor_StaticClassType()
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (m *TMonitor) DisposeOf() {
    Monitor_DisposeOf(m.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMonitor) ClassType() TClass {
    return Monitor_ClassType(m.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMonitor) ClassName() string {
    return Monitor_ClassName(m.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMonitor) InstanceSize() int32 {
    return Monitor_InstanceSize(m.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMonitor) InheritsFrom(AClass TClass) bool {
    return Monitor_InheritsFrom(m.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMonitor) Equals(Obj IObject) bool {
    return Monitor_Equals(m.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMonitor) GetHashCode() int32 {
    return Monitor_GetHashCode(m.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (m *TMonitor) ToString() string {
    return Monitor_ToString(m.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (m *TMonitor) Handle() HMONITOR {
    return Monitor_GetHandle(m.instance)
}

// MonitorNum
func (m *TMonitor) MonitorNum() int32 {
    return Monitor_GetMonitorNum(m.instance)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (m *TMonitor) Left() int32 {
    return Monitor_GetLeft(m.instance)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (m *TMonitor) Height() int32 {
    return Monitor_GetHeight(m.instance)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (m *TMonitor) Top() int32 {
    return Monitor_GetTop(m.instance)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (m *TMonitor) Width() int32 {
    return Monitor_GetWidth(m.instance)
}

// BoundsRect
func (m *TMonitor) BoundsRect() TRect {
    return Monitor_GetBoundsRect(m.instance)
}

// WorkareaRect
func (m *TMonitor) WorkareaRect() TRect {
    return Monitor_GetWorkareaRect(m.instance)
}

// Primary
func (m *TMonitor) Primary() bool {
    return Monitor_GetPrimary(m.instance)
}

// PixelsPerInch
func (m *TMonitor) PixelsPerInch() int32 {
    return Monitor_GetPixelsPerInch(m.instance)
}

