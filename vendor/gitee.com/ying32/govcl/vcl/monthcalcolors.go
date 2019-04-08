
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

type TMonthCalColors struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewMonthCalColors
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMonthCalColors() *TMonthCalColors {
    m := new(TMonthCalColors)
    m.instance = MonthCalColors_Create()
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MonthCalColorsFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func MonthCalColorsFromInst(inst uintptr) *TMonthCalColors {
    m := new(TMonthCalColors)
    m.instance = inst
    m.ptr = unsafe.Pointer(inst)
    return m
}

// MonthCalColorsFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func MonthCalColorsFromObj(obj IObject) *TMonthCalColors {
    m := new(TMonthCalColors)
    m.instance = CheckPtr(obj)
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MonthCalColorsFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func MonthCalColorsFromUnsafePointer(ptr unsafe.Pointer) *TMonthCalColors {
    m := new(TMonthCalColors)
    m.instance = uintptr(ptr)
    m.ptr = ptr
    return m
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (m *TMonthCalColors) Free() {
    if m.instance != 0 {
        MonthCalColors_Free(m.instance)
        m.instance = 0
        m.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMonthCalColors) Instance() uintptr {
    return m.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMonthCalColors) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMonthCalColors) IsValid() bool {
    return m.instance != 0
}

// TMonthCalColorsClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMonthCalColorsClass() TClass {
    return MonthCalColors_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (m *TMonthCalColors) Assign(Source IObject) {
    MonthCalColors_Assign(m.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (m *TMonthCalColors) GetNamePath() string {
    return MonthCalColors_GetNamePath(m.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (m *TMonthCalColors) DisposeOf() {
    MonthCalColors_DisposeOf(m.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMonthCalColors) ClassType() TClass {
    return MonthCalColors_ClassType(m.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMonthCalColors) ClassName() string {
    return MonthCalColors_ClassName(m.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMonthCalColors) InstanceSize() int32 {
    return MonthCalColors_InstanceSize(m.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMonthCalColors) InheritsFrom(AClass TClass) bool {
    return MonthCalColors_InheritsFrom(m.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMonthCalColors) Equals(Obj IObject) bool {
    return MonthCalColors_Equals(m.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMonthCalColors) GetHashCode() int32 {
    return MonthCalColors_GetHashCode(m.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (m *TMonthCalColors) ToString() string {
    return MonthCalColors_ToString(m.instance)
}

// BackColor
func (m *TMonthCalColors) BackColor() TColor {
    return MonthCalColors_GetBackColor(m.instance)
}

// SetBackColor
func (m *TMonthCalColors) SetBackColor(value TColor) {
    MonthCalColors_SetBackColor(m.instance, value)
}

