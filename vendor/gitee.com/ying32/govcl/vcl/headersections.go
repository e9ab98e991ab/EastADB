
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

type THeaderSections struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewHeaderSections
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewHeaderSections() *THeaderSections {
    h := new(THeaderSections)
    h.instance = HeaderSections_Create()
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HeaderSectionsFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func HeaderSectionsFromInst(inst uintptr) *THeaderSections {
    h := new(THeaderSections)
    h.instance = inst
    h.ptr = unsafe.Pointer(inst)
    return h
}

// HeaderSectionsFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func HeaderSectionsFromObj(obj IObject) *THeaderSections {
    h := new(THeaderSections)
    h.instance = CheckPtr(obj)
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HeaderSectionsFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func HeaderSectionsFromUnsafePointer(ptr unsafe.Pointer) *THeaderSections {
    h := new(THeaderSections)
    h.instance = uintptr(ptr)
    h.ptr = ptr
    return h
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (h *THeaderSections) Free() {
    if h.instance != 0 {
        HeaderSections_Free(h.instance)
        h.instance = 0
        h.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (h *THeaderSections) Instance() uintptr {
    return h.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (h *THeaderSections) UnsafeAddr() unsafe.Pointer {
    return h.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (h *THeaderSections) IsValid() bool {
    return h.instance != 0
}

// THeaderSectionsClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func THeaderSectionsClass() TClass {
    return HeaderSections_StaticClassType()
}

// Add
func (h *THeaderSections) Add() *THeaderSection {
    return HeaderSectionFromInst(HeaderSections_Add(h.instance))
}

// AddItem
func (h *THeaderSections) AddItem(Item *THeaderSection, Index int32) *THeaderSection {
    return HeaderSectionFromInst(HeaderSections_AddItem(h.instance, CheckPtr(Item), Index))
}

// Insert
func (h *THeaderSections) Insert(Index int32) *THeaderSection {
    return HeaderSectionFromInst(HeaderSections_Insert(h.instance, Index))
}

// Owner
// CN: 组件所有者。
// EN: component owner.
func (h *THeaderSections) Owner() *TObject {
    return ObjectFromInst(HeaderSections_Owner(h.instance))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (h *THeaderSections) Assign(Source IObject) {
    HeaderSections_Assign(h.instance, CheckPtr(Source))
}

// BeginUpdate
func (h *THeaderSections) BeginUpdate() {
    HeaderSections_BeginUpdate(h.instance)
}

// Clear
// CN: 清除。
// EN: .
func (h *THeaderSections) Clear() {
    HeaderSections_Clear(h.instance)
}

// ClearAndResetID
func (h *THeaderSections) ClearAndResetID() {
    HeaderSections_ClearAndResetID(h.instance)
}

// Delete
func (h *THeaderSections) Delete(Index int32) {
    HeaderSections_Delete(h.instance, Index)
}

// EndUpdate
func (h *THeaderSections) EndUpdate() {
    HeaderSections_EndUpdate(h.instance)
}

// FindItemID
func (h *THeaderSections) FindItemID(ID int32) *TCollectionItem {
    return CollectionItemFromInst(HeaderSections_FindItemID(h.instance, ID))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (h *THeaderSections) GetNamePath() string {
    return HeaderSections_GetNamePath(h.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (h *THeaderSections) DisposeOf() {
    HeaderSections_DisposeOf(h.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (h *THeaderSections) ClassType() TClass {
    return HeaderSections_ClassType(h.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (h *THeaderSections) ClassName() string {
    return HeaderSections_ClassName(h.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (h *THeaderSections) InstanceSize() int32 {
    return HeaderSections_InstanceSize(h.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (h *THeaderSections) InheritsFrom(AClass TClass) bool {
    return HeaderSections_InheritsFrom(h.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (h *THeaderSections) Equals(Obj IObject) bool {
    return HeaderSections_Equals(h.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (h *THeaderSections) GetHashCode() int32 {
    return HeaderSections_GetHashCode(h.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (h *THeaderSections) ToString() string {
    return HeaderSections_ToString(h.instance)
}

// Capacity
func (h *THeaderSections) Capacity() int32 {
    return HeaderSections_GetCapacity(h.instance)
}

// SetCapacity
func (h *THeaderSections) SetCapacity(value int32) {
    HeaderSections_SetCapacity(h.instance, value)
}

// Count
func (h *THeaderSections) Count() int32 {
    return HeaderSections_GetCount(h.instance)
}

// Items
func (h *THeaderSections) Items(Index int32) *THeaderSection {
    return HeaderSectionFromInst(HeaderSections_GetItems(h.instance, Index))
}

// Items
func (h *THeaderSections) SetItems(Index int32, value *THeaderSection) {
    HeaderSections_SetItems(h.instance, Index, CheckPtr(value))
}

