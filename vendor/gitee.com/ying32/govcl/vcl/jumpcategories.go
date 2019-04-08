
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

type TJumpCategories struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// JumpCategoriesFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func JumpCategoriesFromInst(inst uintptr) *TJumpCategories {
    j := new(TJumpCategories)
    j.instance = inst
    j.ptr = unsafe.Pointer(inst)
    return j
}

// JumpCategoriesFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func JumpCategoriesFromObj(obj IObject) *TJumpCategories {
    j := new(TJumpCategories)
    j.instance = CheckPtr(obj)
    j.ptr = unsafe.Pointer(j.instance)
    return j
}

// JumpCategoriesFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func JumpCategoriesFromUnsafePointer(ptr unsafe.Pointer) *TJumpCategories {
    j := new(TJumpCategories)
    j.instance = uintptr(ptr)
    j.ptr = ptr
    return j
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (j *TJumpCategories) Instance() uintptr {
    return j.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (j *TJumpCategories) UnsafeAddr() unsafe.Pointer {
    return j.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (j *TJumpCategories) IsValid() bool {
    return j.instance != 0
}

// TJumpCategoriesClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TJumpCategoriesClass() TClass {
    return JumpCategories_StaticClassType()
}

// GetCategoryIndex
func (j *TJumpCategories) GetCategoryIndex(CategoryName string) int32 {
    return JumpCategories_GetCategoryIndex(j.instance, CategoryName)
}

// Owner
// CN: 组件所有者。
// EN: component owner.
func (j *TJumpCategories) Owner() *TObject {
    return ObjectFromInst(JumpCategories_Owner(j.instance))
}

// Add
func (j *TJumpCategories) Add() *TCollectionItem {
    return CollectionItemFromInst(JumpCategories_Add(j.instance))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (j *TJumpCategories) Assign(Source IObject) {
    JumpCategories_Assign(j.instance, CheckPtr(Source))
}

// BeginUpdate
func (j *TJumpCategories) BeginUpdate() {
    JumpCategories_BeginUpdate(j.instance)
}

// Clear
// CN: 清除。
// EN: .
func (j *TJumpCategories) Clear() {
    JumpCategories_Clear(j.instance)
}

// ClearAndResetID
func (j *TJumpCategories) ClearAndResetID() {
    JumpCategories_ClearAndResetID(j.instance)
}

// Delete
func (j *TJumpCategories) Delete(Index int32) {
    JumpCategories_Delete(j.instance, Index)
}

// EndUpdate
func (j *TJumpCategories) EndUpdate() {
    JumpCategories_EndUpdate(j.instance)
}

// FindItemID
func (j *TJumpCategories) FindItemID(ID int32) *TCollectionItem {
    return CollectionItemFromInst(JumpCategories_FindItemID(j.instance, ID))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (j *TJumpCategories) GetNamePath() string {
    return JumpCategories_GetNamePath(j.instance)
}

// Insert
func (j *TJumpCategories) Insert(Index int32) *TCollectionItem {
    return CollectionItemFromInst(JumpCategories_Insert(j.instance, Index))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (j *TJumpCategories) DisposeOf() {
    JumpCategories_DisposeOf(j.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (j *TJumpCategories) ClassType() TClass {
    return JumpCategories_ClassType(j.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (j *TJumpCategories) ClassName() string {
    return JumpCategories_ClassName(j.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (j *TJumpCategories) InstanceSize() int32 {
    return JumpCategories_InstanceSize(j.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (j *TJumpCategories) InheritsFrom(AClass TClass) bool {
    return JumpCategories_InheritsFrom(j.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (j *TJumpCategories) Equals(Obj IObject) bool {
    return JumpCategories_Equals(j.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (j *TJumpCategories) GetHashCode() int32 {
    return JumpCategories_GetHashCode(j.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (j *TJumpCategories) ToString() string {
    return JumpCategories_ToString(j.instance)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (j *TJumpCategories) SetOnChange(fn TNotifyEvent) {
    JumpCategories_SetOnChange(j.instance, fn)
}

// Capacity
func (j *TJumpCategories) Capacity() int32 {
    return JumpCategories_GetCapacity(j.instance)
}

// SetCapacity
func (j *TJumpCategories) SetCapacity(value int32) {
    JumpCategories_SetCapacity(j.instance, value)
}

// Count
func (j *TJumpCategories) Count() int32 {
    return JumpCategories_GetCount(j.instance)
}

// Items
func (j *TJumpCategories) Items(Index int32) *TJumpCategoryItem {
    return JumpCategoryItemFromInst(JumpCategories_GetItems(j.instance, Index))
}

// Items
func (j *TJumpCategories) SetItems(Index int32, value *TJumpCategoryItem) {
    JumpCategories_SetItems(j.instance, Index, CheckPtr(value))
}

