
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

type TThumbBarButtonList struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewThumbBarButtonList
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewThumbBarButtonList() *TThumbBarButtonList {
    t := new(TThumbBarButtonList)
    t.instance = ThumbBarButtonList_Create()
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// ThumbBarButtonListFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ThumbBarButtonListFromInst(inst uintptr) *TThumbBarButtonList {
    t := new(TThumbBarButtonList)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// ThumbBarButtonListFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ThumbBarButtonListFromObj(obj IObject) *TThumbBarButtonList {
    t := new(TThumbBarButtonList)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// ThumbBarButtonListFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ThumbBarButtonListFromUnsafePointer(ptr unsafe.Pointer) *TThumbBarButtonList {
    t := new(TThumbBarButtonList)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TThumbBarButtonList) Free() {
    if t.instance != 0 {
        ThumbBarButtonList_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TThumbBarButtonList) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TThumbBarButtonList) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TThumbBarButtonList) IsValid() bool {
    return t.instance != 0
}

// TThumbBarButtonListClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TThumbBarButtonListClass() TClass {
    return ThumbBarButtonList_StaticClassType()
}

// Add
func (t *TThumbBarButtonList) Add() *TThumbBarButton {
    return ThumbBarButtonFromInst(ThumbBarButtonList_Add(t.instance))
}

// Owner
// CN: 组件所有者。
// EN: component owner.
func (t *TThumbBarButtonList) Owner() *TObject {
    return ObjectFromInst(ThumbBarButtonList_Owner(t.instance))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TThumbBarButtonList) Assign(Source IObject) {
    ThumbBarButtonList_Assign(t.instance, CheckPtr(Source))
}

// BeginUpdate
func (t *TThumbBarButtonList) BeginUpdate() {
    ThumbBarButtonList_BeginUpdate(t.instance)
}

// Clear
// CN: 清除。
// EN: .
func (t *TThumbBarButtonList) Clear() {
    ThumbBarButtonList_Clear(t.instance)
}

// ClearAndResetID
func (t *TThumbBarButtonList) ClearAndResetID() {
    ThumbBarButtonList_ClearAndResetID(t.instance)
}

// Delete
func (t *TThumbBarButtonList) Delete(Index int32) {
    ThumbBarButtonList_Delete(t.instance, Index)
}

// EndUpdate
func (t *TThumbBarButtonList) EndUpdate() {
    ThumbBarButtonList_EndUpdate(t.instance)
}

// FindItemID
func (t *TThumbBarButtonList) FindItemID(ID int32) *TCollectionItem {
    return CollectionItemFromInst(ThumbBarButtonList_FindItemID(t.instance, ID))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TThumbBarButtonList) GetNamePath() string {
    return ThumbBarButtonList_GetNamePath(t.instance)
}

// Insert
func (t *TThumbBarButtonList) Insert(Index int32) *TCollectionItem {
    return CollectionItemFromInst(ThumbBarButtonList_Insert(t.instance, Index))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TThumbBarButtonList) DisposeOf() {
    ThumbBarButtonList_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TThumbBarButtonList) ClassType() TClass {
    return ThumbBarButtonList_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TThumbBarButtonList) ClassName() string {
    return ThumbBarButtonList_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TThumbBarButtonList) InstanceSize() int32 {
    return ThumbBarButtonList_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TThumbBarButtonList) InheritsFrom(AClass TClass) bool {
    return ThumbBarButtonList_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TThumbBarButtonList) Equals(Obj IObject) bool {
    return ThumbBarButtonList_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TThumbBarButtonList) GetHashCode() int32 {
    return ThumbBarButtonList_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TThumbBarButtonList) ToString() string {
    return ThumbBarButtonList_ToString(t.instance)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (t *TThumbBarButtonList) SetOnChange(fn TNotifyEvent) {
    ThumbBarButtonList_SetOnChange(t.instance, fn)
}

// Capacity
func (t *TThumbBarButtonList) Capacity() int32 {
    return ThumbBarButtonList_GetCapacity(t.instance)
}

// SetCapacity
func (t *TThumbBarButtonList) SetCapacity(value int32) {
    ThumbBarButtonList_SetCapacity(t.instance, value)
}

// Count
func (t *TThumbBarButtonList) Count() int32 {
    return ThumbBarButtonList_GetCount(t.instance)
}

// Items
func (t *TThumbBarButtonList) Items(Index int32) *TThumbBarButton {
    return ThumbBarButtonFromInst(ThumbBarButtonList_GetItems(t.instance, Index))
}

// Items
func (t *TThumbBarButtonList) SetItems(Index int32, value *TThumbBarButton) {
    ThumbBarButtonList_SetItems(t.instance, Index, CheckPtr(value))
}

