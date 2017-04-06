//+build !noasm
//+build !appengine

package gocvsimd

import "unsafe"

//go:noescape
func _ClibFloor32(fl float32) float32

//go:noescape
func _ClibFloor64(fl float64) float64

//go:noescape
func _ClibMemcpy(dst, src unsafe.Pointer, n uint) unsafe.Pointer
