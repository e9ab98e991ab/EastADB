
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

type TPreviewClipRegion struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPreviewClipRegion
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPreviewClipRegion() *TPreviewClipRegion {
    p := new(TPreviewClipRegion)
    p.instance = PreviewClipRegion_Create()
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PreviewClipRegionFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PreviewClipRegionFromInst(inst uintptr) *TPreviewClipRegion {
    p := new(TPreviewClipRegion)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PreviewClipRegionFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PreviewClipRegionFromObj(obj IObject) *TPreviewClipRegion {
    p := new(TPreviewClipRegion)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PreviewClipRegionFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PreviewClipRegionFromUnsafePointer(ptr unsafe.Pointer) *TPreviewClipRegion {
    p := new(TPreviewClipRegion)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPreviewClipRegion) Free() {
    if p.instance != 0 {
        PreviewClipRegion_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPreviewClipRegion) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPreviewClipRegion) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPreviewClipRegion) IsValid() bool {
    return p.instance != 0
}

// TPreviewClipRegionClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPreviewClipRegionClass() TClass {
    return PreviewClipRegion_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPreviewClipRegion) Assign(Source IObject) {
    PreviewClipRegion_Assign(p.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPreviewClipRegion) GetNamePath() string {
    return PreviewClipRegion_GetNamePath(p.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPreviewClipRegion) DisposeOf() {
    PreviewClipRegion_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPreviewClipRegion) ClassType() TClass {
    return PreviewClipRegion_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPreviewClipRegion) ClassName() string {
    return PreviewClipRegion_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPreviewClipRegion) InstanceSize() int32 {
    return PreviewClipRegion_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPreviewClipRegion) InheritsFrom(AClass TClass) bool {
    return PreviewClipRegion_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPreviewClipRegion) Equals(Obj IObject) bool {
    return PreviewClipRegion_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPreviewClipRegion) GetHashCode() int32 {
    return PreviewClipRegion_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPreviewClipRegion) ToString() string {
    return PreviewClipRegion_ToString(p.instance)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (p *TPreviewClipRegion) SetOnChange(fn TNotifyEvent) {
    PreviewClipRegion_SetOnChange(p.instance, fn)
}

// Bounds
func (p *TPreviewClipRegion) Bounds() TRect {
    return PreviewClipRegion_GetBounds(p.instance)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (p *TPreviewClipRegion) Left() int32 {
    return PreviewClipRegion_GetLeft(p.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (p *TPreviewClipRegion) SetLeft(value int32) {
    PreviewClipRegion_SetLeft(p.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (p *TPreviewClipRegion) Top() int32 {
    return PreviewClipRegion_GetTop(p.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (p *TPreviewClipRegion) SetTop(value int32) {
    PreviewClipRegion_SetTop(p.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (p *TPreviewClipRegion) Height() int32 {
    return PreviewClipRegion_GetHeight(p.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (p *TPreviewClipRegion) SetHeight(value int32) {
    PreviewClipRegion_SetHeight(p.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (p *TPreviewClipRegion) Width() int32 {
    return PreviewClipRegion_GetWidth(p.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (p *TPreviewClipRegion) SetWidth(value int32) {
    PreviewClipRegion_SetWidth(p.instance, value)
}

