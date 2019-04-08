
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

type TActionList struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewActionList
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewActionList(owner IComponent) *TActionList {
    a := new(TActionList)
    a.instance = ActionList_Create(CheckPtr(owner))
    a.ptr = unsafe.Pointer(a.instance)
    return a
}

// ActionListFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ActionListFromInst(inst uintptr) *TActionList {
    a := new(TActionList)
    a.instance = inst
    a.ptr = unsafe.Pointer(inst)
    return a
}

// ActionListFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ActionListFromObj(obj IObject) *TActionList {
    a := new(TActionList)
    a.instance = CheckPtr(obj)
    a.ptr = unsafe.Pointer(a.instance)
    return a
}

// ActionListFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ActionListFromUnsafePointer(ptr unsafe.Pointer) *TActionList {
    a := new(TActionList)
    a.instance = uintptr(ptr)
    a.ptr = ptr
    return a
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (a *TActionList) Free() {
    if a.instance != 0 {
        ActionList_Free(a.instance)
        a.instance = 0
        a.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (a *TActionList) Instance() uintptr {
    return a.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (a *TActionList) UnsafeAddr() unsafe.Pointer {
    return a.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (a *TActionList) IsValid() bool {
    return a.instance != 0
}

// TActionListClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TActionListClass() TClass {
    return ActionList_StaticClassType()
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (a *TActionList) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ActionList_FindComponent(a.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (a *TActionList) GetNamePath() string {
    return ActionList_GetNamePath(a.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (a *TActionList) HasParent() bool {
    return ActionList_HasParent(a.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (a *TActionList) Assign(Source IObject) {
    ActionList_Assign(a.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (a *TActionList) DisposeOf() {
    ActionList_DisposeOf(a.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (a *TActionList) ClassType() TClass {
    return ActionList_ClassType(a.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (a *TActionList) ClassName() string {
    return ActionList_ClassName(a.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (a *TActionList) InstanceSize() int32 {
    return ActionList_InstanceSize(a.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (a *TActionList) InheritsFrom(AClass TClass) bool {
    return ActionList_InheritsFrom(a.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (a *TActionList) Equals(Obj IObject) bool {
    return ActionList_Equals(a.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (a *TActionList) GetHashCode() int32 {
    return ActionList_GetHashCode(a.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (a *TActionList) ToString() string {
    return ActionList_ToString(a.instance)
}

// Images
// CN: 获取图标索引列表对象。
// EN: .
func (a *TActionList) Images() *TImageList {
    return ImageListFromInst(ActionList_GetImages(a.instance))
}

// SetImages
// CN: 设置图标索引列表对象。
// EN: .
func (a *TActionList) SetImages(value IComponent) {
    ActionList_SetImages(a.instance, CheckPtr(value))
}

// State
func (a *TActionList) State() TActionListState {
    return ActionList_GetState(a.instance)
}

// SetState
func (a *TActionList) SetState(value TActionListState) {
    ActionList_SetState(a.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (a *TActionList) SetOnChange(fn TNotifyEvent) {
    ActionList_SetOnChange(a.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (a *TActionList) ComponentCount() int32 {
    return ActionList_GetComponentCount(a.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (a *TActionList) ComponentIndex() int32 {
    return ActionList_GetComponentIndex(a.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (a *TActionList) SetComponentIndex(value int32) {
    ActionList_SetComponentIndex(a.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (a *TActionList) Owner() *TComponent {
    return ComponentFromInst(ActionList_GetOwner(a.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (a *TActionList) Name() string {
    return ActionList_GetName(a.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (a *TActionList) SetName(value string) {
    ActionList_SetName(a.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (a *TActionList) Tag() int {
    return ActionList_GetTag(a.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (a *TActionList) SetTag(value int) {
    ActionList_SetTag(a.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (a *TActionList) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ActionList_GetComponents(a.instance, AIndex))
}

