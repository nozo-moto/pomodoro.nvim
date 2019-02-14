[![CircleCI](https://circleci.com/gh/nozo-moto/pomodoro.nvim.svg?style=svg)](https://circleci.com/gh/nozo-moto/pomodoro.nvim)

# Pomodoro.Nvim


``` vim
" use lightline 

let g:lightline = {
       \ 'active': { 
       \   'right': [
       \      [ 'pomodoro']
       \   ]
       \ },
       \ 'component_function': {
       \   'pomodoro': 'LightLinePmdrStatus'
       \ },
       \ }
function! LightLinePmdrStatus()
	return winwidth(0) > 70 ? PmdrStatus() : ''
endfunction
```

