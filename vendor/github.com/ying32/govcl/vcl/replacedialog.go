
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

type TReplaceDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewReplaceDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewReplaceDialog(owner IComponent) *TReplaceDialog {
    r := new(TReplaceDialog)
    r.instance = ReplaceDialog_Create(CheckPtr(owner))
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// ReplaceDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ReplaceDialogFromInst(inst uintptr) *TReplaceDialog {
    r := new(TReplaceDialog)
    r.instance = inst
    r.ptr = unsafe.Pointer(inst)
    return r
}

// ReplaceDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ReplaceDialogFromObj(obj IObject) *TReplaceDialog {
    r := new(TReplaceDialog)
    r.instance = CheckPtr(obj)
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// ReplaceDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ReplaceDialogFromUnsafePointer(ptr unsafe.Pointer) *TReplaceDialog {
    r := new(TReplaceDialog)
    r.instance = uintptr(ptr)
    r.ptr = ptr
    return r
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (r *TReplaceDialog) Free() {
    if r.instance != 0 {
        ReplaceDialog_Free(r.instance)
        r.instance = 0
        r.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (r *TReplaceDialog) Instance() uintptr {
    return r.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (r *TReplaceDialog) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (r *TReplaceDialog) IsValid() bool {
    return r.instance != 0
}

// TReplaceDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TReplaceDialogClass() TClass {
    return ReplaceDialog_StaticClassType()
}

// CloseDialog
func (r *TReplaceDialog) CloseDialog() {
    ReplaceDialog_CloseDialog(r.instance)
}

// Execute
// CN: 执行。
// EN: .
func (r *TReplaceDialog) Execute() bool {
    return ReplaceDialog_Execute(r.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (r *TReplaceDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ReplaceDialog_FindComponent(r.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (r *TReplaceDialog) GetNamePath() string {
    return ReplaceDialog_GetNamePath(r.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (r *TReplaceDialog) HasParent() bool {
    return ReplaceDialog_HasParent(r.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (r *TReplaceDialog) Assign(Source IObject) {
    ReplaceDialog_Assign(r.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (r *TReplaceDialog) DisposeOf() {
    ReplaceDialog_DisposeOf(r.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (r *TReplaceDialog) ClassType() TClass {
    return ReplaceDialog_ClassType(r.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (r *TReplaceDialog) ClassName() string {
    return ReplaceDialog_ClassName(r.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (r *TReplaceDialog) InstanceSize() int32 {
    return ReplaceDialog_InstanceSize(r.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (r *TReplaceDialog) InheritsFrom(AClass TClass) bool {
    return ReplaceDialog_InheritsFrom(r.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (r *TReplaceDialog) Equals(Obj IObject) bool {
    return ReplaceDialog_Equals(r.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (r *TReplaceDialog) GetHashCode() int32 {
    return ReplaceDialog_GetHashCode(r.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (r *TReplaceDialog) ToString() string {
    return ReplaceDialog_ToString(r.instance)
}

// ReplaceText
func (r *TReplaceDialog) ReplaceText() string {
    return ReplaceDialog_GetReplaceText(r.instance)
}

// SetReplaceText
func (r *TReplaceDialog) SetReplaceText(value string) {
    ReplaceDialog_SetReplaceText(r.instance, value)
}

// SetOnReplace
func (r *TReplaceDialog) SetOnReplace(fn TNotifyEvent) {
    ReplaceDialog_SetOnReplace(r.instance, fn)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (r *TReplaceDialog) Left() int32 {
    return ReplaceDialog_GetLeft(r.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (r *TReplaceDialog) SetLeft(value int32) {
    ReplaceDialog_SetLeft(r.instance, value)
}

// Position
func (r *TReplaceDialog) Position() TPoint {
    return ReplaceDialog_GetPosition(r.instance)
}

// SetPosition
func (r *TReplaceDialog) SetPosition(value TPoint) {
    ReplaceDialog_SetPosition(r.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (r *TReplaceDialog) Top() int32 {
    return ReplaceDialog_GetTop(r.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (r *TReplaceDialog) SetTop(value int32) {
    ReplaceDialog_SetTop(r.instance, value)
}

// FindText
func (r *TReplaceDialog) FindText() string {
    return ReplaceDialog_GetFindText(r.instance)
}

// SetFindText
func (r *TReplaceDialog) SetFindText(value string) {
    ReplaceDialog_SetFindText(r.instance, value)
}

// Options
func (r *TReplaceDialog) Options() TFindOptions {
    return ReplaceDialog_GetOptions(r.instance)
}

// SetOptions
func (r *TReplaceDialog) SetOptions(value TFindOptions) {
    ReplaceDialog_SetOptions(r.instance, value)
}

// SetOnFind
func (r *TReplaceDialog) SetOnFind(fn TNotifyEvent) {
    ReplaceDialog_SetOnFind(r.instance, fn)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (r *TReplaceDialog) Handle() HWND {
    return ReplaceDialog_GetHandle(r.instance)
}

// SetOnClose
func (r *TReplaceDialog) SetOnClose(fn TNotifyEvent) {
    ReplaceDialog_SetOnClose(r.instance, fn)
}

// SetOnShow
// CN: 设置显示事件。
// EN: .
func (r *TReplaceDialog) SetOnShow(fn TNotifyEvent) {
    ReplaceDialog_SetOnShow(r.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (r *TReplaceDialog) ComponentCount() int32 {
    return ReplaceDialog_GetComponentCount(r.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (r *TReplaceDialog) ComponentIndex() int32 {
    return ReplaceDialog_GetComponentIndex(r.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (r *TReplaceDialog) SetComponentIndex(value int32) {
    ReplaceDialog_SetComponentIndex(r.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (r *TReplaceDialog) Owner() *TComponent {
    return ComponentFromInst(ReplaceDialog_GetOwner(r.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (r *TReplaceDialog) Name() string {
    return ReplaceDialog_GetName(r.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (r *TReplaceDialog) SetName(value string) {
    ReplaceDialog_SetName(r.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (r *TReplaceDialog) Tag() int {
    return ReplaceDialog_GetTag(r.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (r *TReplaceDialog) SetTag(value int) {
    ReplaceDialog_SetTag(r.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (r *TReplaceDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ReplaceDialog_GetComponents(r.instance, AIndex))
}

