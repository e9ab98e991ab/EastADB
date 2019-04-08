
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

type THeaderSection struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewHeaderSection
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewHeaderSection() *THeaderSection {
    h := new(THeaderSection)
    h.instance = HeaderSection_Create()
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HeaderSectionFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func HeaderSectionFromInst(inst uintptr) *THeaderSection {
    h := new(THeaderSection)
    h.instance = inst
    h.ptr = unsafe.Pointer(inst)
    return h
}

// HeaderSectionFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func HeaderSectionFromObj(obj IObject) *THeaderSection {
    h := new(THeaderSection)
    h.instance = CheckPtr(obj)
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HeaderSectionFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func HeaderSectionFromUnsafePointer(ptr unsafe.Pointer) *THeaderSection {
    h := new(THeaderSection)
    h.instance = uintptr(ptr)
    h.ptr = ptr
    return h
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (h *THeaderSection) Free() {
    if h.instance != 0 {
        HeaderSection_Free(h.instance)
        h.instance = 0
        h.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (h *THeaderSection) Instance() uintptr {
    return h.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (h *THeaderSection) UnsafeAddr() unsafe.Pointer {
    return h.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (h *THeaderSection) IsValid() bool {
    return h.instance != 0
}

// THeaderSectionClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func THeaderSectionClass() TClass {
    return HeaderSection_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (h *THeaderSection) Assign(Source IObject) {
    HeaderSection_Assign(h.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (h *THeaderSection) GetNamePath() string {
    return HeaderSection_GetNamePath(h.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (h *THeaderSection) DisposeOf() {
    HeaderSection_DisposeOf(h.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (h *THeaderSection) ClassType() TClass {
    return HeaderSection_ClassType(h.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (h *THeaderSection) ClassName() string {
    return HeaderSection_ClassName(h.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (h *THeaderSection) InstanceSize() int32 {
    return HeaderSection_InstanceSize(h.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (h *THeaderSection) InheritsFrom(AClass TClass) bool {
    return HeaderSection_InheritsFrom(h.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (h *THeaderSection) Equals(Obj IObject) bool {
    return HeaderSection_Equals(h.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (h *THeaderSection) GetHashCode() int32 {
    return HeaderSection_GetHashCode(h.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (h *THeaderSection) ToString() string {
    return HeaderSection_ToString(h.instance)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (h *THeaderSection) Left() int32 {
    return HeaderSection_GetLeft(h.instance)
}

// Right
func (h *THeaderSection) Right() int32 {
    return HeaderSection_GetRight(h.instance)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (h *THeaderSection) Alignment() TAlignment {
    return HeaderSection_GetAlignment(h.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (h *THeaderSection) SetAlignment(value TAlignment) {
    HeaderSection_SetAlignment(h.instance, value)
}

// AllowClick
func (h *THeaderSection) AllowClick() bool {
    return HeaderSection_GetAllowClick(h.instance)
}

// SetAllowClick
func (h *THeaderSection) SetAllowClick(value bool) {
    HeaderSection_SetAllowClick(h.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (h *THeaderSection) AutoSize() bool {
    return HeaderSection_GetAutoSize(h.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (h *THeaderSection) SetAutoSize(value bool) {
    HeaderSection_SetAutoSize(h.instance, value)
}

// BiDiMode
func (h *THeaderSection) BiDiMode() TBiDiMode {
    return HeaderSection_GetBiDiMode(h.instance)
}

// SetBiDiMode
func (h *THeaderSection) SetBiDiMode(value TBiDiMode) {
    HeaderSection_SetBiDiMode(h.instance, value)
}

// CheckBox
func (h *THeaderSection) CheckBox() bool {
    return HeaderSection_GetCheckBox(h.instance)
}

// SetCheckBox
func (h *THeaderSection) SetCheckBox(value bool) {
    HeaderSection_SetCheckBox(h.instance, value)
}

// Checked
// CN: 获取是否选中。
// EN: .
func (h *THeaderSection) Checked() bool {
    return HeaderSection_GetChecked(h.instance)
}

// SetChecked
// CN: 设置是否选中。
// EN: .
func (h *THeaderSection) SetChecked(value bool) {
    HeaderSection_SetChecked(h.instance, value)
}

// FixedWidth
func (h *THeaderSection) FixedWidth() bool {
    return HeaderSection_GetFixedWidth(h.instance)
}

// SetFixedWidth
func (h *THeaderSection) SetFixedWidth(value bool) {
    HeaderSection_SetFixedWidth(h.instance, value)
}

// ImageIndex
// CN: 获取图像在images中的索引。
// EN: .
func (h *THeaderSection) ImageIndex() int32 {
    return HeaderSection_GetImageIndex(h.instance)
}

// SetImageIndex
// CN: 设置图像在images中的索引。
// EN: .
func (h *THeaderSection) SetImageIndex(value int32) {
    HeaderSection_SetImageIndex(h.instance, value)
}

// MaxWidth
func (h *THeaderSection) MaxWidth() int32 {
    return HeaderSection_GetMaxWidth(h.instance)
}

// SetMaxWidth
func (h *THeaderSection) SetMaxWidth(value int32) {
    HeaderSection_SetMaxWidth(h.instance, value)
}

// MinWidth
func (h *THeaderSection) MinWidth() int32 {
    return HeaderSection_GetMinWidth(h.instance)
}

// SetMinWidth
func (h *THeaderSection) SetMinWidth(value int32) {
    HeaderSection_SetMinWidth(h.instance, value)
}

// Style
func (h *THeaderSection) Style() THeaderSectionStyle {
    return HeaderSection_GetStyle(h.instance)
}

// SetStyle
func (h *THeaderSection) SetStyle(value THeaderSectionStyle) {
    HeaderSection_SetStyle(h.instance, value)
}

// Text
// CN: 获取文本。
// EN: .
func (h *THeaderSection) Text() string {
    return HeaderSection_GetText(h.instance)
}

// SetText
// CN: 设置文本。
// EN: .
func (h *THeaderSection) SetText(value string) {
    HeaderSection_SetText(h.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (h *THeaderSection) Width() int32 {
    return HeaderSection_GetWidth(h.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (h *THeaderSection) SetWidth(value int32) {
    HeaderSection_SetWidth(h.instance, value)
}

// Collection
func (h *THeaderSection) Collection() *TCollection {
    return CollectionFromInst(HeaderSection_GetCollection(h.instance))
}

// SetCollection
func (h *THeaderSection) SetCollection(value *TCollection) {
    HeaderSection_SetCollection(h.instance, CheckPtr(value))
}

// Index
func (h *THeaderSection) Index() int32 {
    return HeaderSection_GetIndex(h.instance)
}

// SetIndex
func (h *THeaderSection) SetIndex(value int32) {
    HeaderSection_SetIndex(h.instance, value)
}

