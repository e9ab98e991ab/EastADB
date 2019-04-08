
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

type TListItem struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewListItem
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewListItem() *TListItem {
    l := new(TListItem)
    l.instance = ListItem_Create()
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListItemFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ListItemFromInst(inst uintptr) *TListItem {
    l := new(TListItem)
    l.instance = inst
    l.ptr = unsafe.Pointer(inst)
    return l
}

// ListItemFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ListItemFromObj(obj IObject) *TListItem {
    l := new(TListItem)
    l.instance = CheckPtr(obj)
    l.ptr = unsafe.Pointer(l.instance)
    return l
}

// ListItemFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ListItemFromUnsafePointer(ptr unsafe.Pointer) *TListItem {
    l := new(TListItem)
    l.instance = uintptr(ptr)
    l.ptr = ptr
    return l
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (l *TListItem) Free() {
    if l.instance != 0 {
        ListItem_Free(l.instance)
        l.instance = 0
        l.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (l *TListItem) Instance() uintptr {
    return l.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (l *TListItem) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (l *TListItem) IsValid() bool {
    return l.instance != 0
}

// TListItemClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TListItemClass() TClass {
    return ListItem_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (l *TListItem) Assign(Source IObject) {
    ListItem_Assign(l.instance, CheckPtr(Source))
}

// CancelEdit
func (l *TListItem) CancelEdit() {
    ListItem_CancelEdit(l.instance)
}

// Delete
func (l *TListItem) Delete() {
    ListItem_Delete(l.instance)
}

// DisplayRect
func (l *TListItem) DisplayRect(Code TDisplayCode) TRect {
    return ListItem_DisplayRect(l.instance, Code)
}

// EditCaption
func (l *TListItem) EditCaption() bool {
    return ListItem_EditCaption(l.instance)
}

// MakeVisible
func (l *TListItem) MakeVisible(PartialOK bool) {
    ListItem_MakeVisible(l.instance, PartialOK)
}

// Update
// CN: 控件更新。
// EN: Update.
func (l *TListItem) Update() {
    ListItem_Update(l.instance)
}

// WorkArea
func (l *TListItem) WorkArea() int32 {
    return ListItem_WorkArea(l.instance)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (l *TListItem) GetNamePath() string {
    return ListItem_GetNamePath(l.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (l *TListItem) DisposeOf() {
    ListItem_DisposeOf(l.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (l *TListItem) ClassType() TClass {
    return ListItem_ClassType(l.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (l *TListItem) ClassName() string {
    return ListItem_ClassName(l.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (l *TListItem) InstanceSize() int32 {
    return ListItem_InstanceSize(l.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (l *TListItem) InheritsFrom(AClass TClass) bool {
    return ListItem_InheritsFrom(l.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (l *TListItem) Equals(Obj IObject) bool {
    return ListItem_Equals(l.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (l *TListItem) GetHashCode() int32 {
    return ListItem_GetHashCode(l.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (l *TListItem) ToString() string {
    return ListItem_ToString(l.instance)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (l *TListItem) Caption() string {
    return ListItem_GetCaption(l.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (l *TListItem) SetCaption(value string) {
    ListItem_SetCaption(l.instance, value)
}

// Checked
// CN: 获取是否选中。
// EN: .
func (l *TListItem) Checked() bool {
    return ListItem_GetChecked(l.instance)
}

// SetChecked
// CN: 设置是否选中。
// EN: .
func (l *TListItem) SetChecked(value bool) {
    ListItem_SetChecked(l.instance, value)
}

// Cut
func (l *TListItem) Cut() bool {
    return ListItem_GetCut(l.instance)
}

// SetCut
func (l *TListItem) SetCut(value bool) {
    ListItem_SetCut(l.instance, value)
}

// Data
func (l *TListItem) Data() unsafe.Pointer {
    return ListItem_GetData(l.instance)
}

// SetData
func (l *TListItem) SetData(value unsafe.Pointer) {
    ListItem_SetData(l.instance, value)
}

// Deleting
func (l *TListItem) Deleting() bool {
    return ListItem_GetDeleting(l.instance)
}

// DropTarget
func (l *TListItem) DropTarget() bool {
    return ListItem_GetDropTarget(l.instance)
}

// SetDropTarget
func (l *TListItem) SetDropTarget(value bool) {
    ListItem_SetDropTarget(l.instance, value)
}

// Focused
// CN: 获取返回是否获取焦点。
// EN: Get Return to get focus.
func (l *TListItem) Focused() bool {
    return ListItem_GetFocused(l.instance)
}

// SetFocused
// CN: 设置返回是否获取焦点。
// EN: Set Return to get focus.
func (l *TListItem) SetFocused(value bool) {
    ListItem_SetFocused(l.instance, value)
}

// GroupID
func (l *TListItem) GroupID() int32 {
    return ListItem_GetGroupID(l.instance)
}

// SetGroupID
func (l *TListItem) SetGroupID(value int32) {
    ListItem_SetGroupID(l.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (l *TListItem) Handle() HWND {
    return ListItem_GetHandle(l.instance)
}

// ImageIndex
// CN: 获取图像在images中的索引。
// EN: .
func (l *TListItem) ImageIndex() int32 {
    return ListItem_GetImageIndex(l.instance)
}

// SetImageIndex
// CN: 设置图像在images中的索引。
// EN: .
func (l *TListItem) SetImageIndex(value int32) {
    ListItem_SetImageIndex(l.instance, value)
}

// Indent
func (l *TListItem) Indent() int32 {
    return ListItem_GetIndent(l.instance)
}

// SetIndent
func (l *TListItem) SetIndent(value int32) {
    ListItem_SetIndent(l.instance, value)
}

// Index
func (l *TListItem) Index() int32 {
    return ListItem_GetIndex(l.instance)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (l *TListItem) Left() int32 {
    return ListItem_GetLeft(l.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (l *TListItem) SetLeft(value int32) {
    ListItem_SetLeft(l.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (l *TListItem) Owner() *TListItems {
    return ListItemsFromInst(ListItem_GetOwner(l.instance))
}

// OverlayIndex
func (l *TListItem) OverlayIndex() int32 {
    return ListItem_GetOverlayIndex(l.instance)
}

// SetOverlayIndex
func (l *TListItem) SetOverlayIndex(value int32) {
    ListItem_SetOverlayIndex(l.instance, value)
}

// Position
func (l *TListItem) Position() TPoint {
    return ListItem_GetPosition(l.instance)
}

// SetPosition
func (l *TListItem) SetPosition(value TPoint) {
    ListItem_SetPosition(l.instance, value)
}

// Selected
func (l *TListItem) Selected() bool {
    return ListItem_GetSelected(l.instance)
}

// SetSelected
func (l *TListItem) SetSelected(value bool) {
    ListItem_SetSelected(l.instance, value)
}

// StateIndex
func (l *TListItem) StateIndex() int32 {
    return ListItem_GetStateIndex(l.instance)
}

// SetStateIndex
func (l *TListItem) SetStateIndex(value int32) {
    ListItem_SetStateIndex(l.instance, value)
}

// SubItems
func (l *TListItem) SubItems() *TStrings {
    return StringsFromInst(ListItem_GetSubItems(l.instance))
}

// SetSubItems
func (l *TListItem) SetSubItems(value IObject) {
    ListItem_SetSubItems(l.instance, CheckPtr(value))
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (l *TListItem) Top() int32 {
    return ListItem_GetTop(l.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (l *TListItem) SetTop(value int32) {
    ListItem_SetTop(l.instance, value)
}

// SubItemImages
func (l *TListItem) SubItemImages(Index int32) int32 {
    return ListItem_GetSubItemImages(l.instance, Index)
}

// SubItemImages
func (l *TListItem) SetSubItemImages(Index int32, value int32) {
    ListItem_SetSubItemImages(l.instance, Index, value)
}

