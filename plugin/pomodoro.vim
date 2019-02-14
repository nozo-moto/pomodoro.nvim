if exists('g:loaded_pomodoro')
  finish
endif
let g:loaded_pomodoro = 1

function! s:RequirePomodoro(host) abort
  return jobstart(['pomodoro.nvim'], {'rpc': v:true})
endfunction

call remote#host#Register('pomodoro.nvim', 'x', function('s:RequirePomodoro'))
call remote#host#RegisterPlugin('pomodoro.nvim', '0', [
    \ {'type': 'function', 'name': 'pmdrStart', 'sync': 0, 'opts': {}},
    \ {'type': 'function', 'name': 'pmdrStop', 'sync': 0, 'opts': {}},
    \ {'type': 'function', 'name': 'pmdrCancel', 'sync': 0, 'opts': {}},
    \ ])


