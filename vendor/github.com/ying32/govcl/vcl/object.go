
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

type TObject struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewObject
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewObject() *TObject {
    o := new(TObject)
    o.instance = Object_Create()
    o.ptr = unsafe.Pointer(o.instance)
    return o
}

// ObjectFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ObjectFromInst(inst uintptr) *TObject {
    o := new(TObject)
    o.instance = inst
    o.ptr = unsafe.Pointer(inst)
    return o
}

// ObjectFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ObjectFromObj(obj IObject) *TObject {
    o := new(TObject)
    o.instance = CheckPtr(obj)
    o.ptr = unsafe.Pointer(o.instance)
    return o
}

// ObjectFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ObjectFromUnsafePointer(ptr unsafe.Pointer) *TObject {
    o := new(TObject)
    o.instance = uintptr(ptr)
    o.ptr = ptr
    return o
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (o *TObject) Free() {
    if o.instance != 0 {
        Object_Free(o.instance)
        o.instance = 0
        o.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (o *TObject) Instance() uintptr {
    return o.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (o *TObject) UnsafeAddr() unsafe.Pointer {
    return o.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (o *TObject) IsValid() bool {
    return o.instance != 0
}

// TObjectClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TObjectClass() TClass {
    return Object_StaticClassType()
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (o *TObject) DisposeOf() {
    Object_DisposeOf(o.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (o *TObject) ClassType() TClass {
    return Object_ClassType(o.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (o *TObject) ClassName() string {
    return Object_ClassName(o.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (o *TObject) InstanceSize() int32 {
    return Object_InstanceSize(o.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (o *TObject) InheritsFrom(AClass TClass) bool {
    return Object_InheritsFrom(o.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (o *TObject) Equals(Obj IObject) bool {
    return Object_Equals(o.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (o *TObject) GetHashCode() int32 {
    return Object_GetHashCode(o.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (o *TObject) ToString() string {
    return Object_ToString(o.instance)
}

