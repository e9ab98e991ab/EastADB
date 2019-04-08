
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

type TFontDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewFontDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewFontDialog(owner IComponent) *TFontDialog {
    f := new(TFontDialog)
    f.instance = FontDialog_Create(CheckPtr(owner))
    f.ptr = unsafe.Pointer(f.instance)
    return f
}

// FontDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func FontDialogFromInst(inst uintptr) *TFontDialog {
    f := new(TFontDialog)
    f.instance = inst
    f.ptr = unsafe.Pointer(inst)
    return f
}

// FontDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func FontDialogFromObj(obj IObject) *TFontDialog {
    f := new(TFontDialog)
    f.instance = CheckPtr(obj)
    f.ptr = unsafe.Pointer(f.instance)
    return f
}

// FontDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func FontDialogFromUnsafePointer(ptr unsafe.Pointer) *TFontDialog {
    f := new(TFontDialog)
    f.instance = uintptr(ptr)
    f.ptr = ptr
    return f
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (f *TFontDialog) Free() {
    if f.instance != 0 {
        FontDialog_Free(f.instance)
        f.instance = 0
        f.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (f *TFontDialog) Instance() uintptr {
    return f.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (f *TFontDialog) UnsafeAddr() unsafe.Pointer {
    return f.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (f *TFontDialog) IsValid() bool {
    return f.instance != 0
}

// TFontDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TFontDialogClass() TClass {
    return FontDialog_StaticClassType()
}

// Execute
// CN: 执行。
// EN: .
func (f *TFontDialog) Execute() bool {
    return FontDialog_Execute(f.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (f *TFontDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(FontDialog_FindComponent(f.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (f *TFontDialog) GetNamePath() string {
    return FontDialog_GetNamePath(f.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (f *TFontDialog) HasParent() bool {
    return FontDialog_HasParent(f.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (f *TFontDialog) Assign(Source IObject) {
    FontDialog_Assign(f.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (f *TFontDialog) DisposeOf() {
    FontDialog_DisposeOf(f.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (f *TFontDialog) ClassType() TClass {
    return FontDialog_ClassType(f.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (f *TFontDialog) ClassName() string {
    return FontDialog_ClassName(f.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (f *TFontDialog) InstanceSize() int32 {
    return FontDialog_InstanceSize(f.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (f *TFontDialog) InheritsFrom(AClass TClass) bool {
    return FontDialog_InheritsFrom(f.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (f *TFontDialog) Equals(Obj IObject) bool {
    return FontDialog_Equals(f.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (f *TFontDialog) GetHashCode() int32 {
    return FontDialog_GetHashCode(f.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (f *TFontDialog) ToString() string {
    return FontDialog_ToString(f.instance)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (f *TFontDialog) Font() *TFont {
    return FontFromInst(FontDialog_GetFont(f.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (f *TFontDialog) SetFont(value *TFont) {
    FontDialog_SetFont(f.instance, CheckPtr(value))
}

// Options
func (f *TFontDialog) Options() TFontDialogOptions {
    return FontDialog_GetOptions(f.instance)
}

// SetOptions
func (f *TFontDialog) SetOptions(value TFontDialogOptions) {
    FontDialog_SetOptions(f.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (f *TFontDialog) Handle() HWND {
    return FontDialog_GetHandle(f.instance)
}

// SetOnClose
func (f *TFontDialog) SetOnClose(fn TNotifyEvent) {
    FontDialog_SetOnClose(f.instance, fn)
}

// SetOnShow
// CN: 设置显示事件。
// EN: .
func (f *TFontDialog) SetOnShow(fn TNotifyEvent) {
    FontDialog_SetOnShow(f.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (f *TFontDialog) ComponentCount() int32 {
    return FontDialog_GetComponentCount(f.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (f *TFontDialog) ComponentIndex() int32 {
    return FontDialog_GetComponentIndex(f.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (f *TFontDialog) SetComponentIndex(value int32) {
    FontDialog_SetComponentIndex(f.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (f *TFontDialog) Owner() *TComponent {
    return ComponentFromInst(FontDialog_GetOwner(f.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (f *TFontDialog) Name() string {
    return FontDialog_GetName(f.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (f *TFontDialog) SetName(value string) {
    FontDialog_SetName(f.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (f *TFontDialog) Tag() int {
    return FontDialog_GetTag(f.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (f *TFontDialog) SetTag(value int) {
    FontDialog_SetTag(f.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (f *TFontDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(FontDialog_GetComponents(f.instance, AIndex))
}

