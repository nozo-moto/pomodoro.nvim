scriptencoding utf-8

if exists('g:loaded_pomodoro')
  finish
endif
let g:loaded_pomodoro = 1

let s:save_cpo = &cpo
set cpo&vim


function! s:RequirePomodoro(host) abort
  return jobstart(['pomodoro.nvim'], {'rpc': v:true})
endfunction

call remote#host#Register('pomodoro.nvim', '0', function('s:RequirePomodoro'))
call remote#host#RegisterPlugin('pomodoro.nvim', '0', [
    \ {'type': 'function', 'name': 'PmdrStart', 'sync': 0, 'opts': {}},
    \ {'type': 'function', 'name': 'PmdrStop', 'sync': 0, 'opts': {}},
    \ {'type': 'function', 'name': 'PmdrCancel', 'sync': 0, 'opts': {}},
    \ {'type': 'function', 'name': 'PmdrStatus', 'sync': 0, 'opts': {}},
    \ ])


let &cpo = s:save_cpo
unlet s:save_cpo
