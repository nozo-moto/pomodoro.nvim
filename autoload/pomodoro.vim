scriptencoding utf-8

if exists('g:loaded_pomodoro')
  finish
endif
let g:loaded_pomodoro = 1

let s:save_cpo = &cpo
set cpo&vim

"

let &cpo = s:save_cpo
unlet s:save_cpo
