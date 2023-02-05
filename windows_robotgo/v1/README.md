# ROBOT 文档

## 方法文档

- ActiveName
> ActivePID通过PID激活窗口，

- CaptureImg
> 捕获屏幕并返回图像。图像

- CaptureScreen
> CaptureScreen捕获屏幕返回位图（c结构）， 
> 使用`defer robotgo.FreeBitmap（位图）`释放位图

- CharCodeAt
> utf-8处理字符代码

- CheckMouse
> 检查鼠标按钮

- Click
> 单击鼠标按钮

- CloseWindow 
> 关闭窗口

- Drag
> 将鼠标拖动到（x，y）；不推荐：使用DragSmooth（），

- DragSmooth
> 将鼠标拖动到（x，y）

- FreeBitmap
> FreeBitmap释放并释放C位图

- GetActive
> 获取活动窗口

- GetHandle
> 拿到窗户句柄

- GetHandPid
> 通过pid获取句柄mdata

- GetMouseColor
> 获取鼠标位置的颜色

- GetMousePos
> 获取鼠标的位置返回x，y

- GetPID
> 获取进程id返回int32

- GetPixelColor
> 获取像素颜色返回字符串

- GetPxColor
> 获取像素颜色返回C.MMRGBHex

- GetScaleSize
> 获取屏幕比例大小

- GetScreenRect
> 获取屏幕矩形（x，y，w，h）

- GetScreenSize
> 获取屏幕尺寸

- GetTitle
> 获取窗口标题

- GetVersion
> 获取robotgo版本

- GetXDisplayName
> 获取XDisplay名称（Linux）

- GoCaptureScreen
> 捕获屏幕并返回位图（go结构）

- GoString
> 转化 C.char to string

- HexToRgb
> 转换 hex to rgb

- Is64Bit
> 确定系统是否为64位

- IsValid
> 判断窗口是否有效

- KeyDown
> 按下一个键

- KeyPress
> 按键字符串

- KeySleep
> 设置密钥默认毫秒睡眠时间

- KeyTap
> 点击键盘代码

- KeyToggle
> 切换键盘，如果没有参数，默认值为“down”

- KeyUp
> 按下一个键

- MaxWindow
> 设置窗口最大化

- MicroSleep
> time C.microsleep(tm)

- MilliSleep
> sleep tm milli second

- MinWindow
> 设置窗口最小化

- Move
> 移动鼠标至 (x, y)

- MoveArgs
> 获取相对位置 (x, y)的实际位置

- MoveClick
> 移动鼠标至 (x, y) 并点击

- MoveRelative
> 移动鼠标至相对位置 (x, y)

- MovesClick
> 移动鼠标至 (x, y) 并点击

- MoveSmooth
> 平滑移动鼠标至 (x, y)

- MoveSmoothRelative
> 平滑移动鼠标至相对位置 (x, y)

- PadHex
> trans C.MMRGBHex to string

- PasteStr
> 粘贴字符串，支持UTF-8

- ReadAll
> 从剪贴板读取字符串

- RgbToHex
> trans rgb to hex

- Scale【已删除】

- Scaled
> 获取屏幕缩放大小

- Scaled0(x int, f float64)
> return int(x * f)

- Scroll
> scroll the mouse to (x, y) 将鼠标滚动到（x，y）

- ScrollRelative
> 将鼠标滚动到相对位置（x，y）

- ScrollSmooth
> 平滑滚动鼠标至(to int, args ...int)

- SetActive
> 将窗口设置为活动

- SetDelay
> 设置键和鼠标延迟

- SetHandle
> set the window handle 设置窗口句柄

- SetHandlePid
> set the window handle by pid 通过PID设置窗口句柄

- SetKeyDelay
> 设置键盘延迟

- SetMouseDelay
> 设置鼠标延迟

- SetXDisplayName
> set XDisplay name (Linux)

- ShowAlert
> show a alert window，显示一个窗口

- Sleep
> 等待几秒

- SysScale
> 系统屏幕比例

- ToBitmap
> trans C.MMBitmapRef to Bitmap

- Toggle
> 切换鼠标，支持按钮：

- ToImage
> 将C.MMBitmapRef转换为标准图像

- ToMMRGBHex
> trans CHex to C.MMRGBHex

- ToRGBA
> convert C.MMBitmapRef to standard image.RGBA

- ToUC
> trans string to unicode []string

- Try
> handler(err)

- TypeStr
> send a string, support UTF-8

- TypeStrDelay
> type string delayed

- U8ToHex
> trans *C.uint8_t to C.MMRGBHex

- U32ToHex
> trans C.uint32_t to C.MMRGBHex

- UintToHex
> trans uint32 to robotgo.CHex

- UnicodeType
> tap uint32 unicode

- WriteAll
> 将字符串写入剪贴板










