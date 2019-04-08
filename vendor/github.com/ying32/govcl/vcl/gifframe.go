
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

type TGIFFrame struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewGIFFrame
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewGIFFrame() *TGIFFrame {
    g := new(TGIFFrame)
    g.instance = GIFFrame_Create()
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GIFFrameFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func GIFFrameFromInst(inst uintptr) *TGIFFrame {
    g := new(TGIFFrame)
    g.instance = inst
    g.ptr = unsafe.Pointer(inst)
    return g
}

// GIFFrameFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func GIFFrameFromObj(obj IObject) *TGIFFrame {
    g := new(TGIFFrame)
    g.instance = CheckPtr(obj)
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GIFFrameFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func GIFFrameFromUnsafePointer(ptr unsafe.Pointer) *TGIFFrame {
    g := new(TGIFFrame)
    g.instance = uintptr(ptr)
    g.ptr = ptr
    return g
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (g *TGIFFrame) Free() {
    if g.instance != 0 {
        GIFFrame_Free(g.instance)
        g.instance = 0
        g.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (g *TGIFFrame) Instance() uintptr {
    return g.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (g *TGIFFrame) UnsafeAddr() unsafe.Pointer {
    return g.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (g *TGIFFrame) IsValid() bool {
    return g.instance != 0
}

// TGIFFrameClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TGIFFrameClass() TClass {
    return GIFFrame_StaticClassType()
}

// Clear
// CN: 清除。
// EN: .
func (g *TGIFFrame) Clear() {
    GIFFrame_Clear(g.instance)
}

// SaveToStream
// CN: 保存至流。
// EN: .
func (g *TGIFFrame) SaveToStream(Stream IObject) {
    GIFFrame_SaveToStream(g.instance, CheckPtr(Stream))
}

// LoadFromStream
// CN: 文件流加载。
// EN: .
func (g *TGIFFrame) LoadFromStream(Stream IObject) {
    GIFFrame_LoadFromStream(g.instance, CheckPtr(Stream))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (g *TGIFFrame) Assign(Source IObject) {
    GIFFrame_Assign(g.instance, CheckPtr(Source))
}

// SaveToFile
// CN: 保存至文件。
// EN: .
func (g *TGIFFrame) SaveToFile(Filename string) {
    GIFFrame_SaveToFile(g.instance, Filename)
}

// LoadFromFile
// CN: 从文件加载。
// EN: .
func (g *TGIFFrame) LoadFromFile(Filename string) {
    GIFFrame_LoadFromFile(g.instance, Filename)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (g *TGIFFrame) GetNamePath() string {
    return GIFFrame_GetNamePath(g.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (g *TGIFFrame) DisposeOf() {
    GIFFrame_DisposeOf(g.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (g *TGIFFrame) ClassType() TClass {
    return GIFFrame_ClassType(g.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (g *TGIFFrame) ClassName() string {
    return GIFFrame_ClassName(g.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (g *TGIFFrame) InstanceSize() int32 {
    return GIFFrame_InstanceSize(g.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (g *TGIFFrame) InheritsFrom(AClass TClass) bool {
    return GIFFrame_InheritsFrom(g.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (g *TGIFFrame) Equals(Obj IObject) bool {
    return GIFFrame_Equals(g.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (g *TGIFFrame) GetHashCode() int32 {
    return GIFFrame_GetHashCode(g.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (g *TGIFFrame) ToString() string {
    return GIFFrame_ToString(g.instance)
}

// HasBitmap
func (g *TGIFFrame) HasBitmap() bool {
    return GIFFrame_GetHasBitmap(g.instance)
}

// SetHasBitmap
func (g *TGIFFrame) SetHasBitmap(value bool) {
    GIFFrame_SetHasBitmap(g.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (g *TGIFFrame) Left() uint16 {
    return GIFFrame_GetLeft(g.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (g *TGIFFrame) SetLeft(value uint16) {
    GIFFrame_SetLeft(g.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (g *TGIFFrame) Top() uint16 {
    return GIFFrame_GetTop(g.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (g *TGIFFrame) SetTop(value uint16) {
    GIFFrame_SetTop(g.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (g *TGIFFrame) Width() uint16 {
    return GIFFrame_GetWidth(g.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (g *TGIFFrame) SetWidth(value uint16) {
    GIFFrame_SetWidth(g.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (g *TGIFFrame) Height() uint16 {
    return GIFFrame_GetHeight(g.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (g *TGIFFrame) SetHeight(value uint16) {
    GIFFrame_SetHeight(g.instance, value)
}

// BoundsRect
func (g *TGIFFrame) BoundsRect() TRect {
    return GIFFrame_GetBoundsRect(g.instance)
}

// SetBoundsRect
func (g *TGIFFrame) SetBoundsRect(value TRect) {
    GIFFrame_SetBoundsRect(g.instance, value)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (g *TGIFFrame) ClientRect() TRect {
    return GIFFrame_GetClientRect(g.instance)
}

// Data
func (g *TGIFFrame) Data() uintptr {
    return GIFFrame_GetData(g.instance)
}

// DataSize
func (g *TGIFFrame) DataSize() int32 {
    return GIFFrame_GetDataSize(g.instance)
}

// Version
func (g *TGIFFrame) Version() TGIFVersion {
    return GIFFrame_GetVersion(g.instance)
}

// BitsPerPixel
func (g *TGIFFrame) BitsPerPixel() int32 {
    return GIFFrame_GetBitsPerPixel(g.instance)
}

// Bitmap
func (g *TGIFFrame) Bitmap() *TBitmap {
    return BitmapFromInst(GIFFrame_GetBitmap(g.instance))
}

// SetBitmap
func (g *TGIFFrame) SetBitmap(value *TBitmap) {
    GIFFrame_SetBitmap(g.instance, CheckPtr(value))
}

// Palette
func (g *TGIFFrame) Palette() HPALETTE {
    return GIFFrame_GetPalette(g.instance)
}

// SetPalette
func (g *TGIFFrame) SetPalette(value HPALETTE) {
    GIFFrame_SetPalette(g.instance, value)
}

// Empty
func (g *TGIFFrame) Empty() bool {
    return GIFFrame_GetEmpty(g.instance)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (g *TGIFFrame) Transparent() bool {
    return GIFFrame_GetTransparent(g.instance)
}

