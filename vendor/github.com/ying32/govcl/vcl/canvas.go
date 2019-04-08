
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

type TCanvas struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCanvas
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCanvas() *TCanvas {
    c := new(TCanvas)
    c.instance = Canvas_Create()
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CanvasFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CanvasFromInst(inst uintptr) *TCanvas {
    c := new(TCanvas)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CanvasFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CanvasFromObj(obj IObject) *TCanvas {
    c := new(TCanvas)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CanvasFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CanvasFromUnsafePointer(ptr unsafe.Pointer) *TCanvas {
    c := new(TCanvas)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCanvas) Free() {
    if c.instance != 0 {
        Canvas_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCanvas) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCanvas) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCanvas) IsValid() bool {
    return c.instance != 0
}

// TCanvasClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCanvasClass() TClass {
    return Canvas_StaticClassType()
}

// Arc
func (c *TCanvas) Arc(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Arc(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

// ArcTo
func (c *TCanvas) ArcTo(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_ArcTo(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

// AngleArc
func (c *TCanvas) AngleArc(X int32, Y int32, Radius uint32, StartAngle float32, SweepAngle float32) {
    Canvas_AngleArc(c.instance, X , Y , Radius , StartAngle , SweepAngle)
}

// Chord
func (c *TCanvas) Chord(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Chord(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

// Ellipse
func (c *TCanvas) Ellipse(X1 int32, Y1 int32, X2 int32, Y2 int32) {
    Canvas_Ellipse(c.instance, X1 , Y1 , X2 , Y2)
}

// FloodFill
func (c *TCanvas) FloodFill(X int32, Y int32, Color TColor, FillStyle TFillStyle) {
    Canvas_FloodFill(c.instance, X , Y , Color , FillStyle)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (c *TCanvas) HandleAllocated() bool {
    return Canvas_HandleAllocated(c.instance)
}

// LineTo
func (c *TCanvas) LineTo(X int32, Y int32) {
    Canvas_LineTo(c.instance, X , Y)
}

// MoveTo
func (c *TCanvas) MoveTo(X int32, Y int32) {
    Canvas_MoveTo(c.instance, X , Y)
}

// Pie
func (c *TCanvas) Pie(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32) {
    Canvas_Pie(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3 , X4 , Y4)
}

// Rectangle
func (c *TCanvas) Rectangle(X1 int32, Y1 int32, X2 int32, Y2 int32) {
    Canvas_Rectangle(c.instance, X1 , Y1 , X2 , Y2)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (c *TCanvas) Refresh() {
    Canvas_Refresh(c.instance)
}

// RoundRect
func (c *TCanvas) RoundRect(X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32) {
    Canvas_RoundRect(c.instance, X1 , Y1 , X2 , Y2 , X3 , Y3)
}

// TextExtent
func (c *TCanvas) TextExtent(Text string) TSize {
    return Canvas_TextExtent(c.instance, Text)
}

// TextOut
func (c *TCanvas) TextOut(X int32, Y int32, Text string) {
    Canvas_TextOut(c.instance, X , Y , Text)
}

// Lock
func (c *TCanvas) Lock() {
    Canvas_Lock(c.instance)
}

// TextHeight
func (c *TCanvas) TextHeight(Text string) int32 {
    return Canvas_TextHeight(c.instance, Text)
}

// TextWidth
func (c *TCanvas) TextWidth(Text string) int32 {
    return Canvas_TextWidth(c.instance, Text)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TCanvas) Assign(Source IObject) {
    Canvas_Assign(c.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TCanvas) GetNamePath() string {
    return Canvas_GetNamePath(c.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCanvas) DisposeOf() {
    Canvas_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCanvas) ClassType() TClass {
    return Canvas_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCanvas) ClassName() string {
    return Canvas_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCanvas) InstanceSize() int32 {
    return Canvas_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCanvas) InheritsFrom(AClass TClass) bool {
    return Canvas_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCanvas) Equals(Obj IObject) bool {
    return Canvas_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCanvas) GetHashCode() int32 {
    return Canvas_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCanvas) ToString() string {
    return Canvas_ToString(c.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TCanvas) Handle() HDC {
    return Canvas_GetHandle(c.instance)
}

// SetHandle
// CN: 设置控件句柄。
// EN: Set Control handle.
func (c *TCanvas) SetHandle(value HDC) {
    Canvas_SetHandle(c.instance, value)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (c *TCanvas) Brush() *TBrush {
    return BrushFromInst(Canvas_GetBrush(c.instance))
}

// SetBrush
// CN: 设置画刷对象。
// EN: Set Brush.
func (c *TCanvas) SetBrush(value *TBrush) {
    Canvas_SetBrush(c.instance, CheckPtr(value))
}

// CopyMode
func (c *TCanvas) CopyMode() int32 {
    return Canvas_GetCopyMode(c.instance)
}

// SetCopyMode
func (c *TCanvas) SetCopyMode(value int32) {
    Canvas_SetCopyMode(c.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (c *TCanvas) Font() *TFont {
    return FontFromInst(Canvas_GetFont(c.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (c *TCanvas) SetFont(value *TFont) {
    Canvas_SetFont(c.instance, CheckPtr(value))
}

// Pen
func (c *TCanvas) Pen() *TPen {
    return PenFromInst(Canvas_GetPen(c.instance))
}

// SetPen
func (c *TCanvas) SetPen(value *TPen) {
    Canvas_SetPen(c.instance, CheckPtr(value))
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (c *TCanvas) SetOnChange(fn TNotifyEvent) {
    Canvas_SetOnChange(c.instance, fn)
}

// SetOnChanging
func (c *TCanvas) SetOnChanging(fn TNotifyEvent) {
    Canvas_SetOnChanging(c.instance, fn)
}

