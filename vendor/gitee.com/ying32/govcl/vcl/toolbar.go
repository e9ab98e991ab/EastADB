
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

type TToolBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewToolBar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewToolBar(owner IComponent) *TToolBar {
    t := new(TToolBar)
    t.instance = ToolBar_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// ToolBarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ToolBarFromInst(inst uintptr) *TToolBar {
    t := new(TToolBar)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// ToolBarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ToolBarFromObj(obj IObject) *TToolBar {
    t := new(TToolBar)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// ToolBarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ToolBarFromUnsafePointer(ptr unsafe.Pointer) *TToolBar {
    t := new(TToolBar)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TToolBar) Free() {
    if t.instance != 0 {
        ToolBar_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TToolBar) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TToolBar) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TToolBar) IsValid() bool {
    return t.instance != 0
}

// TToolBarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TToolBarClass() TClass {
    return ToolBar_StaticClassType()
}

// FlipChildren
func (t *TToolBar) FlipChildren(AllLevels bool) {
    ToolBar_FlipChildren(t.instance, AllLevels)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (t *TToolBar) CanFocus() bool {
    return ToolBar_CanFocus(t.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (t *TToolBar) ContainsControl(Control IControl) bool {
    return ToolBar_ContainsControl(t.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (t *TToolBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(ToolBar_ControlAtPos(t.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (t *TToolBar) DisableAlign() {
    ToolBar_DisableAlign(t.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (t *TToolBar) EnableAlign() {
    ToolBar_EnableAlign(t.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (t *TToolBar) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(ToolBar_FindChildControl(t.instance, ControlName))
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (t *TToolBar) Focused() bool {
    return ToolBar_Focused(t.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (t *TToolBar) HandleAllocated() bool {
    return ToolBar_HandleAllocated(t.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (t *TToolBar) InsertControl(AControl IControl) {
    ToolBar_InsertControl(t.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (t *TToolBar) Invalidate() {
    ToolBar_Invalidate(t.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (t *TToolBar) PaintTo(DC HDC, X int32, Y int32) {
    ToolBar_PaintTo(t.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (t *TToolBar) RemoveControl(AControl IControl) {
    ToolBar_RemoveControl(t.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (t *TToolBar) Realign() {
    ToolBar_Realign(t.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (t *TToolBar) Repaint() {
    ToolBar_Repaint(t.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (t *TToolBar) ScaleBy(M int32, D int32) {
    ToolBar_ScaleBy(t.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (t *TToolBar) ScrollBy(DeltaX int32, DeltaY int32) {
    ToolBar_ScrollBy(t.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (t *TToolBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ToolBar_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (t *TToolBar) SetFocus() {
    ToolBar_SetFocus(t.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (t *TToolBar) Update() {
    ToolBar_Update(t.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (t *TToolBar) UpdateControlState() {
    ToolBar_UpdateControlState(t.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (t *TToolBar) BringToFront() {
    ToolBar_BringToFront(t.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (t *TToolBar) ClientToScreen(Point TPoint) TPoint {
    return ToolBar_ClientToScreen(t.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (t *TToolBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ToolBar_ClientToParent(t.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (t *TToolBar) Dragging() bool {
    return ToolBar_Dragging(t.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (t *TToolBar) HasParent() bool {
    return ToolBar_HasParent(t.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (t *TToolBar) Hide() {
    ToolBar_Hide(t.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (t *TToolBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ToolBar_Perform(t.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (t *TToolBar) Refresh() {
    ToolBar_Refresh(t.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (t *TToolBar) ScreenToClient(Point TPoint) TPoint {
    return ToolBar_ScreenToClient(t.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (t *TToolBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ToolBar_ParentToClient(t.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (t *TToolBar) SendToBack() {
    ToolBar_SendToBack(t.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (t *TToolBar) Show() {
    ToolBar_Show(t.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (t *TToolBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ToolBar_GetTextBuf(t.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (t *TToolBar) GetTextLen() int32 {
    return ToolBar_GetTextLen(t.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (t *TToolBar) SetTextBuf(Buffer string) {
    ToolBar_SetTextBuf(t.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (t *TToolBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ToolBar_FindComponent(t.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TToolBar) GetNamePath() string {
    return ToolBar_GetNamePath(t.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TToolBar) Assign(Source IObject) {
    ToolBar_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TToolBar) DisposeOf() {
    ToolBar_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TToolBar) ClassType() TClass {
    return ToolBar_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TToolBar) ClassName() string {
    return ToolBar_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TToolBar) InstanceSize() int32 {
    return ToolBar_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TToolBar) InheritsFrom(AClass TClass) bool {
    return ToolBar_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TToolBar) Equals(Obj IObject) bool {
    return ToolBar_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TToolBar) GetHashCode() int32 {
    return ToolBar_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TToolBar) ToString() string {
    return ToolBar_ToString(t.instance)
}

// ButtonCount
func (t *TToolBar) ButtonCount() int32 {
    return ToolBar_GetButtonCount(t.instance)
}

// Canvas
// CN: 获取画布。
// EN: .
func (t *TToolBar) Canvas() *TCanvas {
    return CanvasFromInst(ToolBar_GetCanvas(t.instance))
}

// CustomizeKeyName
func (t *TToolBar) CustomizeKeyName() string {
    return ToolBar_GetCustomizeKeyName(t.instance)
}

// SetCustomizeKeyName
func (t *TToolBar) SetCustomizeKeyName(value string) {
    ToolBar_SetCustomizeKeyName(t.instance, value)
}

// CustomizeValueName
func (t *TToolBar) CustomizeValueName() string {
    return ToolBar_GetCustomizeValueName(t.instance)
}

// SetCustomizeValueName
func (t *TToolBar) SetCustomizeValueName(value string) {
    ToolBar_SetCustomizeValueName(t.instance, value)
}

// RowCount
func (t *TToolBar) RowCount() int32 {
    return ToolBar_GetRowCount(t.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (t *TToolBar) Align() TAlign {
    return ToolBar_GetAlign(t.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (t *TToolBar) SetAlign(value TAlign) {
    ToolBar_SetAlign(t.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (t *TToolBar) Anchors() TAnchors {
    return ToolBar_GetAnchors(t.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (t *TToolBar) SetAnchors(value TAnchors) {
    ToolBar_SetAnchors(t.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (t *TToolBar) AutoSize() bool {
    return ToolBar_GetAutoSize(t.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (t *TToolBar) SetAutoSize(value bool) {
    ToolBar_SetAutoSize(t.instance, value)
}

// BorderWidth
// CN: 获取边框的宽度。
// EN: .
func (t *TToolBar) BorderWidth() int32 {
    return ToolBar_GetBorderWidth(t.instance)
}

// SetBorderWidth
// CN: 设置边框的宽度。
// EN: .
func (t *TToolBar) SetBorderWidth(value int32) {
    ToolBar_SetBorderWidth(t.instance, value)
}

// ButtonHeight
func (t *TToolBar) ButtonHeight() int32 {
    return ToolBar_GetButtonHeight(t.instance)
}

// SetButtonHeight
func (t *TToolBar) SetButtonHeight(value int32) {
    ToolBar_SetButtonHeight(t.instance, value)
}

// ButtonWidth
func (t *TToolBar) ButtonWidth() int32 {
    return ToolBar_GetButtonWidth(t.instance)
}

// SetButtonWidth
func (t *TToolBar) SetButtonWidth(value int32) {
    ToolBar_SetButtonWidth(t.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (t *TToolBar) Caption() string {
    return ToolBar_GetCaption(t.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (t *TToolBar) SetCaption(value string) {
    ToolBar_SetCaption(t.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (t *TToolBar) Color() TColor {
    return ToolBar_GetColor(t.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (t *TToolBar) SetColor(value TColor) {
    ToolBar_SetColor(t.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (t *TToolBar) DoubleBuffered() bool {
    return ToolBar_GetDoubleBuffered(t.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (t *TToolBar) SetDoubleBuffered(value bool) {
    ToolBar_SetDoubleBuffered(t.instance, value)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (t *TToolBar) DockSite() bool {
    return ToolBar_GetDockSite(t.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (t *TToolBar) SetDockSite(value bool) {
    ToolBar_SetDockSite(t.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (t *TToolBar) DragCursor() TCursor {
    return ToolBar_GetDragCursor(t.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (t *TToolBar) SetDragCursor(value TCursor) {
    ToolBar_SetDragCursor(t.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (t *TToolBar) DragKind() TDragKind {
    return ToolBar_GetDragKind(t.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (t *TToolBar) SetDragKind(value TDragKind) {
    ToolBar_SetDragKind(t.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (t *TToolBar) DragMode() TDragMode {
    return ToolBar_GetDragMode(t.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (t *TToolBar) SetDragMode(value TDragMode) {
    ToolBar_SetDragMode(t.instance, value)
}

// DrawingStyle
func (t *TToolBar) DrawingStyle() TTBDrawingStyle {
    return ToolBar_GetDrawingStyle(t.instance)
}

// SetDrawingStyle
func (t *TToolBar) SetDrawingStyle(value TTBDrawingStyle) {
    ToolBar_SetDrawingStyle(t.instance, value)
}

// EdgeBorders
func (t *TToolBar) EdgeBorders() TEdgeBorders {
    return ToolBar_GetEdgeBorders(t.instance)
}

// SetEdgeBorders
func (t *TToolBar) SetEdgeBorders(value TEdgeBorders) {
    ToolBar_SetEdgeBorders(t.instance, value)
}

// EdgeInner
func (t *TToolBar) EdgeInner() TEdgeStyle {
    return ToolBar_GetEdgeInner(t.instance)
}

// SetEdgeInner
func (t *TToolBar) SetEdgeInner(value TEdgeStyle) {
    ToolBar_SetEdgeInner(t.instance, value)
}

// EdgeOuter
func (t *TToolBar) EdgeOuter() TEdgeStyle {
    return ToolBar_GetEdgeOuter(t.instance)
}

// SetEdgeOuter
func (t *TToolBar) SetEdgeOuter(value TEdgeStyle) {
    ToolBar_SetEdgeOuter(t.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TToolBar) Enabled() bool {
    return ToolBar_GetEnabled(t.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TToolBar) SetEnabled(value bool) {
    ToolBar_SetEnabled(t.instance, value)
}

// Flat
// CN: 获取平面样式。
// EN: .
func (t *TToolBar) Flat() bool {
    return ToolBar_GetFlat(t.instance)
}

// SetFlat
// CN: 设置平面样式。
// EN: .
func (t *TToolBar) SetFlat(value bool) {
    ToolBar_SetFlat(t.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (t *TToolBar) Font() *TFont {
    return FontFromInst(ToolBar_GetFont(t.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (t *TToolBar) SetFont(value *TFont) {
    ToolBar_SetFont(t.instance, CheckPtr(value))
}

// GradientEndColor
// CN: 获取渐变结束颜色, 仅VCL。
// EN: .
func (t *TToolBar) GradientEndColor() TColor {
    return ToolBar_GetGradientEndColor(t.instance)
}

// SetGradientEndColor
// CN: 设置渐变结束颜色, 仅VCL。
// EN: .
func (t *TToolBar) SetGradientEndColor(value TColor) {
    ToolBar_SetGradientEndColor(t.instance, value)
}

// GradientStartColor
// CN: 获取渐变起始颜色，仅VCL。
// EN: .
func (t *TToolBar) GradientStartColor() TColor {
    return ToolBar_GetGradientStartColor(t.instance)
}

// SetGradientStartColor
// CN: 设置渐变起始颜色，仅VCL。
// EN: .
func (t *TToolBar) SetGradientStartColor(value TColor) {
    ToolBar_SetGradientStartColor(t.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (t *TToolBar) Height() int32 {
    return ToolBar_GetHeight(t.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (t *TToolBar) SetHeight(value int32) {
    ToolBar_SetHeight(t.instance, value)
}

// HideClippedButtons
func (t *TToolBar) HideClippedButtons() bool {
    return ToolBar_GetHideClippedButtons(t.instance)
}

// SetHideClippedButtons
func (t *TToolBar) SetHideClippedButtons(value bool) {
    ToolBar_SetHideClippedButtons(t.instance, value)
}

// HotImages
func (t *TToolBar) HotImages() *TImageList {
    return ImageListFromInst(ToolBar_GetHotImages(t.instance))
}

// SetHotImages
func (t *TToolBar) SetHotImages(value IComponent) {
    ToolBar_SetHotImages(t.instance, CheckPtr(value))
}

// Images
// CN: 获取图标索引列表对象。
// EN: .
func (t *TToolBar) Images() *TImageList {
    return ImageListFromInst(ToolBar_GetImages(t.instance))
}

// SetImages
// CN: 设置图标索引列表对象。
// EN: .
func (t *TToolBar) SetImages(value IComponent) {
    ToolBar_SetImages(t.instance, CheckPtr(value))
}

// Indent
func (t *TToolBar) Indent() int32 {
    return ToolBar_GetIndent(t.instance)
}

// SetIndent
func (t *TToolBar) SetIndent(value int32) {
    ToolBar_SetIndent(t.instance, value)
}

// List
func (t *TToolBar) List() bool {
    return ToolBar_GetList(t.instance)
}

// SetList
func (t *TToolBar) SetList(value bool) {
    ToolBar_SetList(t.instance, value)
}

// Menu
func (t *TToolBar) Menu() *TMainMenu {
    return MainMenuFromInst(ToolBar_GetMenu(t.instance))
}

// SetMenu
func (t *TToolBar) SetMenu(value IComponent) {
    ToolBar_SetMenu(t.instance, CheckPtr(value))
}

// GradientDirection
// CN: 获取渐变颜色方向。
// EN: .
func (t *TToolBar) GradientDirection() TGradientDirection {
    return ToolBar_GetGradientDirection(t.instance)
}

// SetGradientDirection
// CN: 设置渐变颜色方向。
// EN: .
func (t *TToolBar) SetGradientDirection(value TGradientDirection) {
    ToolBar_SetGradientDirection(t.instance, value)
}

// GradientDrawingOptions
// CN: 获取渐变绘制选项。
// EN: .
func (t *TToolBar) GradientDrawingOptions() TTBGradientDrawingOptions {
    return ToolBar_GetGradientDrawingOptions(t.instance)
}

// SetGradientDrawingOptions
// CN: 设置渐变绘制选项。
// EN: .
func (t *TToolBar) SetGradientDrawingOptions(value TTBGradientDrawingOptions) {
    ToolBar_SetGradientDrawingOptions(t.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (t *TToolBar) ParentColor() bool {
    return ToolBar_GetParentColor(t.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (t *TToolBar) SetParentColor(value bool) {
    ToolBar_SetParentColor(t.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (t *TToolBar) ParentDoubleBuffered() bool {
    return ToolBar_GetParentDoubleBuffered(t.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (t *TToolBar) SetParentDoubleBuffered(value bool) {
    ToolBar_SetParentDoubleBuffered(t.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (t *TToolBar) ParentFont() bool {
    return ToolBar_GetParentFont(t.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (t *TToolBar) SetParentFont(value bool) {
    ToolBar_SetParentFont(t.instance, value)
}

// ParentShowHint
func (t *TToolBar) ParentShowHint() bool {
    return ToolBar_GetParentShowHint(t.instance)
}

// SetParentShowHint
func (t *TToolBar) SetParentShowHint(value bool) {
    ToolBar_SetParentShowHint(t.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (t *TToolBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ToolBar_GetPopupMenu(t.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (t *TToolBar) SetPopupMenu(value IComponent) {
    ToolBar_SetPopupMenu(t.instance, CheckPtr(value))
}

// ShowCaptions
// CN: 获取显示标题。
// EN: .
func (t *TToolBar) ShowCaptions() bool {
    return ToolBar_GetShowCaptions(t.instance)
}

// SetShowCaptions
// CN: 设置显示标题。
// EN: .
func (t *TToolBar) SetShowCaptions(value bool) {
    ToolBar_SetShowCaptions(t.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (t *TToolBar) ShowHint() bool {
    return ToolBar_GetShowHint(t.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (t *TToolBar) SetShowHint(value bool) {
    ToolBar_SetShowHint(t.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (t *TToolBar) TabOrder() TTabOrder {
    return ToolBar_GetTabOrder(t.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (t *TToolBar) SetTabOrder(value TTabOrder) {
    ToolBar_SetTabOrder(t.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (t *TToolBar) TabStop() bool {
    return ToolBar_GetTabStop(t.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (t *TToolBar) SetTabStop(value bool) {
    ToolBar_SetTabStop(t.instance, value)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (t *TToolBar) Transparent() bool {
    return ToolBar_GetTransparent(t.instance)
}

// SetTransparent
// CN: 设置透明。
// EN: Set transparent.
func (t *TToolBar) SetTransparent(value bool) {
    ToolBar_SetTransparent(t.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (t *TToolBar) Visible() bool {
    return ToolBar_GetVisible(t.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (t *TToolBar) SetVisible(value bool) {
    ToolBar_SetVisible(t.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (t *TToolBar) StyleElements() TStyleElements {
    return ToolBar_GetStyleElements(t.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (t *TToolBar) SetStyleElements(value TStyleElements) {
    ToolBar_SetStyleElements(t.instance, value)
}

// Wrapable
func (t *TToolBar) Wrapable() bool {
    return ToolBar_GetWrapable(t.instance)
}

// SetWrapable
func (t *TToolBar) SetWrapable(value bool) {
    ToolBar_SetWrapable(t.instance, value)
}

// SetOnAdvancedCustomDraw
func (t *TToolBar) SetOnAdvancedCustomDraw(fn TTBAdvancedCustomDrawEvent) {
    ToolBar_SetOnAdvancedCustomDraw(t.instance, fn)
}

// SetOnAdvancedCustomDrawButton
func (t *TToolBar) SetOnAdvancedCustomDrawButton(fn TTBAdvancedCustomDrawBtnEvent) {
    ToolBar_SetOnAdvancedCustomDrawButton(t.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (t *TToolBar) SetOnClick(fn TNotifyEvent) {
    ToolBar_SetOnClick(t.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (t *TToolBar) SetOnContextPopup(fn TContextPopupEvent) {
    ToolBar_SetOnContextPopup(t.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (t *TToolBar) SetOnDblClick(fn TNotifyEvent) {
    ToolBar_SetOnDblClick(t.instance, fn)
}

// SetOnDockDrop
func (t *TToolBar) SetOnDockDrop(fn TDockDropEvent) {
    ToolBar_SetOnDockDrop(t.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (t *TToolBar) SetOnDragDrop(fn TDragDropEvent) {
    ToolBar_SetOnDragDrop(t.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (t *TToolBar) SetOnDragOver(fn TDragOverEvent) {
    ToolBar_SetOnDragOver(t.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (t *TToolBar) SetOnEndDock(fn TEndDragEvent) {
    ToolBar_SetOnEndDock(t.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (t *TToolBar) SetOnEndDrag(fn TEndDragEvent) {
    ToolBar_SetOnEndDrag(t.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (t *TToolBar) SetOnEnter(fn TNotifyEvent) {
    ToolBar_SetOnEnter(t.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (t *TToolBar) SetOnExit(fn TNotifyEvent) {
    ToolBar_SetOnExit(t.instance, fn)
}

// SetOnGesture
func (t *TToolBar) SetOnGesture(fn TGestureEvent) {
    ToolBar_SetOnGesture(t.instance, fn)
}

// SetOnGetSiteInfo
func (t *TToolBar) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    ToolBar_SetOnGetSiteInfo(t.instance, fn)
}

// SetOnMouseActivate
func (t *TToolBar) SetOnMouseActivate(fn TMouseActivateEvent) {
    ToolBar_SetOnMouseActivate(t.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (t *TToolBar) SetOnMouseDown(fn TMouseEvent) {
    ToolBar_SetOnMouseDown(t.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (t *TToolBar) SetOnMouseEnter(fn TNotifyEvent) {
    ToolBar_SetOnMouseEnter(t.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (t *TToolBar) SetOnMouseLeave(fn TNotifyEvent) {
    ToolBar_SetOnMouseLeave(t.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (t *TToolBar) SetOnMouseMove(fn TMouseMoveEvent) {
    ToolBar_SetOnMouseMove(t.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (t *TToolBar) SetOnMouseUp(fn TMouseEvent) {
    ToolBar_SetOnMouseUp(t.instance, fn)
}

// SetOnResize
// CN: 设置大小被改变事件。
// EN: .
func (t *TToolBar) SetOnResize(fn TNotifyEvent) {
    ToolBar_SetOnResize(t.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (t *TToolBar) SetOnStartDock(fn TStartDockEvent) {
    ToolBar_SetOnStartDock(t.instance, fn)
}

// SetOnUnDock
func (t *TToolBar) SetOnUnDock(fn TUnDockEvent) {
    ToolBar_SetOnUnDock(t.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (t *TToolBar) DockClientCount() int32 {
    return ToolBar_GetDockClientCount(t.instance)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (t *TToolBar) AlignDisabled() bool {
    return ToolBar_GetAlignDisabled(t.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (t *TToolBar) MouseInClient() bool {
    return ToolBar_GetMouseInClient(t.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (t *TToolBar) VisibleDockClientCount() int32 {
    return ToolBar_GetVisibleDockClientCount(t.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (t *TToolBar) Brush() *TBrush {
    return BrushFromInst(ToolBar_GetBrush(t.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (t *TToolBar) ControlCount() int32 {
    return ToolBar_GetControlCount(t.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (t *TToolBar) Handle() HWND {
    return ToolBar_GetHandle(t.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (t *TToolBar) ParentWindow() HWND {
    return ToolBar_GetParentWindow(t.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (t *TToolBar) SetParentWindow(value HWND) {
    ToolBar_SetParentWindow(t.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (t *TToolBar) UseDockManager() bool {
    return ToolBar_GetUseDockManager(t.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (t *TToolBar) SetUseDockManager(value bool) {
    ToolBar_SetUseDockManager(t.instance, value)
}

// Action
func (t *TToolBar) Action() *TAction {
    return ActionFromInst(ToolBar_GetAction(t.instance))
}

// SetAction
func (t *TToolBar) SetAction(value IComponent) {
    ToolBar_SetAction(t.instance, CheckPtr(value))
}

// BiDiMode
func (t *TToolBar) BiDiMode() TBiDiMode {
    return ToolBar_GetBiDiMode(t.instance)
}

// SetBiDiMode
func (t *TToolBar) SetBiDiMode(value TBiDiMode) {
    ToolBar_SetBiDiMode(t.instance, value)
}

// BoundsRect
func (t *TToolBar) BoundsRect() TRect {
    return ToolBar_GetBoundsRect(t.instance)
}

// SetBoundsRect
func (t *TToolBar) SetBoundsRect(value TRect) {
    ToolBar_SetBoundsRect(t.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (t *TToolBar) ClientHeight() int32 {
    return ToolBar_GetClientHeight(t.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (t *TToolBar) SetClientHeight(value int32) {
    ToolBar_SetClientHeight(t.instance, value)
}

// ClientOrigin
func (t *TToolBar) ClientOrigin() TPoint {
    return ToolBar_GetClientOrigin(t.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (t *TToolBar) ClientRect() TRect {
    return ToolBar_GetClientRect(t.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (t *TToolBar) ClientWidth() int32 {
    return ToolBar_GetClientWidth(t.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (t *TToolBar) SetClientWidth(value int32) {
    ToolBar_SetClientWidth(t.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (t *TToolBar) ControlState() TControlState {
    return ToolBar_GetControlState(t.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (t *TToolBar) SetControlState(value TControlState) {
    ToolBar_SetControlState(t.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (t *TToolBar) ControlStyle() TControlStyle {
    return ToolBar_GetControlStyle(t.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (t *TToolBar) SetControlStyle(value TControlStyle) {
    ToolBar_SetControlStyle(t.instance, value)
}

// ExplicitLeft
func (t *TToolBar) ExplicitLeft() int32 {
    return ToolBar_GetExplicitLeft(t.instance)
}

// ExplicitTop
func (t *TToolBar) ExplicitTop() int32 {
    return ToolBar_GetExplicitTop(t.instance)
}

// ExplicitWidth
func (t *TToolBar) ExplicitWidth() int32 {
    return ToolBar_GetExplicitWidth(t.instance)
}

// ExplicitHeight
func (t *TToolBar) ExplicitHeight() int32 {
    return ToolBar_GetExplicitHeight(t.instance)
}

// Floating
func (t *TToolBar) Floating() bool {
    return ToolBar_GetFloating(t.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TToolBar) Parent() *TWinControl {
    return WinControlFromInst(ToolBar_GetParent(t.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (t *TToolBar) SetParent(value IWinControl) {
    ToolBar_SetParent(t.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (t *TToolBar) AlignWithMargins() bool {
    return ToolBar_GetAlignWithMargins(t.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (t *TToolBar) SetAlignWithMargins(value bool) {
    ToolBar_SetAlignWithMargins(t.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (t *TToolBar) Left() int32 {
    return ToolBar_GetLeft(t.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (t *TToolBar) SetLeft(value int32) {
    ToolBar_SetLeft(t.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (t *TToolBar) Top() int32 {
    return ToolBar_GetTop(t.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (t *TToolBar) SetTop(value int32) {
    ToolBar_SetTop(t.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (t *TToolBar) Width() int32 {
    return ToolBar_GetWidth(t.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (t *TToolBar) SetWidth(value int32) {
    ToolBar_SetWidth(t.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (t *TToolBar) Cursor() TCursor {
    return ToolBar_GetCursor(t.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (t *TToolBar) SetCursor(value TCursor) {
    ToolBar_SetCursor(t.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (t *TToolBar) Hint() string {
    return ToolBar_GetHint(t.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (t *TToolBar) SetHint(value string) {
    ToolBar_SetHint(t.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (t *TToolBar) Margins() *TMargins {
    return MarginsFromInst(ToolBar_GetMargins(t.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (t *TToolBar) SetMargins(value *TMargins) {
    ToolBar_SetMargins(t.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (t *TToolBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(ToolBar_GetCustomHint(t.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (t *TToolBar) SetCustomHint(value IComponent) {
    ToolBar_SetCustomHint(t.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TToolBar) ComponentCount() int32 {
    return ToolBar_GetComponentCount(t.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (t *TToolBar) ComponentIndex() int32 {
    return ToolBar_GetComponentIndex(t.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (t *TToolBar) SetComponentIndex(value int32) {
    ToolBar_SetComponentIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TToolBar) Owner() *TComponent {
    return ComponentFromInst(ToolBar_GetOwner(t.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TToolBar) Name() string {
    return ToolBar_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TToolBar) SetName(value string) {
    ToolBar_SetName(t.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TToolBar) Tag() int {
    return ToolBar_GetTag(t.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TToolBar) SetTag(value int) {
    ToolBar_SetTag(t.instance, value)
}

// Buttons
func (t *TToolBar) Buttons(Index int32) *TToolButton {
    return ToolButtonFromInst(ToolBar_GetButtons(t.instance, Index))
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (t *TToolBar) DockClients(Index int32) *TControl {
    return ControlFromInst(ToolBar_GetDockClients(t.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (t *TToolBar) Controls(Index int32) *TControl {
    return ControlFromInst(ToolBar_GetControls(t.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TToolBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ToolBar_GetComponents(t.instance, AIndex))
}

