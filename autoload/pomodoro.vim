if exists('g:loaded_pomodoro')
  finish
endif
let g:loaded_pomodoro = 1

let s:save_cpo = &cpoptions
set cpoptions&vim

let &cpoptions = s:save_cpo
unlet s:save_cpo
