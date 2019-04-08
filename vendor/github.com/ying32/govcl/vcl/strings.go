
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

type TStrings struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewStrings
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewStrings() *TStrings {
    s := new(TStrings)
    s.instance = Strings_Create()
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StringsFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func StringsFromInst(inst uintptr) *TStrings {
    s := new(TStrings)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// StringsFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func StringsFromObj(obj IObject) *TStrings {
    s := new(TStrings)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StringsFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func StringsFromUnsafePointer(ptr unsafe.Pointer) *TStrings {
    s := new(TStrings)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TStrings) Free() {
    if s.instance != 0 {
        Strings_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TStrings) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TStrings) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TStrings) IsValid() bool {
    return s.instance != 0
}

// TStringsClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TStringsClass() TClass {
    return Strings_StaticClassType()
}

// Add
func (s *TStrings) Add(S string) int32 {
    return Strings_Add(s.instance, S)
}

// AddObject
func (s *TStrings) AddObject(S string, AObject IObject) int32 {
    return Strings_AddObject(s.instance, S , CheckPtr(AObject))
}

// Append
func (s *TStrings) Append(S string) {
    Strings_Append(s.instance, S)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TStrings) Assign(Source IObject) {
    Strings_Assign(s.instance, CheckPtr(Source))
}

// BeginUpdate
func (s *TStrings) BeginUpdate() {
    Strings_BeginUpdate(s.instance)
}

// Clear
// CN: 清除。
// EN: .
func (s *TStrings) Clear() {
    Strings_Clear(s.instance)
}

// Delete
func (s *TStrings) Delete(Index int32) {
    Strings_Delete(s.instance, Index)
}

// EndUpdate
func (s *TStrings) EndUpdate() {
    Strings_EndUpdate(s.instance)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TStrings) Equals(Strings IObject) bool {
    return Strings_Equals(s.instance, CheckPtr(Strings))
}

// IndexOf
func (s *TStrings) IndexOf(S string) int32 {
    return Strings_IndexOf(s.instance, S)
}

// IndexOfName
func (s *TStrings) IndexOfName(Name string) int32 {
    return Strings_IndexOfName(s.instance, Name)
}

// IndexOfObject
func (s *TStrings) IndexOfObject(AObject IObject) int32 {
    return Strings_IndexOfObject(s.instance, CheckPtr(AObject))
}

// Insert
func (s *TStrings) Insert(Index int32, S string) {
    Strings_Insert(s.instance, Index , S)
}

// InsertObject
func (s *TStrings) InsertObject(Index int32, S string, AObject IObject) {
    Strings_InsertObject(s.instance, Index , S , CheckPtr(AObject))
}

// LoadFromFile
// CN: 从文件加载。
// EN: .
func (s *TStrings) LoadFromFile(FileName string) {
    Strings_LoadFromFile(s.instance, FileName)
}

// LoadFromStream
// CN: 文件流加载。
// EN: .
func (s *TStrings) LoadFromStream(Stream IObject) {
    Strings_LoadFromStream(s.instance, CheckPtr(Stream))
}

// Move
func (s *TStrings) Move(CurIndex int32, NewIndex int32) {
    Strings_Move(s.instance, CurIndex , NewIndex)
}

// SaveToFile
// CN: 保存至文件。
// EN: .
func (s *TStrings) SaveToFile(FileName string) {
    Strings_SaveToFile(s.instance, FileName)
}

// SaveToStream
// CN: 保存至流。
// EN: .
func (s *TStrings) SaveToStream(Stream IObject) {
    Strings_SaveToStream(s.instance, CheckPtr(Stream))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TStrings) GetNamePath() string {
    return Strings_GetNamePath(s.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TStrings) DisposeOf() {
    Strings_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TStrings) ClassType() TClass {
    return Strings_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TStrings) ClassName() string {
    return Strings_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TStrings) InstanceSize() int32 {
    return Strings_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TStrings) InheritsFrom(AClass TClass) bool {
    return Strings_InheritsFrom(s.instance, AClass)
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TStrings) GetHashCode() int32 {
    return Strings_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TStrings) ToString() string {
    return Strings_ToString(s.instance)
}

// Capacity
func (s *TStrings) Capacity() int32 {
    return Strings_GetCapacity(s.instance)
}

// SetCapacity
func (s *TStrings) SetCapacity(value int32) {
    Strings_SetCapacity(s.instance, value)
}

// CommaText
func (s *TStrings) CommaText() string {
    return Strings_GetCommaText(s.instance)
}

// SetCommaText
func (s *TStrings) SetCommaText(value string) {
    Strings_SetCommaText(s.instance, value)
}

// Count
func (s *TStrings) Count() int32 {
    return Strings_GetCount(s.instance)
}

// Delimiter
func (s *TStrings) Delimiter() uint16 {
    return Strings_GetDelimiter(s.instance)
}

// SetDelimiter
func (s *TStrings) SetDelimiter(value uint16) {
    Strings_SetDelimiter(s.instance, value)
}

// Text
// CN: 获取文本。
// EN: .
func (s *TStrings) Text() string {
    return Strings_GetText(s.instance)
}

// SetText
// CN: 设置文本。
// EN: .
func (s *TStrings) SetText(value string) {
    Strings_SetText(s.instance, value)
}

// WriteBOM
func (s *TStrings) WriteBOM() bool {
    return Strings_GetWriteBOM(s.instance)
}

// SetWriteBOM
func (s *TStrings) SetWriteBOM(value bool) {
    Strings_SetWriteBOM(s.instance, value)
}

// Options
func (s *TStrings) Options() TStringsOptions {
    return Strings_GetOptions(s.instance)
}

// SetOptions
func (s *TStrings) SetOptions(value TStringsOptions) {
    Strings_SetOptions(s.instance, value)
}

// Objects
func (s *TStrings) Objects(Index int32) *TObject {
    return ObjectFromInst(Strings_GetObjects(s.instance, Index))
}

// Objects
func (s *TStrings) SetObjects(Index int32, value IObject) {
    Strings_SetObjects(s.instance, Index, CheckPtr(value))
}

// Values
func (s *TStrings) Values(Name string) string {
    return Strings_GetValues(s.instance, Name)
}

// Values
func (s *TStrings) SetValues(Name string, value string) {
    Strings_SetValues(s.instance, Name, value)
}

// ValueFromIndex
func (s *TStrings) ValueFromIndex(Index int32) string {
    return Strings_GetValueFromIndex(s.instance, Index)
}

// ValueFromIndex
func (s *TStrings) SetValueFromIndex(Index int32, value string) {
    Strings_SetValueFromIndex(s.instance, Index, value)
}

// Strings
func (s *TStrings) Strings(Index int32) string {
    return Strings_GetStrings(s.instance, Index)
}

// Strings
func (s *TStrings) SetStrings(Index int32, value string) {
    Strings_SetStrings(s.instance, Index, value)
}

