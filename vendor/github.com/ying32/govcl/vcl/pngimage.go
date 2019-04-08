
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

type TPngImage struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPngImage
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPngImage() *TPngImage {
    p := new(TPngImage)
    p.instance = PngImage_Create()
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PngImageFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PngImageFromInst(inst uintptr) *TPngImage {
    p := new(TPngImage)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PngImageFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PngImageFromObj(obj IObject) *TPngImage {
    p := new(TPngImage)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PngImageFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PngImageFromUnsafePointer(ptr unsafe.Pointer) *TPngImage {
    p := new(TPngImage)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPngImage) Free() {
    if p.instance != 0 {
        PngImage_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPngImage) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPngImage) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPngImage) IsValid() bool {
    return p.instance != 0
}

// TPngImageClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPngImageClass() TClass {
    return PngImage_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPngImage) Assign(Source IObject) {
    PngImage_Assign(p.instance, CheckPtr(Source))
}

// LoadFromStream
// CN: 文件流加载。
// EN: .
func (p *TPngImage) LoadFromStream(Stream IObject) {
    PngImage_LoadFromStream(p.instance, CheckPtr(Stream))
}

// SaveToStream
// CN: 保存至流。
// EN: .
func (p *TPngImage) SaveToStream(Stream IObject) {
    PngImage_SaveToStream(p.instance, CheckPtr(Stream))
}

// LoadFromResourceName
func (p *TPngImage) LoadFromResourceName(Instance uintptr, Name string) {
    PngImage_LoadFromResourceName(p.instance, Instance , Name)
}

// LoadFromResourceID
func (p *TPngImage) LoadFromResourceID(Instance uintptr, ResID int32) {
    PngImage_LoadFromResourceID(p.instance, Instance , ResID)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPngImage) Equals(Obj IObject) bool {
    return PngImage_Equals(p.instance, CheckPtr(Obj))
}

// LoadFromFile
// CN: 从文件加载。
// EN: .
func (p *TPngImage) LoadFromFile(Filename string) {
    PngImage_LoadFromFile(p.instance, Filename)
}

// SaveToFile
// CN: 保存至文件。
// EN: .
func (p *TPngImage) SaveToFile(Filename string) {
    PngImage_SaveToFile(p.instance, Filename)
}

// SetSize
func (p *TPngImage) SetSize(AWidth int32, AHeight int32) {
    PngImage_SetSize(p.instance, AWidth , AHeight)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPngImage) GetNamePath() string {
    return PngImage_GetNamePath(p.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPngImage) DisposeOf() {
    PngImage_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPngImage) ClassType() TClass {
    return PngImage_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPngImage) ClassName() string {
    return PngImage_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPngImage) InstanceSize() int32 {
    return PngImage_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPngImage) InheritsFrom(AClass TClass) bool {
    return PngImage_InheritsFrom(p.instance, AClass)
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPngImage) GetHashCode() int32 {
    return PngImage_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPngImage) ToString() string {
    return PngImage_ToString(p.instance)
}

// TransparentColor
// CN: 获取使用颜色透明。
// EN: .
func (p *TPngImage) TransparentColor() TColor {
    return PngImage_GetTransparentColor(p.instance)
}

// SetTransparentColor
// CN: 设置使用颜色透明。
// EN: .
func (p *TPngImage) SetTransparentColor(value TColor) {
    PngImage_SetTransparentColor(p.instance, value)
}

// Canvas
// CN: 获取画布。
// EN: .
func (p *TPngImage) Canvas() *TCanvas {
    return CanvasFromInst(PngImage_GetCanvas(p.instance))
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (p *TPngImage) Width() int32 {
    return PngImage_GetWidth(p.instance)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (p *TPngImage) Height() int32 {
    return PngImage_GetHeight(p.instance)
}

// MaxIdatSize
func (p *TPngImage) MaxIdatSize() int32 {
    return PngImage_GetMaxIdatSize(p.instance)
}

// SetMaxIdatSize
func (p *TPngImage) SetMaxIdatSize(value int32) {
    PngImage_SetMaxIdatSize(p.instance, value)
}

// Empty
func (p *TPngImage) Empty() bool {
    return PngImage_GetEmpty(p.instance)
}

// CompressionLevel
func (p *TPngImage) CompressionLevel() TCompressionLevel {
    return PngImage_GetCompressionLevel(p.instance)
}

// SetCompressionLevel
func (p *TPngImage) SetCompressionLevel(value TCompressionLevel) {
    PngImage_SetCompressionLevel(p.instance, value)
}

// Version
func (p *TPngImage) Version() string {
    return PngImage_GetVersion(p.instance)
}

// Modified
// CN: 获取修改。
// EN: Get modified.
func (p *TPngImage) Modified() bool {
    return PngImage_GetModified(p.instance)
}

// SetModified
// CN: 设置修改。
// EN: Set modified.
func (p *TPngImage) SetModified(value bool) {
    PngImage_SetModified(p.instance, value)
}

// Palette
func (p *TPngImage) Palette() HPALETTE {
    return PngImage_GetPalette(p.instance)
}

// SetPalette
func (p *TPngImage) SetPalette(value HPALETTE) {
    PngImage_SetPalette(p.instance, value)
}

// PaletteModified
func (p *TPngImage) PaletteModified() bool {
    return PngImage_GetPaletteModified(p.instance)
}

// SetPaletteModified
func (p *TPngImage) SetPaletteModified(value bool) {
    PngImage_SetPaletteModified(p.instance, value)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (p *TPngImage) Transparent() bool {
    return PngImage_GetTransparent(p.instance)
}

// SetTransparent
// CN: 设置透明。
// EN: Set transparent.
func (p *TPngImage) SetTransparent(value bool) {
    PngImage_SetTransparent(p.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (p *TPngImage) SetOnChange(fn TNotifyEvent) {
    PngImage_SetOnChange(p.instance, fn)
}

