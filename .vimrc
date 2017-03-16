" Enable line numbers
set number

" Colorscheme
colorscheme desert

" Syntax highlighting
syntax on

" Turn off Mac OSX stupid backspace behavior
set backspace=2

" Remap <Tab> to be switch to next buffer
:nnoremap <Tab> :bNext<CR>

" Remap <C-X> to close current buffer
:nnoremap <C-X> :bdelete<CR>

" Highlight trailing whitespace
highlight ExtraWhitespace ctermbg=red guibg=red
match ExtraWhitespace /\s\+$/
augroup whitespace
    autocmd!
    autocmd BufWinEnter * match ExtraWhitespace /\s\+$/
    autocmd InsertEnter * match ExtraWhitespace /\s\+\%#\@<!$/
    autocmd InsertLeave * match ExtraWhitespace /\s\+$/
    autocmd BufWinLeave * call clearmatches()
augroup END


" Enable tabexpand except when editing Makefiles
let _curfile = expand("%:t")
if _curfile =~ "Makefile" || _curfile =~ "makefile" || _curfile =~ ".*\.mk"
    set noexpandtab
else
    set expandtab
    set tabstop=4
    set shiftwidth=4
endif


" RainbowParentheses
let g:rbpt_colorpairs = [
    \ ['brown',       'RoyalBlue3'],
    \ ['Darkblue',    'SeaGreen3'],
    \ ['darkgray',    'DarkOrchid3'],
    \ ['darkgreen',   'firebrick3'],
    \ ['darkcyan',    'RoyalBlue3'],
    \ ['darkred',     'SeaGreen3'],
    \ ['darkmagenta', 'DarkOrchid3'],
    \ ['brown',       'firebrick3'],
    \ ['gray',        'RoyalBlue3'],
    \ ['black',       'SeaGreen3'],
    \ ['darkmagenta', 'DarkOrchid3'],
    \ ['Darkblue',    'firebrick3'],
    \ ['darkgreen',   'RoyalBlue3'],
    \ ['darkcyan',    'SeaGreen3'],
    \ ['darkred',     'DarkOrchid3'],
    \ ['red',         'firebrick3'],
    \ ]


" Configure for language
augroup langgroup
    autocmd!
    " Python
    autocmd Filetype python set foldmethod=indent
    autocmd FileType python let g:SimplyFold_docstring_preview=1
    " C, C++, CUDA
    autocmd Filetype c,cpp,cuda set foldmethod=syntax
    " Clojure
    autocmd FileType clojure RainbowParenthesesToggle
    autocmd FileType clojure RainbowParenthesesLoadRound
    autocmd FileType clojure set tabstop=2
    autocmd FileType clojure set shiftwidth=2
    autocmd FileType clojure set autoindent
    " Elixir - some of this is turned off by the plugin TODO: Fix it
    autocmd Filetype elixir set foldmethod=indent
    autocmd FileType elixir set tabstop=2
    autocmd FileType elixir set shiftwidth=2
    autocmd FileType elixir set autoindent
augroup END


" Code folding
set foldlevel=99

" Powerline - disabled: I am using vim-airline instead
"set rtp+=/Users/maxst/Library/Python/2.7/lib/python/site-packages/powerline/bindings/vim/

" Use 256 colors
set t_Co=256

" Airline
let g:airline#extensions#tabline#enabled = 1
let g:airline_powerline_fonts = 1
set laststatus=2

" Pathogen
" With Pathogen, you can now just cd into ~/.vim/bundle and
" clone whatever plugin repo you want to install:
" e.g.: git clone https://github.com/tpope/vim-sensible.git when in
" ~/.vim/bundle and voila, you have sensible.vim installed
execute pathogen#infect()
