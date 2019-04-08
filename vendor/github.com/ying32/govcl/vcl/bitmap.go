
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

type TBitmap struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewBitmap
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBitmap() *TBitmap {
    b := new(TBitmap)
    b.instance = Bitmap_Create()
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BitmapFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func BitmapFromInst(inst uintptr) *TBitmap {
    b := new(TBitmap)
    b.instance = inst
    b.ptr = unsafe.Pointer(inst)
    return b
}

// BitmapFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func BitmapFromObj(obj IObject) *TBitmap {
    b := new(TBitmap)
    b.instance = CheckPtr(obj)
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BitmapFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func BitmapFromUnsafePointer(ptr unsafe.Pointer) *TBitmap {
    b := new(TBitmap)
    b.instance = uintptr(ptr)
    b.ptr = ptr
    return b
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (b *TBitmap) Free() {
    if b.instance != 0 {
        Bitmap_Free(b.instance)
        b.instance = 0
        b.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBitmap) Instance() uintptr {
    return b.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBitmap) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBitmap) IsValid() bool {
    return b.instance != 0
}

// TBitmapClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBitmapClass() TClass {
    return Bitmap_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (b *TBitmap) Assign(Source IObject) {
    Bitmap_Assign(b.instance, CheckPtr(Source))
}

// FreeImage
func (b *TBitmap) FreeImage() {
    Bitmap_FreeImage(b.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (b *TBitmap) HandleAllocated() bool {
    return Bitmap_HandleAllocated(b.instance)
}

// LoadFromStream
// CN: 文件流加载。
// EN: .
func (b *TBitmap) LoadFromStream(Stream IObject) {
    Bitmap_LoadFromStream(b.instance, CheckPtr(Stream))
}

// SaveToStream
// CN: 保存至流。
// EN: .
func (b *TBitmap) SaveToStream(Stream IObject) {
    Bitmap_SaveToStream(b.instance, CheckPtr(Stream))
}

// SetSize
func (b *TBitmap) SetSize(AWidth int32, AHeight int32) {
    Bitmap_SetSize(b.instance, AWidth , AHeight)
}

// LoadFromResourceName
func (b *TBitmap) LoadFromResourceName(Instance uintptr, ResName string) {
    Bitmap_LoadFromResourceName(b.instance, Instance , ResName)
}

// LoadFromResourceID
func (b *TBitmap) LoadFromResourceID(Instance uintptr, ResID int32) {
    Bitmap_LoadFromResourceID(b.instance, Instance , ResID)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBitmap) Equals(Obj IObject) bool {
    return Bitmap_Equals(b.instance, CheckPtr(Obj))
}

// LoadFromFile
// CN: 从文件加载。
// EN: .
func (b *TBitmap) LoadFromFile(Filename string) {
    Bitmap_LoadFromFile(b.instance, Filename)
}

// SaveToFile
// CN: 保存至文件。
// EN: .
func (b *TBitmap) SaveToFile(Filename string) {
    Bitmap_SaveToFile(b.instance, Filename)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (b *TBitmap) GetNamePath() string {
    return Bitmap_GetNamePath(b.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (b *TBitmap) DisposeOf() {
    Bitmap_DisposeOf(b.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBitmap) ClassType() TClass {
    return Bitmap_ClassType(b.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBitmap) ClassName() string {
    return Bitmap_ClassName(b.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBitmap) InstanceSize() int32 {
    return Bitmap_InstanceSize(b.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBitmap) InheritsFrom(AClass TClass) bool {
    return Bitmap_InheritsFrom(b.instance, AClass)
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBitmap) GetHashCode() int32 {
    return Bitmap_GetHashCode(b.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (b *TBitmap) ToString() string {
    return Bitmap_ToString(b.instance)
}

// Canvas
// CN: 获取画布。
// EN: .
func (b *TBitmap) Canvas() *TCanvas {
    return CanvasFromInst(Bitmap_GetCanvas(b.instance))
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (b *TBitmap) Handle() HBITMAP {
    return Bitmap_GetHandle(b.instance)
}

// SetHandle
// CN: 设置控件句柄。
// EN: Set Control handle.
func (b *TBitmap) SetHandle(value HBITMAP) {
    Bitmap_SetHandle(b.instance, value)
}

// HandleType
func (b *TBitmap) HandleType() TBitmapHandleType {
    return Bitmap_GetHandleType(b.instance)
}

// SetHandleType
func (b *TBitmap) SetHandleType(value TBitmapHandleType) {
    Bitmap_SetHandleType(b.instance, value)
}

// MaskHandle
func (b *TBitmap) MaskHandle() HBITMAP {
    return Bitmap_GetMaskHandle(b.instance)
}

// SetMaskHandle
func (b *TBitmap) SetMaskHandle(value HBITMAP) {
    Bitmap_SetMaskHandle(b.instance, value)
}

// PixelFormat
func (b *TBitmap) PixelFormat() TPixelFormat {
    return Bitmap_GetPixelFormat(b.instance)
}

// SetPixelFormat
func (b *TBitmap) SetPixelFormat(value TPixelFormat) {
    Bitmap_SetPixelFormat(b.instance, value)
}

// TransparentColor
// CN: 获取使用颜色透明。
// EN: .
func (b *TBitmap) TransparentColor() TColor {
    return Bitmap_GetTransparentColor(b.instance)
}

// SetTransparentColor
// CN: 设置使用颜色透明。
// EN: .
func (b *TBitmap) SetTransparentColor(value TColor) {
    Bitmap_SetTransparentColor(b.instance, value)
}

// Empty
func (b *TBitmap) Empty() bool {
    return Bitmap_GetEmpty(b.instance)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (b *TBitmap) Height() int32 {
    return Bitmap_GetHeight(b.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (b *TBitmap) SetHeight(value int32) {
    Bitmap_SetHeight(b.instance, value)
}

// Modified
// CN: 获取修改。
// EN: Get modified.
func (b *TBitmap) Modified() bool {
    return Bitmap_GetModified(b.instance)
}

// SetModified
// CN: 设置修改。
// EN: Set modified.
func (b *TBitmap) SetModified(value bool) {
    Bitmap_SetModified(b.instance, value)
}

// Palette
func (b *TBitmap) Palette() HPALETTE {
    return Bitmap_GetPalette(b.instance)
}

// SetPalette
func (b *TBitmap) SetPalette(value HPALETTE) {
    Bitmap_SetPalette(b.instance, value)
}

// PaletteModified
func (b *TBitmap) PaletteModified() bool {
    return Bitmap_GetPaletteModified(b.instance)
}

// SetPaletteModified
func (b *TBitmap) SetPaletteModified(value bool) {
    Bitmap_SetPaletteModified(b.instance, value)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (b *TBitmap) Transparent() bool {
    return Bitmap_GetTransparent(b.instance)
}

// SetTransparent
// CN: 设置透明。
// EN: Set transparent.
func (b *TBitmap) SetTransparent(value bool) {
    Bitmap_SetTransparent(b.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (b *TBitmap) Width() int32 {
    return Bitmap_GetWidth(b.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (b *TBitmap) SetWidth(value int32) {
    Bitmap_SetWidth(b.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (b *TBitmap) SetOnChange(fn TNotifyEvent) {
    Bitmap_SetOnChange(b.instance, fn)
}

// ScanLine
func (b *TBitmap) ScanLine(Row int32) uintptr {
    return Bitmap_GetScanLine(b.instance, Row)
}

