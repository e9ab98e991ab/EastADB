
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

type TRichEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewRichEdit
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewRichEdit(owner IComponent) *TRichEdit {
    r := new(TRichEdit)
    r.instance = RichEdit_Create(CheckPtr(owner))
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RichEditFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func RichEditFromInst(inst uintptr) *TRichEdit {
    r := new(TRichEdit)
    r.instance = inst
    r.ptr = unsafe.Pointer(inst)
    return r
}

// RichEditFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func RichEditFromObj(obj IObject) *TRichEdit {
    r := new(TRichEdit)
    r.instance = CheckPtr(obj)
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RichEditFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func RichEditFromUnsafePointer(ptr unsafe.Pointer) *TRichEdit {
    r := new(TRichEdit)
    r.instance = uintptr(ptr)
    r.ptr = ptr
    return r
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (r *TRichEdit) Free() {
    if r.instance != 0 {
        RichEdit_Free(r.instance)
        r.instance = 0
        r.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (r *TRichEdit) Instance() uintptr {
    return r.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (r *TRichEdit) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (r *TRichEdit) IsValid() bool {
    return r.instance != 0
}

// TRichEditClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TRichEditClass() TClass {
    return RichEdit_StaticClassType()
}

// Clear
// CN: 清除。
// EN: .
func (r *TRichEdit) Clear() {
    RichEdit_Clear(r.instance)
}

// FindText
func (r *TRichEdit) FindText(SearchStr string, StartPos int32, Length int32, Options TSearchTypes) int32 {
    return RichEdit_FindText(r.instance, SearchStr , StartPos , Length , Options)
}

// Print
func (r *TRichEdit) Print(Caption string) {
    RichEdit_Print(r.instance, Caption)
}

// GetSelTextBuf
func (r *TRichEdit) GetSelTextBuf(Buffer string, BufSize int32) int32 {
    return RichEdit_GetSelTextBuf(r.instance, Buffer , BufSize)
}

// ClearSelection
// CN: 清除选择。
// EN: .
func (r *TRichEdit) ClearSelection() {
    RichEdit_ClearSelection(r.instance)
}

// CopyToClipboard
// CN: 复制到粘贴板。
// EN: .
func (r *TRichEdit) CopyToClipboard() {
    RichEdit_CopyToClipboard(r.instance)
}

// CutToClipboard
// CN: 剪切到粘贴板。
// EN: .
func (r *TRichEdit) CutToClipboard() {
    RichEdit_CutToClipboard(r.instance)
}

// PasteFromClipboard
// CN: 从剪切板粘贴。
// EN: .
func (r *TRichEdit) PasteFromClipboard() {
    RichEdit_PasteFromClipboard(r.instance)
}

// Undo
// CN: 撤销上一次操作。
// EN: .
func (r *TRichEdit) Undo() {
    RichEdit_Undo(r.instance)
}

// ClearUndo
// CN: 清除撤销。
// EN: .
func (r *TRichEdit) ClearUndo() {
    RichEdit_ClearUndo(r.instance)
}

// SelectAll
// CN: 全选。
// EN: .
func (r *TRichEdit) SelectAll() {
    RichEdit_SelectAll(r.instance)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (r *TRichEdit) CanFocus() bool {
    return RichEdit_CanFocus(r.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (r *TRichEdit) ContainsControl(Control IControl) bool {
    return RichEdit_ContainsControl(r.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (r *TRichEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(RichEdit_ControlAtPos(r.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (r *TRichEdit) DisableAlign() {
    RichEdit_DisableAlign(r.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (r *TRichEdit) EnableAlign() {
    RichEdit_EnableAlign(r.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (r *TRichEdit) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(RichEdit_FindChildControl(r.instance, ControlName))
}

// FlipChildren
func (r *TRichEdit) FlipChildren(AllLevels bool) {
    RichEdit_FlipChildren(r.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (r *TRichEdit) Focused() bool {
    return RichEdit_Focused(r.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (r *TRichEdit) HandleAllocated() bool {
    return RichEdit_HandleAllocated(r.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (r *TRichEdit) InsertControl(AControl IControl) {
    RichEdit_InsertControl(r.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (r *TRichEdit) Invalidate() {
    RichEdit_Invalidate(r.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (r *TRichEdit) PaintTo(DC HDC, X int32, Y int32) {
    RichEdit_PaintTo(r.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (r *TRichEdit) RemoveControl(AControl IControl) {
    RichEdit_RemoveControl(r.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (r *TRichEdit) Realign() {
    RichEdit_Realign(r.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (r *TRichEdit) Repaint() {
    RichEdit_Repaint(r.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (r *TRichEdit) ScaleBy(M int32, D int32) {
    RichEdit_ScaleBy(r.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (r *TRichEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    RichEdit_ScrollBy(r.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (r *TRichEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RichEdit_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (r *TRichEdit) SetFocus() {
    RichEdit_SetFocus(r.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (r *TRichEdit) Update() {
    RichEdit_Update(r.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (r *TRichEdit) UpdateControlState() {
    RichEdit_UpdateControlState(r.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (r *TRichEdit) BringToFront() {
    RichEdit_BringToFront(r.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (r *TRichEdit) ClientToScreen(Point TPoint) TPoint {
    return RichEdit_ClientToScreen(r.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (r *TRichEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return RichEdit_ClientToParent(r.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (r *TRichEdit) Dragging() bool {
    return RichEdit_Dragging(r.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (r *TRichEdit) HasParent() bool {
    return RichEdit_HasParent(r.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (r *TRichEdit) Hide() {
    RichEdit_Hide(r.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (r *TRichEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RichEdit_Perform(r.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (r *TRichEdit) Refresh() {
    RichEdit_Refresh(r.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (r *TRichEdit) ScreenToClient(Point TPoint) TPoint {
    return RichEdit_ScreenToClient(r.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (r *TRichEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return RichEdit_ParentToClient(r.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (r *TRichEdit) SendToBack() {
    RichEdit_SendToBack(r.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (r *TRichEdit) Show() {
    RichEdit_Show(r.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (r *TRichEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    return RichEdit_GetTextBuf(r.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (r *TRichEdit) GetTextLen() int32 {
    return RichEdit_GetTextLen(r.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (r *TRichEdit) SetTextBuf(Buffer string) {
    RichEdit_SetTextBuf(r.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (r *TRichEdit) FindComponent(AName string) *TComponent {
    return ComponentFromInst(RichEdit_FindComponent(r.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (r *TRichEdit) GetNamePath() string {
    return RichEdit_GetNamePath(r.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (r *TRichEdit) Assign(Source IObject) {
    RichEdit_Assign(r.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (r *TRichEdit) DisposeOf() {
    RichEdit_DisposeOf(r.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (r *TRichEdit) ClassType() TClass {
    return RichEdit_ClassType(r.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (r *TRichEdit) ClassName() string {
    return RichEdit_ClassName(r.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (r *TRichEdit) InstanceSize() int32 {
    return RichEdit_InstanceSize(r.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (r *TRichEdit) InheritsFrom(AClass TClass) bool {
    return RichEdit_InheritsFrom(r.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (r *TRichEdit) Equals(Obj IObject) bool {
    return RichEdit_Equals(r.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (r *TRichEdit) GetHashCode() int32 {
    return RichEdit_GetHashCode(r.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (r *TRichEdit) ToString() string {
    return RichEdit_ToString(r.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (r *TRichEdit) Align() TAlign {
    return RichEdit_GetAlign(r.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (r *TRichEdit) SetAlign(value TAlign) {
    RichEdit_SetAlign(r.instance, value)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (r *TRichEdit) Alignment() TAlignment {
    return RichEdit_GetAlignment(r.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (r *TRichEdit) SetAlignment(value TAlignment) {
    RichEdit_SetAlignment(r.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (r *TRichEdit) Anchors() TAnchors {
    return RichEdit_GetAnchors(r.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (r *TRichEdit) SetAnchors(value TAnchors) {
    RichEdit_SetAnchors(r.instance, value)
}

// BevelEdges
func (r *TRichEdit) BevelEdges() TBevelEdges {
    return RichEdit_GetBevelEdges(r.instance)
}

// SetBevelEdges
func (r *TRichEdit) SetBevelEdges(value TBevelEdges) {
    RichEdit_SetBevelEdges(r.instance, value)
}

// BevelInner
func (r *TRichEdit) BevelInner() TBevelCut {
    return RichEdit_GetBevelInner(r.instance)
}

// SetBevelInner
func (r *TRichEdit) SetBevelInner(value TBevelCut) {
    RichEdit_SetBevelInner(r.instance, value)
}

// BevelOuter
func (r *TRichEdit) BevelOuter() TBevelCut {
    return RichEdit_GetBevelOuter(r.instance)
}

// SetBevelOuter
func (r *TRichEdit) SetBevelOuter(value TBevelCut) {
    RichEdit_SetBevelOuter(r.instance, value)
}

// BevelKind
func (r *TRichEdit) BevelKind() TBevelKind {
    return RichEdit_GetBevelKind(r.instance)
}

// SetBevelKind
func (r *TRichEdit) SetBevelKind(value TBevelKind) {
    RichEdit_SetBevelKind(r.instance, value)
}

// BiDiMode
func (r *TRichEdit) BiDiMode() TBiDiMode {
    return RichEdit_GetBiDiMode(r.instance)
}

// SetBiDiMode
func (r *TRichEdit) SetBiDiMode(value TBiDiMode) {
    RichEdit_SetBiDiMode(r.instance, value)
}

// BorderStyle
// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (r *TRichEdit) BorderStyle() TBorderStyle {
    return RichEdit_GetBorderStyle(r.instance)
}

// SetBorderStyle
// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (r *TRichEdit) SetBorderStyle(value TBorderStyle) {
    RichEdit_SetBorderStyle(r.instance, value)
}

// BorderWidth
// CN: 获取边框的宽度。
// EN: .
func (r *TRichEdit) BorderWidth() int32 {
    return RichEdit_GetBorderWidth(r.instance)
}

// SetBorderWidth
// CN: 设置边框的宽度。
// EN: .
func (r *TRichEdit) SetBorderWidth(value int32) {
    RichEdit_SetBorderWidth(r.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (r *TRichEdit) Color() TColor {
    return RichEdit_GetColor(r.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (r *TRichEdit) SetColor(value TColor) {
    RichEdit_SetColor(r.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (r *TRichEdit) DragCursor() TCursor {
    return RichEdit_GetDragCursor(r.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (r *TRichEdit) SetDragCursor(value TCursor) {
    RichEdit_SetDragCursor(r.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (r *TRichEdit) DragKind() TDragKind {
    return RichEdit_GetDragKind(r.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (r *TRichEdit) SetDragKind(value TDragKind) {
    RichEdit_SetDragKind(r.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (r *TRichEdit) DragMode() TDragMode {
    return RichEdit_GetDragMode(r.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (r *TRichEdit) SetDragMode(value TDragMode) {
    RichEdit_SetDragMode(r.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (r *TRichEdit) Enabled() bool {
    return RichEdit_GetEnabled(r.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (r *TRichEdit) SetEnabled(value bool) {
    RichEdit_SetEnabled(r.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (r *TRichEdit) Font() *TFont {
    return FontFromInst(RichEdit_GetFont(r.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (r *TRichEdit) SetFont(value *TFont) {
    RichEdit_SetFont(r.instance, CheckPtr(value))
}

// HideSelection
// CN: 获取隐藏选择。
// EN: .
func (r *TRichEdit) HideSelection() bool {
    return RichEdit_GetHideSelection(r.instance)
}

// SetHideSelection
// CN: 设置隐藏选择。
// EN: .
func (r *TRichEdit) SetHideSelection(value bool) {
    RichEdit_SetHideSelection(r.instance, value)
}

// HideScrollBars
func (r *TRichEdit) HideScrollBars() bool {
    return RichEdit_GetHideScrollBars(r.instance)
}

// SetHideScrollBars
func (r *TRichEdit) SetHideScrollBars(value bool) {
    RichEdit_SetHideScrollBars(r.instance, value)
}

// Lines
func (r *TRichEdit) Lines() *TStrings {
    return StringsFromInst(RichEdit_GetLines(r.instance))
}

// SetLines
func (r *TRichEdit) SetLines(value IObject) {
    RichEdit_SetLines(r.instance, CheckPtr(value))
}

// MaxLength
// CN: 获取最大长度。
// EN: .
func (r *TRichEdit) MaxLength() int32 {
    return RichEdit_GetMaxLength(r.instance)
}

// SetMaxLength
// CN: 设置最大长度。
// EN: .
func (r *TRichEdit) SetMaxLength(value int32) {
    RichEdit_SetMaxLength(r.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (r *TRichEdit) ParentColor() bool {
    return RichEdit_GetParentColor(r.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (r *TRichEdit) SetParentColor(value bool) {
    RichEdit_SetParentColor(r.instance, value)
}

// ParentCtl3D
func (r *TRichEdit) ParentCtl3D() bool {
    return RichEdit_GetParentCtl3D(r.instance)
}

// SetParentCtl3D
func (r *TRichEdit) SetParentCtl3D(value bool) {
    RichEdit_SetParentCtl3D(r.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (r *TRichEdit) ParentFont() bool {
    return RichEdit_GetParentFont(r.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (r *TRichEdit) SetParentFont(value bool) {
    RichEdit_SetParentFont(r.instance, value)
}

// ParentShowHint
func (r *TRichEdit) ParentShowHint() bool {
    return RichEdit_GetParentShowHint(r.instance)
}

// SetParentShowHint
func (r *TRichEdit) SetParentShowHint(value bool) {
    RichEdit_SetParentShowHint(r.instance, value)
}

// PlainText
func (r *TRichEdit) PlainText() bool {
    return RichEdit_GetPlainText(r.instance)
}

// SetPlainText
func (r *TRichEdit) SetPlainText(value bool) {
    RichEdit_SetPlainText(r.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (r *TRichEdit) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(RichEdit_GetPopupMenu(r.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (r *TRichEdit) SetPopupMenu(value IComponent) {
    RichEdit_SetPopupMenu(r.instance, CheckPtr(value))
}

// ReadOnly
// CN: 获取只读。
// EN: .
func (r *TRichEdit) ReadOnly() bool {
    return RichEdit_GetReadOnly(r.instance)
}

// SetReadOnly
// CN: 设置只读。
// EN: .
func (r *TRichEdit) SetReadOnly(value bool) {
    RichEdit_SetReadOnly(r.instance, value)
}

// ScrollBars
func (r *TRichEdit) ScrollBars() TScrollStyle {
    return RichEdit_GetScrollBars(r.instance)
}

// SetScrollBars
func (r *TRichEdit) SetScrollBars(value TScrollStyle) {
    RichEdit_SetScrollBars(r.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (r *TRichEdit) ShowHint() bool {
    return RichEdit_GetShowHint(r.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (r *TRichEdit) SetShowHint(value bool) {
    RichEdit_SetShowHint(r.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (r *TRichEdit) TabOrder() TTabOrder {
    return RichEdit_GetTabOrder(r.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (r *TRichEdit) SetTabOrder(value TTabOrder) {
    RichEdit_SetTabOrder(r.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (r *TRichEdit) TabStop() bool {
    return RichEdit_GetTabStop(r.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (r *TRichEdit) SetTabStop(value bool) {
    RichEdit_SetTabStop(r.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (r *TRichEdit) Visible() bool {
    return RichEdit_GetVisible(r.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (r *TRichEdit) SetVisible(value bool) {
    RichEdit_SetVisible(r.instance, value)
}

// WantTabs
func (r *TRichEdit) WantTabs() bool {
    return RichEdit_GetWantTabs(r.instance)
}

// SetWantTabs
func (r *TRichEdit) SetWantTabs(value bool) {
    RichEdit_SetWantTabs(r.instance, value)
}

// WantReturns
func (r *TRichEdit) WantReturns() bool {
    return RichEdit_GetWantReturns(r.instance)
}

// SetWantReturns
func (r *TRichEdit) SetWantReturns(value bool) {
    RichEdit_SetWantReturns(r.instance, value)
}

// WordWrap
// CN: 获取自动换行。
// EN: Get Automatic line break.
func (r *TRichEdit) WordWrap() bool {
    return RichEdit_GetWordWrap(r.instance)
}

// SetWordWrap
// CN: 设置自动换行。
// EN: Set Automatic line break.
func (r *TRichEdit) SetWordWrap(value bool) {
    RichEdit_SetWordWrap(r.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (r *TRichEdit) StyleElements() TStyleElements {
    return RichEdit_GetStyleElements(r.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (r *TRichEdit) SetStyleElements(value TStyleElements) {
    RichEdit_SetStyleElements(r.instance, value)
}

// Zoom
func (r *TRichEdit) Zoom() int32 {
    return RichEdit_GetZoom(r.instance)
}

// SetZoom
func (r *TRichEdit) SetZoom(value int32) {
    RichEdit_SetZoom(r.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (r *TRichEdit) SetOnChange(fn TNotifyEvent) {
    RichEdit_SetOnChange(r.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (r *TRichEdit) SetOnClick(fn TNotifyEvent) {
    RichEdit_SetOnClick(r.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (r *TRichEdit) SetOnContextPopup(fn TContextPopupEvent) {
    RichEdit_SetOnContextPopup(r.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (r *TRichEdit) SetOnDblClick(fn TNotifyEvent) {
    RichEdit_SetOnDblClick(r.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (r *TRichEdit) SetOnDragDrop(fn TDragDropEvent) {
    RichEdit_SetOnDragDrop(r.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (r *TRichEdit) SetOnDragOver(fn TDragOverEvent) {
    RichEdit_SetOnDragOver(r.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (r *TRichEdit) SetOnEndDock(fn TEndDragEvent) {
    RichEdit_SetOnEndDock(r.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (r *TRichEdit) SetOnEndDrag(fn TEndDragEvent) {
    RichEdit_SetOnEndDrag(r.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (r *TRichEdit) SetOnEnter(fn TNotifyEvent) {
    RichEdit_SetOnEnter(r.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (r *TRichEdit) SetOnExit(fn TNotifyEvent) {
    RichEdit_SetOnExit(r.instance, fn)
}

// SetOnGesture
func (r *TRichEdit) SetOnGesture(fn TGestureEvent) {
    RichEdit_SetOnGesture(r.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (r *TRichEdit) SetOnKeyDown(fn TKeyEvent) {
    RichEdit_SetOnKeyDown(r.instance, fn)
}

// SetOnKeyPress
func (r *TRichEdit) SetOnKeyPress(fn TKeyPressEvent) {
    RichEdit_SetOnKeyPress(r.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (r *TRichEdit) SetOnKeyUp(fn TKeyEvent) {
    RichEdit_SetOnKeyUp(r.instance, fn)
}

// SetOnMouseActivate
func (r *TRichEdit) SetOnMouseActivate(fn TMouseActivateEvent) {
    RichEdit_SetOnMouseActivate(r.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (r *TRichEdit) SetOnMouseDown(fn TMouseEvent) {
    RichEdit_SetOnMouseDown(r.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (r *TRichEdit) SetOnMouseEnter(fn TNotifyEvent) {
    RichEdit_SetOnMouseEnter(r.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (r *TRichEdit) SetOnMouseLeave(fn TNotifyEvent) {
    RichEdit_SetOnMouseLeave(r.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (r *TRichEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    RichEdit_SetOnMouseMove(r.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (r *TRichEdit) SetOnMouseUp(fn TMouseEvent) {
    RichEdit_SetOnMouseUp(r.instance, fn)
}

// SetOnMouseWheel
// CN: 设置鼠标滚轮事件。
// EN: .
func (r *TRichEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
    RichEdit_SetOnMouseWheel(r.instance, fn)
}

// SetOnMouseWheelDown
// CN: 设置鼠标滚轮按下事件。
// EN: .
func (r *TRichEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    RichEdit_SetOnMouseWheelDown(r.instance, fn)
}

// SetOnMouseWheelUp
// CN: 设置鼠标滚轮抬起事件。
// EN: .
func (r *TRichEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    RichEdit_SetOnMouseWheelUp(r.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (r *TRichEdit) SetOnStartDock(fn TStartDockEvent) {
    RichEdit_SetOnStartDock(r.instance, fn)
}

// ActiveLineNo
func (r *TRichEdit) ActiveLineNo() uint32 {
    return RichEdit_GetActiveLineNo(r.instance)
}

// DefAttributes
func (r *TRichEdit) DefAttributes() *TTextAttributes {
    return TextAttributesFromInst(RichEdit_GetDefAttributes(r.instance))
}

// SetDefAttributes
func (r *TRichEdit) SetDefAttributes(value *TTextAttributes) {
    RichEdit_SetDefAttributes(r.instance, CheckPtr(value))
}

// SelAttributes
func (r *TRichEdit) SelAttributes() *TTextAttributes {
    return TextAttributesFromInst(RichEdit_GetSelAttributes(r.instance))
}

// SetSelAttributes
func (r *TRichEdit) SetSelAttributes(value *TTextAttributes) {
    RichEdit_SetSelAttributes(r.instance, CheckPtr(value))
}

// PageRect
func (r *TRichEdit) PageRect() TRect {
    return RichEdit_GetPageRect(r.instance)
}

// SetPageRect
func (r *TRichEdit) SetPageRect(value TRect) {
    RichEdit_SetPageRect(r.instance, value)
}

// Paragraph
func (r *TRichEdit) Paragraph() *TParaAttributes {
    return ParaAttributesFromInst(RichEdit_GetParagraph(r.instance))
}

// CaretPos
func (r *TRichEdit) CaretPos() TPoint {
    return RichEdit_GetCaretPos(r.instance)
}

// SetCaretPos
func (r *TRichEdit) SetCaretPos(value TPoint) {
    RichEdit_SetCaretPos(r.instance, value)
}

// CanUndo
// CN: 获取能否撤销。
// EN: .
func (r *TRichEdit) CanUndo() bool {
    return RichEdit_GetCanUndo(r.instance)
}

// Modified
// CN: 获取修改。
// EN: Get modified.
func (r *TRichEdit) Modified() bool {
    return RichEdit_GetModified(r.instance)
}

// SetModified
// CN: 设置修改。
// EN: Set modified.
func (r *TRichEdit) SetModified(value bool) {
    RichEdit_SetModified(r.instance, value)
}

// SelLength
// CN: 获取选择的长度。
// EN: .
func (r *TRichEdit) SelLength() int32 {
    return RichEdit_GetSelLength(r.instance)
}

// SetSelLength
// CN: 设置选择的长度。
// EN: .
func (r *TRichEdit) SetSelLength(value int32) {
    RichEdit_SetSelLength(r.instance, value)
}

// SelStart
// CN: 获取选择的启始位置。
// EN: .
func (r *TRichEdit) SelStart() int32 {
    return RichEdit_GetSelStart(r.instance)
}

// SetSelStart
// CN: 设置选择的启始位置。
// EN: .
func (r *TRichEdit) SetSelStart(value int32) {
    RichEdit_SetSelStart(r.instance, value)
}

// SelText
// CN: 获取选择的文本。
// EN: .
func (r *TRichEdit) SelText() string {
    return RichEdit_GetSelText(r.instance)
}

// SetSelText
// CN: 设置选择的文本。
// EN: .
func (r *TRichEdit) SetSelText(value string) {
    RichEdit_SetSelText(r.instance, value)
}

// Text
// CN: 获取文本。
// EN: .
func (r *TRichEdit) Text() string {
    return RichEdit_GetText(r.instance)
}

// SetText
// CN: 设置文本。
// EN: .
func (r *TRichEdit) SetText(value string) {
    RichEdit_SetText(r.instance, value)
}

// TextHint
// CN: 获取提示文本。
// EN: .
func (r *TRichEdit) TextHint() string {
    return RichEdit_GetTextHint(r.instance)
}

// SetTextHint
// CN: 设置提示文本。
// EN: .
func (r *TRichEdit) SetTextHint(value string) {
    RichEdit_SetTextHint(r.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (r *TRichEdit) DockClientCount() int32 {
    return RichEdit_GetDockClientCount(r.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (r *TRichEdit) DockSite() bool {
    return RichEdit_GetDockSite(r.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (r *TRichEdit) SetDockSite(value bool) {
    RichEdit_SetDockSite(r.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (r *TRichEdit) DoubleBuffered() bool {
    return RichEdit_GetDoubleBuffered(r.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (r *TRichEdit) SetDoubleBuffered(value bool) {
    RichEdit_SetDoubleBuffered(r.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (r *TRichEdit) AlignDisabled() bool {
    return RichEdit_GetAlignDisabled(r.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (r *TRichEdit) MouseInClient() bool {
    return RichEdit_GetMouseInClient(r.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (r *TRichEdit) VisibleDockClientCount() int32 {
    return RichEdit_GetVisibleDockClientCount(r.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (r *TRichEdit) Brush() *TBrush {
    return BrushFromInst(RichEdit_GetBrush(r.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (r *TRichEdit) ControlCount() int32 {
    return RichEdit_GetControlCount(r.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (r *TRichEdit) Handle() HWND {
    return RichEdit_GetHandle(r.instance)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (r *TRichEdit) ParentDoubleBuffered() bool {
    return RichEdit_GetParentDoubleBuffered(r.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (r *TRichEdit) SetParentDoubleBuffered(value bool) {
    RichEdit_SetParentDoubleBuffered(r.instance, value)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (r *TRichEdit) ParentWindow() HWND {
    return RichEdit_GetParentWindow(r.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (r *TRichEdit) SetParentWindow(value HWND) {
    RichEdit_SetParentWindow(r.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (r *TRichEdit) UseDockManager() bool {
    return RichEdit_GetUseDockManager(r.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (r *TRichEdit) SetUseDockManager(value bool) {
    RichEdit_SetUseDockManager(r.instance, value)
}

// Action
func (r *TRichEdit) Action() *TAction {
    return ActionFromInst(RichEdit_GetAction(r.instance))
}

// SetAction
func (r *TRichEdit) SetAction(value IComponent) {
    RichEdit_SetAction(r.instance, CheckPtr(value))
}

// BoundsRect
func (r *TRichEdit) BoundsRect() TRect {
    return RichEdit_GetBoundsRect(r.instance)
}

// SetBoundsRect
func (r *TRichEdit) SetBoundsRect(value TRect) {
    RichEdit_SetBoundsRect(r.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (r *TRichEdit) ClientHeight() int32 {
    return RichEdit_GetClientHeight(r.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (r *TRichEdit) SetClientHeight(value int32) {
    RichEdit_SetClientHeight(r.instance, value)
}

// ClientOrigin
func (r *TRichEdit) ClientOrigin() TPoint {
    return RichEdit_GetClientOrigin(r.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (r *TRichEdit) ClientRect() TRect {
    return RichEdit_GetClientRect(r.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (r *TRichEdit) ClientWidth() int32 {
    return RichEdit_GetClientWidth(r.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (r *TRichEdit) SetClientWidth(value int32) {
    RichEdit_SetClientWidth(r.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (r *TRichEdit) ControlState() TControlState {
    return RichEdit_GetControlState(r.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (r *TRichEdit) SetControlState(value TControlState) {
    RichEdit_SetControlState(r.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (r *TRichEdit) ControlStyle() TControlStyle {
    return RichEdit_GetControlStyle(r.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (r *TRichEdit) SetControlStyle(value TControlStyle) {
    RichEdit_SetControlStyle(r.instance, value)
}

// ExplicitLeft
func (r *TRichEdit) ExplicitLeft() int32 {
    return RichEdit_GetExplicitLeft(r.instance)
}

// ExplicitTop
func (r *TRichEdit) ExplicitTop() int32 {
    return RichEdit_GetExplicitTop(r.instance)
}

// ExplicitWidth
func (r *TRichEdit) ExplicitWidth() int32 {
    return RichEdit_GetExplicitWidth(r.instance)
}

// ExplicitHeight
func (r *TRichEdit) ExplicitHeight() int32 {
    return RichEdit_GetExplicitHeight(r.instance)
}

// Floating
func (r *TRichEdit) Floating() bool {
    return RichEdit_GetFloating(r.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (r *TRichEdit) Parent() *TWinControl {
    return WinControlFromInst(RichEdit_GetParent(r.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (r *TRichEdit) SetParent(value IWinControl) {
    RichEdit_SetParent(r.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (r *TRichEdit) AlignWithMargins() bool {
    return RichEdit_GetAlignWithMargins(r.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (r *TRichEdit) SetAlignWithMargins(value bool) {
    RichEdit_SetAlignWithMargins(r.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (r *TRichEdit) Left() int32 {
    return RichEdit_GetLeft(r.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (r *TRichEdit) SetLeft(value int32) {
    RichEdit_SetLeft(r.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (r *TRichEdit) Top() int32 {
    return RichEdit_GetTop(r.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (r *TRichEdit) SetTop(value int32) {
    RichEdit_SetTop(r.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (r *TRichEdit) Width() int32 {
    return RichEdit_GetWidth(r.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (r *TRichEdit) SetWidth(value int32) {
    RichEdit_SetWidth(r.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (r *TRichEdit) Height() int32 {
    return RichEdit_GetHeight(r.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (r *TRichEdit) SetHeight(value int32) {
    RichEdit_SetHeight(r.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (r *TRichEdit) Cursor() TCursor {
    return RichEdit_GetCursor(r.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (r *TRichEdit) SetCursor(value TCursor) {
    RichEdit_SetCursor(r.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (r *TRichEdit) Hint() string {
    return RichEdit_GetHint(r.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (r *TRichEdit) SetHint(value string) {
    RichEdit_SetHint(r.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (r *TRichEdit) Margins() *TMargins {
    return MarginsFromInst(RichEdit_GetMargins(r.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (r *TRichEdit) SetMargins(value *TMargins) {
    RichEdit_SetMargins(r.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (r *TRichEdit) CustomHint() *TCustomHint {
    return CustomHintFromInst(RichEdit_GetCustomHint(r.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (r *TRichEdit) SetCustomHint(value IComponent) {
    RichEdit_SetCustomHint(r.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (r *TRichEdit) ComponentCount() int32 {
    return RichEdit_GetComponentCount(r.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (r *TRichEdit) ComponentIndex() int32 {
    return RichEdit_GetComponentIndex(r.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (r *TRichEdit) SetComponentIndex(value int32) {
    RichEdit_SetComponentIndex(r.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (r *TRichEdit) Owner() *TComponent {
    return ComponentFromInst(RichEdit_GetOwner(r.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (r *TRichEdit) Name() string {
    return RichEdit_GetName(r.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (r *TRichEdit) SetName(value string) {
    RichEdit_SetName(r.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (r *TRichEdit) Tag() int {
    return RichEdit_GetTag(r.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (r *TRichEdit) SetTag(value int) {
    RichEdit_SetTag(r.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (r *TRichEdit) DockClients(Index int32) *TControl {
    return ControlFromInst(RichEdit_GetDockClients(r.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (r *TRichEdit) Controls(Index int32) *TControl {
    return ControlFromInst(RichEdit_GetControls(r.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (r *TRichEdit) Components(AIndex int32) *TComponent {
    return ComponentFromInst(RichEdit_GetComponents(r.instance, AIndex))
}

