// Copyright (c) 2018 Beta Kuang
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build darwin

package glfw

/*
#cgo CFLAGS: -x objective-c
#include "glfw/src/cocoa_init.m"
#include "glfw/src/cocoa_joystick.m"
#include "glfw/src/cocoa_monitor.m"
#include "glfw/src/cocoa_window.m"
#include "glfw/src/cocoa_time.c"
#include "glfw/src/posix_thread.c"
#include "glfw/src/nsgl_context.m"
*/
import "C"
