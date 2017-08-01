# This file will install all the repos and all the plugins necessary for
# my vimrc to work.
# It WILL NOT override your current .vimrc. To do that, run this, then simply
# cp my .vimrc over yours.

# Check for dependencies
echo "Checking dependencies..."

curl --version
if [ $? -ne 0 ]; then
    echo "Please install curl."
    exit -1
fi

git --version
if [ $? -ne 0 ]; then
    echo "Please install git."
    exit -1
fi

vim --version
if [ $? -ne 0 ]; then
    echo "Please install vim."
    exit -1
fi

# Pathogen
if [ ! -d ~/.vim/bundle ]; then
    echo "Installing Pathogen..."
    mkdir -p ~/.vim/autoload ~/.vim/bundle
    curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim
fi

if [ $? -ne 0 ]; then
    echo "Something went wrong with installing pathogen. Are you connected to the internet?"
    exit -1
fi

# Rainbow Parens
if [ ! -d ~/.vim/bundle/rainbow_parentheses ]; then
    echo "Installing rainbow parentheses..."
    git clone https://github.com/kien/rainbow_parentheses.vim.git ~/.vim/bundle/rainbow_parentheses
fi
cd ~/.vim/bundle/rainbow_parentheses
vim -c q "helptags doc/"

# Elixir
if [ ! -d ~/.vim/bundle/vim-elixir ]; then
    echo "Installing elixir plugin..."
    git clone https://github.com/elixir-lang/vim-elixir.git ~/.vim/bundle/vim-elixir
fi
cd ~/.vim/bundle/vim-elixir
vim -c q "helptags doc/"

# CtrlP
if [ ! -d ~/.vim/bundle/ctrlp ]; then
    echo "Installing ctrlP..."
    git clone https://github.com/ctrlpvim/ctrlp.vim.git ~/.vim/bundle/ctrlp
fi
cd ~/.vim/bundle/ctrlp
vim -c q "helptags doc/"

# Conque
if [ ! -d ~/.vim/bundle/conque ]; then
    echo "Installing conque..."
    git clone https://github.com/wkentaro/conque.vim ~/.vim/bundle/conque
fi
cd ~/.vim/bundle/conque
vim -c q "helptags doc/"

# Buffergator
if [ ! -d ~/.vim/bundle/buffergator ]; then
    echo "Installing buffergator..."
    git clone https://github.com/jeetsukumaran/vim-buffergator.git ~/.vim/bundle/buffergator
fi
cd ~/.vim/bundle/buffergator
vim -c q "helptags doc/"

# Git fugitive
if [ ! -d ~/.vim/bundle/vim-fugitive ]; then
    echo "Installing git fugitive..."
    git clone git://github.com/tpope/vim-fugitive.git ~/.vim/bundle/vim-fugitive
fi
cd ~/.vim/bundle/vim-fugitive
vim -c q "helptags doc/"

# Simplyfold for python
if [ ! -d ~/.vim/bundle/SimplyFold ]; then
    echo "Installing Simplyfold..."
    git clone https://github.com/tmhedberg/SimpylFold.git ~/.vim/bundle/SimplyFold
fi
cd ~/.vim/bundle/SimplyFold
vim -c q "helptags doc/"

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
cd ~/.vim/bundle/vim-airline
vim -c q "helptags doc/"

echo "Done. Now copy the .vimrc from this directory to your home. Also, you may now need to change your terminal font to something for powerline."
