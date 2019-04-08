
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

type TClipboard struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewClipboard
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewClipboard() *TClipboard {
    c := new(TClipboard)
    c.instance = Clipboard_Create()
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ClipboardFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ClipboardFromInst(inst uintptr) *TClipboard {
    c := new(TClipboard)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// ClipboardFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ClipboardFromObj(obj IObject) *TClipboard {
    c := new(TClipboard)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ClipboardFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ClipboardFromUnsafePointer(ptr unsafe.Pointer) *TClipboard {
    c := new(TClipboard)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TClipboard) Free() {
    if c.instance != 0 {
        Clipboard_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TClipboard) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TClipboard) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TClipboard) IsValid() bool {
    return c.instance != 0
}

// TClipboardClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TClipboardClass() TClass {
    return Clipboard_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TClipboard) Assign(Source IObject) {
    Clipboard_Assign(c.instance, CheckPtr(Source))
}

// Clear
// CN: 清除。
// EN: .
func (c *TClipboard) Clear() {
    Clipboard_Clear(c.instance)
}

// Close
// CN: 关闭。
// EN: .
func (c *TClipboard) Close() {
    Clipboard_Close(c.instance)
}

// GetAsHandle
func (c *TClipboard) GetAsHandle(Format uint16) uintptr {
    return Clipboard_GetAsHandle(c.instance, Format)
}

// HasFormat
func (c *TClipboard) HasFormat(Format uint16) bool {
    return Clipboard_HasFormat(c.instance, Format)
}

// Open
func (c *TClipboard) Open() {
    Clipboard_Open(c.instance)
}

// SetAsHandle
func (c *TClipboard) SetAsHandle(Format uint16, Value uintptr) {
    Clipboard_SetAsHandle(c.instance, Format , Value)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (c *TClipboard) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Clipboard_GetTextBuf(c.instance, Buffer , BufSize)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (c *TClipboard) SetTextBuf(Buffer string) {
    Clipboard_SetTextBuf(c.instance, Buffer)
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TClipboard) GetNamePath() string {
    return Clipboard_GetNamePath(c.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TClipboard) DisposeOf() {
    Clipboard_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TClipboard) ClassType() TClass {
    return Clipboard_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TClipboard) ClassName() string {
    return Clipboard_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TClipboard) InstanceSize() int32 {
    return Clipboard_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TClipboard) InheritsFrom(AClass TClass) bool {
    return Clipboard_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TClipboard) Equals(Obj IObject) bool {
    return Clipboard_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TClipboard) GetHashCode() int32 {
    return Clipboard_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TClipboard) ToString() string {
    return Clipboard_ToString(c.instance)
}

// AsText
func (c *TClipboard) AsText() string {
    return Clipboard_GetAsText(c.instance)
}

// SetAsText
func (c *TClipboard) SetAsText(value string) {
    Clipboard_SetAsText(c.instance, value)
}

// FormatCount
func (c *TClipboard) FormatCount() int32 {
    return Clipboard_GetFormatCount(c.instance)
}

// Formats
func (c *TClipboard) Formats(Index int32) uint16 {
    return Clipboard_GetFormats(c.instance, Index)
}

