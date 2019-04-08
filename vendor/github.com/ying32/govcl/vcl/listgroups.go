
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

type TListGroups struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewListGroups
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewListGroups() *TListGroups {
    l := new(TListGroups)
    l.instance = ListGroups_Create()
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListGroupsFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ListGroupsFromInst(inst uintptr) *TListGroups {
    l := new(TListGroups)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// ListGroupsFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ListGroupsFromObj(obj IObject) *TListGroups {
    l := new(TListGroups)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListGroupsFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ListGroupsFromUnsafePointer(ptr unsafe.Pointer) *TListGroups {
    l := new(TListGroups)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TListGroups) Free() {
    if l.instance != 0 {
        ListGroups_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TListGroups) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TListGroups) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TListGroups) IsValid() bool {
    return l.instance != 0
}

// TListGroupsClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListGroupsClass() TClass {
    return ListGroups_StaticClassType()
}

// Add
func (l *TListGroups) Add() *TListGroup {
    return ListGroupFromInst(ListGroups_Add(l.instance))
}

// Owner
// CN: 组件所有者。
// EN: component owner.
func (l *TListGroups) Owner() *TWinControl {
    return WinControlFromInst(ListGroups_Owner(l.instance))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TListGroups) Assign(Source IObject) {
    ListGroups_Assign(l.instance, CheckPtr(Source))
}

// BeginUpdate
func (l *TListGroups) BeginUpdate() {
    ListGroups_BeginUpdate(l.instance)
}

// Clear
// CN: 清除。
// EN: .
func (l *TListGroups) Clear() {
    ListGroups_Clear(l.instance)
}

// ClearAndResetID
func (l *TListGroups) ClearAndResetID() {
    ListGroups_ClearAndResetID(l.instance)
}

// Delete
func (l *TListGroups) Delete(Index int32) {
    ListGroups_Delete(l.instance, Index)
}

// EndUpdate
func (l *TListGroups) EndUpdate() {
    ListGroups_EndUpdate(l.instance)
}

// FindItemID
func (l *TListGroups) FindItemID(ID int32) *TCollectionItem {
    return CollectionItemFromInst(ListGroups_FindItemID(l.instance, ID))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TListGroups) GetNamePath() string {
    return ListGroups_GetNamePath(l.instance)
}

// Insert
func (l *TListGroups) Insert(Index int32) *TCollectionItem {
    return CollectionItemFromInst(ListGroups_Insert(l.instance, Index))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TListGroups) DisposeOf() {
    ListGroups_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TListGroups) ClassType() TClass {
    return ListGroups_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TListGroups) ClassName() string {
    return ListGroups_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TListGroups) InstanceSize() int32 {
    return ListGroups_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TListGroups) InheritsFrom(AClass TClass) bool {
    return ListGroups_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TListGroups) Equals(Obj IObject) bool {
    return ListGroups_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TListGroups) GetHashCode() int32 {
    return ListGroups_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TListGroups) ToString() string {
    return ListGroups_ToString(l.instance)
}

// Capacity
func (l *TListGroups) Capacity() int32 {
    return ListGroups_GetCapacity(l.instance)
}

// SetCapacity
func (l *TListGroups) SetCapacity(value int32) {
    ListGroups_SetCapacity(l.instance, value)
}

// Count
func (l *TListGroups) Count() int32 {
    return ListGroups_GetCount(l.instance)
}

// Items
func (l *TListGroups) Items(Index int32) *TListGroup {
    return ListGroupFromInst(ListGroups_GetItems(l.instance, Index))
}

// Items
func (l *TListGroups) SetItems(Index int32, value *TListGroup) {
    ListGroups_SetItems(l.instance, Index, CheckPtr(value))
}

