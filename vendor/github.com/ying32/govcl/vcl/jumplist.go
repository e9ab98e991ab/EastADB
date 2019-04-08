
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

type TJumpList struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewJumpList
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewJumpList(owner IComponent) *TJumpList {
    j := new(TJumpList)
    j.instance = JumpList_Create(CheckPtr(owner))
    j.ptr = unsafe.Pointer(j.instance)
    return j
}

// JumpListFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func JumpListFromInst(inst uintptr) *TJumpList {
    j := new(TJumpList)
    j.instance = inst
    j.ptr = unsafe.Pointer(inst)
    return j
}

// JumpListFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func JumpListFromObj(obj IObject) *TJumpList {
    j := new(TJumpList)
    j.instance = CheckPtr(obj)
    j.ptr = unsafe.Pointer(j.instance)
    return j
}

// JumpListFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func JumpListFromUnsafePointer(ptr unsafe.Pointer) *TJumpList {
    j := new(TJumpList)
    j.instance = uintptr(ptr)
    j.ptr = ptr
    return j
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (j *TJumpList) Free() {
    if j.instance != 0 {
        JumpList_Free(j.instance)
        j.instance = 0
        j.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (j *TJumpList) Instance() uintptr {
    return j.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (j *TJumpList) UnsafeAddr() unsafe.Pointer {
    return j.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (j *TJumpList) IsValid() bool {
    return j.instance != 0
}

// TJumpListClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TJumpListClass() TClass {
    return JumpList_StaticClassType()
}

// AddToRecent
func (j *TJumpList) AddToRecent(Path string) {
    JumpList_AddToRecent(j.instance, Path)
}

// RemoveFromRecent
func (j *TJumpList) RemoveFromRecent(Path string, AppModelID string) bool {
    return JumpList_RemoveFromRecent(j.instance, Path , AppModelID)
}

// RemoveAllFromRecent
func (j *TJumpList) RemoveAllFromRecent(AppModelID string) bool {
    return JumpList_RemoveAllFromRecent(j.instance, AppModelID)
}

// AddCategory
func (j *TJumpList) AddCategory(CategoryName string) int32 {
    return JumpList_AddCategory(j.instance, CategoryName)
}

// AddTask
func (j *TJumpList) AddTask(FriendlyName string, Path string, Arguments string, Icon string) *TJumpListItem {
    return JumpListItemFromInst(JumpList_AddTask(j.instance, FriendlyName , Path , Arguments , Icon))
}

// AddTaskSeparator
func (j *TJumpList) AddTaskSeparator() *TJumpListItem {
    return JumpListItemFromInst(JumpList_AddTaskSeparator(j.instance))
}

// AddItemToCategory
func (j *TJumpList) AddItemToCategory(CategoryIndex int32, FriendlyName string, Path string, Arguments string, Icon string) *TJumpListItem {
    return JumpListItemFromInst(JumpList_AddItemToCategory(j.instance, CategoryIndex , FriendlyName , Path , Arguments , Icon))
}

// UpdateList
func (j *TJumpList) UpdateList() bool {
    return JumpList_UpdateList(j.instance)
}

// DeleteList
func (j *TJumpList) DeleteList() bool {
    return JumpList_DeleteList(j.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (j *TJumpList) FindComponent(AName string) *TComponent {
    return ComponentFromInst(JumpList_FindComponent(j.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (j *TJumpList) GetNamePath() string {
    return JumpList_GetNamePath(j.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (j *TJumpList) HasParent() bool {
    return JumpList_HasParent(j.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (j *TJumpList) Assign(Source IObject) {
    JumpList_Assign(j.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (j *TJumpList) DisposeOf() {
    JumpList_DisposeOf(j.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (j *TJumpList) ClassType() TClass {
    return JumpList_ClassType(j.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (j *TJumpList) ClassName() string {
    return JumpList_ClassName(j.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (j *TJumpList) InstanceSize() int32 {
    return JumpList_InstanceSize(j.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (j *TJumpList) InheritsFrom(AClass TClass) bool {
    return JumpList_InheritsFrom(j.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (j *TJumpList) Equals(Obj IObject) bool {
    return JumpList_Equals(j.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (j *TJumpList) GetHashCode() int32 {
    return JumpList_GetHashCode(j.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (j *TJumpList) ToString() string {
    return JumpList_ToString(j.instance)
}

// AutoRefresh
func (j *TJumpList) AutoRefresh() bool {
    return JumpList_GetAutoRefresh(j.instance)
}

// SetAutoRefresh
func (j *TJumpList) SetAutoRefresh(value bool) {
    JumpList_SetAutoRefresh(j.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (j *TJumpList) Enabled() bool {
    return JumpList_GetEnabled(j.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (j *TJumpList) SetEnabled(value bool) {
    JumpList_SetEnabled(j.instance, value)
}

// ApplicationID
func (j *TJumpList) ApplicationID() string {
    return JumpList_GetApplicationID(j.instance)
}

// SetApplicationID
func (j *TJumpList) SetApplicationID(value string) {
    JumpList_SetApplicationID(j.instance, value)
}

// CustomCategories
func (j *TJumpList) CustomCategories() *TJumpCategories {
    return JumpCategoriesFromInst(JumpList_GetCustomCategories(j.instance))
}

// SetCustomCategories
func (j *TJumpList) SetCustomCategories(value *TJumpCategories) {
    JumpList_SetCustomCategories(j.instance, CheckPtr(value))
}

// ShowRecent
func (j *TJumpList) ShowRecent() bool {
    return JumpList_GetShowRecent(j.instance)
}

// SetShowRecent
func (j *TJumpList) SetShowRecent(value bool) {
    JumpList_SetShowRecent(j.instance, value)
}

// ShowFrequent
func (j *TJumpList) ShowFrequent() bool {
    return JumpList_GetShowFrequent(j.instance)
}

// SetShowFrequent
func (j *TJumpList) SetShowFrequent(value bool) {
    JumpList_SetShowFrequent(j.instance, value)
}

// TaskList
func (j *TJumpList) TaskList() *TJumpListCollection {
    return JumpListCollectionFromInst(JumpList_GetTaskList(j.instance))
}

// SetTaskList
func (j *TJumpList) SetTaskList(value *TJumpListCollection) {
    JumpList_SetTaskList(j.instance, CheckPtr(value))
}

// SetOnItemDeleted
func (j *TJumpList) SetOnItemDeleted(fn TItemDeletedByUserEvent) {
    JumpList_SetOnItemDeleted(j.instance, fn)
}

// SetOnListUpdateError
func (j *TJumpList) SetOnListUpdateError(fn TCreatingListErrorEvent) {
    JumpList_SetOnListUpdateError(j.instance, fn)
}

// SetOnItemsLoaded
func (j *TJumpList) SetOnItemsLoaded(fn TNotifyEvent) {
    JumpList_SetOnItemsLoaded(j.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (j *TJumpList) ComponentCount() int32 {
    return JumpList_GetComponentCount(j.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (j *TJumpList) ComponentIndex() int32 {
    return JumpList_GetComponentIndex(j.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (j *TJumpList) SetComponentIndex(value int32) {
    JumpList_SetComponentIndex(j.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (j *TJumpList) Owner() *TComponent {
    return ComponentFromInst(JumpList_GetOwner(j.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (j *TJumpList) Name() string {
    return JumpList_GetName(j.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (j *TJumpList) SetName(value string) {
    JumpList_SetName(j.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (j *TJumpList) Tag() int {
    return JumpList_GetTag(j.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (j *TJumpList) SetTag(value int) {
    JumpList_SetTag(j.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (j *TJumpList) Components(AIndex int32) *TComponent {
    return ComponentFromInst(JumpList_GetComponents(j.instance, AIndex))
}

