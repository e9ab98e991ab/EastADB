
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

type TMargins struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewMargins
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMargins() *TMargins {
    m := new(TMargins)
    m.instance = Margins_Create()
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MarginsFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func MarginsFromInst(inst uintptr) *TMargins {
    m := new(TMargins)
    m.instance = inst
    m.ptr = unsafe.Pointer(inst)
    return m
}

// MarginsFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func MarginsFromObj(obj IObject) *TMargins {
    m := new(TMargins)
    m.instance = CheckPtr(obj)
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MarginsFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func MarginsFromUnsafePointer(ptr unsafe.Pointer) *TMargins {
    m := new(TMargins)
    m.instance = uintptr(ptr)
    m.ptr = ptr
    return m
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (m *TMargins) Free() {
    if m.instance != 0 {
        Margins_Free(m.instance)
        m.instance = 0
        m.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMargins) Instance() uintptr {
    return m.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMargins) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMargins) IsValid() bool {
    return m.instance != 0
}

// TMarginsClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMarginsClass() TClass {
    return Margins_StaticClassType()
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (m *TMargins) SetBounds(ALeft int32, ATop int32, ARight int32, ABottom int32) {
    Margins_SetBounds(m.instance, ALeft , ATop , ARight , ABottom)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (m *TMargins) Assign(Source IObject) {
    Margins_Assign(m.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (m *TMargins) GetNamePath() string {
    return Margins_GetNamePath(m.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (m *TMargins) DisposeOf() {
    Margins_DisposeOf(m.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMargins) ClassType() TClass {
    return Margins_ClassType(m.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMargins) ClassName() string {
    return Margins_ClassName(m.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMargins) InstanceSize() int32 {
    return Margins_InstanceSize(m.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMargins) InheritsFrom(AClass TClass) bool {
    return Margins_InheritsFrom(m.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMargins) Equals(Obj IObject) bool {
    return Margins_Equals(m.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMargins) GetHashCode() int32 {
    return Margins_GetHashCode(m.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (m *TMargins) ToString() string {
    return Margins_ToString(m.instance)
}

// ControlLeft
func (m *TMargins) ControlLeft() int32 {
    return Margins_GetControlLeft(m.instance)
}

// ControlTop
func (m *TMargins) ControlTop() int32 {
    return Margins_GetControlTop(m.instance)
}

// ControlWidth
func (m *TMargins) ControlWidth() int32 {
    return Margins_GetControlWidth(m.instance)
}

// ControlHeight
func (m *TMargins) ControlHeight() int32 {
    return Margins_GetControlHeight(m.instance)
}

// ExplicitLeft
func (m *TMargins) ExplicitLeft() int32 {
    return Margins_GetExplicitLeft(m.instance)
}

// ExplicitTop
func (m *TMargins) ExplicitTop() int32 {
    return Margins_GetExplicitTop(m.instance)
}

// ExplicitWidth
func (m *TMargins) ExplicitWidth() int32 {
    return Margins_GetExplicitWidth(m.instance)
}

// ExplicitHeight
func (m *TMargins) ExplicitHeight() int32 {
    return Margins_GetExplicitHeight(m.instance)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (m *TMargins) SetOnChange(fn TNotifyEvent) {
    Margins_SetOnChange(m.instance, fn)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (m *TMargins) Left() int32 {
    return Margins_GetLeft(m.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (m *TMargins) SetLeft(value int32) {
    Margins_SetLeft(m.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (m *TMargins) Top() int32 {
    return Margins_GetTop(m.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (m *TMargins) SetTop(value int32) {
    Margins_SetTop(m.instance, value)
}

// Right
func (m *TMargins) Right() int32 {
    return Margins_GetRight(m.instance)
}

// SetRight
func (m *TMargins) SetRight(value int32) {
    Margins_SetRight(m.instance, value)
}

// Bottom
func (m *TMargins) Bottom() int32 {
    return Margins_GetBottom(m.instance)
}

// SetBottom
func (m *TMargins) SetBottom(value int32) {
    Margins_SetBottom(m.instance, value)
}

