if exists('g:loaded_pomodoro')
  finish
endif
let g:loaded_pomodoro = 1

function! s:RequirePomodoro(host) abort
  return jobstart(['pomodoro.nvim'], {'rpc': v:true})
endfunction

call remote#host#Register('pomodoro.nvim', '0', function('s:RequirePomodoro'))
call remote#host#RegisterPlugin('pomodoro.nvim', '0', [
    \ {'type': 'function', 'name': 'PomodoroStart', 'sync': 1, 'opts': {}},
    \ {'type': 'function', 'name': 'PomodoroStop', 'sync': 1, 'opts': {}},
    \ {'type': 'function', 'name': 'PomodoroCancel', 'sync': 1, 'opts': {}},
    \ ])

