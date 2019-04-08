
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

type TList struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewList
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewList() *TList {
    l := new(TList)
    l.instance = List_Create()
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ListFromInst(inst uintptr) *TList {
    l := new(TList)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// ListFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ListFromObj(obj IObject) *TList {
    l := new(TList)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ListFromUnsafePointer(ptr unsafe.Pointer) *TList {
    l := new(TList)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TList) Free() {
    if l.instance != 0 {
        List_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TList) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TList) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TList) IsValid() bool {
    return l.instance != 0
}

// TListClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListClass() TClass {
    return List_StaticClassType()
}

// Add
func (l *TList) Add(Item uintptr) int32 {
    return List_Add(l.instance, Item)
}

// Clear
// CN: 清除。
// EN: .
func (l *TList) Clear() {
    List_Clear(l.instance)
}

// Delete
func (l *TList) Delete(Index int32) {
    List_Delete(l.instance, Index)
}

// Expand
func (l *TList) Expand() *TList {
    return ListFromInst(List_Expand(l.instance))
}

// IndexOf
func (l *TList) IndexOf(Item uintptr) int32 {
    return List_IndexOf(l.instance, Item)
}

// Insert
func (l *TList) Insert(Index int32, Item uintptr) {
    List_Insert(l.instance, Index , Item)
}

// Move
func (l *TList) Move(CurIndex int32, NewIndex int32) {
    List_Move(l.instance, CurIndex , NewIndex)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TList) DisposeOf() {
    List_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TList) ClassType() TClass {
    return List_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TList) ClassName() string {
    return List_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TList) InstanceSize() int32 {
    return List_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TList) InheritsFrom(AClass TClass) bool {
    return List_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TList) Equals(Obj IObject) bool {
    return List_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TList) GetHashCode() int32 {
    return List_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TList) ToString() string {
    return List_ToString(l.instance)
}

// Capacity
func (l *TList) Capacity() int32 {
    return List_GetCapacity(l.instance)
}

// SetCapacity
func (l *TList) SetCapacity(value int32) {
    List_SetCapacity(l.instance, value)
}

// Count
// CN: 获取项目总数。
// EN: Get Total number of projects.
func (l *TList) Count() int32 {
    return List_GetCount(l.instance)
}

// SetCount
// CN: 设置项目总数。
// EN: Set Total number of projects.
func (l *TList) SetCount(value int32) {
    List_SetCount(l.instance, value)
}

// List
// CN: 获取获取列表指针。
// EN: Get Get list pointer.
func (l *TList) List() uintptr {
    return List_GetList(l.instance)
}

// Items
// CN: 获取获取指定索引项目。
// EN: Get Get the specified index item.
func (l *TList) Items(Index int32) uintptr {
    return List_GetItems(l.instance, Index)
}

// Items
// CN: 设置获取指定索引项目。
// EN: Set Get the specified index item.
func (l *TList) SetItems(Index int32, value uintptr) {
    List_SetItems(l.instance, Index, value)
}

