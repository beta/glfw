// Copyright (c) 2018 Beta Kuang
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package glfw3

// #cgo pkg-config: glfw3
// #include <GLFW/glfw3.h>
import "C"

// GLFW version constants.
const (
	// VersionMajor : The major version number of the GLFW library.
	//
	// This is incremented when the API is changed in non-compatible ways.
	VersionMajor uint = 3
	// VersionMinor : The minor version number of the GLFW library.
	//
	// This is incremented when features are added to the API but it remains
	// backward-compatible.
	VersionMinor uint = 2
	// VersionRevision : The revision number of the GLFW library.
	//
	// This is incremented when a bug fix release is made that does not contain
	// any API changes.
	VersionRevision uint = 1
)

// Key and button actions.
const (
	// Release : The key or mouse button was released.
	Release = 0
	// Press : The key or mouse button was pressed.
	Press = 1
	// Repeat : The key was held down until it repeated.
	Repeat = 2
)

// Keyboard keys.
//
// These key codes are inspired by the _USB HID Usage Tables v1.12_ (p. 53-60),
// but re-arranged to map to 7-bit ASCII for printable keys (function keys are
// put in the 256+ range).
//
// The naming of the key codes follow these rules:
//  - The US keyboard layout is used
//  - Names of printable alpha-numeric characters are used (e.g. "A", "R",
//    "3", etc.)
//  - For non-alphanumeric characters, Unicode:ish names are used (e.g.
//    "COMMA", "LEFT_SQUARE_BRACKET", etc.). Note that some names do not
//    correspond to the Unicode standard (usually for brevity)
//  - Keys that lack a clear US mapping are named "WORLD_x"
//  - For non-printable keys, custom names are used (e.g. "F4",
//    "BACKSPACE", etc.)
const (
	// KeyUnknown : The unknown key.
	KeyUnknown = -1

	// Printable keys.

	KeySpace = 32
	// KeyApostrophe : "'".
	KeyApostrophe = 39
	// KeyComma : ",".
	KeyComma = 44
	// KeyMinus : "-".
	KeyMinus = 45
	// KeyPeriod : ".".
	KeyPeriod = 46
	// KeySlash : "/".
	KeySlash = 47
	Key0     = 48
	Key1     = 49
	Key2     = 50
	Key3     = 51
	Key4     = 52
	Key5     = 53
	Key6     = 54
	Key7     = 55
	Key8     = 56
	Key9     = 57
	// KeySemicolon : ";".
	KeySemicolon = 59
	// KeyEqual : "=".
	KeyEqual = 61
	KeyA     = 65
	KeyB     = 66
	KeyC     = 67
	KeyD     = 68
	KeyE     = 69
	KeyF     = 70
	KeyG     = 71
	KeyH     = 72
	KeyI     = 73
	KeyJ     = 74
	KeyK     = 75
	KeyL     = 76
	KeyM     = 77
	KeyN     = 78
	KeyO     = 79
	KeyP     = 80
	KeyQ     = 81
	KeyR     = 82
	KeyS     = 83
	KeyT     = 84
	KeyU     = 85
	KeyV     = 86
	KeyW     = 87
	KeyX     = 88
	KeyY     = 89
	KeyZ     = 90
	// KeyLeftBracket : "[".
	KeyLeftBracket = 91
	// KeyBackslash : "\".
	KeyBackslash = 92
	// KeyRightBracket : "]".
	KeyRightBracket = 93
	// KeyGraveAccent : "`".
	KeyGraveAccent = 96
	// KeyWorld1 : non-US #1.
	KeyWorld1 = 161
	// KeyWorld2 : non-US #2.
	KeyWorld2 = 162

	// Function keys.

	KeyEscape       = 256
	KeyEnter        = 257
	KeyTab          = 258
	KeyBackspace    = 259
	KeyInsert       = 260
	KeyDelete       = 261
	KeyRight        = 262
	KeyLeft         = 263
	KeyDown         = 264
	KeyUp           = 265
	KeyPageUp       = 266
	KeyPageDown     = 267
	KeyHome         = 268
	KeyEnd          = 269
	KeyCapsLock     = 280
	KeyScrollLock   = 281
	KeyNumLock      = 282
	KeyPrintScreen  = 283
	KeyPause        = 284
	KeyF1           = 290
	KeyF2           = 291
	KeyF3           = 292
	KeyF4           = 293
	KeyF5           = 294
	KeyF6           = 295
	KeyF7           = 296
	KeyF8           = 297
	KeyF9           = 298
	KeyF10          = 299
	KeyF11          = 300
	KeyF12          = 301
	KeyF13          = 302
	KeyF14          = 303
	KeyF15          = 304
	KeyF16          = 305
	KeyF17          = 306
	KeyF18          = 307
	KeyF19          = 308
	KeyF20          = 309
	KeyF21          = 310
	KeyF22          = 311
	KeyF23          = 312
	KeyF24          = 313
	KeyF25          = 314
	KeyKp0          = 320
	KeyKp1          = 321
	KeyKp2          = 322
	KeyKp3          = 323
	KeyKp4          = 324
	KeyKp5          = 325
	KeyKp6          = 326
	KeyKp7          = 327
	KeyKp8          = 328
	KeyKp9          = 329
	KeyKpDecimal    = 330
	KeyKpDivide     = 331
	KeyKpMultiply   = 332
	KeyKpSubtract   = 333
	KeyKpAdd        = 334
	KeyKpEnter      = 335
	KeyKpEqual      = 336
	KeyLeftShift    = 340
	KeyLeftControl  = 341
	KeyLeftAlt      = 342
	KeyLeftSuper    = 343
	KeyRightShift   = 344
	KeyRightControl = 345
	KeyRightAlt     = 346
	KeyRightSuper   = 347
	KeyMenu         = 348

	KeyLast = KeyMenu
)

// Modifier key flags.
const (
	// ModShift : If this bit is set one or more Shift keys were held down.
	ModShift = 0x0001
	// ModControl : If this bit is set one or more Control keys were held down.
	ModControl = 0x0002
	// ModAlt : If this bit is set one or more Alt keys were held down.
	ModAlt = 0x0004
	// ModSuper : If this bit is set one or more Super keys were held down.
	ModSuper = 0x0008
)

// Mouse buttons.
const (
	MouseButton1      = 0
	MouseButton2      = 1
	MouseButton3      = 2
	MouseButton4      = 3
	MouseButton5      = 4
	MouseButton6      = 5
	MouseButton7      = 6
	MouseButton8      = 7
	MouseButtonLast   = MouseButton8
	MouseButtonLeft   = MouseButton1
	MouseButtonRight  = MouseButton2
	MouseButtonMiddle = MouseButton3
)

// Joysticks.
const (
	Joystick1    = 0
	Joystick2    = 1
	Joystick3    = 2
	Joystick4    = 3
	Joystick5    = 4
	Joystick6    = 5
	Joystick7    = 6
	Joystick8    = 7
	Joystick9    = 8
	Joystick10   = 9
	Joystick11   = 10
	Joystick12   = 11
	Joystick13   = 12
	Joystick14   = 13
	Joystick15   = 14
	Joystick16   = 15
	JoystickLast = Joystick16
)

// Error codes.
const (
	// NotInitialized : GLFW has not been initialized.
	//
	// This occurs if a GLFW function was called that must not be called unless
	// the library is initialized.
	//
	// Analysis:
	//
	// This is an application programmer error. Initialize GLFW before calling
	// any function that requires initialization.
	NotInitialized = 0x00010001
	// NoCurrentContext : No context is current for this thread.
	//
	// This occurs if a GLFW function was called that needs and operates on the
	// current OpenGL or OpenGL ES contest but no context is current on the
	// calling thread. One such function is SwapInterval().
	//
	// Analysis:
	//
	// This is an application programmer error. Ensure a context is current
	// before calling functions that require a current context.
	NoCurrentContext = 0x00010002
	// InvalidEnum : One or the arguments to the function was an invalid enum
	// value, for example requesting RedBits with GetWindowAttrib.
	//
	// Analysis:
	//
	// This is an application programmer error. Fix the offending call.
	InvalidEnum = 0x00010003
	// InvalidValue : One of the arguments to the function was an invalid value,
	// for example requesting a non-existent OpenGL or OpenGL ES version like
	// 2.7.
	//
	// Requesting a valid but unavailable OpenGL or OpenGL ES version will
	// instead result in a VersionUnavailable error.
	//
	// Analysis:
	//
	// This is an application programmer error. Fix the offending call.
	InvalidValue = 0x00010004
	// OutOfMemory : A memory allocation failed.
	//
	// Analysis:
	//
	// This is a bug in GLFW3, GLFW or the underlying operating system. Report
	// the bug to our issue tracker (https://github.com/paperui/glfw3/issues) or
	// GLFW's (https://github.com/glfw/glfw/issues).
	OutOfMemory = 0x00010005
	// APIUnavailable : GLFW could not find support for the requested API on the
	// system.
	//
	// Analysis:
	//
	// Thhe installed graphics driver does not support the requested API, or
	// does not support it via the chosen context creation backend. Below are a
	// few examples.
	//
	// Some pre-installed Windows graphics drivers do not support OpenGL. AMD
	// only supports OpenGL ES via EGL, while Nvidia and Intel only support it
	// via a WGL or GLX extension. OS X does not provide OpenGL ES at all. The
	// Mesa EGL, OpenGL and OpenGL ES libraries do not interface with the
	// Nvidia binary driver. Older graphics drivers do not support Vulkan.
	APIUnavailable = 0x00010006
	// VersionUnavailable : The requested OpenGL or OpenGL ES version (including
	// any requested context or framebuffer hints) is not available on this
	// machine.
	//
	// Analysis:
	//
	// The machine does not support your requirements. If your application is
	// sufficiently flexible, downgrade your requirements and try again.
	// Otherwise, inform the user that their machine does not match your
	// requirements.
	//
	// Future invalid OpenGL and OpenGL ES versions, for example OpenGL 4.8 if
	// 5.0 comes out before the 4.x series gets that far, also fail with this
	// error and not InvalidValue, because GLFW cannot know what future versions
	// will exist.
	VersionUnavailable = 0x00010007
	// PlatformError : A platform-specific error occurred that does not match
	// any of the more specific categories.
	//
	// Analysis:
	//
	// This is a bug or configuration error in GLFW3, GLFW, the underlying
	// operating system or its drivers, or a lack of required resources. Report
	// the issue to our issue tracker (https://github.com/paperui/glfw3/issues)
	// or GLFW's (https://github.com/glfw/glfw/issues).
	PlatformError = 0x00010008
	// FormatUnavailable : The requested format is not supported or available.
	//
	// If emitted during window creation, the requested pixel format is not
	// supported.
	//
	// If emitted when querying the clipboard, the contents of the clipboard
	// could not be converted to the requested format.
	//
	// Analysis:
	//
	// If emmited during window creation, one or more hard constraints
	// (http://www.glfw.org/docs/latest/window_guide.html#window_hints_hard) did
	// not match any of the available pixel formats. If your application is
	// sufficiently flexible, downgrade your requirements and try again.
	// Otherwise, inform the user that their machine does not match your
	// requirements.
	//
	// If emitted when querying the clipboard, ignore the error or report it to
	// the user, as appropriate.
	FormatUnavailable = 0x00010009
	// NoWindowContext : A window that does not have an OpenGL or OpenGL ES was
	// passed to a function that requires it to have one.
	//
	// Analysis:
	//
	// This is an application programmer error. Fix the offending call.
	NoWindowContext = 0x0001000A
)

// Constants.
const (
	Focused     = 0x00020001
	Iconified   = 0x00020002
	Resizable   = 0x00020003
	Visible     = 0x00020004
	Decorated   = 0x00020005
	AutoIconify = 0x00020006
	Floating    = 0x00020007
	Maximized   = 0x00020008

	RedBits        = 0x00021001
	GreenBits      = 0x00021002
	BlueBits       = 0x00021003
	AlphaBits      = 0x00021004
	DepthBits      = 0x00021005
	StencilBits    = 0x00021006
	AccumRedBits   = 0x00021007
	AccumGreenBits = 0x00021008
	AccumBlueBits  = 0x00021009
	AccumAlphaBits = 0x0002100A
	AuxBuffers     = 0x0002100B
	Stereo         = 0x0002100C
	Samples        = 0x0002100D
	SRGBCapable    = 0x0002100E
	RefreshRate    = 0x0002100F
	Doublebuffer   = 0x00021010

	ClientAPI              = 0x00022001
	ContextVersionMajor    = 0x00022002
	ContextVersionMinor    = 0x00022003
	ContextRevision        = 0x00022004
	ContextRobustness      = 0x00022005
	OpenGLForwardCompat    = 0x00022006
	OpenGLDebugContext     = 0x00022007
	OpenGLProfile          = 0x00022008
	ContextReleaseBehavior = 0x00022009
	ContextNoError         = 0x0002200A
	ContextCreationAPI     = 0x0002200B

	NoAPI       = 0
	OpenGLAPI   = 0x00030001
	OpenGLESAPI = 0x00030002

	NoRobustness        = 0
	NoResetNotification = 0x00031001
	LoseContextOnReset  = 0x00031002

	OpenGLAnyProfile    = 0
	OpenGLCoreProfile   = 0x00032001
	OpenGLCompatProfile = 0x00032002

	Cursor             = 0x00033001
	StickyKeys         = 0x00033002
	StickyMouseButtons = 0x00033003

	CursorNormal   = 0x00034001
	CursorHidden   = 0x00034002
	CursorDisabled = 0x00034003

	AnyReleaseBehavior   = 0
	ReleaseBehaviorFlush = 0x00035001
	ReleaseBehaviorNone  = 0x00035002

	NativeContextAPI = 0x00036001
	EGLContextAPI    = 0x00036002

	Connected    = 0x00040001
	Disconnected = 0x00040002

	DontCare = -1
)

// Standard cursor shapes.
const (
	// ArrowCursor : The regular arrow cursor shape.
	ArrowCursor = 0x00036001
	// IBeamCursor : The text input I-beam cursor shape.
	IBeamCursor = 0x00036002
	// CrosshairCursor : The crosshair shape.
	CrosshairCursor = 0x00036003
	// HandCursor : The hand shape.
	HandCursor = 0x00036004
	// HResizeCursor : The horizontal resize arrow shape.
	HResizeCursor = 0x00036005
	// VResizeCursor : The vertical resize arrow shape.
	VResizeCursor = 0x00036006
)
