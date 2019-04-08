
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

type TMouse struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewMouse
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMouse() *TMouse {
    m := new(TMouse)
    m.instance = Mouse_Create()
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MouseFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func MouseFromInst(inst uintptr) *TMouse {
    m := new(TMouse)
    m.instance = inst
    m.ptr = unsafe.Pointer(inst)
    return m
}

// MouseFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func MouseFromObj(obj IObject) *TMouse {
    m := new(TMouse)
    m.instance = CheckPtr(obj)
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MouseFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func MouseFromUnsafePointer(ptr unsafe.Pointer) *TMouse {
    m := new(TMouse)
    m.instance = uintptr(ptr)
    m.ptr = ptr
    return m
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (m *TMouse) Free() {
    if m.instance != 0 {
        Mouse_Free(m.instance)
        m.instance = 0
        m.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMouse) Instance() uintptr {
    return m.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMouse) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMouse) IsValid() bool {
    return m.instance != 0
}

// TMouseClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMouseClass() TClass {
    return Mouse_StaticClassType()
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (m *TMouse) DisposeOf() {
    Mouse_DisposeOf(m.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMouse) ClassType() TClass {
    return Mouse_ClassType(m.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMouse) ClassName() string {
    return Mouse_ClassName(m.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMouse) InstanceSize() int32 {
    return Mouse_InstanceSize(m.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMouse) InheritsFrom(AClass TClass) bool {
    return Mouse_InheritsFrom(m.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMouse) Equals(Obj IObject) bool {
    return Mouse_Equals(m.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMouse) GetHashCode() int32 {
    return Mouse_GetHashCode(m.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (m *TMouse) ToString() string {
    return Mouse_ToString(m.instance)
}

// Capture
func (m *TMouse) Capture() HWND {
    return Mouse_GetCapture(m.instance)
}

// SetCapture
func (m *TMouse) SetCapture(value HWND) {
    Mouse_SetCapture(m.instance, value)
}

// CursorPos
func (m *TMouse) CursorPos() TPoint {
    return Mouse_GetCursorPos(m.instance)
}

// SetCursorPos
func (m *TMouse) SetCursorPos(value TPoint) {
    Mouse_SetCursorPos(m.instance, value)
}

// IsDragging
func (m *TMouse) IsDragging() bool {
    return Mouse_GetIsDragging(m.instance)
}

// IsPanning
func (m *TMouse) IsPanning() bool {
    return Mouse_GetIsPanning(m.instance)
}

// WheelPresent
func (m *TMouse) WheelPresent() bool {
    return Mouse_GetWheelPresent(m.instance)
}

// WheelScrollLines
func (m *TMouse) WheelScrollLines() int32 {
    return Mouse_GetWheelScrollLines(m.instance)
}

