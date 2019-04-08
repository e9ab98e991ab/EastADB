
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

type TListItems struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewListItems
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewListItems() *TListItems {
    l := new(TListItems)
    l.instance = ListItems_Create()
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListItemsFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ListItemsFromInst(inst uintptr) *TListItems {
    l := new(TListItems)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// ListItemsFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ListItemsFromObj(obj IObject) *TListItems {
    l := new(TListItems)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListItemsFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ListItemsFromUnsafePointer(ptr unsafe.Pointer) *TListItems {
    l := new(TListItems)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TListItems) Free() {
    if l.instance != 0 {
        ListItems_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TListItems) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TListItems) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TListItems) IsValid() bool {
    return l.instance != 0
}

// TListItemsClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListItemsClass() TClass {
    return ListItems_StaticClassType()
}

// Add
func (l *TListItems) Add() *TListItem {
    return ListItemFromInst(ListItems_Add(l.instance))
}

// AddItem
func (l *TListItems) AddItem(Item *TListItem, Index int32) *TListItem {
    return ListItemFromInst(ListItems_AddItem(l.instance, CheckPtr(Item), Index))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TListItems) Assign(Source IObject) {
    ListItems_Assign(l.instance, CheckPtr(Source))
}

// BeginUpdate
func (l *TListItems) BeginUpdate() {
    ListItems_BeginUpdate(l.instance)
}

// Clear
// CN: 清除。
// EN: .
func (l *TListItems) Clear() {
    ListItems_Clear(l.instance)
}

// Delete
func (l *TListItems) Delete(Index int32) {
    ListItems_Delete(l.instance, Index)
}

// EndUpdate
func (l *TListItems) EndUpdate() {
    ListItems_EndUpdate(l.instance)
}

// IndexOf
func (l *TListItems) IndexOf(Value *TListItem) int32 {
    return ListItems_IndexOf(l.instance, CheckPtr(Value))
}

// Insert
func (l *TListItems) Insert(Index int32) *TListItem {
    return ListItemFromInst(ListItems_Insert(l.instance, Index))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TListItems) GetNamePath() string {
    return ListItems_GetNamePath(l.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TListItems) DisposeOf() {
    ListItems_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TListItems) ClassType() TClass {
    return ListItems_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TListItems) ClassName() string {
    return ListItems_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TListItems) InstanceSize() int32 {
    return ListItems_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TListItems) InheritsFrom(AClass TClass) bool {
    return ListItems_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TListItems) Equals(Obj IObject) bool {
    return ListItems_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TListItems) GetHashCode() int32 {
    return ListItems_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TListItems) ToString() string {
    return ListItems_ToString(l.instance)
}

// Count
func (l *TListItems) Count() int32 {
    return ListItems_GetCount(l.instance)
}

// SetCount
func (l *TListItems) SetCount(value int32) {
    ListItems_SetCount(l.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (l *TListItems) Handle() HWND {
    return ListItems_GetHandle(l.instance)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TListItems) Owner() *TWinControl {
    return WinControlFromInst(ListItems_GetOwner(l.instance))
}

// Item
func (l *TListItems) Item(Index int32) *TListItem {
    return ListItemFromInst(ListItems_GetItem(l.instance, Index))
}

// Item
func (l *TListItems) SetItem(Index int32, value *TListItem) {
    ListItems_SetItem(l.instance, Index, CheckPtr(value))
}

