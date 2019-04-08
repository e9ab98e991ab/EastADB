
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

type TPicture struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPicture
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPicture() *TPicture {
    p := new(TPicture)
    p.instance = Picture_Create()
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PictureFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PictureFromInst(inst uintptr) *TPicture {
    p := new(TPicture)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PictureFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PictureFromObj(obj IObject) *TPicture {
    p := new(TPicture)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PictureFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PictureFromUnsafePointer(ptr unsafe.Pointer) *TPicture {
    p := new(TPicture)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPicture) Free() {
    if p.instance != 0 {
        Picture_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPicture) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPicture) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPicture) IsValid() bool {
    return p.instance != 0
}

// TPictureClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPictureClass() TClass {
    return Picture_StaticClassType()
}

// LoadFromFile
// CN: 从文件加载。
// EN: .
func (p *TPicture) LoadFromFile(Filename string) {
    Picture_LoadFromFile(p.instance, Filename)
}

// SaveToFile
// CN: 保存至文件。
// EN: .
func (p *TPicture) SaveToFile(Filename string) {
    Picture_SaveToFile(p.instance, Filename)
}

// LoadFromStream
// CN: 文件流加载。
// EN: .
func (p *TPicture) LoadFromStream(Stream IObject) {
    Picture_LoadFromStream(p.instance, CheckPtr(Stream))
}

// SaveToStream
// CN: 保存至流。
// EN: .
func (p *TPicture) SaveToStream(Stream IObject) {
    Picture_SaveToStream(p.instance, CheckPtr(Stream))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPicture) Assign(Source IObject) {
    Picture_Assign(p.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPicture) GetNamePath() string {
    return Picture_GetNamePath(p.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPicture) DisposeOf() {
    Picture_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPicture) ClassType() TClass {
    return Picture_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPicture) ClassName() string {
    return Picture_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPicture) InstanceSize() int32 {
    return Picture_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPicture) InheritsFrom(AClass TClass) bool {
    return Picture_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPicture) Equals(Obj IObject) bool {
    return Picture_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPicture) GetHashCode() int32 {
    return Picture_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPicture) ToString() string {
    return Picture_ToString(p.instance)
}

// Bitmap
func (p *TPicture) Bitmap() *TBitmap {
    return BitmapFromInst(Picture_GetBitmap(p.instance))
}

// SetBitmap
func (p *TPicture) SetBitmap(value *TBitmap) {
    Picture_SetBitmap(p.instance, CheckPtr(value))
}

// Graphic
func (p *TPicture) Graphic() *TGraphic {
    return GraphicFromInst(Picture_GetGraphic(p.instance))
}

// SetGraphic
func (p *TPicture) SetGraphic(value *TGraphic) {
    Picture_SetGraphic(p.instance, CheckPtr(value))
}

// Height
// CN: 获取高度。
// EN: Get height.
func (p *TPicture) Height() int32 {
    return Picture_GetHeight(p.instance)
}

// Icon
// CN: 获取图标。
// EN: Get icon.
func (p *TPicture) Icon() *TIcon {
    return IconFromInst(Picture_GetIcon(p.instance))
}

// SetIcon
// CN: 设置图标。
// EN: Set icon.
func (p *TPicture) SetIcon(value *TIcon) {
    Picture_SetIcon(p.instance, CheckPtr(value))
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (p *TPicture) Width() int32 {
    return Picture_GetWidth(p.instance)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (p *TPicture) SetOnChange(fn TNotifyEvent) {
    Picture_SetOnChange(p.instance, fn)
}

