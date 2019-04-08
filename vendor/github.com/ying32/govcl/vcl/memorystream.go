
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

type TMemoryStream struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewMemoryStream
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMemoryStream() *TMemoryStream {
    m := new(TMemoryStream)
    m.instance = MemoryStream_Create()
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MemoryStreamFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func MemoryStreamFromInst(inst uintptr) *TMemoryStream {
    m := new(TMemoryStream)
    m.instance = inst
    m.ptr = unsafe.Pointer(inst)
    return m
}

// MemoryStreamFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func MemoryStreamFromObj(obj IObject) *TMemoryStream {
    m := new(TMemoryStream)
    m.instance = CheckPtr(obj)
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

// MemoryStreamFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func MemoryStreamFromUnsafePointer(ptr unsafe.Pointer) *TMemoryStream {
    m := new(TMemoryStream)
    m.instance = uintptr(ptr)
    m.ptr = ptr
    return m
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (m *TMemoryStream) Free() {
    if m.instance != 0 {
        MemoryStream_Free(m.instance)
        m.instance = 0
        m.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMemoryStream) Instance() uintptr {
    return m.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMemoryStream) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMemoryStream) IsValid() bool {
    return m.instance != 0
}

// TMemoryStreamClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMemoryStreamClass() TClass {
    return MemoryStream_StaticClassType()
}

// Clear
// CN: 清除。
// EN: .
func (m *TMemoryStream) Clear() {
    MemoryStream_Clear(m.instance)
}

// LoadFromStream
// CN: 文件流加载。
// EN: .
func (m *TMemoryStream) LoadFromStream(Stream IObject) {
    MemoryStream_LoadFromStream(m.instance, CheckPtr(Stream))
}

// LoadFromFile
// CN: 从文件加载。
// EN: .
func (m *TMemoryStream) LoadFromFile(FileName string) {
    MemoryStream_LoadFromFile(m.instance, FileName)
}

// Seek
func (m *TMemoryStream) Seek(Offset int64, Origin TSeekOrigin) int64 {
    return MemoryStream_Seek(m.instance, Offset , Origin)
}

// SaveToStream
// CN: 保存至流。
// EN: .
func (m *TMemoryStream) SaveToStream(Stream IObject) {
    MemoryStream_SaveToStream(m.instance, CheckPtr(Stream))
}

// SaveToFile
// CN: 保存至文件。
// EN: .
func (m *TMemoryStream) SaveToFile(FileName string) {
    MemoryStream_SaveToFile(m.instance, FileName)
}

// CopyFrom
func (m *TMemoryStream) CopyFrom(Source IObject, Count int64) int64 {
    return MemoryStream_CopyFrom(m.instance, CheckPtr(Source), Count)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (m *TMemoryStream) DisposeOf() {
    MemoryStream_DisposeOf(m.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMemoryStream) ClassType() TClass {
    return MemoryStream_ClassType(m.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMemoryStream) ClassName() string {
    return MemoryStream_ClassName(m.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMemoryStream) InstanceSize() int32 {
    return MemoryStream_InstanceSize(m.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMemoryStream) InheritsFrom(AClass TClass) bool {
    return MemoryStream_InheritsFrom(m.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMemoryStream) Equals(Obj IObject) bool {
    return MemoryStream_Equals(m.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMemoryStream) GetHashCode() int32 {
    return MemoryStream_GetHashCode(m.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (m *TMemoryStream) ToString() string {
    return MemoryStream_ToString(m.instance)
}

// Memory
func (m *TMemoryStream) Memory() uintptr {
    return MemoryStream_GetMemory(m.instance)
}

// Position
func (m *TMemoryStream) Position() int64 {
    return MemoryStream_GetPosition(m.instance)
}

// SetPosition
func (m *TMemoryStream) SetPosition(value int64) {
    MemoryStream_SetPosition(m.instance, value)
}

// Size
func (m *TMemoryStream) Size() int64 {
    return MemoryStream_GetSize(m.instance)
}

// SetSize
func (m *TMemoryStream) SetSize(value int64) {
    MemoryStream_SetSize(m.instance, value)
}

