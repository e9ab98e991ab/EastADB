
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

type TButton struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewButton
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewButton(owner IComponent) *TButton {
    b := new(TButton)
    b.instance = Button_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// ButtonFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ButtonFromInst(inst uintptr) *TButton {
    b := new(TButton)
    b.instance = inst
    b.ptr = unsafe.Pointer(inst)
    return b
}

// ButtonFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ButtonFromObj(obj IObject) *TButton {
    b := new(TButton)
    b.instance = CheckPtr(obj)
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// ButtonFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ButtonFromUnsafePointer(ptr unsafe.Pointer) *TButton {
    b := new(TButton)
    b.instance = uintptr(ptr)
    b.ptr = ptr
    return b
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (b *TButton) Free() {
    if b.instance != 0 {
        Button_Free(b.instance)
        b.instance = 0
        b.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TButton) Instance() uintptr {
    return b.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TButton) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TButton) IsValid() bool {
    return b.instance != 0
}

// TButtonClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TButtonClass() TClass {
    return Button_StaticClassType()
}

// Click
// CN: 单击。
// EN: .
func (b *TButton) Click() {
    Button_Click(b.instance)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (b *TButton) CanFocus() bool {
    return Button_CanFocus(b.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (b *TButton) ContainsControl(Control IControl) bool {
    return Button_ContainsControl(b.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (b *TButton) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(Button_ControlAtPos(b.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (b *TButton) DisableAlign() {
    Button_DisableAlign(b.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (b *TButton) EnableAlign() {
    Button_EnableAlign(b.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (b *TButton) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(Button_FindChildControl(b.instance, ControlName))
}

// FlipChildren
func (b *TButton) FlipChildren(AllLevels bool) {
    Button_FlipChildren(b.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (b *TButton) Focused() bool {
    return Button_Focused(b.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (b *TButton) HandleAllocated() bool {
    return Button_HandleAllocated(b.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (b *TButton) InsertControl(AControl IControl) {
    Button_InsertControl(b.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (b *TButton) Invalidate() {
    Button_Invalidate(b.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (b *TButton) PaintTo(DC HDC, X int32, Y int32) {
    Button_PaintTo(b.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (b *TButton) RemoveControl(AControl IControl) {
    Button_RemoveControl(b.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (b *TButton) Realign() {
    Button_Realign(b.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (b *TButton) Repaint() {
    Button_Repaint(b.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (b *TButton) ScaleBy(M int32, D int32) {
    Button_ScaleBy(b.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (b *TButton) ScrollBy(DeltaX int32, DeltaY int32) {
    Button_ScrollBy(b.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (b *TButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Button_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (b *TButton) SetFocus() {
    Button_SetFocus(b.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (b *TButton) Update() {
    Button_Update(b.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (b *TButton) UpdateControlState() {
    Button_UpdateControlState(b.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (b *TButton) BringToFront() {
    Button_BringToFront(b.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (b *TButton) ClientToScreen(Point TPoint) TPoint {
    return Button_ClientToScreen(b.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (b *TButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Button_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (b *TButton) Dragging() bool {
    return Button_Dragging(b.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (b *TButton) HasParent() bool {
    return Button_HasParent(b.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (b *TButton) Hide() {
    Button_Hide(b.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (b *TButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Button_Perform(b.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (b *TButton) Refresh() {
    Button_Refresh(b.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (b *TButton) ScreenToClient(Point TPoint) TPoint {
    return Button_ScreenToClient(b.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (b *TButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Button_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (b *TButton) SendToBack() {
    Button_SendToBack(b.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (b *TButton) Show() {
    Button_Show(b.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (b *TButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Button_GetTextBuf(b.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (b *TButton) GetTextLen() int32 {
    return Button_GetTextLen(b.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (b *TButton) SetTextBuf(Buffer string) {
    Button_SetTextBuf(b.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (b *TButton) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Button_FindComponent(b.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (b *TButton) GetNamePath() string {
    return Button_GetNamePath(b.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (b *TButton) Assign(Source IObject) {
    Button_Assign(b.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (b *TButton) DisposeOf() {
    Button_DisposeOf(b.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TButton) ClassType() TClass {
    return Button_ClassType(b.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TButton) ClassName() string {
    return Button_ClassName(b.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TButton) InstanceSize() int32 {
    return Button_InstanceSize(b.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TButton) InheritsFrom(AClass TClass) bool {
    return Button_InheritsFrom(b.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TButton) Equals(Obj IObject) bool {
    return Button_Equals(b.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TButton) GetHashCode() int32 {
    return Button_GetHashCode(b.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (b *TButton) ToString() string {
    return Button_ToString(b.instance)
}

// Action
func (b *TButton) Action() *TAction {
    return ActionFromInst(Button_GetAction(b.instance))
}

// SetAction
func (b *TButton) SetAction(value IComponent) {
    Button_SetAction(b.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (b *TButton) Align() TAlign {
    return Button_GetAlign(b.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (b *TButton) SetAlign(value TAlign) {
    Button_SetAlign(b.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (b *TButton) Anchors() TAnchors {
    return Button_GetAnchors(b.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (b *TButton) SetAnchors(value TAnchors) {
    Button_SetAnchors(b.instance, value)
}

// BiDiMode
func (b *TButton) BiDiMode() TBiDiMode {
    return Button_GetBiDiMode(b.instance)
}

// SetBiDiMode
func (b *TButton) SetBiDiMode(value TBiDiMode) {
    Button_SetBiDiMode(b.instance, value)
}

// Cancel
func (b *TButton) Cancel() bool {
    return Button_GetCancel(b.instance)
}

// SetCancel
func (b *TButton) SetCancel(value bool) {
    Button_SetCancel(b.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (b *TButton) Caption() string {
    return Button_GetCaption(b.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (b *TButton) SetCaption(value string) {
    Button_SetCaption(b.instance, value)
}

// CommandLinkHint
func (b *TButton) CommandLinkHint() string {
    return Button_GetCommandLinkHint(b.instance)
}

// SetCommandLinkHint
func (b *TButton) SetCommandLinkHint(value string) {
    Button_SetCommandLinkHint(b.instance, value)
}

// Default
func (b *TButton) Default() bool {
    return Button_GetDefault(b.instance)
}

// SetDefault
func (b *TButton) SetDefault(value bool) {
    Button_SetDefault(b.instance, value)
}

// DisabledImageIndex
func (b *TButton) DisabledImageIndex() int32 {
    return Button_GetDisabledImageIndex(b.instance)
}

// SetDisabledImageIndex
func (b *TButton) SetDisabledImageIndex(value int32) {
    Button_SetDisabledImageIndex(b.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (b *TButton) DoubleBuffered() bool {
    return Button_GetDoubleBuffered(b.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (b *TButton) SetDoubleBuffered(value bool) {
    Button_SetDoubleBuffered(b.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (b *TButton) DragCursor() TCursor {
    return Button_GetDragCursor(b.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (b *TButton) SetDragCursor(value TCursor) {
    Button_SetDragCursor(b.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (b *TButton) DragKind() TDragKind {
    return Button_GetDragKind(b.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (b *TButton) SetDragKind(value TDragKind) {
    Button_SetDragKind(b.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (b *TButton) DragMode() TDragMode {
    return Button_GetDragMode(b.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (b *TButton) SetDragMode(value TDragMode) {
    Button_SetDragMode(b.instance, value)
}

// ElevationRequired
func (b *TButton) ElevationRequired() bool {
    return Button_GetElevationRequired(b.instance)
}

// SetElevationRequired
func (b *TButton) SetElevationRequired(value bool) {
    Button_SetElevationRequired(b.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (b *TButton) Enabled() bool {
    return Button_GetEnabled(b.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (b *TButton) SetEnabled(value bool) {
    Button_SetEnabled(b.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (b *TButton) Font() *TFont {
    return FontFromInst(Button_GetFont(b.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (b *TButton) SetFont(value *TFont) {
    Button_SetFont(b.instance, CheckPtr(value))
}

// HotImageIndex
func (b *TButton) HotImageIndex() int32 {
    return Button_GetHotImageIndex(b.instance)
}

// SetHotImageIndex
func (b *TButton) SetHotImageIndex(value int32) {
    Button_SetHotImageIndex(b.instance, value)
}

// ImageAlignment
func (b *TButton) ImageAlignment() TImageAlignment {
    return Button_GetImageAlignment(b.instance)
}

// SetImageAlignment
func (b *TButton) SetImageAlignment(value TImageAlignment) {
    Button_SetImageAlignment(b.instance, value)
}

// ImageIndex
// CN: 获取图像在images中的索引。
// EN: .
func (b *TButton) ImageIndex() int32 {
    return Button_GetImageIndex(b.instance)
}

// SetImageIndex
// CN: 设置图像在images中的索引。
// EN: .
func (b *TButton) SetImageIndex(value int32) {
    Button_SetImageIndex(b.instance, value)
}

// Images
// CN: 获取图标索引列表对象。
// EN: .
func (b *TButton) Images() *TImageList {
    return ImageListFromInst(Button_GetImages(b.instance))
}

// SetImages
// CN: 设置图标索引列表对象。
// EN: .
func (b *TButton) SetImages(value IComponent) {
    Button_SetImages(b.instance, CheckPtr(value))
}

// ModalResult
// CN: 获取模态对话框显示结果。
// EN: .
func (b *TButton) ModalResult() TModalResult {
    return Button_GetModalResult(b.instance)
}

// SetModalResult
// CN: 设置模态对话框显示结果。
// EN: .
func (b *TButton) SetModalResult(value TModalResult) {
    Button_SetModalResult(b.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (b *TButton) ParentDoubleBuffered() bool {
    return Button_GetParentDoubleBuffered(b.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (b *TButton) SetParentDoubleBuffered(value bool) {
    Button_SetParentDoubleBuffered(b.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (b *TButton) ParentFont() bool {
    return Button_GetParentFont(b.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (b *TButton) SetParentFont(value bool) {
    Button_SetParentFont(b.instance, value)
}

// ParentShowHint
func (b *TButton) ParentShowHint() bool {
    return Button_GetParentShowHint(b.instance)
}

// SetParentShowHint
func (b *TButton) SetParentShowHint(value bool) {
    Button_SetParentShowHint(b.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (b *TButton) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Button_GetPopupMenu(b.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (b *TButton) SetPopupMenu(value IComponent) {
    Button_SetPopupMenu(b.instance, CheckPtr(value))
}

// PressedImageIndex
func (b *TButton) PressedImageIndex() int32 {
    return Button_GetPressedImageIndex(b.instance)
}

// SetPressedImageIndex
func (b *TButton) SetPressedImageIndex(value int32) {
    Button_SetPressedImageIndex(b.instance, value)
}

// SelectedImageIndex
func (b *TButton) SelectedImageIndex() int32 {
    return Button_GetSelectedImageIndex(b.instance)
}

// SetSelectedImageIndex
func (b *TButton) SetSelectedImageIndex(value int32) {
    Button_SetSelectedImageIndex(b.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (b *TButton) ShowHint() bool {
    return Button_GetShowHint(b.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (b *TButton) SetShowHint(value bool) {
    Button_SetShowHint(b.instance, value)
}

// Style
func (b *TButton) Style() TButtonStyle {
    return Button_GetStyle(b.instance)
}

// SetStyle
func (b *TButton) SetStyle(value TButtonStyle) {
    Button_SetStyle(b.instance, value)
}

// StylusHotImageIndex
func (b *TButton) StylusHotImageIndex() int32 {
    return Button_GetStylusHotImageIndex(b.instance)
}

// SetStylusHotImageIndex
func (b *TButton) SetStylusHotImageIndex(value int32) {
    Button_SetStylusHotImageIndex(b.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (b *TButton) TabOrder() TTabOrder {
    return Button_GetTabOrder(b.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (b *TButton) SetTabOrder(value TTabOrder) {
    Button_SetTabOrder(b.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (b *TButton) TabStop() bool {
    return Button_GetTabStop(b.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (b *TButton) SetTabStop(value bool) {
    Button_SetTabStop(b.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (b *TButton) Visible() bool {
    return Button_GetVisible(b.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (b *TButton) SetVisible(value bool) {
    Button_SetVisible(b.instance, value)
}

// WordWrap
// CN: 获取自动换行。
// EN: Get Automatic line break.
func (b *TButton) WordWrap() bool {
    return Button_GetWordWrap(b.instance)
}

// SetWordWrap
// CN: 设置自动换行。
// EN: Set Automatic line break.
func (b *TButton) SetWordWrap(value bool) {
    Button_SetWordWrap(b.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (b *TButton) StyleElements() TStyleElements {
    return Button_GetStyleElements(b.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (b *TButton) SetStyleElements(value TStyleElements) {
    Button_SetStyleElements(b.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (b *TButton) SetOnClick(fn TNotifyEvent) {
    Button_SetOnClick(b.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (b *TButton) SetOnContextPopup(fn TContextPopupEvent) {
    Button_SetOnContextPopup(b.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (b *TButton) SetOnDragDrop(fn TDragDropEvent) {
    Button_SetOnDragDrop(b.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (b *TButton) SetOnDragOver(fn TDragOverEvent) {
    Button_SetOnDragOver(b.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (b *TButton) SetOnEndDock(fn TEndDragEvent) {
    Button_SetOnEndDock(b.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (b *TButton) SetOnEndDrag(fn TEndDragEvent) {
    Button_SetOnEndDrag(b.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (b *TButton) SetOnEnter(fn TNotifyEvent) {
    Button_SetOnEnter(b.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (b *TButton) SetOnExit(fn TNotifyEvent) {
    Button_SetOnExit(b.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (b *TButton) SetOnKeyDown(fn TKeyEvent) {
    Button_SetOnKeyDown(b.instance, fn)
}

// SetOnKeyPress
func (b *TButton) SetOnKeyPress(fn TKeyPressEvent) {
    Button_SetOnKeyPress(b.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (b *TButton) SetOnKeyUp(fn TKeyEvent) {
    Button_SetOnKeyUp(b.instance, fn)
}

// SetOnMouseActivate
func (b *TButton) SetOnMouseActivate(fn TMouseActivateEvent) {
    Button_SetOnMouseActivate(b.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (b *TButton) SetOnMouseDown(fn TMouseEvent) {
    Button_SetOnMouseDown(b.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (b *TButton) SetOnMouseEnter(fn TNotifyEvent) {
    Button_SetOnMouseEnter(b.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (b *TButton) SetOnMouseLeave(fn TNotifyEvent) {
    Button_SetOnMouseLeave(b.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (b *TButton) SetOnMouseMove(fn TMouseMoveEvent) {
    Button_SetOnMouseMove(b.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (b *TButton) SetOnMouseUp(fn TMouseEvent) {
    Button_SetOnMouseUp(b.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (b *TButton) SetOnStartDock(fn TStartDockEvent) {
    Button_SetOnStartDock(b.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (b *TButton) DockClientCount() int32 {
    return Button_GetDockClientCount(b.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (b *TButton) DockSite() bool {
    return Button_GetDockSite(b.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (b *TButton) SetDockSite(value bool) {
    Button_SetDockSite(b.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (b *TButton) AlignDisabled() bool {
    return Button_GetAlignDisabled(b.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (b *TButton) MouseInClient() bool {
    return Button_GetMouseInClient(b.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (b *TButton) VisibleDockClientCount() int32 {
    return Button_GetVisibleDockClientCount(b.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (b *TButton) Brush() *TBrush {
    return BrushFromInst(Button_GetBrush(b.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (b *TButton) ControlCount() int32 {
    return Button_GetControlCount(b.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (b *TButton) Handle() HWND {
    return Button_GetHandle(b.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (b *TButton) ParentWindow() HWND {
    return Button_GetParentWindow(b.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (b *TButton) SetParentWindow(value HWND) {
    Button_SetParentWindow(b.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (b *TButton) UseDockManager() bool {
    return Button_GetUseDockManager(b.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (b *TButton) SetUseDockManager(value bool) {
    Button_SetUseDockManager(b.instance, value)
}

// BoundsRect
func (b *TButton) BoundsRect() TRect {
    return Button_GetBoundsRect(b.instance)
}

// SetBoundsRect
func (b *TButton) SetBoundsRect(value TRect) {
    Button_SetBoundsRect(b.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (b *TButton) ClientHeight() int32 {
    return Button_GetClientHeight(b.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (b *TButton) SetClientHeight(value int32) {
    Button_SetClientHeight(b.instance, value)
}

// ClientOrigin
func (b *TButton) ClientOrigin() TPoint {
    return Button_GetClientOrigin(b.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (b *TButton) ClientRect() TRect {
    return Button_GetClientRect(b.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (b *TButton) ClientWidth() int32 {
    return Button_GetClientWidth(b.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (b *TButton) SetClientWidth(value int32) {
    Button_SetClientWidth(b.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (b *TButton) ControlState() TControlState {
    return Button_GetControlState(b.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (b *TButton) SetControlState(value TControlState) {
    Button_SetControlState(b.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (b *TButton) ControlStyle() TControlStyle {
    return Button_GetControlStyle(b.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (b *TButton) SetControlStyle(value TControlStyle) {
    Button_SetControlStyle(b.instance, value)
}

// ExplicitLeft
func (b *TButton) ExplicitLeft() int32 {
    return Button_GetExplicitLeft(b.instance)
}

// ExplicitTop
func (b *TButton) ExplicitTop() int32 {
    return Button_GetExplicitTop(b.instance)
}

// ExplicitWidth
func (b *TButton) ExplicitWidth() int32 {
    return Button_GetExplicitWidth(b.instance)
}

// ExplicitHeight
func (b *TButton) ExplicitHeight() int32 {
    return Button_GetExplicitHeight(b.instance)
}

// Floating
func (b *TButton) Floating() bool {
    return Button_GetFloating(b.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (b *TButton) Parent() *TWinControl {
    return WinControlFromInst(Button_GetParent(b.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (b *TButton) SetParent(value IWinControl) {
    Button_SetParent(b.instance, CheckPtr(value))
}

// SetOnGesture
func (b *TButton) SetOnGesture(fn TGestureEvent) {
    Button_SetOnGesture(b.instance, fn)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (b *TButton) AlignWithMargins() bool {
    return Button_GetAlignWithMargins(b.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (b *TButton) SetAlignWithMargins(value bool) {
    Button_SetAlignWithMargins(b.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (b *TButton) Left() int32 {
    return Button_GetLeft(b.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (b *TButton) SetLeft(value int32) {
    Button_SetLeft(b.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (b *TButton) Top() int32 {
    return Button_GetTop(b.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (b *TButton) SetTop(value int32) {
    Button_SetTop(b.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (b *TButton) Width() int32 {
    return Button_GetWidth(b.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (b *TButton) SetWidth(value int32) {
    Button_SetWidth(b.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (b *TButton) Height() int32 {
    return Button_GetHeight(b.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (b *TButton) SetHeight(value int32) {
    Button_SetHeight(b.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (b *TButton) Cursor() TCursor {
    return Button_GetCursor(b.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (b *TButton) SetCursor(value TCursor) {
    Button_SetCursor(b.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (b *TButton) Hint() string {
    return Button_GetHint(b.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (b *TButton) SetHint(value string) {
    Button_SetHint(b.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (b *TButton) Margins() *TMargins {
    return MarginsFromInst(Button_GetMargins(b.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (b *TButton) SetMargins(value *TMargins) {
    Button_SetMargins(b.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (b *TButton) CustomHint() *TCustomHint {
    return CustomHintFromInst(Button_GetCustomHint(b.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (b *TButton) SetCustomHint(value IComponent) {
    Button_SetCustomHint(b.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TButton) ComponentCount() int32 {
    return Button_GetComponentCount(b.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (b *TButton) ComponentIndex() int32 {
    return Button_GetComponentIndex(b.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (b *TButton) SetComponentIndex(value int32) {
    Button_SetComponentIndex(b.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TButton) Owner() *TComponent {
    return ComponentFromInst(Button_GetOwner(b.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (b *TButton) Name() string {
    return Button_GetName(b.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (b *TButton) SetName(value string) {
    Button_SetName(b.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TButton) Tag() int {
    return Button_GetTag(b.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TButton) SetTag(value int) {
    Button_SetTag(b.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (b *TButton) DockClients(Index int32) *TControl {
    return ControlFromInst(Button_GetDockClients(b.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (b *TButton) Controls(Index int32) *TControl {
    return ControlFromInst(Button_GetControls(b.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TButton) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Button_GetComponents(b.instance, AIndex))
}

