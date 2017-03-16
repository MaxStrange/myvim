# This file will install all the repos and all the plugins necessary for
# my vimrc to work.
# It WILL NOT override your current .vimrc. To do that, run this, then simply
# cp my .vimrc over yours.

# Pathogen
if [ ! -d ~/.vim/bundle ]; then
    echo "Installing Pathogen..."
    mkdir -p ~/.vim/autoload ~/.vim/bundle
    curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim
fi

# Rainbow Parens
if [ ! -d ~/.vim/bundle/rainbow_parentheses ]; then
    echo "Installing rainbow parentheses..."
    git clone https://github.com/kien/rainbow_parentheses.vim.git ~/.vim/bundle/rainbow_parentheses
fi

# Elixir
if [ ! -d ~/.vim/bundle/vim-elixir ]; then
    echo "Installing elixir plugin..."
    git clone https://github.com/elixir-lang/vim-elixir.git ~/.vim/bundle/vim-elixir
fi

# CtrlP
if [ ! -d ~/.vim/bundle/ctrlp ]; then
    echo "Installing ctrlP..."
    git clone https://github.com/ctrlpvim/ctrlp.vim.git ~/.vim/bundle/ctrlp
fi

# Buffergator
if [ ! -d ~/.vim/bundle/buffergator ]; then
    echo "Installing buffergator..."
    git clone https://github.com/jeetsukumaran/vim-buffergator.git ~/.vim/bundle/buffergator
fi

# Git fugitive
if [ ! -d ~/.vim/bundle/vim-fugitive ]; then
    echo "Installing git fugitive..."
    git clone git://github.com/tpope/vim-fugitive.git ~/.vim/bundle/vim-fugitive
fi

# Simplyfold for python
if [ ! -d ~/.vim/bundle/SimplyFold ]; then
    echo "Installing Simplyfold..."
    git clone https://github.com/tmhedberg/SimpylFold.git ~/.vim/bundle/SimplyFold
fi

# Vim airline
if [ ! -d ~/.vim/bundle/vim-airline ]; then
    if [ ! -d "~/fonts" ]; then
        echo "Installing airline..."
        git clone git://github.com/powerline/fonts ~/fonts
        cd ~/fonts
        ./install.sh
    fi
    cd ~/
    git clone https://github.com/vim-airline/vim-airline ~/.vim/bundle/vim-airline
fi

echo "Done. You now need to change your terminal font to something for powerline."
