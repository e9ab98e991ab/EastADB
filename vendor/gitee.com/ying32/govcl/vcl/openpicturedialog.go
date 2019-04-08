
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

type TOpenPictureDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewOpenPictureDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewOpenPictureDialog(owner IComponent) *TOpenPictureDialog {
    o := new(TOpenPictureDialog)
    o.instance = OpenPictureDialog_Create(CheckPtr(owner))
    o.ptr = unsafe.Pointer(o.instance)
    return o
}

// OpenPictureDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func OpenPictureDialogFromInst(inst uintptr) *TOpenPictureDialog {
    o := new(TOpenPictureDialog)
    o.instance = inst
    o.ptr = unsafe.Pointer(inst)
    return o
}

// OpenPictureDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func OpenPictureDialogFromObj(obj IObject) *TOpenPictureDialog {
    o := new(TOpenPictureDialog)
    o.instance = CheckPtr(obj)
    o.ptr = unsafe.Pointer(o.instance)
    return o
}

// OpenPictureDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func OpenPictureDialogFromUnsafePointer(ptr unsafe.Pointer) *TOpenPictureDialog {
    o := new(TOpenPictureDialog)
    o.instance = uintptr(ptr)
    o.ptr = ptr
    return o
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (o *TOpenPictureDialog) Free() {
    if o.instance != 0 {
        OpenPictureDialog_Free(o.instance)
        o.instance = 0
        o.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (o *TOpenPictureDialog) Instance() uintptr {
    return o.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (o *TOpenPictureDialog) UnsafeAddr() unsafe.Pointer {
    return o.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (o *TOpenPictureDialog) IsValid() bool {
    return o.instance != 0
}

// TOpenPictureDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TOpenPictureDialogClass() TClass {
    return OpenPictureDialog_StaticClassType()
}

// Execute
// CN: 执行。
// EN: .
func (o *TOpenPictureDialog) Execute() bool {
    return OpenPictureDialog_Execute(o.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (o *TOpenPictureDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(OpenPictureDialog_FindComponent(o.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (o *TOpenPictureDialog) GetNamePath() string {
    return OpenPictureDialog_GetNamePath(o.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (o *TOpenPictureDialog) HasParent() bool {
    return OpenPictureDialog_HasParent(o.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (o *TOpenPictureDialog) Assign(Source IObject) {
    OpenPictureDialog_Assign(o.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (o *TOpenPictureDialog) DisposeOf() {
    OpenPictureDialog_DisposeOf(o.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (o *TOpenPictureDialog) ClassType() TClass {
    return OpenPictureDialog_ClassType(o.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (o *TOpenPictureDialog) ClassName() string {
    return OpenPictureDialog_ClassName(o.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (o *TOpenPictureDialog) InstanceSize() int32 {
    return OpenPictureDialog_InstanceSize(o.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (o *TOpenPictureDialog) InheritsFrom(AClass TClass) bool {
    return OpenPictureDialog_InheritsFrom(o.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (o *TOpenPictureDialog) Equals(Obj IObject) bool {
    return OpenPictureDialog_Equals(o.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (o *TOpenPictureDialog) GetHashCode() int32 {
    return OpenPictureDialog_GetHashCode(o.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (o *TOpenPictureDialog) ToString() string {
    return OpenPictureDialog_ToString(o.instance)
}

// Filter
func (o *TOpenPictureDialog) Filter() string {
    return OpenPictureDialog_GetFilter(o.instance)
}

// SetFilter
func (o *TOpenPictureDialog) SetFilter(value string) {
    OpenPictureDialog_SetFilter(o.instance, value)
}

// Files
func (o *TOpenPictureDialog) Files() *TStrings {
    return StringsFromInst(OpenPictureDialog_GetFiles(o.instance))
}

// DefaultExt
func (o *TOpenPictureDialog) DefaultExt() string {
    return OpenPictureDialog_GetDefaultExt(o.instance)
}

// SetDefaultExt
func (o *TOpenPictureDialog) SetDefaultExt(value string) {
    OpenPictureDialog_SetDefaultExt(o.instance, value)
}

// FileName
func (o *TOpenPictureDialog) FileName() string {
    return OpenPictureDialog_GetFileName(o.instance)
}

// SetFileName
func (o *TOpenPictureDialog) SetFileName(value string) {
    OpenPictureDialog_SetFileName(o.instance, value)
}

// FilterIndex
func (o *TOpenPictureDialog) FilterIndex() int32 {
    return OpenPictureDialog_GetFilterIndex(o.instance)
}

// SetFilterIndex
func (o *TOpenPictureDialog) SetFilterIndex(value int32) {
    OpenPictureDialog_SetFilterIndex(o.instance, value)
}

// InitialDir
func (o *TOpenPictureDialog) InitialDir() string {
    return OpenPictureDialog_GetInitialDir(o.instance)
}

// SetInitialDir
func (o *TOpenPictureDialog) SetInitialDir(value string) {
    OpenPictureDialog_SetInitialDir(o.instance, value)
}

// Options
func (o *TOpenPictureDialog) Options() TOpenOptions {
    return OpenPictureDialog_GetOptions(o.instance)
}

// SetOptions
func (o *TOpenPictureDialog) SetOptions(value TOpenOptions) {
    OpenPictureDialog_SetOptions(o.instance, value)
}

// OptionsEx
func (o *TOpenPictureDialog) OptionsEx() TOpenOptionsEx {
    return OpenPictureDialog_GetOptionsEx(o.instance)
}

// SetOptionsEx
func (o *TOpenPictureDialog) SetOptionsEx(value TOpenOptionsEx) {
    OpenPictureDialog_SetOptionsEx(o.instance, value)
}

// Title
func (o *TOpenPictureDialog) Title() string {
    return OpenPictureDialog_GetTitle(o.instance)
}

// SetTitle
func (o *TOpenPictureDialog) SetTitle(value string) {
    OpenPictureDialog_SetTitle(o.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (o *TOpenPictureDialog) Handle() HWND {
    return OpenPictureDialog_GetHandle(o.instance)
}

// SetOnClose
func (o *TOpenPictureDialog) SetOnClose(fn TNotifyEvent) {
    OpenPictureDialog_SetOnClose(o.instance, fn)
}

// SetOnShow
// CN: 设置显示事件。
// EN: .
func (o *TOpenPictureDialog) SetOnShow(fn TNotifyEvent) {
    OpenPictureDialog_SetOnShow(o.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (o *TOpenPictureDialog) ComponentCount() int32 {
    return OpenPictureDialog_GetComponentCount(o.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (o *TOpenPictureDialog) ComponentIndex() int32 {
    return OpenPictureDialog_GetComponentIndex(o.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (o *TOpenPictureDialog) SetComponentIndex(value int32) {
    OpenPictureDialog_SetComponentIndex(o.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (o *TOpenPictureDialog) Owner() *TComponent {
    return ComponentFromInst(OpenPictureDialog_GetOwner(o.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (o *TOpenPictureDialog) Name() string {
    return OpenPictureDialog_GetName(o.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (o *TOpenPictureDialog) SetName(value string) {
    OpenPictureDialog_SetName(o.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (o *TOpenPictureDialog) Tag() int {
    return OpenPictureDialog_GetTag(o.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (o *TOpenPictureDialog) SetTag(value int) {
    OpenPictureDialog_SetTag(o.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (o *TOpenPictureDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(OpenPictureDialog_GetComponents(o.instance, AIndex))
}

