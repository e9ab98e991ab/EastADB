
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

type TRadioButton struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewRadioButton
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewRadioButton(owner IComponent) *TRadioButton {
    r := new(TRadioButton)
    r.instance = RadioButton_Create(CheckPtr(owner))
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RadioButtonFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func RadioButtonFromInst(inst uintptr) *TRadioButton {
    r := new(TRadioButton)
    r.instance = inst
    r.ptr = unsafe.Pointer(inst)
    return r
}

// RadioButtonFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func RadioButtonFromObj(obj IObject) *TRadioButton {
    r := new(TRadioButton)
    r.instance = CheckPtr(obj)
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RadioButtonFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func RadioButtonFromUnsafePointer(ptr unsafe.Pointer) *TRadioButton {
    r := new(TRadioButton)
    r.instance = uintptr(ptr)
    r.ptr = ptr
    return r
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (r *TRadioButton) Free() {
    if r.instance != 0 {
        RadioButton_Free(r.instance)
        r.instance = 0
        r.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (r *TRadioButton) Instance() uintptr {
    return r.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (r *TRadioButton) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (r *TRadioButton) IsValid() bool {
    return r.instance != 0
}

// TRadioButtonClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TRadioButtonClass() TClass {
    return RadioButton_StaticClassType()
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (r *TRadioButton) CanFocus() bool {
    return RadioButton_CanFocus(r.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (r *TRadioButton) ContainsControl(Control IControl) bool {
    return RadioButton_ContainsControl(r.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (r *TRadioButton) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(RadioButton_ControlAtPos(r.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (r *TRadioButton) DisableAlign() {
    RadioButton_DisableAlign(r.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (r *TRadioButton) EnableAlign() {
    RadioButton_EnableAlign(r.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (r *TRadioButton) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(RadioButton_FindChildControl(r.instance, ControlName))
}

// FlipChildren
func (r *TRadioButton) FlipChildren(AllLevels bool) {
    RadioButton_FlipChildren(r.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (r *TRadioButton) Focused() bool {
    return RadioButton_Focused(r.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (r *TRadioButton) HandleAllocated() bool {
    return RadioButton_HandleAllocated(r.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (r *TRadioButton) InsertControl(AControl IControl) {
    RadioButton_InsertControl(r.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (r *TRadioButton) Invalidate() {
    RadioButton_Invalidate(r.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (r *TRadioButton) PaintTo(DC HDC, X int32, Y int32) {
    RadioButton_PaintTo(r.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (r *TRadioButton) RemoveControl(AControl IControl) {
    RadioButton_RemoveControl(r.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (r *TRadioButton) Realign() {
    RadioButton_Realign(r.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (r *TRadioButton) Repaint() {
    RadioButton_Repaint(r.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (r *TRadioButton) ScaleBy(M int32, D int32) {
    RadioButton_ScaleBy(r.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (r *TRadioButton) ScrollBy(DeltaX int32, DeltaY int32) {
    RadioButton_ScrollBy(r.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (r *TRadioButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RadioButton_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (r *TRadioButton) SetFocus() {
    RadioButton_SetFocus(r.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (r *TRadioButton) Update() {
    RadioButton_Update(r.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (r *TRadioButton) UpdateControlState() {
    RadioButton_UpdateControlState(r.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (r *TRadioButton) BringToFront() {
    RadioButton_BringToFront(r.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (r *TRadioButton) ClientToScreen(Point TPoint) TPoint {
    return RadioButton_ClientToScreen(r.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (r *TRadioButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return RadioButton_ClientToParent(r.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (r *TRadioButton) Dragging() bool {
    return RadioButton_Dragging(r.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (r *TRadioButton) HasParent() bool {
    return RadioButton_HasParent(r.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (r *TRadioButton) Hide() {
    RadioButton_Hide(r.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (r *TRadioButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RadioButton_Perform(r.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (r *TRadioButton) Refresh() {
    RadioButton_Refresh(r.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (r *TRadioButton) ScreenToClient(Point TPoint) TPoint {
    return RadioButton_ScreenToClient(r.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (r *TRadioButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return RadioButton_ParentToClient(r.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (r *TRadioButton) SendToBack() {
    RadioButton_SendToBack(r.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (r *TRadioButton) Show() {
    RadioButton_Show(r.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (r *TRadioButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    return RadioButton_GetTextBuf(r.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (r *TRadioButton) GetTextLen() int32 {
    return RadioButton_GetTextLen(r.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (r *TRadioButton) SetTextBuf(Buffer string) {
    RadioButton_SetTextBuf(r.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (r *TRadioButton) FindComponent(AName string) *TComponent {
    return ComponentFromInst(RadioButton_FindComponent(r.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (r *TRadioButton) GetNamePath() string {
    return RadioButton_GetNamePath(r.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (r *TRadioButton) Assign(Source IObject) {
    RadioButton_Assign(r.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (r *TRadioButton) DisposeOf() {
    RadioButton_DisposeOf(r.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (r *TRadioButton) ClassType() TClass {
    return RadioButton_ClassType(r.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (r *TRadioButton) ClassName() string {
    return RadioButton_ClassName(r.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (r *TRadioButton) InstanceSize() int32 {
    return RadioButton_InstanceSize(r.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (r *TRadioButton) InheritsFrom(AClass TClass) bool {
    return RadioButton_InheritsFrom(r.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (r *TRadioButton) Equals(Obj IObject) bool {
    return RadioButton_Equals(r.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (r *TRadioButton) GetHashCode() int32 {
    return RadioButton_GetHashCode(r.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (r *TRadioButton) ToString() string {
    return RadioButton_ToString(r.instance)
}

// Action
func (r *TRadioButton) Action() *TAction {
    return ActionFromInst(RadioButton_GetAction(r.instance))
}

// SetAction
func (r *TRadioButton) SetAction(value IComponent) {
    RadioButton_SetAction(r.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (r *TRadioButton) Align() TAlign {
    return RadioButton_GetAlign(r.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (r *TRadioButton) SetAlign(value TAlign) {
    RadioButton_SetAlign(r.instance, value)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (r *TRadioButton) Alignment() TLeftRight {
    return RadioButton_GetAlignment(r.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (r *TRadioButton) SetAlignment(value TLeftRight) {
    RadioButton_SetAlignment(r.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (r *TRadioButton) Anchors() TAnchors {
    return RadioButton_GetAnchors(r.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (r *TRadioButton) SetAnchors(value TAnchors) {
    RadioButton_SetAnchors(r.instance, value)
}

// BiDiMode
func (r *TRadioButton) BiDiMode() TBiDiMode {
    return RadioButton_GetBiDiMode(r.instance)
}

// SetBiDiMode
func (r *TRadioButton) SetBiDiMode(value TBiDiMode) {
    RadioButton_SetBiDiMode(r.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (r *TRadioButton) Caption() string {
    return RadioButton_GetCaption(r.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (r *TRadioButton) SetCaption(value string) {
    RadioButton_SetCaption(r.instance, value)
}

// Checked
// CN: 获取是否选中。
// EN: .
func (r *TRadioButton) Checked() bool {
    return RadioButton_GetChecked(r.instance)
}

// SetChecked
// CN: 设置是否选中。
// EN: .
func (r *TRadioButton) SetChecked(value bool) {
    RadioButton_SetChecked(r.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (r *TRadioButton) Color() TColor {
    return RadioButton_GetColor(r.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (r *TRadioButton) SetColor(value TColor) {
    RadioButton_SetColor(r.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (r *TRadioButton) DoubleBuffered() bool {
    return RadioButton_GetDoubleBuffered(r.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (r *TRadioButton) SetDoubleBuffered(value bool) {
    RadioButton_SetDoubleBuffered(r.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (r *TRadioButton) DragCursor() TCursor {
    return RadioButton_GetDragCursor(r.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (r *TRadioButton) SetDragCursor(value TCursor) {
    RadioButton_SetDragCursor(r.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (r *TRadioButton) DragKind() TDragKind {
    return RadioButton_GetDragKind(r.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (r *TRadioButton) SetDragKind(value TDragKind) {
    RadioButton_SetDragKind(r.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (r *TRadioButton) DragMode() TDragMode {
    return RadioButton_GetDragMode(r.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (r *TRadioButton) SetDragMode(value TDragMode) {
    RadioButton_SetDragMode(r.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (r *TRadioButton) Enabled() bool {
    return RadioButton_GetEnabled(r.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (r *TRadioButton) SetEnabled(value bool) {
    RadioButton_SetEnabled(r.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (r *TRadioButton) Font() *TFont {
    return FontFromInst(RadioButton_GetFont(r.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (r *TRadioButton) SetFont(value *TFont) {
    RadioButton_SetFont(r.instance, CheckPtr(value))
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (r *TRadioButton) ParentColor() bool {
    return RadioButton_GetParentColor(r.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (r *TRadioButton) SetParentColor(value bool) {
    RadioButton_SetParentColor(r.instance, value)
}

// ParentCtl3D
func (r *TRadioButton) ParentCtl3D() bool {
    return RadioButton_GetParentCtl3D(r.instance)
}

// SetParentCtl3D
func (r *TRadioButton) SetParentCtl3D(value bool) {
    RadioButton_SetParentCtl3D(r.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (r *TRadioButton) ParentDoubleBuffered() bool {
    return RadioButton_GetParentDoubleBuffered(r.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (r *TRadioButton) SetParentDoubleBuffered(value bool) {
    RadioButton_SetParentDoubleBuffered(r.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (r *TRadioButton) ParentFont() bool {
    return RadioButton_GetParentFont(r.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (r *TRadioButton) SetParentFont(value bool) {
    RadioButton_SetParentFont(r.instance, value)
}

// ParentShowHint
func (r *TRadioButton) ParentShowHint() bool {
    return RadioButton_GetParentShowHint(r.instance)
}

// SetParentShowHint
func (r *TRadioButton) SetParentShowHint(value bool) {
    RadioButton_SetParentShowHint(r.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (r *TRadioButton) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(RadioButton_GetPopupMenu(r.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (r *TRadioButton) SetPopupMenu(value IComponent) {
    RadioButton_SetPopupMenu(r.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (r *TRadioButton) ShowHint() bool {
    return RadioButton_GetShowHint(r.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (r *TRadioButton) SetShowHint(value bool) {
    RadioButton_SetShowHint(r.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (r *TRadioButton) TabOrder() TTabOrder {
    return RadioButton_GetTabOrder(r.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (r *TRadioButton) SetTabOrder(value TTabOrder) {
    RadioButton_SetTabOrder(r.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (r *TRadioButton) TabStop() bool {
    return RadioButton_GetTabStop(r.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (r *TRadioButton) SetTabStop(value bool) {
    RadioButton_SetTabStop(r.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (r *TRadioButton) Visible() bool {
    return RadioButton_GetVisible(r.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (r *TRadioButton) SetVisible(value bool) {
    RadioButton_SetVisible(r.instance, value)
}

// WordWrap
// CN: 获取自动换行。
// EN: Get Automatic line break.
func (r *TRadioButton) WordWrap() bool {
    return RadioButton_GetWordWrap(r.instance)
}

// SetWordWrap
// CN: 设置自动换行。
// EN: Set Automatic line break.
func (r *TRadioButton) SetWordWrap(value bool) {
    RadioButton_SetWordWrap(r.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (r *TRadioButton) StyleElements() TStyleElements {
    return RadioButton_GetStyleElements(r.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (r *TRadioButton) SetStyleElements(value TStyleElements) {
    RadioButton_SetStyleElements(r.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (r *TRadioButton) SetOnClick(fn TNotifyEvent) {
    RadioButton_SetOnClick(r.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (r *TRadioButton) SetOnContextPopup(fn TContextPopupEvent) {
    RadioButton_SetOnContextPopup(r.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (r *TRadioButton) SetOnDblClick(fn TNotifyEvent) {
    RadioButton_SetOnDblClick(r.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (r *TRadioButton) SetOnDragDrop(fn TDragDropEvent) {
    RadioButton_SetOnDragDrop(r.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (r *TRadioButton) SetOnDragOver(fn TDragOverEvent) {
    RadioButton_SetOnDragOver(r.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (r *TRadioButton) SetOnEndDock(fn TEndDragEvent) {
    RadioButton_SetOnEndDock(r.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (r *TRadioButton) SetOnEndDrag(fn TEndDragEvent) {
    RadioButton_SetOnEndDrag(r.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (r *TRadioButton) SetOnEnter(fn TNotifyEvent) {
    RadioButton_SetOnEnter(r.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (r *TRadioButton) SetOnExit(fn TNotifyEvent) {
    RadioButton_SetOnExit(r.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (r *TRadioButton) SetOnKeyDown(fn TKeyEvent) {
    RadioButton_SetOnKeyDown(r.instance, fn)
}

// SetOnKeyPress
func (r *TRadioButton) SetOnKeyPress(fn TKeyPressEvent) {
    RadioButton_SetOnKeyPress(r.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (r *TRadioButton) SetOnKeyUp(fn TKeyEvent) {
    RadioButton_SetOnKeyUp(r.instance, fn)
}

// SetOnMouseActivate
func (r *TRadioButton) SetOnMouseActivate(fn TMouseActivateEvent) {
    RadioButton_SetOnMouseActivate(r.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (r *TRadioButton) SetOnMouseDown(fn TMouseEvent) {
    RadioButton_SetOnMouseDown(r.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (r *TRadioButton) SetOnMouseEnter(fn TNotifyEvent) {
    RadioButton_SetOnMouseEnter(r.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (r *TRadioButton) SetOnMouseLeave(fn TNotifyEvent) {
    RadioButton_SetOnMouseLeave(r.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (r *TRadioButton) SetOnMouseMove(fn TMouseMoveEvent) {
    RadioButton_SetOnMouseMove(r.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (r *TRadioButton) SetOnMouseUp(fn TMouseEvent) {
    RadioButton_SetOnMouseUp(r.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (r *TRadioButton) SetOnStartDock(fn TStartDockEvent) {
    RadioButton_SetOnStartDock(r.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (r *TRadioButton) DockClientCount() int32 {
    return RadioButton_GetDockClientCount(r.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (r *TRadioButton) DockSite() bool {
    return RadioButton_GetDockSite(r.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (r *TRadioButton) SetDockSite(value bool) {
    RadioButton_SetDockSite(r.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (r *TRadioButton) AlignDisabled() bool {
    return RadioButton_GetAlignDisabled(r.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (r *TRadioButton) MouseInClient() bool {
    return RadioButton_GetMouseInClient(r.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (r *TRadioButton) VisibleDockClientCount() int32 {
    return RadioButton_GetVisibleDockClientCount(r.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (r *TRadioButton) Brush() *TBrush {
    return BrushFromInst(RadioButton_GetBrush(r.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (r *TRadioButton) ControlCount() int32 {
    return RadioButton_GetControlCount(r.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (r *TRadioButton) Handle() HWND {
    return RadioButton_GetHandle(r.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (r *TRadioButton) ParentWindow() HWND {
    return RadioButton_GetParentWindow(r.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (r *TRadioButton) SetParentWindow(value HWND) {
    RadioButton_SetParentWindow(r.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (r *TRadioButton) UseDockManager() bool {
    return RadioButton_GetUseDockManager(r.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (r *TRadioButton) SetUseDockManager(value bool) {
    RadioButton_SetUseDockManager(r.instance, value)
}

// BoundsRect
func (r *TRadioButton) BoundsRect() TRect {
    return RadioButton_GetBoundsRect(r.instance)
}

// SetBoundsRect
func (r *TRadioButton) SetBoundsRect(value TRect) {
    RadioButton_SetBoundsRect(r.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (r *TRadioButton) ClientHeight() int32 {
    return RadioButton_GetClientHeight(r.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (r *TRadioButton) SetClientHeight(value int32) {
    RadioButton_SetClientHeight(r.instance, value)
}

// ClientOrigin
func (r *TRadioButton) ClientOrigin() TPoint {
    return RadioButton_GetClientOrigin(r.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (r *TRadioButton) ClientRect() TRect {
    return RadioButton_GetClientRect(r.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (r *TRadioButton) ClientWidth() int32 {
    return RadioButton_GetClientWidth(r.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (r *TRadioButton) SetClientWidth(value int32) {
    RadioButton_SetClientWidth(r.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (r *TRadioButton) ControlState() TControlState {
    return RadioButton_GetControlState(r.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (r *TRadioButton) SetControlState(value TControlState) {
    RadioButton_SetControlState(r.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (r *TRadioButton) ControlStyle() TControlStyle {
    return RadioButton_GetControlStyle(r.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (r *TRadioButton) SetControlStyle(value TControlStyle) {
    RadioButton_SetControlStyle(r.instance, value)
}

// ExplicitLeft
func (r *TRadioButton) ExplicitLeft() int32 {
    return RadioButton_GetExplicitLeft(r.instance)
}

// ExplicitTop
func (r *TRadioButton) ExplicitTop() int32 {
    return RadioButton_GetExplicitTop(r.instance)
}

// ExplicitWidth
func (r *TRadioButton) ExplicitWidth() int32 {
    return RadioButton_GetExplicitWidth(r.instance)
}

// ExplicitHeight
func (r *TRadioButton) ExplicitHeight() int32 {
    return RadioButton_GetExplicitHeight(r.instance)
}

// Floating
func (r *TRadioButton) Floating() bool {
    return RadioButton_GetFloating(r.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (r *TRadioButton) Parent() *TWinControl {
    return WinControlFromInst(RadioButton_GetParent(r.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (r *TRadioButton) SetParent(value IWinControl) {
    RadioButton_SetParent(r.instance, CheckPtr(value))
}

// SetOnGesture
func (r *TRadioButton) SetOnGesture(fn TGestureEvent) {
    RadioButton_SetOnGesture(r.instance, fn)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (r *TRadioButton) AlignWithMargins() bool {
    return RadioButton_GetAlignWithMargins(r.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (r *TRadioButton) SetAlignWithMargins(value bool) {
    RadioButton_SetAlignWithMargins(r.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (r *TRadioButton) Left() int32 {
    return RadioButton_GetLeft(r.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (r *TRadioButton) SetLeft(value int32) {
    RadioButton_SetLeft(r.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (r *TRadioButton) Top() int32 {
    return RadioButton_GetTop(r.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (r *TRadioButton) SetTop(value int32) {
    RadioButton_SetTop(r.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (r *TRadioButton) Width() int32 {
    return RadioButton_GetWidth(r.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (r *TRadioButton) SetWidth(value int32) {
    RadioButton_SetWidth(r.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (r *TRadioButton) Height() int32 {
    return RadioButton_GetHeight(r.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (r *TRadioButton) SetHeight(value int32) {
    RadioButton_SetHeight(r.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (r *TRadioButton) Cursor() TCursor {
    return RadioButton_GetCursor(r.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (r *TRadioButton) SetCursor(value TCursor) {
    RadioButton_SetCursor(r.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (r *TRadioButton) Hint() string {
    return RadioButton_GetHint(r.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (r *TRadioButton) SetHint(value string) {
    RadioButton_SetHint(r.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (r *TRadioButton) Margins() *TMargins {
    return MarginsFromInst(RadioButton_GetMargins(r.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (r *TRadioButton) SetMargins(value *TMargins) {
    RadioButton_SetMargins(r.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (r *TRadioButton) CustomHint() *TCustomHint {
    return CustomHintFromInst(RadioButton_GetCustomHint(r.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (r *TRadioButton) SetCustomHint(value IComponent) {
    RadioButton_SetCustomHint(r.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (r *TRadioButton) ComponentCount() int32 {
    return RadioButton_GetComponentCount(r.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (r *TRadioButton) ComponentIndex() int32 {
    return RadioButton_GetComponentIndex(r.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (r *TRadioButton) SetComponentIndex(value int32) {
    RadioButton_SetComponentIndex(r.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (r *TRadioButton) Owner() *TComponent {
    return ComponentFromInst(RadioButton_GetOwner(r.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (r *TRadioButton) Name() string {
    return RadioButton_GetName(r.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (r *TRadioButton) SetName(value string) {
    RadioButton_SetName(r.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (r *TRadioButton) Tag() int {
    return RadioButton_GetTag(r.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (r *TRadioButton) SetTag(value int) {
    RadioButton_SetTag(r.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (r *TRadioButton) DockClients(Index int32) *TControl {
    return ControlFromInst(RadioButton_GetDockClients(r.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (r *TRadioButton) Controls(Index int32) *TControl {
    return ControlFromInst(RadioButton_GetControls(r.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (r *TRadioButton) Components(AIndex int32) *TComponent {
    return ComponentFromInst(RadioButton_GetComponents(r.instance, AIndex))
}

