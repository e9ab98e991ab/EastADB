
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

type TJumpListCollection struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// JumpListCollectionFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func JumpListCollectionFromInst(inst uintptr) *TJumpListCollection {
    j := new(TJumpListCollection)
    j.instance = inst
    j.ptr = unsafe.Pointer(inst)
    return j
}

// JumpListCollectionFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func JumpListCollectionFromObj(obj IObject) *TJumpListCollection {
    j := new(TJumpListCollection)
    j.instance = CheckPtr(obj)
    j.ptr = unsafe.Pointer(j.instance)
    return j
}

// JumpListCollectionFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func JumpListCollectionFromUnsafePointer(ptr unsafe.Pointer) *TJumpListCollection {
    j := new(TJumpListCollection)
    j.instance = uintptr(ptr)
    j.ptr = ptr
    return j
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (j *TJumpListCollection) Instance() uintptr {
    return j.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (j *TJumpListCollection) UnsafeAddr() unsafe.Pointer {
    return j.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (j *TJumpListCollection) IsValid() bool {
    return j.instance != 0
}

// TJumpListCollectionClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TJumpListCollectionClass() TClass {
    return JumpListCollection_StaticClassType()
}

// Owner
// CN: 组件所有者。
// EN: component owner.
func (j *TJumpListCollection) Owner() *TObject {
    return ObjectFromInst(JumpListCollection_Owner(j.instance))
}

// Add
func (j *TJumpListCollection) Add() *TCollectionItem {
    return CollectionItemFromInst(JumpListCollection_Add(j.instance))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (j *TJumpListCollection) Assign(Source IObject) {
    JumpListCollection_Assign(j.instance, CheckPtr(Source))
}

// BeginUpdate
func (j *TJumpListCollection) BeginUpdate() {
    JumpListCollection_BeginUpdate(j.instance)
}

// Clear
// CN: 清除。
// EN: .
func (j *TJumpListCollection) Clear() {
    JumpListCollection_Clear(j.instance)
}

// ClearAndResetID
func (j *TJumpListCollection) ClearAndResetID() {
    JumpListCollection_ClearAndResetID(j.instance)
}

// Delete
func (j *TJumpListCollection) Delete(Index int32) {
    JumpListCollection_Delete(j.instance, Index)
}

// EndUpdate
func (j *TJumpListCollection) EndUpdate() {
    JumpListCollection_EndUpdate(j.instance)
}

// FindItemID
func (j *TJumpListCollection) FindItemID(ID int32) *TCollectionItem {
    return CollectionItemFromInst(JumpListCollection_FindItemID(j.instance, ID))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (j *TJumpListCollection) GetNamePath() string {
    return JumpListCollection_GetNamePath(j.instance)
}

// Insert
func (j *TJumpListCollection) Insert(Index int32) *TCollectionItem {
    return CollectionItemFromInst(JumpListCollection_Insert(j.instance, Index))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (j *TJumpListCollection) DisposeOf() {
    JumpListCollection_DisposeOf(j.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (j *TJumpListCollection) ClassType() TClass {
    return JumpListCollection_ClassType(j.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (j *TJumpListCollection) ClassName() string {
    return JumpListCollection_ClassName(j.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (j *TJumpListCollection) InstanceSize() int32 {
    return JumpListCollection_InstanceSize(j.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (j *TJumpListCollection) InheritsFrom(AClass TClass) bool {
    return JumpListCollection_InheritsFrom(j.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (j *TJumpListCollection) Equals(Obj IObject) bool {
    return JumpListCollection_Equals(j.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (j *TJumpListCollection) GetHashCode() int32 {
    return JumpListCollection_GetHashCode(j.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (j *TJumpListCollection) ToString() string {
    return JumpListCollection_ToString(j.instance)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (j *TJumpListCollection) SetOnChange(fn TNotifyEvent) {
    JumpListCollection_SetOnChange(j.instance, fn)
}

// Capacity
func (j *TJumpListCollection) Capacity() int32 {
    return JumpListCollection_GetCapacity(j.instance)
}

// SetCapacity
func (j *TJumpListCollection) SetCapacity(value int32) {
    JumpListCollection_SetCapacity(j.instance, value)
}

// Count
func (j *TJumpListCollection) Count() int32 {
    return JumpListCollection_GetCount(j.instance)
}

// Items
func (j *TJumpListCollection) Items(Index int32) *TJumpListItem {
    return JumpListItemFromInst(JumpListCollection_GetItems(j.instance, Index))
}

// Items
func (j *TJumpListCollection) SetItems(Index int32, value *TJumpListItem) {
    JumpListCollection_SetItems(j.instance, Index, CheckPtr(value))
}

