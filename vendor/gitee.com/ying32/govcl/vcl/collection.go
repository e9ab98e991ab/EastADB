
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

type TCollection struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCollection
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCollection() *TCollection {
    c := new(TCollection)
    c.instance = Collection_Create()
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CollectionFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CollectionFromInst(inst uintptr) *TCollection {
    c := new(TCollection)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CollectionFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CollectionFromObj(obj IObject) *TCollection {
    c := new(TCollection)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CollectionFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CollectionFromUnsafePointer(ptr unsafe.Pointer) *TCollection {
    c := new(TCollection)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCollection) Free() {
    if c.instance != 0 {
        Collection_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCollection) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCollection) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCollection) IsValid() bool {
    return c.instance != 0
}

// TCollectionClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCollectionClass() TClass {
    return Collection_StaticClassType()
}

// Owner
// CN: 组件所有者。
// EN: component owner.
func (c *TCollection) Owner() *TObject {
    return ObjectFromInst(Collection_Owner(c.instance))
}

// Add
func (c *TCollection) Add() *TCollectionItem {
    return CollectionItemFromInst(Collection_Add(c.instance))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TCollection) Assign(Source IObject) {
    Collection_Assign(c.instance, CheckPtr(Source))
}

// BeginUpdate
func (c *TCollection) BeginUpdate() {
    Collection_BeginUpdate(c.instance)
}

// Clear
// CN: 清除。
// EN: .
func (c *TCollection) Clear() {
    Collection_Clear(c.instance)
}

// ClearAndResetID
func (c *TCollection) ClearAndResetID() {
    Collection_ClearAndResetID(c.instance)
}

// Delete
func (c *TCollection) Delete(Index int32) {
    Collection_Delete(c.instance, Index)
}

// EndUpdate
func (c *TCollection) EndUpdate() {
    Collection_EndUpdate(c.instance)
}

// FindItemID
func (c *TCollection) FindItemID(ID int32) *TCollectionItem {
    return CollectionItemFromInst(Collection_FindItemID(c.instance, ID))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TCollection) GetNamePath() string {
    return Collection_GetNamePath(c.instance)
}

// Insert
func (c *TCollection) Insert(Index int32) *TCollectionItem {
    return CollectionItemFromInst(Collection_Insert(c.instance, Index))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCollection) DisposeOf() {
    Collection_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCollection) ClassType() TClass {
    return Collection_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCollection) ClassName() string {
    return Collection_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCollection) InstanceSize() int32 {
    return Collection_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCollection) InheritsFrom(AClass TClass) bool {
    return Collection_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCollection) Equals(Obj IObject) bool {
    return Collection_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCollection) GetHashCode() int32 {
    return Collection_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCollection) ToString() string {
    return Collection_ToString(c.instance)
}

// Capacity
func (c *TCollection) Capacity() int32 {
    return Collection_GetCapacity(c.instance)
}

// SetCapacity
func (c *TCollection) SetCapacity(value int32) {
    Collection_SetCapacity(c.instance, value)
}

// Count
func (c *TCollection) Count() int32 {
    return Collection_GetCount(c.instance)
}

// Items
func (c *TCollection) Items(Index int32) *TCollectionItem {
    return CollectionItemFromInst(Collection_GetItems(c.instance, Index))
}

// Items
func (c *TCollection) SetItems(Index int32, value *TCollectionItem) {
    Collection_SetItems(c.instance, Index, CheckPtr(value))
}

