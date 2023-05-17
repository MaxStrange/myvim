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

# CtrlP
if [ ! -d ~/.vim/bundle/ctrlp ]; then
    echo "Installing ctrlP..."
    git clone https://github.com/ctrlpvim/ctrlp.vim.git ~/.vim/bundle/ctrlp
fi
cd ~/.vim/bundle/ctrlp
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

# TMUX - plugin manager
mkdir -p $HOME/.tmux/plugins
git clone https://github.com/tmux-plugins/tpm $HOME/.tmux/plugins/tpm

echo "Done. Now copy the .vimrc from this directory to your home. Also, you may now need to change your terminal font to something for powerline."
echo "For TMUX save/restore functionality, you need to copy .tmux.conf into your home directory, open tmux and install the plug in with Ctrl-J-I"

