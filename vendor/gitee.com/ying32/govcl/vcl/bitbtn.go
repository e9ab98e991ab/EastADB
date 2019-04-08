
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

type TBitBtn struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewBitBtn
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBitBtn(owner IComponent) *TBitBtn {
    b := new(TBitBtn)
    b.instance = BitBtn_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BitBtnFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func BitBtnFromInst(inst uintptr) *TBitBtn {
    b := new(TBitBtn)
    b.instance = inst
    b.ptr = unsafe.Pointer(inst)
    return b
}

// BitBtnFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func BitBtnFromObj(obj IObject) *TBitBtn {
    b := new(TBitBtn)
    b.instance = CheckPtr(obj)
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BitBtnFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func BitBtnFromUnsafePointer(ptr unsafe.Pointer) *TBitBtn {
    b := new(TBitBtn)
    b.instance = uintptr(ptr)
    b.ptr = ptr
    return b
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (b *TBitBtn) Free() {
    if b.instance != 0 {
        BitBtn_Free(b.instance)
        b.instance = 0
        b.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBitBtn) Instance() uintptr {
    return b.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBitBtn) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBitBtn) IsValid() bool {
    return b.instance != 0
}

// TBitBtnClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBitBtnClass() TClass {
    return BitBtn_StaticClassType()
}

// Click
// CN: 单击。
// EN: .
func (b *TBitBtn) Click() {
    BitBtn_Click(b.instance)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (b *TBitBtn) CanFocus() bool {
    return BitBtn_CanFocus(b.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (b *TBitBtn) ContainsControl(Control IControl) bool {
    return BitBtn_ContainsControl(b.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (b *TBitBtn) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(BitBtn_ControlAtPos(b.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (b *TBitBtn) DisableAlign() {
    BitBtn_DisableAlign(b.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (b *TBitBtn) EnableAlign() {
    BitBtn_EnableAlign(b.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (b *TBitBtn) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(BitBtn_FindChildControl(b.instance, ControlName))
}

// FlipChildren
func (b *TBitBtn) FlipChildren(AllLevels bool) {
    BitBtn_FlipChildren(b.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (b *TBitBtn) Focused() bool {
    return BitBtn_Focused(b.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (b *TBitBtn) HandleAllocated() bool {
    return BitBtn_HandleAllocated(b.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (b *TBitBtn) InsertControl(AControl IControl) {
    BitBtn_InsertControl(b.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (b *TBitBtn) Invalidate() {
    BitBtn_Invalidate(b.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (b *TBitBtn) PaintTo(DC HDC, X int32, Y int32) {
    BitBtn_PaintTo(b.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (b *TBitBtn) RemoveControl(AControl IControl) {
    BitBtn_RemoveControl(b.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (b *TBitBtn) Realign() {
    BitBtn_Realign(b.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (b *TBitBtn) Repaint() {
    BitBtn_Repaint(b.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (b *TBitBtn) ScaleBy(M int32, D int32) {
    BitBtn_ScaleBy(b.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (b *TBitBtn) ScrollBy(DeltaX int32, DeltaY int32) {
    BitBtn_ScrollBy(b.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (b *TBitBtn) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    BitBtn_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (b *TBitBtn) SetFocus() {
    BitBtn_SetFocus(b.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (b *TBitBtn) Update() {
    BitBtn_Update(b.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (b *TBitBtn) UpdateControlState() {
    BitBtn_UpdateControlState(b.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (b *TBitBtn) BringToFront() {
    BitBtn_BringToFront(b.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (b *TBitBtn) ClientToScreen(Point TPoint) TPoint {
    return BitBtn_ClientToScreen(b.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (b *TBitBtn) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return BitBtn_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (b *TBitBtn) Dragging() bool {
    return BitBtn_Dragging(b.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (b *TBitBtn) HasParent() bool {
    return BitBtn_HasParent(b.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (b *TBitBtn) Hide() {
    BitBtn_Hide(b.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (b *TBitBtn) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return BitBtn_Perform(b.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (b *TBitBtn) Refresh() {
    BitBtn_Refresh(b.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (b *TBitBtn) ScreenToClient(Point TPoint) TPoint {
    return BitBtn_ScreenToClient(b.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (b *TBitBtn) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return BitBtn_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (b *TBitBtn) SendToBack() {
    BitBtn_SendToBack(b.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (b *TBitBtn) Show() {
    BitBtn_Show(b.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (b *TBitBtn) GetTextBuf(Buffer string, BufSize int32) int32 {
    return BitBtn_GetTextBuf(b.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (b *TBitBtn) GetTextLen() int32 {
    return BitBtn_GetTextLen(b.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (b *TBitBtn) SetTextBuf(Buffer string) {
    BitBtn_SetTextBuf(b.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (b *TBitBtn) FindComponent(AName string) *TComponent {
    return ComponentFromInst(BitBtn_FindComponent(b.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (b *TBitBtn) GetNamePath() string {
    return BitBtn_GetNamePath(b.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (b *TBitBtn) Assign(Source IObject) {
    BitBtn_Assign(b.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (b *TBitBtn) DisposeOf() {
    BitBtn_DisposeOf(b.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBitBtn) ClassType() TClass {
    return BitBtn_ClassType(b.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBitBtn) ClassName() string {
    return BitBtn_ClassName(b.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBitBtn) InstanceSize() int32 {
    return BitBtn_InstanceSize(b.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBitBtn) InheritsFrom(AClass TClass) bool {
    return BitBtn_InheritsFrom(b.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBitBtn) Equals(Obj IObject) bool {
    return BitBtn_Equals(b.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBitBtn) GetHashCode() int32 {
    return BitBtn_GetHashCode(b.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (b *TBitBtn) ToString() string {
    return BitBtn_ToString(b.instance)
}

// Action
func (b *TBitBtn) Action() *TAction {
    return ActionFromInst(BitBtn_GetAction(b.instance))
}

// SetAction
func (b *TBitBtn) SetAction(value IComponent) {
    BitBtn_SetAction(b.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (b *TBitBtn) Align() TAlign {
    return BitBtn_GetAlign(b.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (b *TBitBtn) SetAlign(value TAlign) {
    BitBtn_SetAlign(b.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (b *TBitBtn) Anchors() TAnchors {
    return BitBtn_GetAnchors(b.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (b *TBitBtn) SetAnchors(value TAnchors) {
    BitBtn_SetAnchors(b.instance, value)
}

// BiDiMode
func (b *TBitBtn) BiDiMode() TBiDiMode {
    return BitBtn_GetBiDiMode(b.instance)
}

// SetBiDiMode
func (b *TBitBtn) SetBiDiMode(value TBiDiMode) {
    BitBtn_SetBiDiMode(b.instance, value)
}

// Cancel
func (b *TBitBtn) Cancel() bool {
    return BitBtn_GetCancel(b.instance)
}

// SetCancel
func (b *TBitBtn) SetCancel(value bool) {
    BitBtn_SetCancel(b.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (b *TBitBtn) Caption() string {
    return BitBtn_GetCaption(b.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (b *TBitBtn) SetCaption(value string) {
    BitBtn_SetCaption(b.instance, value)
}

// Default
func (b *TBitBtn) Default() bool {
    return BitBtn_GetDefault(b.instance)
}

// SetDefault
func (b *TBitBtn) SetDefault(value bool) {
    BitBtn_SetDefault(b.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (b *TBitBtn) DoubleBuffered() bool {
    return BitBtn_GetDoubleBuffered(b.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (b *TBitBtn) SetDoubleBuffered(value bool) {
    BitBtn_SetDoubleBuffered(b.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (b *TBitBtn) DragCursor() TCursor {
    return BitBtn_GetDragCursor(b.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (b *TBitBtn) SetDragCursor(value TCursor) {
    BitBtn_SetDragCursor(b.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (b *TBitBtn) DragKind() TDragKind {
    return BitBtn_GetDragKind(b.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (b *TBitBtn) SetDragKind(value TDragKind) {
    BitBtn_SetDragKind(b.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (b *TBitBtn) DragMode() TDragMode {
    return BitBtn_GetDragMode(b.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (b *TBitBtn) SetDragMode(value TDragMode) {
    BitBtn_SetDragMode(b.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (b *TBitBtn) Enabled() bool {
    return BitBtn_GetEnabled(b.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (b *TBitBtn) SetEnabled(value bool) {
    BitBtn_SetEnabled(b.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (b *TBitBtn) Font() *TFont {
    return FontFromInst(BitBtn_GetFont(b.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (b *TBitBtn) SetFont(value *TFont) {
    BitBtn_SetFont(b.instance, CheckPtr(value))
}

// Glyph
func (b *TBitBtn) Glyph() *TBitmap {
    return BitmapFromInst(BitBtn_GetGlyph(b.instance))
}

// SetGlyph
func (b *TBitBtn) SetGlyph(value *TBitmap) {
    BitBtn_SetGlyph(b.instance, CheckPtr(value))
}

// Kind
func (b *TBitBtn) Kind() TBitBtnKind {
    return BitBtn_GetKind(b.instance)
}

// SetKind
func (b *TBitBtn) SetKind(value TBitBtnKind) {
    BitBtn_SetKind(b.instance, value)
}

// Layout
func (b *TBitBtn) Layout() TButtonLayout {
    return BitBtn_GetLayout(b.instance)
}

// SetLayout
func (b *TBitBtn) SetLayout(value TButtonLayout) {
    BitBtn_SetLayout(b.instance, value)
}

// ModalResult
// CN: 获取模态对话框显示结果。
// EN: .
func (b *TBitBtn) ModalResult() TModalResult {
    return BitBtn_GetModalResult(b.instance)
}

// SetModalResult
// CN: 设置模态对话框显示结果。
// EN: .
func (b *TBitBtn) SetModalResult(value TModalResult) {
    BitBtn_SetModalResult(b.instance, value)
}

// NumGlyphs
func (b *TBitBtn) NumGlyphs() TNumGlyphs {
    return BitBtn_GetNumGlyphs(b.instance)
}

// SetNumGlyphs
func (b *TBitBtn) SetNumGlyphs(value TNumGlyphs) {
    BitBtn_SetNumGlyphs(b.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (b *TBitBtn) ParentDoubleBuffered() bool {
    return BitBtn_GetParentDoubleBuffered(b.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (b *TBitBtn) SetParentDoubleBuffered(value bool) {
    BitBtn_SetParentDoubleBuffered(b.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (b *TBitBtn) ParentFont() bool {
    return BitBtn_GetParentFont(b.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (b *TBitBtn) SetParentFont(value bool) {
    BitBtn_SetParentFont(b.instance, value)
}

// ParentShowHint
func (b *TBitBtn) ParentShowHint() bool {
    return BitBtn_GetParentShowHint(b.instance)
}

// SetParentShowHint
func (b *TBitBtn) SetParentShowHint(value bool) {
    BitBtn_SetParentShowHint(b.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (b *TBitBtn) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(BitBtn_GetPopupMenu(b.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (b *TBitBtn) SetPopupMenu(value IComponent) {
    BitBtn_SetPopupMenu(b.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (b *TBitBtn) ShowHint() bool {
    return BitBtn_GetShowHint(b.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (b *TBitBtn) SetShowHint(value bool) {
    BitBtn_SetShowHint(b.instance, value)
}

// Style
func (b *TBitBtn) Style() TButtonStyle {
    return BitBtn_GetStyle(b.instance)
}

// SetStyle
func (b *TBitBtn) SetStyle(value TButtonStyle) {
    BitBtn_SetStyle(b.instance, value)
}

// Spacing
func (b *TBitBtn) Spacing() int32 {
    return BitBtn_GetSpacing(b.instance)
}

// SetSpacing
func (b *TBitBtn) SetSpacing(value int32) {
    BitBtn_SetSpacing(b.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (b *TBitBtn) TabOrder() TTabOrder {
    return BitBtn_GetTabOrder(b.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (b *TBitBtn) SetTabOrder(value TTabOrder) {
    BitBtn_SetTabOrder(b.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (b *TBitBtn) TabStop() bool {
    return BitBtn_GetTabStop(b.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (b *TBitBtn) SetTabStop(value bool) {
    BitBtn_SetTabStop(b.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (b *TBitBtn) Visible() bool {
    return BitBtn_GetVisible(b.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (b *TBitBtn) SetVisible(value bool) {
    BitBtn_SetVisible(b.instance, value)
}

// WordWrap
// CN: 获取自动换行。
// EN: Get Automatic line break.
func (b *TBitBtn) WordWrap() bool {
    return BitBtn_GetWordWrap(b.instance)
}

// SetWordWrap
// CN: 设置自动换行。
// EN: Set Automatic line break.
func (b *TBitBtn) SetWordWrap(value bool) {
    BitBtn_SetWordWrap(b.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (b *TBitBtn) StyleElements() TStyleElements {
    return BitBtn_GetStyleElements(b.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (b *TBitBtn) SetStyleElements(value TStyleElements) {
    BitBtn_SetStyleElements(b.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (b *TBitBtn) SetOnClick(fn TNotifyEvent) {
    BitBtn_SetOnClick(b.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (b *TBitBtn) SetOnContextPopup(fn TContextPopupEvent) {
    BitBtn_SetOnContextPopup(b.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (b *TBitBtn) SetOnDragDrop(fn TDragDropEvent) {
    BitBtn_SetOnDragDrop(b.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (b *TBitBtn) SetOnDragOver(fn TDragOverEvent) {
    BitBtn_SetOnDragOver(b.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (b *TBitBtn) SetOnEndDock(fn TEndDragEvent) {
    BitBtn_SetOnEndDock(b.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (b *TBitBtn) SetOnEndDrag(fn TEndDragEvent) {
    BitBtn_SetOnEndDrag(b.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (b *TBitBtn) SetOnEnter(fn TNotifyEvent) {
    BitBtn_SetOnEnter(b.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (b *TBitBtn) SetOnExit(fn TNotifyEvent) {
    BitBtn_SetOnExit(b.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (b *TBitBtn) SetOnKeyDown(fn TKeyEvent) {
    BitBtn_SetOnKeyDown(b.instance, fn)
}

// SetOnKeyPress
func (b *TBitBtn) SetOnKeyPress(fn TKeyPressEvent) {
    BitBtn_SetOnKeyPress(b.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (b *TBitBtn) SetOnKeyUp(fn TKeyEvent) {
    BitBtn_SetOnKeyUp(b.instance, fn)
}

// SetOnMouseActivate
func (b *TBitBtn) SetOnMouseActivate(fn TMouseActivateEvent) {
    BitBtn_SetOnMouseActivate(b.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (b *TBitBtn) SetOnMouseDown(fn TMouseEvent) {
    BitBtn_SetOnMouseDown(b.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (b *TBitBtn) SetOnMouseEnter(fn TNotifyEvent) {
    BitBtn_SetOnMouseEnter(b.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (b *TBitBtn) SetOnMouseLeave(fn TNotifyEvent) {
    BitBtn_SetOnMouseLeave(b.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (b *TBitBtn) SetOnMouseMove(fn TMouseMoveEvent) {
    BitBtn_SetOnMouseMove(b.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (b *TBitBtn) SetOnMouseUp(fn TMouseEvent) {
    BitBtn_SetOnMouseUp(b.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (b *TBitBtn) SetOnStartDock(fn TStartDockEvent) {
    BitBtn_SetOnStartDock(b.instance, fn)
}

// CommandLinkHint
func (b *TBitBtn) CommandLinkHint() string {
    return BitBtn_GetCommandLinkHint(b.instance)
}

// SetCommandLinkHint
func (b *TBitBtn) SetCommandLinkHint(value string) {
    BitBtn_SetCommandLinkHint(b.instance, value)
}

// DisabledImageIndex
func (b *TBitBtn) DisabledImageIndex() int32 {
    return BitBtn_GetDisabledImageIndex(b.instance)
}

// SetDisabledImageIndex
func (b *TBitBtn) SetDisabledImageIndex(value int32) {
    BitBtn_SetDisabledImageIndex(b.instance, value)
}

// ElevationRequired
func (b *TBitBtn) ElevationRequired() bool {
    return BitBtn_GetElevationRequired(b.instance)
}

// SetElevationRequired
func (b *TBitBtn) SetElevationRequired(value bool) {
    BitBtn_SetElevationRequired(b.instance, value)
}

// HotImageIndex
func (b *TBitBtn) HotImageIndex() int32 {
    return BitBtn_GetHotImageIndex(b.instance)
}

// SetHotImageIndex
func (b *TBitBtn) SetHotImageIndex(value int32) {
    BitBtn_SetHotImageIndex(b.instance, value)
}

// Images
// CN: 获取图标索引列表对象。
// EN: .
func (b *TBitBtn) Images() *TImageList {
    return ImageListFromInst(BitBtn_GetImages(b.instance))
}

// SetImages
// CN: 设置图标索引列表对象。
// EN: .
func (b *TBitBtn) SetImages(value IComponent) {
    BitBtn_SetImages(b.instance, CheckPtr(value))
}

// ImageAlignment
func (b *TBitBtn) ImageAlignment() TImageAlignment {
    return BitBtn_GetImageAlignment(b.instance)
}

// SetImageAlignment
func (b *TBitBtn) SetImageAlignment(value TImageAlignment) {
    BitBtn_SetImageAlignment(b.instance, value)
}

// ImageIndex
// CN: 获取图像在images中的索引。
// EN: .
func (b *TBitBtn) ImageIndex() int32 {
    return BitBtn_GetImageIndex(b.instance)
}

// SetImageIndex
// CN: 设置图像在images中的索引。
// EN: .
func (b *TBitBtn) SetImageIndex(value int32) {
    BitBtn_SetImageIndex(b.instance, value)
}

// PressedImageIndex
func (b *TBitBtn) PressedImageIndex() int32 {
    return BitBtn_GetPressedImageIndex(b.instance)
}

// SetPressedImageIndex
func (b *TBitBtn) SetPressedImageIndex(value int32) {
    BitBtn_SetPressedImageIndex(b.instance, value)
}

// SelectedImageIndex
func (b *TBitBtn) SelectedImageIndex() int32 {
    return BitBtn_GetSelectedImageIndex(b.instance)
}

// SetSelectedImageIndex
func (b *TBitBtn) SetSelectedImageIndex(value int32) {
    BitBtn_SetSelectedImageIndex(b.instance, value)
}

// StylusHotImageIndex
func (b *TBitBtn) StylusHotImageIndex() int32 {
    return BitBtn_GetStylusHotImageIndex(b.instance)
}

// SetStylusHotImageIndex
func (b *TBitBtn) SetStylusHotImageIndex(value int32) {
    BitBtn_SetStylusHotImageIndex(b.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (b *TBitBtn) DockClientCount() int32 {
    return BitBtn_GetDockClientCount(b.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (b *TBitBtn) DockSite() bool {
    return BitBtn_GetDockSite(b.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (b *TBitBtn) SetDockSite(value bool) {
    BitBtn_SetDockSite(b.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (b *TBitBtn) AlignDisabled() bool {
    return BitBtn_GetAlignDisabled(b.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (b *TBitBtn) MouseInClient() bool {
    return BitBtn_GetMouseInClient(b.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (b *TBitBtn) VisibleDockClientCount() int32 {
    return BitBtn_GetVisibleDockClientCount(b.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (b *TBitBtn) Brush() *TBrush {
    return BrushFromInst(BitBtn_GetBrush(b.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (b *TBitBtn) ControlCount() int32 {
    return BitBtn_GetControlCount(b.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (b *TBitBtn) Handle() HWND {
    return BitBtn_GetHandle(b.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (b *TBitBtn) ParentWindow() HWND {
    return BitBtn_GetParentWindow(b.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (b *TBitBtn) SetParentWindow(value HWND) {
    BitBtn_SetParentWindow(b.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (b *TBitBtn) UseDockManager() bool {
    return BitBtn_GetUseDockManager(b.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (b *TBitBtn) SetUseDockManager(value bool) {
    BitBtn_SetUseDockManager(b.instance, value)
}

// BoundsRect
func (b *TBitBtn) BoundsRect() TRect {
    return BitBtn_GetBoundsRect(b.instance)
}

// SetBoundsRect
func (b *TBitBtn) SetBoundsRect(value TRect) {
    BitBtn_SetBoundsRect(b.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (b *TBitBtn) ClientHeight() int32 {
    return BitBtn_GetClientHeight(b.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (b *TBitBtn) SetClientHeight(value int32) {
    BitBtn_SetClientHeight(b.instance, value)
}

// ClientOrigin
func (b *TBitBtn) ClientOrigin() TPoint {
    return BitBtn_GetClientOrigin(b.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (b *TBitBtn) ClientRect() TRect {
    return BitBtn_GetClientRect(b.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (b *TBitBtn) ClientWidth() int32 {
    return BitBtn_GetClientWidth(b.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (b *TBitBtn) SetClientWidth(value int32) {
    BitBtn_SetClientWidth(b.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (b *TBitBtn) ControlState() TControlState {
    return BitBtn_GetControlState(b.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (b *TBitBtn) SetControlState(value TControlState) {
    BitBtn_SetControlState(b.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (b *TBitBtn) ControlStyle() TControlStyle {
    return BitBtn_GetControlStyle(b.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (b *TBitBtn) SetControlStyle(value TControlStyle) {
    BitBtn_SetControlStyle(b.instance, value)
}

// ExplicitLeft
func (b *TBitBtn) ExplicitLeft() int32 {
    return BitBtn_GetExplicitLeft(b.instance)
}

// ExplicitTop
func (b *TBitBtn) ExplicitTop() int32 {
    return BitBtn_GetExplicitTop(b.instance)
}

// ExplicitWidth
func (b *TBitBtn) ExplicitWidth() int32 {
    return BitBtn_GetExplicitWidth(b.instance)
}

// ExplicitHeight
func (b *TBitBtn) ExplicitHeight() int32 {
    return BitBtn_GetExplicitHeight(b.instance)
}

// Floating
func (b *TBitBtn) Floating() bool {
    return BitBtn_GetFloating(b.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (b *TBitBtn) Parent() *TWinControl {
    return WinControlFromInst(BitBtn_GetParent(b.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (b *TBitBtn) SetParent(value IWinControl) {
    BitBtn_SetParent(b.instance, CheckPtr(value))
}

// SetOnGesture
func (b *TBitBtn) SetOnGesture(fn TGestureEvent) {
    BitBtn_SetOnGesture(b.instance, fn)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (b *TBitBtn) AlignWithMargins() bool {
    return BitBtn_GetAlignWithMargins(b.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (b *TBitBtn) SetAlignWithMargins(value bool) {
    BitBtn_SetAlignWithMargins(b.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (b *TBitBtn) Left() int32 {
    return BitBtn_GetLeft(b.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (b *TBitBtn) SetLeft(value int32) {
    BitBtn_SetLeft(b.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (b *TBitBtn) Top() int32 {
    return BitBtn_GetTop(b.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (b *TBitBtn) SetTop(value int32) {
    BitBtn_SetTop(b.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (b *TBitBtn) Width() int32 {
    return BitBtn_GetWidth(b.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (b *TBitBtn) SetWidth(value int32) {
    BitBtn_SetWidth(b.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (b *TBitBtn) Height() int32 {
    return BitBtn_GetHeight(b.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (b *TBitBtn) SetHeight(value int32) {
    BitBtn_SetHeight(b.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (b *TBitBtn) Cursor() TCursor {
    return BitBtn_GetCursor(b.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (b *TBitBtn) SetCursor(value TCursor) {
    BitBtn_SetCursor(b.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (b *TBitBtn) Hint() string {
    return BitBtn_GetHint(b.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (b *TBitBtn) SetHint(value string) {
    BitBtn_SetHint(b.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (b *TBitBtn) Margins() *TMargins {
    return MarginsFromInst(BitBtn_GetMargins(b.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (b *TBitBtn) SetMargins(value *TMargins) {
    BitBtn_SetMargins(b.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (b *TBitBtn) CustomHint() *TCustomHint {
    return CustomHintFromInst(BitBtn_GetCustomHint(b.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (b *TBitBtn) SetCustomHint(value IComponent) {
    BitBtn_SetCustomHint(b.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TBitBtn) ComponentCount() int32 {
    return BitBtn_GetComponentCount(b.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (b *TBitBtn) ComponentIndex() int32 {
    return BitBtn_GetComponentIndex(b.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (b *TBitBtn) SetComponentIndex(value int32) {
    BitBtn_SetComponentIndex(b.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TBitBtn) Owner() *TComponent {
    return ComponentFromInst(BitBtn_GetOwner(b.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (b *TBitBtn) Name() string {
    return BitBtn_GetName(b.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (b *TBitBtn) SetName(value string) {
    BitBtn_SetName(b.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TBitBtn) Tag() int {
    return BitBtn_GetTag(b.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TBitBtn) SetTag(value int) {
    BitBtn_SetTag(b.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (b *TBitBtn) DockClients(Index int32) *TControl {
    return ControlFromInst(BitBtn_GetDockClients(b.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (b *TBitBtn) Controls(Index int32) *TControl {
    return ControlFromInst(BitBtn_GetControls(b.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TBitBtn) Components(AIndex int32) *TComponent {
    return ComponentFromInst(BitBtn_GetComponents(b.instance, AIndex))
}

